package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saeidraei/go-jwt-auth/implem/json.formatter"
)

func (rH RouterHandler) profileGet(c *gin.Context) {
	log := rH.log(rH.MethodAndPath(c))

	user, follows, err := rH.ucHandler.ProfileGet(rH.getUserNameFromToken(c), c.Param("username"))
	if err != nil {
		log(err)
		c.Status(http.StatusUnprocessableEntity)
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": formatter.NewProfileFromDomain(*user, follows)})
}
