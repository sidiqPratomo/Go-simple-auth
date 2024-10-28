package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sidiqPratomo/DJKI-Pengaduan/appconstant"

	"github.com/google/uuid"
)

func RequestIdHandlerMiddleware(c *gin.Context) {
	uuid := uuid.NewString()
	c.Set(appconstant.RequestId, uuid)

	c.Next()
}
