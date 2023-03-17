package service

import "gocart/models"

func GetProductListingsService() ([]models.Product, error) {
	list := []models.Product{}
	for _, e := range Products {
		list = append(list, e)
	}
	return list, nil
}
