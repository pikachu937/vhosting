package response

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mikerumy/vhosting/internal/models"
)

func ErrorCannotResponseProperly(ctx *gin.Context, err error) {
	Response(ctx, models.Log{
		ErrorLevel: ErrLevelError,
		ErrorCode:  701,
		Message:    fmt.Sprintf("Cannot response properly. Error: %s.", err.Error()),
	})
}
