package recovery

import (
	"fmt"
	"net/http"
	"project-adhyaksa/pkg/res"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Recover(ctx *gin.Context) {
	if r := recover(); r != nil {
		zap.L().Error(fmt.Sprintf("Recovered by error : %v", r))
		ctx.JSON(http.StatusInternalServerError, res.JSON(false, "Something went wrong", nil))
	}
}
