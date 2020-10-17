package controllers

import (
	"gin-rest-microservice/domain/httperrors"
	"gin-rest-microservice/domain/users"
	"gin-rest-microservice/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	UsersController = usersController{}
)

type usersController struct {}

func respond(c *gin.Context, isXML bool, httpCode int, body interface{}) {
	if isXML {
		c.XML(httpCode, body)
		return
	}
	c.JSON(httpCode, body)
}

func (controller usersController) Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		httpErr := httperrors.NewBadRequestError("Invalid json body")
		c.JSON(httpErr.Code, httpErr)
		return
	}
	createdUser, err := services.UsersService.Create(user)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	//return created user
	c.JSON(http.StatusCreated, createdUser)
}

func (controller usersController) Get(c *gin.Context) {
	isXML := c.GetHeader("Accept") == "application/xml"
	userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		httpErr := httperrors.NewBadRequestError("invalid user id")
		respond(c, isXML, httpErr.Code, httpErr)
		return
	}

	user, getErr := services.UsersService.Get(userId)
	if getErr != nil {
		respond(c, isXML, getErr.Code, getErr)
		return
	}

	respond(c, isXML, http.StatusOK, user)

}



