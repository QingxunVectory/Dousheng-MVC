package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm/clause"
)

func UpdateCommentCountOfVideos(video *model.Video) (err error) {
	return DB.Save(video).Error
}

func UpdateCommentCountOfFavorites(favorite *model.Favorite) (err error) {
	return DB.Save(favorite).Error
}

func CreateComment(comment *model.Comment) (*model.Comment, error) {
	return comment, DB.Create(comment).Error
}

func DeleteCommentByCommentId(id int64) (err error) {
	return DB.Where("id = ?", id).Delete(&model.Comment{}).Error
}

func GetCommentByVideoId(id int64) ([]model.Comment, error) {
	comments := []model.Comment{}
	return comments, DB.Preload(clause.Associations).Order("created_at desc").Where("video_id = ?", id).Find(&comments).Error
}

func FindVideoIdByCommentId(id int64) (comment *model.Comment, err error) {
	comment = &model.Comment{}
	result := DB.Where("id = ?", id).First(&comment)
	if result.RowsAffected == 0 {
		return nil, err
	}
	return comment, nil
}

func GetVideoByVideoId(id int64) (video *model.Video, err error) {
	video = &model.Video{}
	result := DB.Where("id = ?", id).First(&video)
	if result.RowsAffected == 0 {
		return nil, err
	}
	return video, err
}

func GetFavoriteByFavoriteId(id int64) (favorite *model.Favorite, err error) {
	favorite = &model.Favorite{}
	result := DB.Where("id = ?", id).First(&favorite)
	if result.RowsAffected == 0 {
		return nil, err
	}
	return favorite, err
}