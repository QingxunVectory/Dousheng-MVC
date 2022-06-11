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
	return favorites, DB.Preload("Author").Order("created_at desc").Where("author_id = ?", &userId).Find(&favorites).Error
}

func FindFvorite(id int64, payUrl string) (favorite *model.Favorite, err error) {

	favorite = &model.Favorite{}
	result := DB.Where(map[string]interface{}{"author_id": id, "play_url": payUrl}).Find(&favorite)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("check your userName,there is no user %d", id)
	}
	return favorite, err
	//favorite = &model.Favorite{}
	//result := DB.Where("author_id=?", id).First(favorite)
	//if result.RowsAffected == 0 {
	//	return nil, fmt.Errorf("check your userName,there is no user %d", id)
	//}
	//return favorite, err
}

func UpdateFavorite(favorite *model.Favorite, isFavorite bool, FavoriteCount int64) (err error) {
	update := DB.Model(favorite).Updates(map[string]interface{}{"is_favorite": isFavorite, "favorite_count": FavoriteCount})
	//update := DB.Model(video).Update("is_favorite", isFavorite)
	if update.RowsAffected == 0 {
		return err
	}
	return nil
}

func DeleteFavorite(favorite *model.Favorite) (err error) {
	return DB.Delete(&favorite).Error
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
