package service

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"gorm.io/gorm"
	"time"
)

//更新videos中的comment_count 并添加评论内容到comments
func AddComment(token string, videoId int64, comment_text string) (comment *model.Comment, err error) {
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
	createdComment := &model.Comment{
		UserID:    user.Id,
		VideoID:   videoId,
		Content:   comment_text,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: gorm.DeletedAt{},
	}
	comment, err = repository.CreateComment(createdComment)
	if err != nil {
		return nil, err
	}
	err = repository.UpdateVideoCommentCountPlus(videoId)
	if err != nil {
		return nil, err
	}
	comment.User = *user
	comment.CreateDate = comment.CreatedAt.Format("01-02")
	return comment, nil
}

func DeleteCommentByCommentId(id int64) (err error) {
	//根据评论id找到video的id和favorite的id 将其中的commentcount减一

	comment, err := repository.FindVideoIdByCommentId(id)
	if err != nil {
		return err
	}

	err = repository.DeleteCommentByCommentId(id)
	if err != nil {
		return err
	}
	err = repository.UpdateVideoCommentCountMinus(comment.VideoID)
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
