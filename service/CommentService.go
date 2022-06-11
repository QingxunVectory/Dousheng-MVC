package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"time"
)

//更新videos中的comment_count 并添加评论内容到comments
func AddComment(token string, id int64, actionType int64, comment_text string) (comment *model.Comment, err error) {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	userName := parseToken.UserName
	fmt.Println(userName)
	//根据姓名找id
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		return nil, err
	}

	if actionType == 1 {
		//comment_text := c.Query("comment_text")
		fmt.Println("actionType:", actionType)
		video, err := repository.GetIsFavoriteByVideoId(id)
		video.CommentCount++
		err = repository.UpdateCommentCountOfVideos(video)
		if err != nil {
			return nil, err
		}
		//查找favorites 中是否有这个视频
		favorite, err := repository.FindFvorite(user.Id, video.PlayUrl)
		if err != nil {
			return nil, err
		}
		favorite.CommentCount++
		err = repository.UpdateCommentCountOfFavorites(favorite)
		if err != nil {
			return nil, err
		}
		createCommnet := &model.Comment{
			UserID:     user.Id,
			VideoID:    video.Id,
			FavoriteId: favorite.Id,
			Content:    comment_text,
			CreatedAt:  time.Time{},
		}
		comment, err := repository.CreateComment(createCommnet)
		if err != nil {
			return nil, err
		}
		return comment, nil
	}
	return comment, nil
}

func DeleteCommentByCommentId(id int64) (err error) {
	//根据评论id找到video的id和favorite的id 将其中的commentcount减一
	comment, err := repository.FindVideoIdByCommentId(id)
	if err != nil {
		return err
	}
	video, err := repository.GetVideoByVideoId(comment.VideoID)
	if err != nil {
		return err
	}
	video.CommentCount--
	err = repository.UpdateCommentCountOfVideos(video)
	if err != nil {
		return err
	}
	favorite, err := repository.GetFavoriteByFavoriteId(comment.FavoriteId)
	if err != nil {
		return err
	}
	favorite.CommentCount--
	err = repository.UpdateCommentCountOfFavorites(favorite)
	if err != nil {
		return err
	}
	err = repository.DeleteCommentByCommentId(id)
	if err != nil {
		return err
	}
	return nil
}

func GetCommentByVideoId(id int64) ([]model.Comment, error) {
	comments, err := repository.GetCommentByVideoId(id)
	fmt.Println(comments)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
