package user

import (
	"aura-test/env"
	"aura-test/pkg/config"
	"aura-test/pkg/log"
	"aura-test/repository/user"
	"context"
	"encoding/json"
	"errors"
	"time"

	"aura-test/helper/utils"

	"aura-test/pkg/redis"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type (
	taskOfLogin struct {
		req     *ReqOfLogin
		storage *storageOfLogin
	}

	ReqOfLogin struct {
		Username string `json:"Username" binding:"required" example:"test"`
		Password string `json:"Password" binding:"required" example:"12345"`
	}

	storageOfLogin struct {
		User *user.ResOfInfo
	}
)

func Login(ctx context.Context, in *ReqOfLogin) error {
	task := newTaskOfLogin(in)
	if err := task.exec(ctx); err != nil {
		return err
	}

	return nil
}

func newTaskOfLogin(in *ReqOfLogin) *taskOfLogin {
	return &taskOfLogin{
		req:     in,
		storage: &storageOfLogin{},
	}
}

func (t *taskOfLogin) exec(ctx context.Context) error {
	if err := t.validate(ctx); err != nil {
		return err
	}

	//if err := t.setData(ctx); err != nil {
	//	return err
	//}

	return nil
}

func (t *taskOfLogin) validate(ctx context.Context) error {
	userInfo, err := user.GetUserInfo(ctx, t.req.Username)
	if err != nil {
		return errors.New("get user info failed")
	}
	if userInfo == nil {
		return errors.New("user not exist")
	}

	if bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(t.req.Password)) != nil {
		return errors.New("wrong password")
	}

	t.storage.User = userInfo

	return nil
}

func (t *taskOfLogin) setData(ctx context.Context) error {
	client := redis.NewClient(ctx, redis.UserDB)

	token := utils.GenerateToken()

	userToJson, err := json.Marshal(t.storage.User)
	if err != nil {
		log.Error(zap.Any("user data to json failed error:", err))
		return errors.New("json encode failed")
	}

	apiToken := "user._." + t.storage.User.UUID + "._." + t.req.Username + token

	client.Set(ctx, apiToken, string(userToJson), config.ConForge().GetDuration(env.TokenExpireTime)*time.Second)

	return nil
}
