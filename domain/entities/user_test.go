package entities_test

import (
	"testing"

	"github.com/JoaoRafa19/plataforma-ead-service-core/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity(t *testing.T) {
	t.Run("Should create a new user successfully", func(t *testing.T) {
		user, err := entities.NewUser("name", "teste@teste", "asdlfkjalsçdkf", "123456789")
		assert.Nil(t, err)
		assert.NotNil(t, user)

		assert.NotEmpty(t, user.ID)
		assert.Equal(t, user.Email, "teste@teste")
		assert.Equal(t, user.Phone, "123456789")
	})
	t.Run("Should return error in name", func(t *testing.T) {
		_, err := entities.NewUser("", "teste@asdklfjð", "asdlfkjalsçdkf", "123456789")
		assert.EqualError(t, err, "User Name invalid")
	})
	t.Run("Should return error in email", func(t *testing.T) {
		_, err := entities.NewUser("teste", "", "asdlfkjalsçdkf", "123456789")
		assert.EqualError(t, err, "User Email invalid")
	})
	t.Run("Should return error in password", func(t *testing.T) {
		_, err := entities.NewUser("teste", "teste", "", "123456789")
		assert.EqualError(t, err, "User Password invalid")
		_, err = entities.NewUser("teste", "teste", "teste", "")
		assert.EqualError(t, err, "User Password invalid")
	})
	t.Run("Should return errror in phone", func(t *testing.T) {
		_, err := entities.NewUser("teste", "teste", "testeste", "")
		assert.EqualError(t, err, "User Phone invalid")
	})
	
}
