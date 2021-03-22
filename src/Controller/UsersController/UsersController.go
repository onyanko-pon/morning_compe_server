package UsersController

import (
  "github.com/gin-gonic/gin"
  "net/http"
  // "log"
  "github.com/gin-gonic/gin/binding"
	"strconv"
	DataBase "../../DataBase"
	"../../Model"
	"../../Table/UsersTable"
)

func GetUsers(c *gin.Context) {
	users := []Model.User{}
	db := DataBase.New()
	db.Find(&users) // 全レコード

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func UpdateUser(c *gin.Context) {
	user_binding_json := Model.UserBindingJson{}
	c.BindWith(&user_binding_json, binding.JSON)

	interface_user, _ := c.Get("LoginUser")
	user := interface_user.(Model.User)


	Model.BindUser(&user, user_binding_json)

	UsersTable.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUser(c *gin.Context) {
	user := Model.User{}
	var user_id int
	user_id, _ = strconv.Atoi(c.Param("id"))

	db := DataBase.New()
	db.Where("id = ?", user_id).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func DeleteUser(c *gin.Context) {
	user := Model.User{}
	var user_id int
	user_id, _ = strconv.Atoi(c.Param("id"))
	user.ID = user_id

	db := DataBase.New()
	db.Delete(&user)
	c.String(http.StatusOK, "ok")
}