package createclient

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

func TestCreateClientUseCase(t *testing.T) {
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Save", mock.Anything).Return(nil)
	uc := NewCreateClientUseCase(clientGatewayMock)

	output, err := uc.Execute(CreateClientInputDTO{
		Name:  "John",
		Email: "h0f7P@example.com",
	})
	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, "h0f7P@example.com", output.Email)
	assert.Equal(t, "John", output.Name)
	clientGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
