package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mocked version of your database client
type MockedDbClient struct {
	mock.Mock
}

func (m *MockedDbClient) DeleteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockedDbClient) GetUser(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockedDbClient) CreateUser(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockedDbClient) UpdateUser(id int, user *User) (*User, error) {
	args := m.Called(id, user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockedDbClient) CheckUserRole(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func TestDeleteUser(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MockedDbClient)

	// Setup expectations
	testObj.On("DeleteUser", 123).Return(nil)

	// Call the function we want to test
	err := testObj.DeleteUser(123)

	// Assert expectations
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestGetUser(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MockedDbClient)

	// Setup expectations
	expectedUser := &User{
		Username: "testuser",
		Email:    "testuser@gmail.com",
		Age:      25,
	}
	testObj.On("GetUser", 123).Return(expectedUser, nil)

	// Call the function we want to test
	user, err := testObj.GetUser(123)

	// Assert expectations
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestCreateUser(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MockedDbClient)

	// Setup expectations
	newUser := &User{
		Username: "testuser",
		Email:    "testuser@example.com",
	}
	testObj.On("CreateUser", newUser).Return(newUser, nil)

	// Call the function we want to test
	user, err := testObj.CreateUser(newUser)

	// Assert expectations
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, newUser, user)
}

func TestUpdateUser(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MockedDbClient)

	// Setup expectations
	updatedUser := &User{Username: "testuser", Email: "testuser@egmail.com"}
	testObj.On("UpdateUser", 123, updatedUser).Return(updatedUser, nil)

	// Call the function we want to test
	user, err := testObj.UpdateUser(123, updatedUser)

	// Assert expectations
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser, user)
}

func TestCheckUserRole(t *testing.T) {
	// Create an instance of our test object
	testObj := new(MockedDbClient)

	// Setup expectations
	expectedRole := "admin"
	testObj.On("CheckUserRole", 123).Return(expectedRole, nil)

	// Call the function we want to test
	role, err := testObj.CheckUserRole(123)

	// Assert expectations
	testObj.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, expectedRole, role)
}
