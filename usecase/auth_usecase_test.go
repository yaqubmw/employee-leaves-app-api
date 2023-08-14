package usecase

// import (
// 	"employeeleave/model"
// 	"employeeleave/model/dto"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type UserAuthUseCase interface {
// 	FindByUsernamePassword(username, password string) (model.UserCredential, error)
// 	// Define other methods of the UserUseCase interface if needed
// }

// type UserAuthUseCaseMock struct {
// 	mock.Mock
// }

// // FindAllUser implements UserUseCase.
// func (*UserAuthUseCaseMock) FindAllUser(requesPaging dto.PaginationParam) ([]model.UserCredential, dto.Paging, error) {
// 	panic("unimplemented")
// }

// // FindByIdUser implements UserUseCase.
// func (*UserAuthUseCaseMock) FindByIdUser(id string) (model.UserCredential, error) {
// 	panic("unimplemented")
// }

// // FindByUsername implements UserUseCase.
// func (*UserAuthUseCaseMock) FindByUsername(username string) (model.UserCredential, error) {
// 	panic("unimplemented")
// }

// // RegisterNewUser implements UserUseCase.
// func (*UserAuthUseCaseMock) RegisterNewUser(payload model.UserCredential) error {
// 	panic("unimplemented")
// }

// // UpdateUser implements UserUseCase.
// func (*UserAuthUseCaseMock) UpdateUser(payload model.UserCredential) error {
// 	panic("unimplemented")
// }

// // MockSecurity is a mock implementation of your security package functions.
// type MockSecurity struct {
// 	mock.Mock
// }

// // CreateAccessToken is a mocked implementation of the CreateAccessToken function.
// func (m *MockSecurity) CreateAccessToken(user model.UserCredential) (string, error) {
// 	args := m.Called(user)
// 	return args.String(0), args.Error(1)
// }

// func (m *UserAuthUseCaseMock) FindByUsernamePassword(username, password string) (model.UserCredential, error) {
// 	args := m.Called(username, password)
// 	return args.Get(0).(model.UserCredential), args.Error(1)
// }

// func TestLogin_Success(t *testing.T) {
// 	// Create a mock for UserUseCase
// 	userUseCaseMock := new(UserAuthUseCaseMock)

// 	// Set up expectations for the mock method
// 	expectedUser := dataDummy[0]
// 	userUseCaseMock.On("FindByUsernamePassword", "agung", "123").Return(expectedUser, nil)

// 	// Create an instance of authUseCase with the mock
// 	authUC := NewAuthUseCase(userUseCaseMock)

// 	// Create a mock for the security package
// 	securityMock := new(MockSecurity)
// 	securityMock.On("CreateAccessToken", expectedUser).Return("your_expected_token", nil)

// 	// Create an instance of authUseCase with the mocks
// 	authUC = userUseCase.NewAuthUseCase(userUseCaseMock, securityMock)

// 	// Call the Login function
// 	token, err := authUC.Login("agung", "123")

// 	// Assert that the returned token and error match the expected values
// 	assert.NoError(t, err)
// 	assert.Equal(t, "your_expected_token", token)

// 	// Verify that the methods on the mocks were called with the expected arguments
// 	userUseCaseMock.AssertExpectations(t)
// 	securityMock.AssertExpectations(t)
// }
