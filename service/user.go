package service

import (
	"context"
	"myapp/dto"
	"myapp/middleware"
	"myapp/model"
	"myapp/tools"

	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, input dto.NewUser) (model.User, error) {
	var (
		password = tools.HashAndSalt([]byte(input.Password))

		user = model.User{
			Name:     input.Name,
			Password: password,
			Email:    input.Email,
		}
	)

	err := DB.Model(&user).Create(&user).Error
	if err != nil {
		panic(err)
	}

	return user, nil

}

func GetUserById(ctx context.Context, id int) (*model.User, error) {
	var (
		user model.User
	)

	err := DB.Model(&user).Where("id = ?", id).Take(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &user, nil
}

func GetAllUser(ctx context.Context) ([]*model.User, error) {
	var (
		users []*model.User
	)

	err := DB.Model(&users).Find(&users).Error
	if err != nil {
		panic(err)
	}

	return users, nil
}

func Deleteuser(ctx context.Context, id int) (string, error) {
	var (
		user []model.User
	)

	err := DB.Model(&user).Where("id = ?", id).Delete(&user).Error
	if err != nil {
		panic(err)
	}

	return "success", nil

}

func Updateuser(ctx context.Context, id int, input dto.Updateuser) (string, error) {
	var (
		user model.User
	)

	err := DB.Model(&user).Where("id = ?", id).Updates(map[string]interface{}{
		"name":  input.Name,
		"email": input.Email,
	}).Error

	if err != nil {
		panic(err)
	}

	return "success", nil

}

func UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		user model.User
	)

	err := DB.Model(&user).Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &user, nil
}

func LoginUser(ctx context.Context, input dto.LoginUser) (*string, error) {
	user, err := UserGetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	err = tools.ComparePasswords(user.Password, input.Password)
	if err != nil {
		return nil, err
	}

	token := middleware.JwtGenerate(user.ID)

	return &token, nil
}
