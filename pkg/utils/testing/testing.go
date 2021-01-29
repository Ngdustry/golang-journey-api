package testing

import "github.com/gorilla/mux"

type TestEnvironment struct {
	Router *mux.Router
}

func InitTestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	return api
}

func InitTestEnvironment() TestEnvironment {
	testRouter := InitTestRouter()

	testEnvironment := TestEnvironment{
		Router: testRouter,
	}

	return testEnvironment
}

func ResetTestRouter(testEnv *TestEnvironment) {
	newRouter := InitTestRouter()

	testEnv.Router = newRouter
}
