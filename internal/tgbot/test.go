package tgbot

import (
	"fmt"
	"github.com/Liki4/Kyoto/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"net/http"
	log "unknwon.dev/clog/v2"
)

func TestTelegramBot(c *gin.Context) (int, int, interface{}) {
	bot, err := tgbotapi.NewBotAPI(conf.Get().TgBot.Token)
	if err != nil {
		log.Error(err.Error())
		return http.StatusInternalServerError, 50000, "TelegramBot Failed"
	}

	bot.Debug = true

	log.Trace(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	chanmsg := tgbotapi.NewMessageToChannel(conf.Get().TgBot.TargetChannel, "test")
	_, err = bot.Send(chanmsg)
	if err != nil {
		log.Error(fmt.Sprintf("send msg err: %v, msg: %v", err, chanmsg))
	}
	//usermsg := tgbotapi.NewMessage(conf.Get().TgBot.TargetUser, "test")
	//_, err = bot.Send(usermsg)
	//if err != nil {
	//	log.Error(fmt.Sprintf("send msg err: %v, msg: %v", err, usermsg))
	//}

	return http.StatusOK, 20000, "Msg Sent"
}
