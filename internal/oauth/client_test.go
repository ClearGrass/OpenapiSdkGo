package oauth

import (
	"fmt"
	"testing"
)

var (
	authPath  = "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey = "8XL4IBGGR"
	secretKey = "b34bc5bfbf5611eabd0200163e2c48b3"
)

func TestClient_GetToken(t *testing.T) {
	client := NewClient(authPath, accessKey, secretKey)
	for i := 0; i < 10; i++ {
		token, err := client.GetToken()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(token.AccessToken)
	}
}
