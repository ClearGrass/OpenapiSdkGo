package oauth

import (
	"context"
	"sync"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func NewClient(authPath, accessKey, secretKey string, useAgent bool) *Client {
	client := new(Client)
	client.authPath = authPath
	client.accessKey = accessKey
	client.secretKey = secretKey
	client.useAgent = useAgent
	return client
}

type Client struct {
	authPath   string
	accessKey  string
	secretKey  string
	useAgent   bool
	tokenStore sync.Map // 缓存token
}

func (c *Client) GetToken() (*oauth2.Token, error) {
	if token, exists := c.getTokenFromCache(); exists {
		if token.Expiry.Unix()-time.Now().Unix() > 0 {
			return token, nil
		}
	}

	token, err := c.generateToken()
	if err != nil {
		return nil, errors.Wrap(err, "auth token")
	}

	c.storeTokenInCache(token)
	return token, nil
}

func (c *Client) getTokenFromCache() (*oauth2.Token, bool) {
	token, exists := c.tokenStore.Load(authTokenKey)
	if !exists {
		return nil, false
	}

	return token.(*oauth2.Token), true
}

func (c *Client) storeTokenInCache(token *oauth2.Token) {
	c.tokenStore.Store(authTokenKey, token)
}

func (c *Client) generateToken() (*oauth2.Token, error) {
	scope := []string{scopeDeviceManage}
	if c.useAgent {
		scope = []string{scopeDeviceManage, scopeAgentManage}
	}
	config := clientcredentials.Config{
		ClientID:     c.accessKey,
		ClientSecret: c.secretKey,
		TokenURL:     c.authPath,
		Scopes:       scope,
	}

	return config.Token(context.Background())
}
