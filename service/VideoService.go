package service

import (
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/url"
	"time"
)

//后续优化相关逻辑
//test branch
func UploadVideo(ctx *gin.Context, data *multipart.FileHeader) error {
	fmt.Println(data.Filename)
	fileKey := utils.GenerateVideoKey(data.Filename)
	fmt.Println("fileKey:" + fileKey)
	open, err := data.Open()
	if err != nil {
		return err
	}
	err = utils.Upload(fileKey, open)
	if err != nil {
		return err
	}
	claim, err := utils.GetClaimInfoByCtx(ctx)
	if err != nil {
		return err
	}
	user, err := repository.GetUserByUserName(claim.UserName)
	if err != nil {
		return err
	}
	fmt.Println(utils.GetVideoUrl(fileKey))
	createdVideo := &model.Video{
		AuthorID:      user.Id,
		PlayUrl:       utils.GetVideoUrl(fileKey),
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		VideoName:     fileKey,
		IsFavorite:    false,
		CommentList:   nil,
	}
	_, err = repository.CreateVideo(createdVideo)
	if err != nil {
		return err
	}
	return nil
}

func GetVideos(queryTime time.Time) ([]model.Video, int64, error) {
	videos, err := repository.GetVideosByTime(queryTime)
	if err != nil {
		return nil, -1, err
	}
	if len(videos) == 0 {
		return videos, time.Now().UnixMilli(), nil
	}
	lastVideo := videos[len(videos)-1]
	return videos, lastVideo.CreatedAt.Unix(), nil
}

func GetVideosByUserId(id int64) ([]model.Video, error) {
	return repository.GetVideosByUserId(id)
}

func UpdateVideoImgUrl(jsonStr []byte) error {

	var j *simplejson.Json
	unescape, err := url.QueryUnescape(string(jsonStr))
	if err != nil {
		return err
	}
	j, err = simplejson.NewJson([]byte(unescape))
	if err != nil {
		return err
	}
	workflowTaskEvent := j.Get("WorkflowTaskEvent")
	name := workflowTaskEvent.Get("InputInfo").Get("CosInputInfo").Get("Object")
	path := workflowTaskEvent.Get("MediaProcessResultSet").GetIndex(0).Get("SnapshotByTimeOffsetTask").Get("Output").Get("PicInfoSet").GetIndex(0).Get("Path")

	if name == nil || path == nil {
		return errors.New("name or path is nil")
	}
	nameStr, err := name.String()
	if err != nil {
		return err
	}
	pathStr, err := path.String()
	if err != nil {
		return err
	}
	if len(nameStr) == 0 || len(pathStr) == 0 {
		return errors.New("name or path's len is 0")
	}
	imgPath := utils.GetVideoUrl(pathStr[1:])
	nameStrs := nameStr[1:]
	_, err = repository.UpdateVideosByUrl(nameStrs, imgPath)
	if err != nil {
		return err
	}
	return nil
}


func UpdateIsFavorite(token string, videos []model.Video) ([]model.Video, error) {
	claim, err := utils.ParseToken(token)
	if err != nil {
		return nil, err
	}
	user, err := repository.GetUserByUserName(claim.UserName)
	if err != nil {
		return nil, err
	}
	favoriteVideos, err := repository.GetFavoritesByUserId(user.Id)
	if err != nil {
		return nil, err
	}
	videoSet := make(map[int64]interface{})
	for _, favorite := range favoriteVideos {
		videoSet[favorite.VideoID] = true
	}
	retList := []model.Video{}
	for _, video := range videos {
		_, ok := videoSet[video.Id]
		if ok {
			video.IsFavorite = true
		}
		retList = append(retList, video)
	}
	return retList, nil
}