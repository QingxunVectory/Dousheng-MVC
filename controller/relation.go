package controller

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	//token := c.Query("token")
	c.JSON(http.StatusOK, model.Response{StatusCode: 0})

}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: []model.User{DemoUser},
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, model.UserListResponse{
		Response: model.Response{
			StatusCode: 0,
		},
		UserList: []model.User{DemoUser},
	})
}
