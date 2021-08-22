package birthday

import (
	"github.com/Liki4/Kyoto/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CountModel struct {
	Count int64
}

func CountBirthday(c *gin.Context) (int, int, interface{}) {
	counts := &CountModel{db.CountBirthday()}

	return http.StatusOK, 20000, counts
}
