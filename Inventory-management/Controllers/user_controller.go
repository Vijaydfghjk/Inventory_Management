package controllers

import (
	models "inventory_management/Models"
	token_stuff "inventory_management/Token_stuff"
	"inventory_management/dbrepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User_controller struct {
	service dbrepository.User_interect
}

func User_con(service dbrepository.User_interect) *User_controller {

	return &User_controller{service: service}
}

func (a *User_controller) Signup(c *gin.Context) {

	var temp models.User

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err, userdb := a.service.Register_new_user(temp)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully Register",

		"User_id": userdb.ID,
	})
}

func (a *User_controller) Login(c *gin.Context) {

	var temp models.Userlogin

	if err := c.ShouldBindJSON(&temp); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err, db_user := a.service.Logging(temp.ID, temp.Password)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	token, err := token_stuff.GenerateJWT(int(db_user.ID), db_user.Email, db_user.Role)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"M3essage": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"Status": "Successfully login",
		"Token":  token,
	})

}
