package repo

import (
	"fmt"
	"github.com/federicoleon/golang-restclient/rest"
	//"github.com/federicoleon/golang-restclient/rest"
	"github.com/guebu/common-utils/logger"
	"github.com/stretchr/testify/assert"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start the tests!!!!!!!!!")
	logger.Info("About to start mockup server for test!")
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetGetTokenByStringWithValidToken(t *testing.T) {

	fmt.Println("about to start the tests!!!!!!!!!")
	rest.FlushMockups()

	// Important: you have to include the PORT in the URL. Otherwise you get an error later on during the retreival
	rest.AddMockups(&rest.Mock{
		URL: "http://localhost:8081/oauth/access_token/eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxNzMwMDYsImp0aSI6IjQ3MTEiLCJpYXQiOjE5NzAwODE0LCJpc3MiOiJmaW1Ab3V0bG9vay5kZSIsInN1YiI6IkfDvEJ1cyBPQXV0aC1TZXJ2aWNlIn0.r9BgB9CPG694NgmURPDmWdltTQg4nPo82ziCnbp6rn8",
		HTTPMethod: http.MethodGet,
		//ReqBody: `{"email":"fim@outlook.de","the-password"}`,
		RespHTTPCode: 200,
		RespBody: `{"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxNjIyNDAsImlhdCI6MTk3MDA4MTQsImlzcyI6ImZpbUBvdXRsb29rLmRlIiwic3ViIjoiR8O8QnVzIE9BdXRoLVNlcnZpY2UifQ.UlErfk7fcSr2wh6DVTcZ7XsEEyWHwDf34Waz_8_A_H8", 
					"user_id": 4711, 
					"client_id": 100, 
					"expires_string": "2020-12-28 13:30:40.924092 +0000 UTC", 
					"expires_time": "2020-12-28T13:30:40.924Z", 
					"expires_int": 1609162240, 
					"is_already_expired": true}`,
		})

	repo := NewOauthRepository()

	// Access token should be available
	atString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxNzMwMDYsImp0aSI6IjQ3MTEiLCJpYXQiOjE5NzAwODE0LCJpc3MiOiJmaW1Ab3V0bG9vay5kZSIsInN1YiI6IkfDvEJ1cyBPQXV0aC1TZXJ2aWNlIn0.r9BgB9CPG694NgmURPDmWdltTQg4nPo82ziCnbp6rn8"
	at, err := repo.GetTokenByString(atString)

	assert.Nil(t, err, "Access token should be retreivable!")
	assert.NotNil(t, at, "Access token result should be available!")

	// Access token should not be available
	atString = "eyJhbGciOiJIUzI1NiIsInReHAiOjE2MDkxNzMwMDYsImp0aSI6IjQ3MTEiLCJpYXQiOjE5NzAwODE0LCJpc3MiOiJmaW1Ab3V0bG9vay5kZSIsInN1YiI6IkfDvEJ1cyBPQXV0aC1TZXJ2aWNlIn0.r9BgB9CPG694NgmURPDmWdltTQg4nPo82ziCnbp6rn8"
	at, err = repo.GetTokenByString(atString)

	assert.NotNil(t, err, "Access token should NOT be retreivable!")
	assert.Nil(t, at, "Access token result should NOT be available!")
}

func TestGetGetTokenByStringWithInvalidToken(t *testing.T) {

	fmt.Println("about to start the tests!!!!!!!!!")
	rest.FlushMockups()

	// Important: you have to include the PORT in the URL. Otherwise you get an error later on during the retreival
	rest.AddMockups(&rest.Mock{
		URL: "http://localhost:8081/oauth/access_token/gheyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxNzMwMDYsImp0aSI6IjQ3MTEiLCJpYXQiOjE5NzAwODE0LCJpc3MiOiJmaW1Ab3V0bG9vay5kZSIsInN1YiI6IkfDvEJ1cyBPQXV0aC1TZXJ2aWNlIn0.r9BgB9CPG694NgmURPDmWdltTQg4nPo82ziCnbp6rn8",
		HTTPMethod: http.MethodGet,
		//ReqBody: `{"email":"fim@outlook.de","the-password"}`,
		RespHTTPCode: 500,
		RespBody: `{
					"message": "Couldn't find access_token with ID gheyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MDkxNjIyNDAsImlhdCI6MTk3MDA4MTQsImlzcyI6ImZpbUBvdXRsb29rLmRlIiwic3ViIjoiR8O8QnVzIE9BdXRoLVNlcnZpY2UifQ.UlErfk7fcSr2wh6DVTcZ7XsEEyWHwDf34Waz_8_A_H8",
					"status": 500, 
					"code": "internal server error"
					}`,
	})

	repo := NewOauthRepository()

	// Access token should not be available
	atString := "eyJhbGciOiJIUzI1NiIsInReHAiOjE2MDkxNzMwMDYsImp0aSI6IjQ3MTEiLCJpYXQiOjE5NzAwODE0LCJpc3MiOiJmaW1Ab3V0bG9vay5kZSIsInN1YiI6IkfDvEJ1cyBPQXV0aC1TZXJ2aWNlIn0.r9BgB9CPG694NgmURPDmWdltTQg4nPo82ziCnbp6rn8"
	at, err := repo.GetTokenByString(atString)

	assert.NotNil(t, err, "Access token should NOT be retreivable!")
	assert.Nil(t, at, "Access token result should NOT be available!")
}
