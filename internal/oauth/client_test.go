package oauth

import (
	"fmt"
	"testing"
)

var (
	authPath  = "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey = "GhTBXTZGg"
	secretKey = "f4cfd224b43d11ea8bf400163e2c48b3"
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
