package userstate

import (
	"context"
	"strings"

	"github.com/Simplewallethq/tg-bot/ent"
	"github.com/Simplewallethq/tg-bot/ent/user"
	"github.com/sirupsen/logrus"
)

type State struct {
	DB     *ent.Client
	logger *logrus.Logger
}

func NewState(DB *ent.Client, logger *logrus.Logger) *State {
	return &State{
		DB:     DB,
		logger: logger,
	}
}

func (s *State) GetUserState(id int64) ([]string, error) {
	u, err := s.DB.User.Query().Where(user.IDEQ(id)).Only(context.Background())
	if err != nil {
		s.logger.Error(err)
		return nil, err
	}
	stateExist, err := u.QueryState().Exist(context.Background())
	if err != nil {
		return nil, nil
	}
	if !stateExist {
		return nil, nil
	}
	state, err := u.QueryState().Only(context.Background())
	if err != nil {
		return nil, err
	}
	result := strings.Split(state.State, "/")
	return result, nil
}

func (s *State) SetUserState(id int64, newState []string) error {
	u, err := s.DB.User.Query().Where(user.IDEQ(id)).Only(context.Background())
	if err != nil {
		s.logger.Error(err)
		return err
	}
	stateExist, err := u.QueryState().Exist(context.Background())
	if err != nil {
		return nil
	}
	if stateExist {
		state, err := u.QueryState().Only(context.Background())
		if err != nil {
			return err
		}
		_, err = state.Update().SetState(strings.Join(newState, "/")).Save(context.Background())
		if err != nil {
			return err
		}
	} else {
		_, err := s.DB.UserState.Create().SetState(strings.Join(newState, "/")).SetOwner(u).Save(context.Background())
		if err != nil {
			return err
		}

	}
	return nil

}

func (s *State) DeleteUserState(id int64) error {
	u, err := s.DB.User.Query().Where(user.IDEQ(id)).Only(context.Background())
	if err != nil {
		s.logger.Error(err)
		return err
	}
	stateExist, err := u.QueryState().Exist(context.Background())
	if err != nil {
		return nil
	}
	if stateExist {
		state, err := u.QueryState().Only(context.Background())
		if err != nil {
			return err
		}
		err = s.DB.UserState.DeleteOne(state).Exec(context.Background())
		if err != nil {
			return err
		}
	}
	return nil

}
