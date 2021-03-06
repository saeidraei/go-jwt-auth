package uc_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/saeidraei/go-jwt-auth/implem/uc.mock"
	"github.com/saeidraei/go-jwt-auth/testData"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate_happyCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	authToken := "token"
	rick := testData.User("rick")
	i := mock.NewMockedInteractor(mockCtrl)
	i.UserRW.EXPECT().Create(rick.Name, rick.Email, rick.Password).Return(&rick, nil).Times(1)
	i.AuthHandler.EXPECT().GenUserToken(rick.Name).Return(authToken, nil)
	retUser, token, err := i.GetUCHandler().UserCreate(rick.Name, rick.Email, rick.Password)

	assert.NoError(t, err)
	assert.Equal(t, authToken, token)
	assert.Equal(t, rick, *retUser)
}
