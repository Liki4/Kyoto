package birthday

import (
	"github.com/Liki4/Kyoto/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetModel struct {
	Birthdays []*db.Birthday
}

func GetBirthday(c *gin.Context) (int, int, interface{}) {
	birthdays := &GetModel{db.GetBirthday()}

	return http.StatusOK, 20000, birthdays
}

func GetTodayBirthday(c *gin.Context) (int, int, interface{}) {
	birthdays := &GetModel{db.GetTodayBirthday()}

	return http.StatusOK, 20000, birthdays
}
