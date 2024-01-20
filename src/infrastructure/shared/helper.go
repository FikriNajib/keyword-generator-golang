package infrashared

import (
	"math/rand"
	"strings"
)

func GetPlainToken(token string) string {
	if strings.HasPrefix(token, "Bearer") == true {
		tokenStringArr := strings.Split(token, " ")
		if len(tokenStringArr) > 1 {
			token = tokenStringArr[1]
		} else {
			token = tokenStringArr[0]
		}
	}
	return token
}

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(Charset[rand.Intn(len(Charset))])
	}
	return sb.String()
}

const NumSet = "1234567890"

func RandomNum(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(NumSet[rand.Intn(len(NumSet))])
	}
	return sb.String()
}
