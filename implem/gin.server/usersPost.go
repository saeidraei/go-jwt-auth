package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saeidraei/go-jwt-auth/implem/json.formatter"
)

type userPostRequest struct {
	User struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	} `json:"user" binding:"required"`
}

func (rH RouterHandler) userPost(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	body := &userPostRequest{}
	if err := c.BindJSON(body); err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err: err, Type: gin.ErrorTypePublic})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	user, token, err := rH.ucHandler.UserCreate(body.User.Username, body.User.Email, body.User.Password)
	if err != nil {
		log(err)
		c.Errors = append(c.Errors, &gin.Error{Err: err, Type: gin.ErrorTypePublic})
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": formatter.NewUserResp(*user, token)})
}
