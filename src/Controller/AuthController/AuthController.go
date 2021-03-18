package AuthController

import (
  "github.com/gin-gonic/gin"
  "net/http"
  // "log"
  "github.com/gin-gonic/gin/binding"
	// "strconv"
	DataBase "../../DataBase"
	"golang.org/x/crypto/bcrypt"
	Auth "../../Auth"
	"../../GmailClient"
	"os"
	"strconv"
	Model "../../Model"
	"../../Cookie"
	"../../Table/UsersTable"
)

type SignInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(c *gin.Context) {
	user_bind_json := Model.UserBindingJson{}
	c.BindWith(&user_bind_json, binding.JSON)

	user := Model.User{}
	var username string
	username = user_bind_json.Username
	db := DataBase.New()
	db.Where("username = ?", username).First(&user)

	err := Auth.VerifyPassword(user.PasswordHash, user_bind_json.Password)
  if err != nil {
		c.String(http.StatusUnauthorized, "Not Auth")
		return
	}

	token := Auth.GenerateJwtToken(user.ID)
	Cookie.Set(c, "token", token)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
		"jwt_token": token,
	})
}

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

	hash, _ := bcrypt.GenerateFromPassword([]byte(user_bind_json.Password), bcrypt.DefaultCost)
	user.PasswordHash = string(hash)

	response := CreateUserResponse{}
	db := DataBase.New()

	errors := []Error{}
	new_error := Error{}
	find_users := []Model.User{}

	if len(user_bind_json.Password) < 9 {
		new_error.Name = "password"
		new_error.Type = "short"
		new_error.Message = "8文字以上パスワードを入力してください"
		errors = append(errors, new_error)
	}

	db.Table("users").Where("username = ?",  user.Username).Find(&find_users)
	if len(find_users) > 0 {
		new_error.Name = "username"
		new_error.Type = "duplicate"
		new_error.Message = "既に登録されたユーザ名です"
		errors = append(errors, new_error)
	}

	db.Table("users").Where("email = ?",  user.Email).Find(&find_users)
	if len(find_users) > 0 {
		new_error.Name = "email"
		new_error.Type = "duplicate"
		new_error.Message = "既に登録されたメールアドレスです"
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
	domain := os.Getenv("FRONTEND_DOMAIN")
	user_id_str := strconv.Itoa(user.ID)
	url := "http://" + domain + "/email_authorize_user?authrize_token=" + authorize_token + "&user_id=" + user_id_str

	mail := GmailClient.Mail{}
	mail.To = user.Email
	mail.Body = "登録が完了したのでお知らせします\n\r <a href=\"" + url +  "\">本登録する</a>"
	mail.Subject = "登録完了のお知らせ"

	GmailClient.Send(mail)
}


type EmailVerifyUserJsonParam struct {
	UserID string `json:"user_id"`
	AuthorizeToken string `json:"authorize_token"`
}

func EmailVerifyUser(c *gin.Context) {

	json_param := EmailVerifyUserJsonParam{}
	c.BindWith(&json_param, binding.JSON)

	user_id := json_param.UserID

	user :=  Model.User{}
	db := DataBase.New()
	result := db.Where("id = ?", user_id).First(&user)

	if result.Error != nil {
		c.String(http.StatusUnauthorized, "Not Auth")
		return
	}

	verify_password_err := Auth.VerifyAuthorizeToken(user.AuthorizeTokenHash, json_param.AuthorizeToken)
	if verify_password_err != nil {
		c.String(http.StatusUnauthorized, "Not Auth")
		return
	}

	db.Model(&user).Update("status", 2)
	db.Where("id = ?", user_id).First(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func GetSession(c *gin.Context) {
	interface_user, _ := c.Get("LoginUser")
	user := interface_user.(Model.User)

	c.JSON(http.StatusOK, gin.H{"user": user})
}