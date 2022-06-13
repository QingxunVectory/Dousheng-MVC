package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
)

func CreateUser(user *model.User) (int64, error) {
	return user.Id, DB.Create(user).Error
}

func GetUserByUserName(userName string) (user *model.User, err error) {
	findingUser := &model.User{}
	result := DB.Where("user_name=?", userName).First(findingUser)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %s", userName)
	}
	return findingUser, err
}

func UpdateFollowCountPlus(userId int64) error {
	return DB.Model(model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count+?", 1)).Error
}

func UpdateFollowCountMinus(userId int64) error {
	return DB.Model(model.User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count-?", 1)).Error
}

func UpdateFollowerCountPlus(userId int64) error {
	return DB.Model(model.User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count+?", 1)).Error
}

func UpdateFollowerCountMinus(userId int64) error {
	return DB.Model(model.User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count-?", 1)).Error
}
