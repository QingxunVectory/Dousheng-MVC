package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/sirupsen/logrus"
	"time"
)

func Suscribe(toUserId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		logrus.Errorf("[Suscribe] ParseToken failed ,the error is %s", err)
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		logrus.Errorf("[Suscribe] GetUserByUserName failed ,the error is %s", err)
		return err
	}
	if user.Id == toUserId {
		return errors.New("can not suscribe yourself")
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
		logrus.Errorf("[Suscribe] CreateRelation failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateFollowerCountPlus(toUserId)
	if err != nil {
		logrus.Errorf("[Suscribe] CreateRelation failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateFollowCountPlus(userId)
	if err != nil {
		logrus.Errorf("[Suscribe] UpdateFollowCountPlus failed ,the error is %s", err)
		return err
	}
	return nil
}

func CancelSuscribe(toUserId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		logrus.Errorf("[CancelSuscribe] ParseToken failed ,the error is %s", err)
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		logrus.Errorf("[CancelSuscribe] GetUserByUserName failed ,the error is %s", err)
		return err
	}
	userId := user.Id
	if user.Id == toUserId {
		return errors.New("can not cancelSuscribe yourself")
	}
	err = repository.DeleteFavoriteByUserIDAndFollowerId(user.Id, toUserId)
	if err != nil {
		logrus.Errorf("[CancelSuscribe] DeleteFavoriteByUserIDAndFollowerId failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateFollowerCountMinus(toUserId)
	if err != nil {
		logrus.Errorf("[CancelSuscribe] UpdateFollowerCountMinus failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateFollowCountMinus(userId)
	if err != nil {
		logrus.Errorf("[UpdateFollowCountMinus] UpdateFollowCountMinus failed ,the error is %s", err)
		return err
	}
	return nil
}

func GetConcernsByUserId(userId int64) ([]model.User, error) {
	ToUsers, err := repository.GetToUserIdByUserId(userId)
	if err != nil {
		logrus.Errorf("[GetToUserIdByUserId] GetToUserIdByUserId failed ,the error is %s", err)
		return nil, err
	}
	users := []model.User{}
	for _, user := range ToUsers {
		toUser := user.ToUser
		toUser.IsFollow = true
		users = append(users, toUser)
	}
	return users, nil
}

func GetFansByTouserId(toUserId int64) ([]model.User, error) {
	users, err := repository.GetUserIdByToUserId(toUserId)
	if err != nil {
		logrus.Errorf("[GetFansByTouserId] GetToUserIdByUserId failed ,the error is %s", err)
		return nil, err
	}
	//check 当前用户是否同时关注粉丝
	followers, err := repository.GetToUserIdByUserId(toUserId)
	if err != nil {
		logrus.Errorf("[GetFansByTouserId] GetToUserIdByUserId failed ,the error is %s", err)
		return nil, err
	}
	fansSet := make(map[int64]interface{})
	for _, follower := range followers {
		fansSet[follower.ToUserID] = true
	}
	tousers := []model.User{}
	for _, user := range users {
		myUser := user.User
		_, ok := fansSet[myUser.Id]
		if ok {
			myUser.IsFollow = true
		}
		tousers = append(tousers, myUser)
	}
	return tousers, nil
}
