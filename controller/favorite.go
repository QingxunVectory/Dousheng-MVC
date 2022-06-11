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
	//
	err = service.GetIsFavoriteByVideoId(videoID, actionType, token)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: -1})
	}

	c.JSON(http.StatusOK, model.Response{StatusCode: 0})

	//token := c.Query("token")
	//
	//if _, exist := usersLoginInfo[token]; exist {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, model.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
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
	favorites, err := service.GetFavoritesByUserId(userId)
	fmt.Println("favorites：", favorites)
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
		FavoriteList: favorites,
	})

	//c.JSON(http.StatusOK, model.VideoListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: favorites,
	//})
}
