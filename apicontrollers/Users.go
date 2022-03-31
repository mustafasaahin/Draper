package apicontrollers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/libs"
	"github.com/mustafasaahin/Draper/models"
	"net/http"
)

func POSTRegister(c *gin.Context) {
	var form models.User
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := config.DB.
		Where("mail=?", form.Mail).
		Find(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	openPassword := form.Password
	form.Password = config.HashPassword(form.Password)
	if err := config.DB.Create(&form).Error; err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		mailHtml := `
<html><h1>
	Merhaba %s %s,
</h1>
<p>Sisteme girmek için %s şifresini kullanabilirsiniz.</p>
</html>
`
		mailHtml = fmt.Sprintf(mailHtml, form.Name, form.Surname, openPassword)
		libs.DraperMail{}.SendMail(form.Mail, "Kaydınızı aldık", mailHtml)
		c.JSON(http.StatusOK, form)
		return
	}
}

func POSTLogin(c *gin.Context) {
	type LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var form LoginForm
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(form.Username, config.HashPassword(form.Password))
	var user models.User
	if err := config.DB.Where("mail=? and password=?", form.Username,
		config.HashPassword(form.Password)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}
	if token, err := libs.CreateToken(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		exp := libs.ParseToken(token, "expd")

		c.JSON(http.StatusOK, gin.H{
			"user":    user,
			"token":   token,
			"exp":     exp,
			"ID":      user.ID,
			"Name":    user.Name,
			"Surname": user.Surname,
			"Phone":   user.Phone,
		})
		return
	}
}
