package service

import (
	"context"
	"myapp/dto"
	"myapp/model"
)

func CreateProduct(ctx context.Context, input dto.NewProduct) (model.Product, error) {
	var (
		product = model.Product{
			Code:        input.Code,
			Description: input.Description,
		}
	)

	err := DB.Model(&product).Create(&product).Error
	if err != nil {
		panic(err)
	}

	return product, nil

}

func GetAllProduct(ctx context.Context) ([]*model.Product, error) {
	var (
		products []*model.Product
	)

	err := DB.Model(&products).Find(&products).Error
	if err != nil {
		panic(err)
	}

	return products, nil
}
