package util

import (
	"testing"

	"github.com/google/uuid"
	"github.com/haji-sudo/ShabehRoshan/models"
	"github.com/haji-sudo/ShabehRoshan/util"
	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	t.Run("Token Generation", func(t *testing.T) {
		userid := "74987e29-4141-42f3-bc15-963fc9baeefa"

		// Generate a token using the util.GenerateToken function
		token, err := util.GenerateToken(models.User{ID: uuid.MustParse(userid)})
		assert.NoError(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("Token Validation", func(t *testing.T) {
		userid := "74987e29-4141-42f3-bc15-963fc9baeefa"

		// Generate a token using the util.GenerateToken function
		token, err := util.GenerateToken(models.User{ID: uuid.MustParse(userid)})
		assert.NoError(t, err)
		assert.NotEmpty(t, token)

		// Validate the generated token using the util.ValidateToken function
		resID, err := util.ValidateToken(token)
		assert.NoError(t, err)
		assert.Equal(t, userid, resID)
	})
}
