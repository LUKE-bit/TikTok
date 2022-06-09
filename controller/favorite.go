package controller

import (
	"TikTok/service"
	"TikTok/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			err = service.AddLike(videoId, actionType)
			if err != nil {
				c.JSON(http.StatusOK, Response{
					StatusCode: 1,
					StatusMsg:  err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, Response{
					StatusCode: 0,
					StatusMsg:  "AllRight",
				})
			}
			return
		}
	}
	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	if token != "" {
		if claim, err := utils.ParseToken(token); claim != nil && err == nil {
			video_list := service.GetFavoriteList()
			c.JSON(http.StatusOK, VideoListResponse{
				Response: Response{
					StatusCode: 0,
				},
				VideoList: video_list,
			})
		}
	}

}
