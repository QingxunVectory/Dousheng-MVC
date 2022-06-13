package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"time"
)

func Suscribe(toUserId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		return err
	}
	userId := user.Id
	createdRelation := &model.UserRelation{
		UserID:    user.Id,
		ToUserID:  toUserId,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err = repository.CreateRelation(createdRelation)
	if err != nil {
		return err
	}
	err = repository.UpdateFollowerCountPlus(toUserId)
	if err != nil {
		return err
	}
	err = repository.UpdateFollowCountPlus(userId)
	if err != nil {
		return err
	}
	return nil
}

func CancelSuscribe(toUserId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		return err
	}
	userId := user.Id

	err = repository.DeleteFavoriteByUserIDAndFollowerId(user.Id, toUserId)
	if err != nil {
		return err
	}
	err = repository.UpdateFollowerCountMinus(toUserId)
	if err != nil {
		return err
	}
	err = repository.UpdateFollowCountMinus(userId)
	if err != nil {
		return err
	}
	return nil
}

func GetConcernsByUserId(userId int64) ([]model.User, error){
	ToUsers, err := repository.GetToUserIdByUserId(userId)
	if err != nil {
		return nil, err
	}
	users := []model.User{}
	for _, toUser := range ToUsers {
		user, err := repository.GetUsersByToUserId(toUser.ToUserID)
		if user == nil {
			panic("user为空")
		}
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func GetFansByTouserId(toUserId int64) ([]model.User, error){
	users, err := repository.GetUserIdByToUserId(toUserId)
	if err != nil {
		return nil, err
	}
	tousers := []model.User{}
	for _, toUser := range users {
		user, err := repository.GetUsersByToUserId(toUser.UserID)
		if user == nil {
			panic("user为空")
		}
		if err != nil {
			return nil, err
		}
		tousers = append(tousers, *user)
	}
	return tousers, nil
}
