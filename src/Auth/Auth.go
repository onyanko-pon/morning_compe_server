package Auth

import (
    "time"
    "strconv"
    "../Jwt"
)

func GenerateJwtToken(user_id int) string {
    claims := Jwt.Claims{
        Sub: strconv.Itoa(user_id),
        Exp: int(time.Now().Add(time.Hour * 24).Unix()),
    }
    return Jwt.GenerateJwtToken(&claims)
}

func GetUserIDFromJwt(token string) string {
    return Jwt.TokenToClaims(token).Sub
}