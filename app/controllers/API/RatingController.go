package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/surajNirala/rating_services/app/commons"
	"github.com/surajNirala/rating_services/app/config"
	"github.com/surajNirala/rating_services/app/models"
	"github.com/surajNirala/rating_services/app/validation"
)

var DB = config.DB

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type userResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func RatingList(c *gin.Context) {

	var ratings []models.Rating
	DB.Select("id", "rating", "user_id", "hotel_id").Order("created_at DESC").Find(&ratings)
	commons.ResponseSuccess(c, 200, "Get all rating list.", ratings)
	// return
}

func RatingStore(c *gin.Context) {
	var request models.Rating
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Request binding error: %v", err)
		customErrors := validation.TranslateValidationErrors(err)
		res := gin.H{
			"status":  400,
			"message": "Invalid Request",
			"errors":  customErrors,
		}
		c.JSON(400, res)
		return
	}

	ratingdata := models.Rating{
		Rating:  request.Rating,
		UserID:  request.UserID,
		HotelID: request.HotelID,
	}

	if err := DB.Create(&ratingdata).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Rating not created successfully",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}
	res := gin.H{
		"status":  201,
		"message": "Rating Created Successfully",
		"data":    ratingdata,
	}
	c.JSON(201, res)
}

func RatingDetail(c *gin.Context) {
	var rating models.Rating
	rating_id := c.Param("rating_id")
	result := DB.Select("id", "name", "email").Where("id = ?", rating_id).Find(&rating)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Rate not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}
	res := Response{
		Status:  200,
		Message: "Fetch User Detail",
		Data:    rating,
	}
	c.JSON(200, res)
}

func RatingUpdate(c *gin.Context) {
	var rating models.Rating
	rating_id := c.Param("rating_id")
	result := DB.Select("id", "rating", "user_id", "hotel_id").Where("id = ?", rating_id).Find(&rating)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Rate not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}
	var request models.Rating
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("Request binding error: %v", err)
		customErrors := validation.TranslateValidationErrors(err)
		res := gin.H{
			"status":  400,
			"message": "Invalid Request",
			"errors":  customErrors,
		}
		c.JSON(400, res)
		return
	}
	ratingdata := models.Rating{
		Rating:  request.Rating,
		UserID:  request.UserID,
		HotelID: request.HotelID,
	}

	if err := DB.Where("id = ?", rating_id).Updates(&ratingdata).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Rating is not updated.",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}
	res := Response{
		Status:  200,
		Message: "Rating Detail Updated Successfully.",
		Data:    ratingdata,
	}
	c.JSON(200, res)
}

func RatingDelete(c *gin.Context) {
	var rating models.Rating
	rating_id := c.Param("rating_id")
	result := DB.Select("id", "name", "email").Where("id = ?", rating_id).Find(&rating)
	if result.RowsAffected == 0 {
		res := Response{
			Status:  409,
			Message: "Rating not found.",
			Data:    nil,
		}
		c.JSON(409, res)
		return
	}

	if err := DB.Where("id = ?", rating_id).Delete(&rating).Error; err != nil {
		res := gin.H{
			"status":  500,
			"message": "Rating is not deleted.",
			"error":   err.Error(),
			"data":    nil,
		}
		c.JSON(500, res)
		return
	}

	res := Response{
		Status:  200,
		Message: "Rating Deleted Successfully.",
		Data:    nil,
	}
	c.JSON(200, res)
}
