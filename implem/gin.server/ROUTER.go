package server

import (
	"fmt"
	"net/http"

	"strings"

	"errors"

	"github.com/gin-gonic/gin"
	"github.com/saeidraei/go-jwt-auth/uc"
)

type RouterHandler struct {
	ucHandler   uc.Handler
	authHandler uc.AuthHandler
	Logger      uc.Logger
}

func NewRouter(i uc.Handler, auth uc.AuthHandler) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
	}
}

func NewRouterWithLogger(i uc.Handler, auth uc.AuthHandler, logger uc.Logger) RouterHandler {
	return RouterHandler{
		ucHandler:   i,
		authHandler: auth,
		Logger:      logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(rH.errorCatcher())

	rH.usersRoutes(api)

}

func (rH RouterHandler) usersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	users.POST("", rH.userPost)            // Register a new user
	users.POST("/login", rH.userLoginPost) // Login for existing user

	user := api.Group("/user")
	user.GET("", rH.jwtMiddleware(), rH.userGet) // Gets the currently logged-in user
}

const userNameKey = "userNameKey"

func (rH RouterHandler) jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwt, err := getJWT(c.GetHeader("Authorization"))
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		userName, err := rH.authHandler.GetUserName(jwt)
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.SetAccepted()
		c.Set(userNameKey, userName)
		c.Next()
	}
}

func (rH RouterHandler) getUserNameFromToken(c *gin.Context) string {
	jwt, err := getJWT(c.GetHeader("Authorization"))
	if err != nil {
		return ""
	}

	userName, err := rH.authHandler.GetUserName(jwt)
	if err != nil {
		return ""
	}

	return userName
}

func getJWT(authHeader string) (string, error) {
	splitted := strings.Split(authHeader, "Token ")
	if len(splitted) != 2 {
		return "", errors.New("malformed header")
	}
	return splitted[1], nil
}

func (rH RouterHandler) errorCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Status() > 399 {
			rb := c.Errors.JSON()
			//if there was no error message show this message
			if c.Errors == nil {
				rb = gin.H{"error": "something went wrong"}
			}
			c.JSON(c.Writer.Status(), rb)
		}
	}
}

func (RouterHandler) getUserName(c *gin.Context) string {
	if userName, ok := c.Keys[userNameKey].(string); ok {
		return userName
	}
	return ""
}

// log is used to "partially apply" the title to the rH.logger.Log function
// so we can see in the logs from which route the log comes from
func (rH RouterHandler) log(title string) func(...interface{}) {
	return func(logs ...interface{}) {
		rH.Logger.Log(title, logs)
	}
}

func (RouterHandler) MethodAndPath(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.Path)
}
