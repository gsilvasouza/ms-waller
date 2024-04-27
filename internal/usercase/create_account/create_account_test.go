package createaccount

import (
	"testing"

	"github.com/gsilvasouza/ms-waller/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCase(t *testing.T) {
	client, _ := entity.NewClient("John", "h0f7P@example.com")
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	createAccountUseCase := NewCreateAccountUseCase(clientGatewayMock, accountGatewayMock)
	input := &CreateAccountInputDTO{
		ClientID: client.ID,
	}
	output, err := createAccountUseCase.Execute(input)
	assert.Nil(t, err)
	assert.NotNil(t, output)
	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
