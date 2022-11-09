package service

import (
	"context"
	"myapp/dto"
	"myapp/model"
)

func CreatePost(ctx context.Context, input dto.NewPost) (model.Post, error) {
	var (
		post = model.Post{
			Title:       input.Title,
			Description: input.Description,
			UserId:      input.UserId,
		}
	)

	err := DB.Model(&post).Create(&post).Error
	if err != nil {
		panic(err)
	}

	return post, nil
}

func GetAllPost(ctx context.Context) ([]*model.Post, error) {
	var (
		posts []*model.Post
	)

	err := DB.Model(&posts).Find(&posts).Error
	if err != nil {
		panic(err)
	}

	return posts, nil
}

func GetPostById(ctx context.Context, id int) (model.Post, error) {
	var (
		post model.Post
	)
	err := DB.Model(&post).Where("id = ? ", id).Take(&post).Error
	if err != nil {
		panic(err)
	}

	return post, err
}

func DeletePost(ctx context.Context, id int, userId int) (string, error) {
	var (
		post []model.Post
	)

	_, err := GetUserById(ctx, userId)
	if err != nil {
		panic(err)
	}
	err = DB.Model(&post).Where("id = ?", id).Delete(&post).Error
	if err != nil {
		panic(err)
	}

	return "success", nil
}

func UpdatePost(ctx context.Context, id int, input dto.Updatepost, userId int) (string, error) {
	var (
		post model.Post
	)

	_, err := GetUserById(ctx, userId)
	if err != nil {
		panic(err)
	}

	err = DB.Model(&post).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       input.Title,
		"description": input.Description,
	}).Error

	if err != nil {
		panic(err)
	}

	return "success", nil
}
