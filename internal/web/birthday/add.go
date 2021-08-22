package birthday

import (
	"fmt"
	"github.com/Liki4/Kyoto/internal/db"
	"github.com/gin-gonic/gin"
	"net/http"
	log "unknwon.dev/clog/v2"
)

type AddModel struct {
	Animate string `json:"animate"`
	Name    string `json:"name"`
	Date    string `json:"date"`
}

func AddBirthday(c *gin.Context) (int, int, interface{}) {
	addModel := new(AddModel)
	err := c.ShouldBindJSON(addModel)
	//log.Trace("body: %+v", addModel)
	if err != nil {
		log.Warn("can't bind json")
		return http.StatusBadRequest, 40000, "can't bind json"
	}

	newBirthday := &db.Birthday{Animate: addModel.Animate, Name: addModel.Name, Date: addModel.Date}

	if err := db.AddBirthday(newBirthday); err != nil {
		if db.IsErrBirthdayAlreadyExist(err) {
			log.Error(fmt.Sprintf("%s AlreadyExist", newBirthday.Name))
			return http.StatusInternalServerError, 50000, fmt.Sprintf("%s AlreadyExist", newBirthday.Name)
		}
		log.Error(fmt.Sprintf("Add %s Failed", newBirthday.Name))
		return http.StatusInternalServerError, 50000, fmt.Sprintf("Add %s Failed", newBirthday.Name)
	}
	log.Trace(fmt.Sprintf("%s Added", newBirthday.Name))
	return http.StatusOK, 20000, fmt.Sprintf("%s Added", newBirthday.Name)
}
