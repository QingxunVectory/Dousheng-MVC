package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	//token := c.Query("token")
	token := c.Query("token")
	to_user_id := c.Query("to_user_id")
	//token := c.Query("token")
	action_type := c.Query("action_type")
	toUserId, err := strconv.ParseInt(to_user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "parse user Id failed",
		})
		return
	}
	actionType, err := strconv.ParseInt(action_type, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
		return
	}

	if actionType == 1 {
		err = service.Suscribe(toUserId, token)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: -1, StatusMsg: err.Error()})
		}
	} else if actionType == 2 {
		err = service.CancelSuscribe(toUserId, token)
		if err != nil {
			c.JSON(http.StatusOK, model.Response{StatusCode: -1, StatusMsg: err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, model.Response{StatusCode: -1, StatusMsg: "check your request"})
	}
	//
	c.JSON(http.StatusOK, model.Response{StatusCode: 0})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
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
	follows, err := service.GetConcernsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: -1})
	}
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{StatusCode: 0},
		UserList: follows,
	})
	//c.JSON(http.StatusOK, model.UserListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []model.User{DemoUser},
	//})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
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
	fans, err := service.GetFansByTouserId(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.Response{StatusCode: -1})
	}
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{StatusCode: 0},
		UserList: fans,
	})
	//c.JSON(http.StatusOK, model.UserListResponse{
	//	Response: model.Response{
	//		StatusCode: 0,
	//	},
	//	UserList: []model.User{DemoUser},
	//})
}
