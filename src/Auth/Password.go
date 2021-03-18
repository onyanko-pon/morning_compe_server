package Auth

import (
    "errors"
    "golang.org/x/crypto/bcrypt"
    "crypto/rand"
)

func HashingPassword(password string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(hash)
}

func VerifyPassword(hash, pw string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}

func VerifyAuthorizeToken(hash, token string) error {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(token))
    return err
}

func HashingAuthorizeToken(authorize_token string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(authorize_token), bcrypt.DefaultCost)
    return string(hash)
}

func GenarateAuthorizeToken() (string, error) {
    return makeRandomStr(10)
}

func makeRandomStr(digit uint32) (string, error) {
    const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

    // 乱数を生成
    b := make([]byte, digit)
    if _, err := rand.Read(b); err != nil {
        return "", errors.New("unexpected error...")
    }

    // letters からランダムに取り出して文字列を生成
    var result string
    for _, v := range b {
        // index が letters の長さに収まるように調整
        result += string(letters[int(v)%len(letters)])
    }
    return result, nil
}
