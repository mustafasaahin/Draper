package apicontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
	"gorm.io/gorm"
	"net/http"
)

func GETYarns(c *gin.Context) {
	var yarns []models.Yarns
	if err := config.DB.
		Find(&yarns).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, yarns)
}
func GETYarnsByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var yarns models.Yarns
	config.DB.
		Where("id=?", id).
		First(&yarns)
	c.JSON(http.StatusOK, yarns)
}
func POSTYarns(c *gin.Context) {
	var yarns models.Yarns
	if err := c.Bind(&yarns); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.Create(&yarns).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		c.JSON(http.StatusOK, yarns)
		return
	}
}
func PUTYarns(c *gin.Context) {
	var yarns models.Yarns
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&yarns).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadGateway, err.Error())
			return
		}
	} else {
		_ = c.Bind(&yarns)
		config.DB.Save(&yarns)
		c.JSON(http.StatusOK, yarns)
		return
	}
}
func DELETEYarns(c *gin.Context) {
	var yarns models.Yarns
	id := c.Params.ByName("id")
	if err := config.DB.Where("id=?", id).First(&yarns).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, err.Error())
			return
		} else {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	} else {
		if err := config.DB.Where("id=?", yarns.ID).Delete(&yarns).Error; err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, "Done")
	}
}
