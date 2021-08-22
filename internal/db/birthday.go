package db

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"
)

// Birthday 数据库表结构
type Birthday struct {
	gorm.Model

	Animate string
	Name    string
	Date    string
}

type ErrBirthdayAlreadyExist struct {
	arg interface{}
}

func IsBirthdayExist(animate string, name string) bool {
	if name == "" {
		return false
	}
	var b Birthday
	db.Model(&Birthday{}).Where(&Birthday{Animate: animate, Name: name}).Find(&b)
	return b.ID != 0
}

func (err ErrBirthdayAlreadyExist) Error() string {
	return fmt.Sprintf("生日已经收录啦: %v", err.arg)
}

func IsErrBirthdayAlreadyExist(err error) bool {
	_, ok := err.(ErrBirthdayAlreadyExist)
	return ok
}

// AddBirthday 添加生日
func AddBirthday(b *Birthday) error {
	isExist := IsBirthdayExist(b.Animate, b.Name)
	if isExist {
		return ErrBirthdayAlreadyExist{arg: b.Name}
	}

	tx := db.Begin()
	if tx.Create(b).RowsAffected != 1 {
		tx.Rollback()
		return errors.Errorf("数据库错误")
	}
	tx.Commit()
	return nil
}

// GetBirthday 批量获取生日
// options[0] offset
// options[1] limit
func GetBirthday() []*Birthday {
	var birthdays []*Birthday

	db.Model(&Birthday{}).Find(&birthdays)

	return birthdays
}

// GetTodayBirthday 获取当日生日
func GetTodayBirthday() []*Birthday {
	var birthdays []*Birthday

	TodayString := fmt.Sprintf("%02d-%02d", time.Now().Month(), time.Now().Day())
	db.Model(&Birthday{}).Where(&Birthday{Date: TodayString}).Find(&birthdays)

	return birthdays
}

// CountBirthday 返回生日的总数
func CountBirthday() int64 {
	var count int64
	db.Model(&Birthday{}).Count(&count)
	return count
}
