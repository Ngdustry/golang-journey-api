package router

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
		} else {
			bearer := r.Header.Get("Authorization")
			token := strings.Replace(bearer, "Bearer ", "", 1)
			googleURL := "https://www.googleapis.com/oauth2/v3/tokeninfo?id_token=" + token

			res, err := http.Get(googleURL)
			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			var data map[string]interface{}

			if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
				panic(err)
			}

			ctx := context.WithValue(r.Context(), "Email", data["email"])

			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
