package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"gorm.io/gorm"
	"net/http"
)

func GETProduct(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}
func GETProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product models.Product
	config.DB.Where("id=?", id).
		First(&product)
	c.JSON(http.StatusOK, product)
}
func POSTProduct(c *gin.Context) {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, product)
		return
	}
}
func PUTProduct(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&product)
		config.DB.Save(&product)
		c.JSON(http.StatusOK, product)
		return
	}
}
func DELETEProduct(c *gin.Context) {
	var product models.Product
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Where("id=?", id, product.ID).Delete(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Done")
	}
}
