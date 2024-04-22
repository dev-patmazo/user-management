package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dev-patmazo/user-management/models"
	"github.com/dev-patmazo/user-management/utils"
	"github.com/gorilla/mux"
)

// AuthenticationRBAC is a middleware function that performs role-based access control (RBAC) authentication.
// It checks if the request has valid basic authentication credentials and if the user has the necessary roles
// to access the requested route. If the authentication or authorization fails, it returns an HTTP unauthorized error.
// Otherwise, it calls the next handler in the chain.
//
// Parameters:
//   - next: The next http.Handler in the middleware chain.
//
// Returns:
//   - http.Handler: The middleware handler function.
//
// Example usage:
//
//	router := mux.NewRouter()
//	router.Use(AuthenticationRBAC)
//	router.HandleFunc("/protected", handleProtectedEndpoint)
//
//	func handleProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
//	  // Handle the protected endpoint logic here
//	}
func AuthenticationRBAC(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user, pass, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		result, err := models.CheckUserRole(user, utils.BasicAuthGenerator(user, pass))
		if err != nil {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		// Get the URL path pattern from the request.
		route := mux.CurrentRoute(r)
		pathPattern, _ := route.GetPathTemplate()

		// Check if the user has roles access for the route
		fmt.Println(pathPattern, r.Method)
		if !checkRouteRoles(result.Role, pathPattern, r.Method) {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func checkRouteRoles(role string, route string, method string) bool {

	roles := map[string]map[string][]string{
		"admin": {
			"GET": []string{
				"/users",
				"/users/{id}",
			},
			"POST": []string{
				"/users",
			},
			"PUT": []string{
				"/users/{id}",
			},
			"DELETE": []string{
				"/users/{id}",
			},
		},
		"editor": {
			"GET": []string{
				"/users",
				"/users/{id}",
			},
			"PUT": []string{
				"/users/{id}",
			},
		},
		"viewer": {
			"GET": []string{
				"/users/{id}",
			},
		},
	}

	for _, r := range roles[role][method] {
		if r == route {
			return true
		}
	}

	return false
}
