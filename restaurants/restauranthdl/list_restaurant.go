package restauranthdl

import (
	"fooddlv/common"
	"fooddlv/restaurants/restaurantmodel"
	"fooddlv/restaurants/restaurantrepo"
	"fooddlv/restaurants/restaurantstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListRestaurant(appCtx common.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		var p restaurantmodel.ListParam

		if err := c.ShouldBind(&p); err != nil && err.Error() != "EOF" {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		p.Fulfill()

		db := appCtx.GetDBConnection()

		store := restaurantstorage.NewMySQLStore(db)
		repo := restaurantrepo.NewListRestaurantRepo(store)

		result, err := repo.ListRestaurants(c.Request.Context(), &(p.Paging), p.ListFilter)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, p.Paging, p.ListFilter))
	}
}
