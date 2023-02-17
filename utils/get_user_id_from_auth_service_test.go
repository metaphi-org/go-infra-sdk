package utils_test

import (
	"context"
	"os"
	"testing"

	"github.com/metaphi-org/go-infra-sdk/config"
	"github.com/metaphi-org/go-infra-sdk/constants"
	"github.com/metaphi-org/go-infra-sdk/utils"
	"github.com/stretchr/testify/assert"
)

var REGION = os.Getenv("AWS_REGION")
var LAMBDA_NAME = os.Getenv("TEST_LAMBDA")

func TestGetUserIdFromAuthService(t *testing.T) {
	user_id, err := utils.GetUserIdFromAuthService(context.Background(), config.CreateAWSConfig(REGION), LAMBDA_NAME, "41941e6345fb2a2c33f037f84d33b7f1fbae")
	assert.Nil(t, err)
	assert.NotNil(t, user_id)

	_, err = utils.GetUserIdFromAuthService(context.Background(), config.CreateAWSConfig(REGION), LAMBDA_NAME, "wrong_session_id")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), constants.ErrorInvalidSessionId)
}
