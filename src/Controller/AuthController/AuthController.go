package AuthController

import (
  "github.com/gin-gonic/gin"
  "net/http"
  // "log"
  "github.com/gin-gonic/gin/binding"
	// "strconv"
	DataBase "../../DataBase"
	// "golang.org/x/crypto/bcrypt"
	Auth "../../Auth"
	// "os"
	// "strconv"
	Model "../../Model"
	// "../../Cookie"
	"../../Table/UsersTable"
)

type CreateUserResponse struct {
	User *Model.User `json:"user"`
	Errors []Error `json:"errors"`
}

type Error struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Message string `json:"message"`
}

func SignUp(c *gin.Context) {
	user := Model.User{}
	user_bind_json := Model.UserBindingJson{}
	c.ShouldBindWith(&user_bind_json, binding.JSON)

	Model.BindUser(&user, user_bind_json)

	response := CreateUserResponse{}
	db := DataBase.New()

	errors := []Error{}
	new_error := Error{}
	find_users := []Model.User{}

	db.Table("users").Where("username = ?",  user.Username).Find(&find_users)
	if len(find_users) > 0 {
		new_error.Name = "username"
		new_error.Type = "duplicate"
		new_error.Message = "既に登録されたユーザ名です"
		errors = append(errors, new_error)
	}

	db.Table("users").Where("twitter_user_id = ?",  user.TwitterUserID).Find(&find_users)
	if len(find_users) > 0 {
		new_error.Name = "email"
		new_error.Type = "duplicate"
		new_error.Message = "既に登録されたユーザーです"
		errors = append(errors, new_error)
	}

	if len(errors) > 0 {
		response.User = &user
		response.Errors = errors
		c.JSON(http.StatusOK, response)
		return
	}

	UsersTable.Create(&user)
	response.User = &user
	response.Errors = errors

	c.JSON(http.StatusOK, response)

	authorize_token, _ :=  Auth.GenarateAuthorizeToken()
	authorize_token_hash := Auth.HashingAuthorizeToken(authorize_token)

	db.Model(&user).Update("authorize_token_hash", authorize_token_hash)

}


func GetSession(c *gin.Context) {
	interface_user, _ := c.Get("loginUser")
	user := interface_user.(Model.User)

	c.JSON(http.StatusOK, gin.H{"user": user})
}