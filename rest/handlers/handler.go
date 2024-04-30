package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Simplewallethq/rest-api/blockchain/casper"

	"github.com/Simplewallethq/rest-api/middleware/mw_logger"
	"github.com/Simplewallethq/rest-api/middleware/mwauth"
	"github.com/Simplewallethq/rest-api/middleware/timeout"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	// swagger embed files
	// gin-swagger middleware
)

type GinConfig struct {
	Port string
	Mode string
}

type Auth struct {
	Auth  bool
	Token string
}
type Handler struct {
	router        *gin.Engine
	caspermainnet *casper.Casper
	caspertestnet *casper.Casper
	logger        *logrus.Logger
	auth          Auth
}

func NewHandler(logger *logrus.Logger, csprmain *casper.Casper, csprtest *casper.Casper, swagger gin.HandlerFunc, auth Auth) *Handler {
	handler := &Handler{
		router:        nil,
		caspermainnet: csprmain,
		caspertestnet: csprtest,
		logger:        logger,
		auth:          auth,
	}
	handler.InitRoutes(swagger)
	return handler
}

func (h *Handler) InitRoutes(swaggerdoc gin.HandlerFunc) {

	r := gin.New()

	mwlogger := mw_logger.New(h.logger) //using middleware to log gin events
	r.Use(mwlogger.LoggingMiddleware())
	r.Use(timeout.TimeoutMiddleware(10 * time.Minute))
	r.Use(gin.Recovery()) //NEED use after all middlewares
	//fmt.Println(h.auth)
	mwauth := mwauth.AuthMiddleware(h.auth.Auth, h.auth.Token)

	apiV1 := r.Group("/api/v1", mwauth)
	{
		mainnet_mw_checkrpc := h.caspermainnet.CheckChainMainnet()
		testnet_mw_checkrpc := h.caspertestnet.CheckChainTestnet()
		mainnet_mw_check_adress := h.caspermainnet.CheckAddress()
		testnet_mw_check_adress := h.caspertestnet.CheckAddress()
		cspr := apiV1.Group("/cspr", mainnet_mw_checkrpc, mainnet_mw_check_adress)
		{
			cspr.GET("/state", h.caspermainnet.GetStateHandler)
			cspr.GET("/is_address", h.caspermainnet.IsAddressHandler)
			cspr.GET("/get_balance_main", h.caspermainnet.GetAccountBalanceHandler)
			cspr.GET("/get_balance_delegated", h.caspermainnet.GetDelegatedBalanceHandler)
			cspr.GET("/get_history_transfers", h.caspermainnet.GetHistoryTransfers)
			cspr.GET("/get_rewards_by_blocks", h.caspermainnet.GetRewardsByBlock)
			cspr.GET("/get_rewards_by_era", h.caspermainnet.GetRewardsByEra)
			cspr.GET("/get_timestamp_by_block", h.caspermainnet.GetTimestampByBlock)
			cspr.GET("/get_timestamp_by_era", h.caspermainnet.GetTimestampByEra)
			cspr.GET("/get_block_by_timestamp", h.caspermainnet.GetBlockByTimestamp)
			cspr.GET("/get_era_by_timestamp", h.caspermainnet.GetEraByTimestamp)
			cspr.GET("/get_balance_being_undelegated", h.caspermainnet.GetBalanceBeingUndelegated)
			cspr.GET("/get_history_undelegate", h.caspermainnet.GetHistoryUndelegate)
			cspr.GET("/get_price_main_coin", h.caspermainnet.GetPriceMainCoin)
			cspr.GET("/get_history_delegate", h.caspermainnet.GetHistoryDelegate)
			cspr.GET("/get_apr_by_era", h.caspermainnet.GetAPRByEra)
			cspr.GET("/get_balance_being_delegated", h.caspermainnet.GetBalanceBeingdelegated)
			cspr.GET("/calculate_current_chain_apy", h.caspermainnet.CalculateCurrentChainAPY)
			cspr.GET("/get_validators", h.caspermainnet.GetValidators)
			cspr.GET("/get_block_events", h.caspermainnet.GetBlockEvents)
			cspr.POST("/put_deploy", h.caspermainnet.PutDeploy)
		}
		cspr_test := apiV1.Group("/cspr-testnet", testnet_mw_checkrpc, testnet_mw_check_adress)
		{
			cspr_test.GET("/state", h.caspertestnet.GetStateHandler)
			cspr_test.GET("/is_address", h.caspertestnet.IsAddressHandler)
			cspr_test.GET("/get_balance_main", h.caspertestnet.GetAccountBalanceHandler)
			cspr_test.GET("/get_balance_delegated", h.caspertestnet.GetDelegatedBalanceHandler)
			cspr_test.GET("/get_history_transfers", h.caspertestnet.GetHistoryTransfers)
			cspr_test.GET("/get_rewards_by_blocks", h.caspertestnet.GetRewardsByBlock)
			cspr_test.GET("/get_rewards_by_era", h.caspertestnet.GetRewardsByEra)
			cspr_test.GET("/get_timestamp_by_block", h.caspertestnet.GetTimestampByBlock)
			cspr_test.GET("/get_timestamp_by_era", h.caspertestnet.GetTimestampByEra)
			cspr_test.GET("/get_block_by_timestamp", h.caspertestnet.GetBlockByTimestamp)
			cspr_test.GET("/get_era_by_timestamp", h.caspertestnet.GetEraByTimestamp)
			cspr_test.GET("/get_balance_being_undelegated", h.caspertestnet.GetBalanceBeingUndelegated)
			cspr_test.GET("/get_history_undelegate", h.caspertestnet.GetHistoryUndelegate)
			cspr_test.GET("/get_price_main_coin", h.caspertestnet.GetPriceMainCoin)
			cspr_test.GET("/get_history_delegate", h.caspertestnet.GetHistoryDelegate)
			cspr_test.GET("/get_apr_by_era", h.caspertestnet.GetAPRByEra)
			cspr_test.GET("/get_balance_being_delegated", h.caspertestnet.GetBalanceBeingdelegated)
			cspr_test.GET("/calculate_current_chain_apy", h.caspertestnet.CalculateCurrentChainAPY)
			cspr_test.GET("/get_validators", h.caspertestnet.GetValidators)
			cspr_test.GET("/get_block_events", h.caspertestnet.GetBlockEvents)
			cspr_test.POST("/put_deploy", h.caspertestnet.PutDeploy)
		}

	}
	r.GET("/swagger/*any", swaggerdoc)
	h.router = r

}

func (h *Handler) Start(ctx context.Context, config GinConfig) error {
	gin.SetMode(config.Mode)
	srv := &http.Server{
		Addr:    config.Port,
		Handler: h.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			h.logger.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	h.logger.Println("Shutting down server...")

	// Set a timeout for the graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		h.logger.Fatalf("Server forced to shutdown: %s\n", err)
	}

	h.logger.Println("Server exiting")
	return nil
}
