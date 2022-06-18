package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpdateCommentCountOfVideos(video *model.Video) (err error) {
	return DB.Save(video).Error
}

func UpdateCommentCountOfFavorites(favorite *model.Favorite) (err error) {
	return DB.Save(favorite).Error
}

func CreateComment(comment *model.Comment, videoId int64) (*model.Comment, error) {
	tx := DB.Begin()
	defer tx.Callback()
	if err := tx.Create(comment).Error; err != nil {
		logrus.Errorf("[CreateComment] CreateComment failed error is %s", err)
		tx.Callback()
		return nil, err
	}
	if err := tx.Model(model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count+?", 1)).Error; err != nil {
		logrus.Errorf("[CreateComment] UpdateCount  error is %s", err)
		tx.Callback()
		return nil, err
	}
	tx.Commit()
	return comment, nil
}

func DeleteCommentByCommentId(id int64, videoId int64) (err error) {
	tx := DB.Begin()
	defer tx.Callback()
	if err := tx.Where("id = ?", id).Delete(&model.Comment{}).Error; err != nil {
		logrus.Errorf("[DeleteCommentByCommentId] DeleteComment failed error is %s", err)
		tx.Callback()
		return err
	}
	if err := DB.Model(model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count-?", 1)).Error; err != nil {
		logrus.Errorf("[DeleteCommentByCommentId] UpdateComment failed error is %s", err)
		tx.Callback()
		return err
	}
	tx.Commit()
	return nil
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
