package service_test

import (
	"errors"
	"testing"
	"wanderer/features/users"
	"wanderer/features/users/mocks"
	"wanderer/features/users/service"
	encMock "wanderer/helpers/encrypt/mocks"

	"github.com/stretchr/testify/assert"
)

func TestUserServiceRegister(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewBcryptHash(t)
	var srv = service.NewUserService(repo, enc)

	t.Run("invalid name", func(t *testing.T) {
		var caseData = users.User{
			Name:     "",
			Phone:    "08123456789",
			Email:    "galih@mail.com",
			Password: "test",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "name")
	})

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "08123456789",
			Email:    "",
			Password: "test",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "email")
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "08123456789",
			Email:    "galih@gmail.com",
			Password: "",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "password")
	})

	t.Run("invalid phone", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "",
			Email:    "galih@gmail.com",
			Password: "test",
		}

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "phone")
	})

	t.Run("error from encrypt", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "08123456789",
			Email:    "galih@gmail.com",
			Password: "test",
		}

		enc.On("Hash", caseData.Password).Return("", errors.New("some error from encrypt")).Once()

		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "some error from encrypt")

		enc.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "08123456789",
			Email:    "galih@gmail.com",
			Password: "test",
			Role:     "user",
			ImageUrl: "default",
		}

		enc.On("Hash", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", caseData).Return(errors.New("some error from repository")).Once()

		caseData.Password = "test"
		err := srv.Register(caseData)

		assert.ErrorContains(t, err, "some error from repository")

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Name:     "Galih",
			Phone:    "08123456789",
			Email:    "galih@gmail.com",
			Password: "test",
			Role:     "user",
			ImageUrl: "default",
		}

		enc.On("Hash", caseData.Password).Return("secret", nil).Once()

		caseData.Password = "secret"
		repo.On("Register", caseData).Return(nil).Once()

		caseData.Password = "test"
		err := srv.Register(caseData)

		assert.NoError(t, err)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)
	})
}

func TestUserServiceLogin(t *testing.T) {
	var repo = mocks.NewRepository(t)
	var enc = encMock.NewBcryptHash(t)
	var srv = service.NewUserService(repo, enc)

	t.Run("invalid email", func(t *testing.T) {
		var caseData = users.User{
			Email:    "",
			Password: "test",
		}

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "email")
		assert.Nil(t, result)
	})

	t.Run("invalid password", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "",
		}

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "password")
		assert.Nil(t, result)
	})

	t.Run("error from repository", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "test",
		}

		repo.On("Login", caseData.Email).Return(nil, errors.New("some error from repository")).Once()

		result, err := srv.Login(caseData.Email, caseData.Password)

		assert.ErrorContains(t, err, "some error from repository")
		assert.Nil(t, result)

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		var caseData = users.User{
			Email:    "galih@gmail.com",
			Password: "test",
		}

		var caseResult = users.User{
			Id:       1,
			Name:     "Galih",
			ImageUrl: "default",
			Password: "test",
			Role:     "user",
		}

		repo.On("Login", caseData.Email).Return(&caseResult, nil).Once()
		enc.On("Compare", caseResult.Password, caseData.Password).Return(nil).Once()
		res, err := srv.Login(caseData.Email, caseData.Password)

		enc.AssertExpectations(t)
		repo.AssertExpectations(t)

		assert.Nil(t, err)
		assert.Equal(t, "Galih", res.Name)
		assert.Equal(t, "user", res.Role)
	})
}