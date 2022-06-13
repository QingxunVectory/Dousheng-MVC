package repository

import "github.com/RaymondCode/simple-demo/model"

func CreateRelation(relation *model.UserRelation) (int64, error) {
	return relation.Id, DB.Create(relation).Error
}

func DeleteFavoriteByUserIDAndFollowerId(userId int64, followerID int64) (err error) {
	return DB.Where("user_id = ?", userId).Where("to_user_id =?", followerID).Delete(&model.UserRelation{}).Error
}
