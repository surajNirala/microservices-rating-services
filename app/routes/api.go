package routes

import (
	"github.com/gin-gonic/gin"
	api "github.com/surajNirala/rating_services/app/controllers/API"
)

func ApiRoutes(apiRouter *gin.Engine) {
	route := apiRouter.Group("/api")
	{
		route.GET("/ratings", api.RatingList)
		route.POST("/rating/store", api.RatingStore)
		route.GET("/rating/:user_id", api.RatingDetail)
		route.PUT("/rating/:user_id", api.RatingUpdate)
		route.DELETE("/rating/:user_id", api.RatingDelete)
	}
}
