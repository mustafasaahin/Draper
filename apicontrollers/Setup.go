package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"gorm.io/gorm"
	"net/http"
)

func GETSetup(c *gin.Context) {
	var setup []models.Setup
	if err := config.DB.Find(&setup).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, setup)
}
func GETSetupByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var setup models.Setup
	config.DB.Where("id=?", id).
		First(&setup)
	c.JSON(http.StatusOK, setup)
}
func POSTSetup(c *gin.Context) {
	var setup models.Setup
	if err := c.Bind(&setup); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&setup).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, setup)
	return
}
func PUTSetup(c *gin.Context) {
	var setup models.Setup
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&setup).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&setup)
		config.DB.Save(&setup)
		c.JSON(http.StatusOK, setup)
		return
	}
}
func DELETESetup(c *gin.Context) {
	var setup models.Setup
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&setup).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Where("id=?", setup.ID).Delete(&setup).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Done")
	}
}
