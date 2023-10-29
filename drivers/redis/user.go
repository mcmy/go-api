package redis

import (
	"api/config"
	"api/drivers/db/model"
	"api/utils"
	"context"
	"time"
)

func SetUser(token string, user model.User) error {
	var rdb = Redis

	serializedUser, err := utils.ByteSerialize(user)
	if err != nil {
		return err
	}

	liveTime := time.Second * time.Duration(config.T.Http.SessionLiveSecond)
	err = rdb.Set(context.Background(), GenKey(token, "user"), serializedUser, liveTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUser(token string) (model.User, error) {
	var rdb = Redis

	var user model.User
	serializedUser, err := rdb.Get(context.Background(), GenKey(token, "user")).Bytes()
	if err != nil {
		return user, err
	}

	err = utils.ByteDeserialize(serializedUser, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
