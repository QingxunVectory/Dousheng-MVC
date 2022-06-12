package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// CommentAction no practical effect, just check if token is valid
//func CommentAction(c *gin.Context) {
//	token := c.Query("token")
//	video_id := c.Query("video_id")
//	action_type := c.Query("action_type")
//	comment_text := c.Query("comment_text")
//	comment_id := c.Query("comment_id")
//	videoId, err := strconv.ParseInt(video_id, 10, 64)
//	if err != nil {
//		return
//	}
//	actionType, err := strconv.ParseInt(action_type, 10, 64)
//	if err != nil {
//		return
//	}
//	commentId, err := strconv.ParseInt(comment_id, 10, 64)
//
//	if err != nil {
//		comment, err := service.AddComment(token, videoId, actionType, comment_text)
//		if err != nil {
//			c.JSON(http.StatusOK, model.Response{StatusCode: 1})
//		}
//
//		c.JSON(http.StatusOK, model.CommentResponse{
//			Response: model.Response{
//				StatusCode: 0,
//				StatusMsg:  "发表成功",
//			},
//			Comment: model.Comment{
//				UserID:     comment.UserID,
//				VideoID:    comment.VideoID,
//				FavoriteId: comment.FavoriteId,
//				Content:    comment.Content,
//				CreatedAt:  time.Time{},
//			},
//		})
//	} else {
//		err := service.DeleteCommentByCommentId(commentId)
//		if err != nil {
//			c.JSON(http.StatusOK, model.Response{
//				StatusCode: 1,
//				StatusMsg:  "删除失败",
//			})
//		}
//		c.JSON(http.StatusOK, model.Response{
//			StatusCode: 0,
//			StatusMsg:  "删除成功",
//		})
//	}
//
//	//token := c.Query("token")
//	//
//	//if _, exist := usersLoginInfo[token]; exist {
//	//	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
//	//} else {
//	//	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	//}
//}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	//根据视频id找到
	video_id := c.Query("video_id")
	videoId, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	comments, err := service.GetCommentByVideoId(videoId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: 1})
	}
	c.JSON(http.StatusOK, model.CommentListResponse{
		Response:    model.Response{StatusCode: 0},
		CommentList: comments,
	})

	//c.JSON(http.StatusOK, model.CommentListResponse{
	//	Response:    model.Response{StatusCode: 0},
	//	CommentList: DemoComments,
	//})
}
