package oauth

import (
	"github.com/guebu/common-utils/errors"
	"github.com/guebu/common-utils/logger"
	"github.com/guebu/oauth-go/config"
	rest "github.com/guebu/oauth-go/repo"
	"net/http"
	"strings"
)

type oauthClient struct {

}

type oauthInterface interface {

}

func IsPublic(request *http.Request) bool {
	logger.Info("Checking if request is public or not...", "app:ouath-go", "Layer:helper-service", "Func:IsPublic", "Status:Start")
	if request == nil {
		return true
	}
	isPublic := request.Header.Get(config.HeaderFieldXPublic) == config.HeaderFieldXPublicValuePublic


	if isPublic {
		logger.Info("Is request public? -- Answer: YES")
	} else {
		logger.Info("Is request public? -- Answer: NO")
	}
	logger.Info("Checking if request is public or not...", "app:ouath-go", "Layer:helper-service", "Func:IsPublic", "Status:Start")
	return isPublic

}

func AuthenticateRequest(request *http.Request) *errors.ApplicationError {
	logger.Info("Authenticating request is starting...", "app:ouath-go", "Layer:helper-service", "Func:AuthenticateRequest", "Status:Start")
	if request == nil {
		err := errors.NewInternalServerError("No Request information was transmitted!")
		logger.Error("No Request information was transmitted!", err, "app:ouath-go", "Layer:helper-service", "Func:AuthenticateRequest", "Status:Error")
	}

	// try to read token in request...
	bearerToken := request.Header.Get("Authorization")
	if (bearerToken == "") {
		logger.Info("No token information transmitted in request...", "app:ouath-go", "Layer:helper-service", "Func:GetUser", "Level:Info",  "State:Error")
		apiErr := errors.NewBadRequestError("Token information must be transmitted inside the request!")
		// c.JSON(apiErr.AStatusCode, apiErr)
		// log.Printf("Controller-Error: BAD REQUEST - No token information transmitted!")
		return apiErr
	}

	logger.Info("Bearer-Token with ID read - Before Trimoperation --- " + bearerToken)
	// delete the präfix "Bearer " from bearer token string
	bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
	logger.Info("Bearer-Token with ID read - After Trimoperation ---  " + bearerToken)

	// token transmitted in request...
	// Now we have to check if the token is valid
	oauthRepo := rest.NewOauthRepository()
	accessToken, oauthErr := oauthRepo.GetTokenByString(bearerToken)

	if oauthErr != nil {
		logger.Info("Detailed token information couldn't be retreived form oauth API...", "Layer:Controller", "Func:GetUser", "Level:Info",  "State:Error")
		apiErr := errors.NewInternalServerError("Detailed token information couldn't be retreived form oauth API...")
		// c.JSON(apiErr.AStatusCode, apiErr)
		// log.Printf("Controller-Error: BAD REQUEST - No token information retreivable...!")
		return apiErr
	}

	if accessToken.IsAlreadyExpired {
		logger.Info("Access Token already expired!", "Layer:Controller", "Func:GetUser", "Level:Info",  "State:Error")
		apiErr := errors.NewTokenExpiredError("Access Token already expired")
		// c.JSON(apiErr.AStatusCode, apiErr)
		// log.Printf("Controller-Error: BAD REQUEST - No token information retreivable...!")
		return apiErr
	}

	logger.Info("-----------------------------------------------------------------")
	logger.Info("Access-Token-String: " + accessToken.AccessToken)
	logger.Info("-----------------------------------------------------------------")

	logger.Info("Authenticating request is successfully finished", "app:ouath-go", "Func:AuthenticateRequest", "Status:End")

	return nil
}
