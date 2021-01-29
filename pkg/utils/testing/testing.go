package testing

import "github.com/gorilla/mux"

// TestEnvironment aggregates all entities required for testing.
type TestEnvironment struct {
	Router *mux.Router
}

// InitTestRouter creates a new router for testing.
func InitTestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	return api
}

// InitTestEnvironment creates a new environment for testing.
func InitTestEnvironment() TestEnvironment {
	testRouter := InitTestRouter()

	testEnvironment := TestEnvironment{
		Router: testRouter,
	}

	return testEnvironment
}

// ResetTestRouter clears test router of endpoints.
func ResetTestRouter(testEnv *TestEnvironment) {
	newRouter := InitTestRouter()

	testEnv.Router = newRouter
}
