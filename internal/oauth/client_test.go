package oauth

import (
	"fmt"
	"testing"
)

var (
	authPath  = "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey = "LlBKqo7Sg"
	secretKey = "9b5233fe714c11eeba8300163e3260ae"
)

func TestClient_GetToken(t *testing.T) {
	client := NewClient(authPath, accessKey, secretKey, false)
	for i := 0; i < 10; i++ {
		token, err := client.GetToken()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(token.AccessToken)
	}
}
