package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
  Auth "./Auth"
  "github.com/gin-contrib/cors"
  // "log"
  "os"
  // DataBase "./DataBase"
  UsersController "./Controller/UsersController"
  AuthController "./Controller/AuthController"
  "./Jwt"
  "errors"
  "./DataBase"
  "./Model"
  // Model "./Model"
)

// 認証が必要なリクエストの時に走らせる処理
func AuthRequiredMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    var err error
    var token string
    user := Model.User{}
    token, err = c.Cookie("token")

    if err != nil {
      c.String(http.StatusUnauthorized, err.Error())
      c.Abort()
    }

    valid := Jwt.Valid(token)
    if !valid {
      err = errors.New("invalid token")
      c.String(http.StatusUnauthorized, err.Error())
      c.Abort()
    }

    user_id := Auth.GetUserIDFromJwt(token)
    db := DataBase.New()
    result := db.Model("users").Where("id = ?", user_id).First(&user)
    err = result.Error

    if err != nil {
        c.String(http.StatusUnauthorized, err.Error())
        c.Abort()
    }
    c.Set("loginUser", user)
  }
}

func main() {

  engine := gin.Default()
  config := cors.DefaultConfig()
  // config.AllowOrigins = []string{"http://localhost:3000"}
  config.AllowOrigins = []string{"*"}
  config.AllowHeaders = []string{"*"}
  // Origin, X-Requested-With, Content-Type, Accept'
  config.AllowHeaders = []string{"Origin, X-Requested-With, Content-Type, Accept"}
  config.AllowCredentials = true
  engine.Use(cors.New(config))

  authNeedEngine := engine.Group("/")
  authNeedEngine.Use(AuthRequiredMiddleware())

  engine.GET("/hello", func(c *gin.Context) {
		message := "hello world"
    c.JSON(http.StatusOK, gin.H{"message": message})
  })

  engine.GET("/get_request_token", AuthController.GetRequestToken)
  engine.GET("/get_access_token", AuthController.GetAccessToken)
  engine.GET("/get_user_info", AuthController.GetUserInfo)

  engine.GET("/users", UsersController.GetUsers)
  authNeedEngine.PUT("/users/:id", UsersController.UpdateUser)
  // router.GET("/users/:id", UsersController.GetUser)
  // router.DELETE("/users/:id",  UsersController.DeleteUser)

  // engine.POST("/signup", AuthController.SignUp)
  // engine.POST("/signin", AuthController.SignIn)
  // authNeedEngine.POST("/get_session", AuthController.GetSession)
  // engine.POST("/email_authorize_user", AuthController.EmailVerifyUser)

  port := os.Getenv("PORT")
  engine.Run(":" + port)
}