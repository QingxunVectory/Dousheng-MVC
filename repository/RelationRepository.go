package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateRelation(relation *model.UserRelation, toUserId int64, userId int64) (int64, error) {
	tx := DB.Begin()
	defer tx.Callback()
	if err := tx.Create(relation).Error; err != nil {
		logrus.Errorf("[CreateRelation] CreateRelation failed error is %s", err)
		tx.Callback()
		return -1, err
	}
	if err := tx.Model(model.User{}).Where("id = ?", toUserId).Update("follower_count", gorm.Expr("follower_count+?", 1)).Error; err != nil {
		logrus.Errorf("[CreateRelation] CreateRelation failed error is %s", err)
		tx.Callback()
		return -1, err
	}
	if err := tx.Model(model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count+?", 1)).Error; err != nil {
		logrus.Errorf("[CreateRelation] CreateRelation failed error is %s", err)
		tx.Callback()
		return -1, err
	}
	tx.Commit()
	return relation.Id, nil
}

func DeleteFavoriteByUserIDAndFollowerId(userId int64, followerID int64) (err error) {
	tx := DB.Begin()
	defer tx.Callback()
	if err := tx.Where("user_id = ?", userId).Where("to_user_id =?", followerID).Delete(&model.UserRelation{}).Error; err != nil {
		logrus.Errorf("[DeleteFavoriteByUserIDAndFollowerId] DeleteFavoriteByUserIDAndFollowerId failed error is %s", err)
		tx.Callback()
		return err
	}
	if err := tx.Model(model.User{}).Where("id = ?", followerID).Update("follower_count", gorm.Expr("follower_count-?", 1)).Error; err != nil {
		logrus.Errorf("[DeleteFavoriteByUserIDAndFollowerId] DeleteFavoriteByUserIDAndFollowerId failed error is %s", err)
		tx.Callback()
		return err
	}
	if err := tx.Model(model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count-?", 1)).Error; err != nil {
		logrus.Errorf("[DeleteFavoriteByUserIDAndFollowerId] DeleteFavoriteByUserIDAndFollowerId failed error is %s", err)
		tx.Callback()
		return err
	}
	tx.Commit()
	return nil
}

func GetToUserIdByUserId(userId int64) ([]model.UserRelation, error) {
	ToUsers := []model.UserRelation{}
	return ToUsers, DB.Preload("User").Preload("ToUser").Order("created_at desc").Where("user_id = ?", &userId).Find(&ToUsers).Error
}

func GetUsersByToUserId(toUserId int64) (user *model.User, err error) {
	findingUser := &model.User{}
	result := DB.Where("id=?", toUserId).First(findingUser)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %s", toUserId)
	}
	return findingUser, err
}

func GetUserIdByToUserId(toUserId int64) ([]model.UserRelation, error) {
	ToUsers := []model.UserRelation{}
	return ToUsers, DB.Preload("User").Preload("ToUser").Order("created_at desc").Where("to_user_id = ?", &toUserId).Find(&ToUsers).Error
}
