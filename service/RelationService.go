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
