package Auth

import (
	// "time"
	"strconv"
	"encoding/base64"
	"net/url"
	"net/http"
	"bytes"
	"sort"
	"strings"
	"fmt"
	"crypto/hmac"
  "crypto/sha1"
	"encoding/json"
	"io"
)

type OAuth1 struct {
	ConsumerKey string
	ConsumerSecret string
  TokenSecret string
	Timestamp string
  Nonce string
  SignatureMethod string
  Version string
	Signature string
	Header map[string]string
}

func InitOAuth1 (
	ConsumerKey string,
	ConsumerSecret string,
	TokenSecret string,
) OAuth1 {

	oauth1 := OAuth1{}
	oauth1.Version = "1.0"
	oauth1.SignatureMethod = "HMAC-SHA1"
	oauth1.ConsumerKey = ConsumerKey
	oauth1.ConsumerSecret = ConsumerSecret
	oauth1.TokenSecret = TokenSecret

	// timestamp := int(time.Now().Unix())
	timestamp := 1616338781
	oauth1.Timestamp = strconv.Itoa(timestamp)
	oauth1.Nonce = oauth1.Timestamp
	// oauth1.Nonce = base64.StdEncoding.EncodeToString([]byte(string(timestamp))) // 文分けたい

	header := map[string]string{}
	header["oauth_version"] = oauth1.Version
	header["oauth_signature_method"] = oauth1.SignatureMethod
	header["oauth_consumer_key"] = oauth1.ConsumerKey
	// header["oauth_consumer_secret"] = oauth1.ConsumerSecret
	// header["oauth_token_secret"] = oauth1.TokenSecret
	header["oauth_timestamp"] = oauth1.Timestamp
	header["oauth_nonce"] = oauth1.Nonce

	oauth1.Header = header

	return oauth1
}

// hmac sha1 base64
func getHmacSha1(input, key string) string {
  // key_for_sign := []byte(key)
  // h := hmac.New(sha1.New, key_for_sign)
  // h.Write([]byte(input))
  // return base64.StdEncoding.EncodeToString(h.Sum(nil))

	key_for_sign := ([]byte)(key)
	hash := hmac.New(sha1.New, key_for_sign)
	io.WriteString(hash, input)
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}

func (oauth1 *OAuth1) generateSignature(uri, method string) string {
	keys := []string{}

	for key, _ := range oauth1.Header {
    keys = append(keys, key)
	}

	sort.Strings(keys)

	res := []string{}

	for i := 0; i < len(keys); i++ {
		key := keys[i]
		value := oauth1.Header[key]
		res = append(res, key + "=" + url.QueryEscape(value))
	}

	parameter_string := url.QueryEscape(strings.Join(res, "&"))
	// parameter_string := strings.Join(res, "&")
	uri = url.QueryEscape(uri)

	signature_base := method + "&" + uri + "&" + parameter_string
	signature_key := oauth1.ConsumerSecret + "&" + oauth1.TokenSecret

	fmt.Println("parameter_string", parameter_string)
	fmt.Println("signature_key", signature_key)
	fmt.Println("signature_base", signature_base)

	signature := getHmacSha1(signature_base, signature_key)
	fmt.Println("signature", signature)
	// signature = Base64.stringify(hash)
	return url.QueryEscape(signature)
}

func (oauth1 *OAuth1) generateOauthString(uri, method string) string {
	// keys := []string{}
	res := []string{}

	for key, value := range oauth1.Header {
		res = append(res, key + "=" + url.QueryEscape(value))
	}

	signature := oauth1.generateSignature(uri, method)
	fmt.Println("signature_escaped", signature)
	res = append(res, "oauth_signature=" + signature)

	return strings.Join(res, ",")
}

func (oauth1 *OAuth1) Post(url string, form map[string]string) http.Response {

	value, _ := json.Marshal(form)

	req, _ := http.NewRequest(
			"POST",
			url,
			bytes.NewBuffer(value),
	)

	// Content-Type 設定
	oauth := oauth1.generateOauthString(url, "POST")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauth)

	fmt.Println("req", req)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	// defer res.Body.Close()

	return *res
}

func (oauth1 *OAuth1) Get(uri string, params map[string]string) http.Response {
	values := url.Values{}
	for key, value := range params{
		values.Add(key, value)
	}

	uri = uri + "?" + values.Encode()

	req, _ := http.NewRequest(
		"GET",
		uri,
		nil,
	)

	fmt.Println("req", req)

	// Content-Type 設定
	oauth := oauth1.generateOauthString(uri, "GET")
	// req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "OAuth " + oauth)
	// req.Header.Set("Authorization", "OAuth oauth_consumer_key=\"laikETaGy65sjG3RPchRC5xRy\",oauth_token=\"1132314412637753344-284Vha6f2RwvBplwWRWocHpiqXOCL7\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1616338781\",oauth_nonce=\"1616338781\",oauth_version=\"1.0\",oauth_signature=\"tKOyMbbwGOgqgyeBYUYCq2DkLYk%3D\"")

	fmt.Println("oauth", oauth)
	fmt.Println("request_header", req.Header)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)
	// defer res.Body.Close()

	return *res
}
