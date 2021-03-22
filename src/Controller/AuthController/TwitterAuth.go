package AuthController

import (
  "github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	Auth "../../Auth"
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
	"encoding/json"
)


type UserInfoJson struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ScreenName string `json:"screen_name"`
	Description string `json:"description"`
	ProfileImageUrl string `json:"profile_image_url"`
	ProfileBannerUrl string `json:"profile_banner_url"`
}

func GetRequestToken(c *gin.Context) {
	res := Auth.GetRequestToken()
	body_output, _ := ioutil.ReadAll(res.Body)
	body := string(body_output)
	values := strings.Split(body, "&")
	oauth_token := strings.Split(values[0], "=")[1]
	oauth_token_secret := strings.Split(values[1], "=")[1]

	c.JSON(http.StatusOK, gin.H{
		"oauth_token": oauth_token,
		"oauth_token_secret": oauth_token_secret,
		"url": "https://api.twitter.com/oauth/authorize?oauth_token=" + oauth_token,
	})
}

// type OAuthJson struct {
// 	OAuthToken string `json:"oauth_token"`
// 	OAuthVerifier string `json:"oauth_verifier"`
// }

func GetAccessToken(c *gin.Context) {
	// oauth_json := OAuthJson{}
	// c.BindWith(&oauth_json, binding.JSON)
	// res := Auth.GetAccessToken(oauth_json.OAuthToken, oauth_json.OAuthVerifier)
	oauth_token := c.Query("oauth_token")
	oauth_verifier := c.Query("oauth_verifier")
	res := Auth.GetAccessToken(oauth_token, oauth_verifier)

	body_output, _ := ioutil.ReadAll(res.Body)
	body := string(body_output)

	values := strings.Split(body, "&")

	// oauth_token := strings.Split(values[0], "=")[1]
	oauth_token_secret := strings.Split(values[1], "=")[1]
	user_id := strings.Split(values[2], "=")[1]
	screen_name := strings.Split(values[3], "=")[1]

	c.JSON(http.StatusOK, gin.H{
		"oauth_token": oauth_token,
		"oauth_token_secret": oauth_token_secret,
		"user_id": user_id,
		"screen_name": screen_name,
	})
}

func GetUserInfo(c *gin.Context) {
	oauth_token := c.Query("oauth_token")
	oauth_verifier := c.Query("oauth_verifier")
	res := Auth.GetAccessToken(oauth_token, oauth_verifier)

	body_output, _ := ioutil.ReadAll(res.Body)
	body := string(body_output)

	values := strings.Split(body, "&")

	fmt.Println("values", values)

	// oauth_token := strings.Split(values[0], "=")[1]
	// oauth_token_secret := strings.Split(values[1], "=")[1]
	// user_id := strings.Split(values[2], "=")[1]
	screen_name := strings.Split(values[3], "=")[1]

	res = Auth.GetUserInfo(screen_name)
	body_output, _ = ioutil.ReadAll(res.Body)

	var d UserInfoJson
	json.Unmarshal(body_output, &d)

	c.JSON(http.StatusOK, d)

	// c.JSON(http.StatusOK, gin.H{
	// 	"oauth_token": oauth_token,
	// 	"oauth_token_secret": oauth_token_secret,
	// 	"user_id": user_id,
	// 	"screen_name": screen_name,
	// })
}

// func GetUserInfo(c *gin.Context) {
// 	screen_name := c.Query("screen_name")

// 	fmt.Println("screen_name", screen_name)

// 	res := Auth.GetUserInfo(screen_name)
// 	body, _ := ioutil.ReadAll(res.Body)
// 	// body := string(body_output)
// 	// values := strings.Split(body, "&")

// 	// oauth_token := strings.Split(values[0], "=")[1]
// 	// oauth_token_secret := strings.Split(values[1], "=")[1]

// 	var d UserInfoJson
// 	json.Unmarshal(body, &d)

// 	c.JSON(http.StatusOK, d)

// 	// c.JSON(http.StatusOK, gin.H{
// 	// 	"body": body,
// 	// })
// }