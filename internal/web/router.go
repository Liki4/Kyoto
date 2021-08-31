package web

import (
	"fmt"
	"github.com/Liki4/Kyoto/internal/conf"
	"github.com/Liki4/Kyoto/internal/web/birthday"
	. "github.com/Liki4/Kyoto/toolkit"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	b := r.Group("/birthday")
	b.POST("/add", Entry(birthday.AddBirthday))
	//b.GET("/get", Entry(birthday.GetBirthday))
	//b.GET("/getToday", Entry(birthday.GetTodayBirthday))
	//b.GET("/count", Entry(birthday.CountBirthday))

	//t := r.Group("/tgbot")
	//t.GET("/test", Entry(tgbot.TestTelegramBot))

	err := r.Run(fmt.Sprintf(":%d", conf.Get().Site.Port))
	if err != nil {
		panic(err)
	}
}
