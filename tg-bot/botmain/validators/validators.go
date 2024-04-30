package validators

import (
	"context"
	"log"
	"time"

	"github.com/Simplewallethq/tg-bot/botmain/restclient"
	"github.com/Simplewallethq/tg-bot/ent"
	entval "github.com/Simplewallethq/tg-bot/ent/validators"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Crawler struct {
	DB         *ent.Client
	Restclient *restclient.Client
	logger     *logrus.Logger
	RPCNode    string
}

func NewValidatorsCrawler(DB *ent.Client, resthost string, rpcNode string, logger *logrus.Logger) *Crawler {
	return &Crawler{
		DB:         DB,
		Restclient: restclient.NewClient(resthost),
		logger:     logger,
		RPCNode:    rpcNode,
	}
}
func NewValidatorsCrawlerWithToken(DB *ent.Client, resthost string, rpcNode string, logger *logrus.Logger, RESTtoken string) *Crawler {
	return &Crawler{
		DB:         DB,
		Restclient: restclient.NewClientWithToken(resthost, RESTtoken),
		logger:     logger,
		RPCNode:    rpcNode,
	}
}

func (v *Crawler) Start() {
	time.Sleep(5 * time.Second)
	err := v.ScanNetwork()
	if err != nil {
		v.logger.Error(err)
	}
	log.Println("start validators crawler")
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		err := v.ScanNetwork()
		if err != nil {
			v.logger.Error(err)
		}
	}
}

func (v *Crawler) ScanNetwork() error {
	state, err := v.Restclient.GetState(v.RPCNode)
	if err != nil {
		return err
	}

	optcount, err := v.DB.Settings.Query().Count(context.Background())
	if err != nil {
		return err
	}
	if optcount == 0 {
		_, err = v.DB.Settings.Create().SetLastScannedEraValidators(int64(state.CurrentEra)).Save(context.Background())
		if err != nil {
			return err
		}
		err = v.GetValidators()
		if err != nil {
			return err
		}
	}

	settings, err := v.DB.Settings.Query().Only(context.Background())
	if err != nil {
		return err
	}

	if settings.LastScannedEraValidators < int64(state.CurrentEra) {
		err = v.GetValidators()
		if err != nil {
			return err
		}
		settings, err = settings.Update().SetLastScannedEraValidators(int64(state.CurrentEra)).Save(context.Background())
		if err != nil {
			return err
		}
	}

	log.Println("scan validators. Last era in db: ", settings.LastScannedEraValidators, "now era: ", state.CurrentEra)
	return nil
}

func (v *Crawler) GetValidators() error {
	validators, err := v.Restclient.GetValidators(v.RPCNode)
	if err != nil {
		return errors.Wrap(err, "failed to get validators")
	}
	if len(validators.Validators) > 0 {
		err = v.DB.Validators.Update().SetActive(false).Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed trim validators activity")
		}
	} else {
		return errors.Wrap(err, "error with validators response")
	}

	for _, validator := range validators.Validators {
		err = v.DB.Validators.Create().SetAddress(validator.Address).
			SetActive(validator.Active).
			SetFee(int8(validator.Fee)).
			SetDelegators(validator.Delegators).OnConflictColumns(entval.FieldAddress).UpdateNewValues().Exec(context.Background())
		if err != nil {
			return errors.Wrap(err, "failed update validators")
		}
	}
	return nil
}
