package restaurantrepo

import (
	"context"
	"fooddlv/common"
	"fooddlv/restaurants/restaurantmodel"
)

type ListRestaurantStorage interface {
	List(ctx context.Context, paging *common.Paging, filter *restaurantmodel.ListFilter) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantRepo struct {
	store ListRestaurantStorage
}

func NewListRestaurantRepo(store ListRestaurantStorage) *listRestaurantRepo {
	return &listRestaurantRepo{store: store}
}

func (repo *listRestaurantRepo) ListRestaurants(ctx context.Context, paging *common.Paging, filter *restaurantmodel.ListFilter) ([]restaurantmodel.Restaurant, error) {
	restaurants, err := repo.store.List(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrCannotListEntity(restaurantmodel.EntityName, err)
	}

	return restaurants, nil
}
