package restaurantstorage

import (
	"context"
	"fooddlv/common"
	"fooddlv/restaurants/restaurantmodel"
)

type a struct {
}

func (s *storeMySQL) List(
	ctx context.Context,
	paging *common.Paging,
	filter *restaurantmodel.ListFilter,
) ([]restaurantmodel.Restaurant, error) {
	var listRestaurants []restaurantmodel.Restaurant

	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status = 1")

	if err := db.Count(&(paging.Total)).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := s.db.
		Limit(paging.Limit).Offset((paging.Page - 1) * paging.Limit).
		Find(&listRestaurants).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return listRestaurants, nil
}
