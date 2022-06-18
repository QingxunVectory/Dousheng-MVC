package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/sirupsen/logrus"
	"time"
)

func LikeVideo(videoId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		logrus.Errorf("[LikeVideo] ParseToken failed ,the error is %s", err)
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		logrus.Errorf("[LikeVideo] ParseToken failed ,the error is %s", err)
		return err
	}
	createdFavorite := &model.Favorite{
		UserID:    user.Id,
		VideoID:   videoId,
		Video:     model.Video{},
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
	_, err = repository.CreateFavorites(createdFavorite)
	if err != nil {
		logrus.Errorf("[LikeVideo] CreateFavorites failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateVideoLikeCountPlus(videoId)
	if err != nil {
		logrus.Errorf("[LikeVideo] UpdateVideoLikeCountPlus failed ,the error is %s", err)
		return err
	}
	return nil
}

func DislikeVideo(videoId int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		logrus.Errorf("[DislikeVideo] ParseToken failed ,the error is %s", err)
		return err
	}
	userName := parseToken.UserName
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		logrus.Errorf("[DislikeVideo] GetUserByUserName failed ,the error is %s", err)
		return err
	}
	err = repository.DeleteFavoriteByUserIDAndVideoID(user.Id, videoId)
	if err != nil {
		logrus.Errorf("[DislikeVideo] DeleteFavoriteByUserIDAndVideoID failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateVideoLikeCountMinus(videoId)
	if err != nil {
		logrus.Errorf("[DislikeVideo] UpdateVideoLikeCountMinus failed ,the error is %s", err)
		return err
	}
	return nil
}

func GetFavoritesByUserId(userId int64) ([]model.Video, error) {
	favorite, err := repository.GetFavoritesByUserId(userId)
	if err != nil {
		logrus.Errorf("[GetFavoritesByUserId] GetFavoritesByUserId failed ,the error is %s", err)
		return nil, err
	}
	videos := []model.Video{}
	for _, favorite := range favorite {
		favorite.Video.IsFavorite = true
		videos = append(videos, favorite.Video)
	}
	return videos, nil
}
