package services

import (
	"api-instagram/app/models"
	"api-instagram/db"
)

func InsetPost(post *models.Posts) (*models.Posts, error) {
	if err := db.Instance.Table("Posts").Create(&post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func GetPosts(filter *models.PostFilter) (*[]models.Posts, error) {
	var posts *[]models.Posts

	//Add query
	query := db.Instance.Table("Posts")

	if filter != nil {
		if filter.User_id != 0 {
			query = query.Where("user_id = ?", filter.User_id)
		}
	}

	if filter.Page > 0 && filter.Limit > 0 {
		offset := (filter.Page - 1) * filter.Limit
		query.Offset(offset).Limit(filter.Limit)
	}

	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
