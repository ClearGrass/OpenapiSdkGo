package main

import (
	"context"
	"fmt"

	"github.com/ClearGrass/OpenapiSdkGo/openapi"
)

func main() {
	apiHost := "http://api.test.cleargrass.com:9181"
	authPath := "http://oauth.test.cleargrass.com/oauth2/token"
	accessKey := "8XL4IBGGR"
	secretKey := "b34bc5bfbf5611eabd0200163e2c48b3"
	client := openapi.NewClient(apiHost, authPath, accessKey, secretKey)

	res, err := client.DeviceList(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
