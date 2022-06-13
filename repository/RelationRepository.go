package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
)

func CreateRelation(relation *model.UserRelation) (int64, error) {
	return relation.Id, DB.Create(relation).Error
}

func DeleteFavoriteByUserIDAndFollowerId(userId int64, followerID int64) (err error) {
	return DB.Where("user_id = ?", userId).Where("to_user_id =?", followerID).Delete(&model.UserRelation{}).Error
}

func GetToUserIdByUserId(userId int64) ([]model.UserRelation, error) {
	ToUsers := []model.UserRelation{}
	return ToUsers, DB.Preload("User").Preload("User").Order("created_at desc").Where("user_id = ?", &userId).Find(&ToUsers).Error
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
	return ToUsers, DB.Preload("User").Preload("User").Order("created_at desc").Where("to_user_id = ?", &toUserId).Find(&ToUsers).Error
}
