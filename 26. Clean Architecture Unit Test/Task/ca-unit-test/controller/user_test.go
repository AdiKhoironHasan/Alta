package controller_test

import (
	"belajar-go-echo/controller/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mocksUser := new(mocks.User)
	// mocksUserData := user.User{
	// 	Username: "myusername",
	// 	Name:     "myname",
	// }

	t.Run("success", func(t *testing.T) {
		mocksUser.On("Create", mock.Anything).Return(nil).Once()

	})

	t.Run("failed", func(t *testing.T) {
		mocksUser.On("Create", mock.Anything).Return(errors.New("error unit test")).Once()

	})
}
