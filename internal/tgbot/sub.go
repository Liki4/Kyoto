package tgbot

import (
	"fmt"
	"github.com/Liki4/Kyoto/internal/conf"
	"github.com/Liki4/Kyoto/internal/db"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
	"time"
	log "unknwon.dev/clog/v2"
)

func SendTodayBirthdaysToChannel() {
	bot, err := tgbotapi.NewBotAPI(conf.Get().TgBot.Token)
	if err != nil {
		log.Error(err.Error())
		return
	}

	bot.Debug = true

	//log.Trace(fmt.Sprintf("Authorized on account %s", bot.Self.UserName))

	birthdays := db.GetTodayBirthday()
	if len(birthdays) == 0 {
		return
	}

	msg := fmt.Sprintf(MsgMainTpl, time.Now().Month(), time.Now().Day())
	for _, birthday := range birthdays {
		msg += fmt.Sprintf(BirthdayTpl, birthday.Animate, birthday.Name)
	}
	msg = msg + "\n```"

	chanmsg := tgbotapi.NewMessageToChannel(conf.Get().TgBot.TargetChannel, msg)
	chanmsg.ParseMode = "Markdown"
	_, err = bot.Send(chanmsg)
	if err != nil {
		log.Error(fmt.Sprintf("send msg err: %v, msg: %v", err, chanmsg))
	}
}

func Run() {
	c := cron.New()
	//if err := c.AddFunc("@every 1m", SendTodayBirthdaysToChannel); err != nil {
	if err := c.AddFunc("@daily", SendTodayBirthdaysToChannel); err != nil {
		log.Error("Cron Failed")
		return
	}
	c.Start()
}
