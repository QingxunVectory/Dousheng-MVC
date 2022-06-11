package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
)

func GetIsFavoriteByVideoId(videoId int64, actionType int64, token string) error {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return err
	}
	userName := parseToken.UserName
	//根据姓名找id
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		return err
	}
	//找到被点赞视频 更新isfavorite 和favoriteCount
	video, err := repository.GetIsFavoriteByVideoId(videoId)

	if err != nil {
		return err
	}

	if actionType == 2 {
		fmt.Println("action", actionType)
		if video.FavoriteCount > 0 {
			video.FavoriteCount--
		} else {
			video.FavoriteCount = 0
		}
		repository.UpdateIsFavorite(video, 0, video.FavoriteCount) // 0可以改成false也行
		favorites := &model.Favorite{
			AuthorID:      user.Id, //这边是谁的id 现在是点赞者的id
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			VideoName:     video.VideoName,
			IsFavorite:    video.IsFavorite,
			CommentList:   nil,
		}

		favorite, err := repository.FindFvorite(favorites.AuthorID, favorites.PlayUrl)
		if err != nil {
			return nil
		} else {
			err = repository.DeleteFavorite(favorite)
			if err != nil {
				return err
			}
		}

	} else {
		video.FavoriteCount++
		//
		repository.UpdateIsFavorite(video, 1, video.FavoriteCount)
		favorites := &model.Favorite{
			AuthorID:      user.Id,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			VideoName:     video.VideoName,
			IsFavorite:    video.IsFavorite,
			CommentList:   nil,
		}

		favorite, err := repository.FindFvorite(favorites.AuthorID, favorites.PlayUrl) //判断这个视频播放地址和上传视频作者相同？
		if err != nil {
			_, err = repository.CreateFavorites(favorites) //如果是同一个就不创建了 改一下
			if err != nil {
				return err
			}
		} else {
			//如果相同则更新
			repository.UpdateFavorite(favorite, favorites.IsFavorite, favorites.FavoriteCount)
		}

	}

	return nil
}

func GetFavoritesByUserId(userId int64) ([]model.Favorite, error) {
	return repository.GetFavoritesByUserId(userId)
}
