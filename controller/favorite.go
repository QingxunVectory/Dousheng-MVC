package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	//token := c.Query("token")
	action_type := c.Query("action_type")
	fmt.Println("action_type")
	videoID, err := strconv.ParseInt(video_id, 10, 64)
	actionType, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}
	if actionType == 1 {
		err = service.LikeVideo(videoID, token)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: -1})
		}
	} else if actionType == 2 {
		err = service.DislikeVideo(videoID, token)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: -1})
		}
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: -1, StatusMsg: "check your request"})
	}
	//
	c.JSON(http.StatusOK, model.Response{StatusCode: 0})

}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	idStr := c.Query("user_id")
	if idStr == "" {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "some params is missing",
		})
		return
	}
	userId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videoList, err := service.GetFavoritesByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, model.FavoriteListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		FavoriteList: videoList,
	})
}
