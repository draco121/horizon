package clients

import (
	"github.com/draco121/common/constants"
	"github.com/draco121/common/models"
	"github.com/draco121/common/utils"
	"github.com/go-resty/resty/v2"
)

type IAuthorizationServiceApiClient interface {
	Authorize(input models.AuthorizationInput) (*models.AuthorizationOutput, error)
}

type authorizationServiceApiClient struct {
	IAuthorizationServiceApiClient
	client *resty.Client
}

func NewAuthorizationServiceApiClient(baseUrl string) IAuthorizationServiceApiClient {
	utils.Logger.Info("initializing authorization service client")
	c := resty.New()
	c.BaseURL = baseUrl
	client := authorizationServiceApiClient{
		client: c,
	}
	utils.Logger.Info("authorization service client initialized")
	return &client
}

func (c *authorizationServiceApiClient) Authorize(input models.AuthorizationInput) (*models.AuthorizationOutput, error) {
	utils.Logger.Info("attempting to authorize user")
	var result models.AuthorizationOutput
	_, err := c.client.R().
		SetResult(&result).
		SetBody(input).
		ForceContentType("application/json").
		Post("v1/authorize")
	if err != nil {
		return nil, err
	} else {
		if result.Grant == constants.Rejected {
			utils.Logger.Error("failed to authorize user")
			return &result, nil
		} else {
			utils.Logger.Info("authorized user successfully")
			return &result, nil
		}
	}

}
