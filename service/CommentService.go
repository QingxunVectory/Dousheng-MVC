package service

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

//更新videos中的comment_count 并添加评论内容到comments
func AddComment(token string, videoId int64, comment_text string) (comment *model.Comment, err error) {
	parseToken, err := utils.ParseToken(token)
	if err != nil {
		logrus.Errorf("[AddComment] Add comment failed ,the error is %s", err)
		return nil, err
	}
	userName := parseToken.UserName
	//根据姓名找id
	user, err := repository.GetUserByUserName(userName)
	if err != nil {
		logrus.Errorf("[AddComment] GetUserByUserName failed ,the error is %s", err)
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
		logrus.Errorf("[AddComment] GetUserByUserName failed ,the error is %s", err)
		return nil, err
	}
	err = repository.UpdateVideoCommentCountPlus(videoId)
	if err != nil {
		logrus.Errorf("[AddComment] UpdateVideoCommentCountPlus failed ,the error is %s", err)
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
		logrus.Errorf("[DeleteCommentByCommentId] UpdateVideoCommentCountPlus failed ,the error is %s", err)
		return err
	}

	err = repository.DeleteCommentByCommentId(id)
	if err != nil {
		logrus.Errorf("[DeleteCommentByCommentId] DeleteCommentByCommentId failed ,the error is %s", err)
		return err
	}
	err = repository.UpdateVideoCommentCountMinus(comment.VideoID)
	if err != nil {
		logrus.Errorf("[DeleteCommentByCommentId] UpdateVideoCommentCountMinus failed ,the error is %s", err)
		return err
	}

	return nil
}

func GetCommentByVideoId(id int64) ([]model.Comment, error) {
	comments, err := repository.GetCommentByVideoId(id)
	if err != nil {
		logrus.Errorf("[GetCommentByVideoId] GetCommentByVideoId failed ,the error is %s", err)
		return nil, err
	}
	return comments, nil
}
