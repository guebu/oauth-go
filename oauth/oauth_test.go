package oauth

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAuthenticateRequest(t *testing.T) {
	request := http.Request{
	}

	error := AuthenticateRequest(&request)

	assert.NotNil(t, error, "Func AuthenticateRequest with an empty request as argument should result in an error!")
}
