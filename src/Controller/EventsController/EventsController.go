package EventsController

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "log"
  "github.com/gin-gonic/gin/binding"
	"strconv"
	DataBase "../../DataBase"
	"time"
	Model "../../Model"
)

func newEvent() *Model.Event {
	var event Model.Event
	event.StartTime = time.Now()
	event.EndTime = time.Now()
	return &event
}

func GetEvents(c *gin.Context) {
	events := []Model.Event{}

	db := DataBase.New()
	db.Preload("Users").Preload("HostUser").Find(&events) // 全レコード

	token, _ := c.Cookie("token")
	log.Print(token)

	c.JSON(http.StatusOK, gin.H{
		"events": events,
	})
}

func CreateEvent(c *gin.Context) {
	interface_user, _ := c.Get("LoginUser")
	user := interface_user.(Model.User)

	event := *newEvent()

	c.BindWith(&event, binding.JSON)
	event.HostUserID = user.ID
	db := DataBase.New()
	db.Create(&event)

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func UpdateEvent(c *gin.Context) {
	event := *newEvent()
	c.BindWith(&event, binding.JSON)

	var event_id int
	event_id, _ = strconv.Atoi(c.Param("id"))
	event.ID = event_id

	db := DataBase.New()
	db.Model(&event).Updates(&event)

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func GetEvent(c *gin.Context) {
	event := Model.Event{}
	var event_id int
	event_id, _ = strconv.Atoi(c.Param("id"))

	db := DataBase.New()
	db.Where("id = ?", event_id).First(&event)

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}

func DeleteEvent(c *gin.Context) {
	event := Model.Event{}
	var event_id int
	event_id, _ = strconv.Atoi(c.Param("id"))
	event.ID = event_id

	db := DataBase.New()
	db.Delete(&event)
	c.String(http.StatusOK, "ok")
}