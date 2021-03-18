package Cookie

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Set(c *gin.Context, name string, value string) {
	go_env := os.Getenv("GO_ENV")

	if go_env == "production" {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     name,
			Value:    value,
			MaxAge:   100000,
			Path:     "",
			Domain:   "sagyo-with-me-server.herokuapp.com",
			SameSite: http.SameSiteNoneMode,
			Secure:   true,
			HttpOnly: true,
		})
	} else {
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     name,
			Value:    value,
			MaxAge:   100000,
			Path:     "",
			Domain:   "192.168.11.5:8080",
			SameSite: http.SameSiteNoneMode,
			Secure:   false,
			HttpOnly: true,
		})
	}
}