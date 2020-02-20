package server

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	formatter "github.com/saeidraei/go-jwt-auth/implem/json.formatter"
)

type userLoginPostBody struct {
	User struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

func (rH RouterHandler) userLoginPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	body := &userLoginPostBody{}
	if err := c.BindJSON(body); err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err: errors.New("please provide email and password"), Type: gin.ErrorTypePrivate})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	user, token, err := rH.ucHandler.UserLogin(body.User.Email, body.User.Password)
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err: errors.New("wrong username and password"), Type: gin.ErrorTypePrivate})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": formatter.NewUserResp(*user, token)})
}
