package repository

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gorm"
	"time"
)

func CreateVideo(video *model.Video) (int64, error) {
	return video.Id, DB.Create(video).Error
}

func GetVideosByTime(lastTime time.Time) ([]model.Video, error) {
	videos := []model.Video{}
	return videos, DB.Preload("Author").Order("created_at desc").Where("created_at < ?", &lastTime).Limit(30).Find(&videos).Error
}

func GetVideosByUserId(id int64) ([]model.Video, error) {
	videos := []model.Video{}
	return videos, DB.Preload("Author").Order("created_at desc").Where("author_id = ?", &id).Find(&videos).Error
}

func UpdateVideo(video *model.Video) (int64, error) {
	return DB.RowsAffected, DB.Save(video).Error
}

func UpdateVideoLikeCountPlus(videoId int64) error {
	return DB.Model(model.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count+?", 1)).Error
}

func UpdateVideoLikeCountMinus(videoId int64) error {
	return DB.Model(model.Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count-?", 1)).Error
}

func UpdateVideoCommentCountPlus(videoId int64) error {
	return DB.Model(model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count+?", 1)).Error
}

func UpdateVideoCommentCountMinus(videoId int64) error {
	return DB.Model(model.Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count-?", 1)).Error
}
