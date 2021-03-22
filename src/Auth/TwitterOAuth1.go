package Auth

import (
	"os"
	"net/http"
	"net/url"
	"fmt"
)

func GetRequestToken() http.Response {
	ConsumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	ConsumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	oauth1 := InitOAuth1(ConsumerKey, ConsumerSecret, "")
	oauth_callback := os.Getenv("TWITTER_CALLBACK_URL")
	oauth1.Header["oauth_callback"] = oauth_callback

	url := "https://api.twitter.com/oauth/request_token"
	res := oauth1.Post(url, map[string]string{"oauth_callback": oauth_callback})

	return res
}

func GetAccessToken(oauth_token, oauth_verifier string) http.Response {
	uri := "https://api.twitter.com/oauth/access_token"

	args := url.Values{}
	args.Add("oauth_token", oauth_token)
	args.Add("oauth_verifier", oauth_verifier)
	res, _ := http.PostForm(uri, args)

	return *res
}


func GetUserInfo(screen_name string) http.Response {
	uri := "https://api.twitter.com/1.1/users/show.json?screen_name="+screen_name

	// ConsumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	// ConsumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")

	BearerToken := os.Getenv("TWITTER_BEARER_TOKEN")

	req, _ := http.NewRequest(
		"GET",
		uri,
		nil,
	)

	req.Header.Set("Authorization", "Bearer " + BearerToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	return *res
}