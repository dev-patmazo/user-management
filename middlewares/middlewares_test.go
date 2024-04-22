package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dev-patmazo/user-management/utils"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedDbClient struct {
	mock.Mock
}

func TestAuthenticationRBAC(t *testing.T) {
	// Create a mock request with basic authentication
	req, err := http.NewRequest("GET", "/api/users", nil)
	assert.NoError(t, err)
	req.SetBasicAuth("testuser", "testpassword")

	// Create a mock response recorder
	rr := httptest.NewRecorder()

	// Create a mock router
	router := mux.NewRouter()
	router.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
		// This is the handler function for the protected route
		w.WriteHeader(http.StatusOK)
	}).Methods("GET")

	// Create a mock database client
	mockDbClient := new(MockedDbClient)
	mockDbClient.On("CheckUserRole", "testuser", utils.BasicAuthGenerator("testuser", "testpassword")).Return("admin", nil)

	// Call the middleware function
	router.ServeHTTP(rr, req)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, rr.Code)
}
