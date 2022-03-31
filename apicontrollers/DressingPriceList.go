package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"gorm.io/gorm"
	"net/http"
)

func GETDressingPriceList(c *gin.Context) {
	var form []models.DressingPriceList
	if err := config.DB.
		Find(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, form)
}
func GETDressingPriceListByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var form models.DressingPriceList
	config.DB.Where("id=?", id).
		First(&form)
	c.JSON(http.StatusOK, form)
}
func POSTDressingPriceList(c *gin.Context) {
	var form models.DressingPriceList
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusCreated, form)
		return
	}
}
func PUTDressingPriceList(c *gin.Context) {
	var form models.DressingPriceList
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&form)
		config.DB.Save(&form)
		return
	}
}
func DELETEDressingPriceList(c *gin.Context) {
	var form models.DressingPriceList
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&form).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())

			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Where("id=?", form.ID).Delete(&form).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Done")
	}
}
