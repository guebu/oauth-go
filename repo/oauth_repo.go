package repo

import (
	"encoding/json"
	"fmt"
	"github.com/federicoleon/golang-restclient/rest"
	"github.com/guebu/common-utils/errors"
	"github.com/guebu/common-utils/logger"
	"github.com/guebu/oauth-go/domain"
	"time"
)

var (
	oauthRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost"  + ":8081",
		Timeout: 1000 * time.Millisecond,
	}
)

type RestOAuthRepository interface {
	GetTokenByString(string) (*domain.AccessTokenResult, *errors.ApplicationError)
}

type oauthRepository struct{

}

func init() {
	logger.Info("In Init-Function of oauth_repository.go...", "Layer:Repo", "App:helper-service", "Func:init", "Status:Open")
	// ToDo: Clarifiy mess with BaseURL...
	oauthRestClient.BaseURL = "http://localhost"  + ":8081"
	oauthRestClient.Timeout = 1000 * time.Millisecond
	logger.Info("In Init-Function of oauth_repository.go...", "Layer:Repo", "App:users", "Func:init", "Status:End")
}

func NewOauthRepository() RestOAuthRepository{
	return &oauthRepository{}
}

func (ur *oauthRepository) GetTokenByString(tokenString string) (*domain.AccessTokenResult, *errors.ApplicationError) {
	logger.Info("Start to get access token by given string...", "Layer:Repo", "Func:GetTokenByString", "Status:Start", "app:user")

	response := oauthRestClient.Get("/oauth/access_token/" + tokenString)

	fmt.Println( "Body: " + response.String() )
	fmt.Println("----------------------")

	if response == nil || response.Response == nil {
		// we run in a technical error situation, e.g. a timeout situation...
		err := errors.NewInternalServerError("Invalid restclient response when trying to retrieve Access Token from oauth API!", nil)
		logger.Error("Invalid restclient response when trying to retrieve Access Token from oauth API!", err, "Layer:Repo", "Func:Login", "Status:Error", "app:user")
		return nil, err
	}

	fmt.Println("----------------------")
	fmt.Println(fmt.Sprintf("Error-Status-Code: %d", response.StatusCode))
	fmt.Println("----------------------")

	if response.StatusCode > 299 {
		// we run in an other error situation, which have to be handled
		var apiErr errors.ApplicationError
		err := json.Unmarshal(response.Bytes(), &apiErr)
		if err != nil {
			logger.Error("Invalid error interface when trying to retrieve Access Token Information from oauth API!", err, "app:user", "Layer:Repo", "Func:Login", "Status:Error")
			return nil, errors.NewInternalServerError("Invalid error interface when trying to retrieve Access Token Information from oauth API!", err)
		}
		logger.Info("Attempt to retreive Access Token Information from oauth API results in an error! Message: "+apiErr.AMessage, "app:user", "Layer:Repo", "Func:Login", "Status:Error")
		return nil, &apiErr
	}

	var at domain.AccessTokenResult

	if err:=json.Unmarshal(response.Bytes(), &at); err != nil {
		logger.Error("Invalid Access Token Interface when trying to retreive Access Token Information...!", err, "app:user", "Layer:Repo", "Func:Login", "Status:Error")
		return nil, errors.NewInternalServerError("Invalid Access Token Interface when trying to retreive Access Token Information...!", err)
	}
	logger.Info("End of retreiving Access Token information...", "app:user", "Layer:Repo", "Func:Login", "Status:End")
	return &at, nil
}
