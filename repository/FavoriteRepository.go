package repository

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/model"
)

func CreateFavorites(favorite *model.Favorite) (int64, error) {
	return favorite.Id, DB.Create(favorite).Error
}

func GetFavoritesByUserId(userId int64) ([]model.Favorite, error) {
	favorites := []model.Favorite{}
	//"author_id"改成喜爱人的id
	return favorites, DB.Preload("User").Preload("Video").Order("created_at desc").Where("user_id = ?", &userId).Find(&favorites).Error
}

func FindFvorite(id int64, payUrl string) (favorite *model.Favorite, err error) {

	favorite = &model.Favorite{}
	result := DB.Where(map[string]interface{}{"author_id": id, "play_url": payUrl}).Find(&favorite)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %d", id)
	}
	return favorite, err
}

func UpdateFavorite(favorite *model.Favorite, isFavorite bool, FavoriteCount int64) (err error) {
	update := DB.Model(favorite).Updates(map[string]interface{}{"is_favorite": isFavorite, "favorite_count": FavoriteCount})
  
	if update.RowsAffected == 0 {
		return err
	}
	return nil
}

func DeleteFavoriteByUserIDAndVideoID(userId int64, videoId int64) (err error) {
	return DB.Where("user_id = ?", userId).Where("video_id =?", videoId).Delete(&model.Favorite{}).Error
}

func GetIsFavoriteByVideoId(videoId int64) (video *model.Video, err error) {
	findVideo := &model.Video{}
	result := DB.Where("id=?", videoId).First(findVideo)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %d", videoId)
	}
	return findVideo, err
}

func UpdateIsFavorite(video *model.Video, isFavorite int64, FavoriteCount int64) (err error) {
	update := DB.Model(video).Updates(map[string]interface{}{"is_favorite": isFavorite, "favorite_count": FavoriteCount})
	//update := DB.Model(video).Update("is_favorite", isFavorite)
	if update.RowsAffected == 0 {
		return err
	}
	return nil
}
