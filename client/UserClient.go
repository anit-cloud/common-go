package client

import (
	"net/http"

	"github.com/anit-cloud/common-go/client/resource"
)

type UserClient struct {
	EntryPoint string
	*BaseClient
}

func NewUserClient(host string, timeout int) *UserClient {
	return &UserClient{"/v1/user", NewBaseClient(host, timeout)}
}

func (c *UserClient) GetUser(requestResource *resource.RequestResource) (*http.Response, error) {
	url := c.Host + c.EntryPoint

	return c.DoPostRequest(url, requestResource)
}
