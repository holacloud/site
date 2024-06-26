package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fulldump/box"
)

type HttpError struct {
	Status  int
	Name    string
	Message string
}

func (h HttpError) Error() string {
	return h.Message
}

func PrettyErrorInterceptor(next box.H) box.H {
	return func(ctx context.Context) {

		next(ctx)

		err := box.GetError(ctx)
		if err == nil {
			return
		}
		w := box.GetResponse(ctx)

		// if err == glueauth.ErrUnauthorized {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	json.NewEncoder(w).Encode(map[string]interface{}{
		// 		"error": map[string]interface{}{
		// 			"message":     err.Error(),
		// 			"description": fmt.Sprintf("user is not authenticated"),
		// 		},
		// 	})
		// 	return
		// }

		if err == box.ErrResourceNotFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": map[string]interface{}{
					"message":     err.Error(),
					"description": fmt.Sprintf("resource '%s' not found", box.GetRequest(ctx).URL.String()),
				},
			})
			return
		}

		if err == box.ErrMethodNotAllowed {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": map[string]interface{}{
					"message":     err.Error(),
					"description": fmt.Sprintf("method '%s' not allowed", box.GetRequest(ctx).Method),
				},
			})
			return
		}

		if _, ok := err.(*json.SyntaxError); ok {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": map[string]interface{}{
					"message":     err.Error(),
					"description": "Malformed JSON",
				},
			})
			return
		}

		if httperr, ok := err.(HttpError); ok {
			w.WriteHeader(httperr.Status)
			json.NewEncoder(w).Encode(map[string]any{
				"error": map[string]any{
					"name":    httperr.Name,
					"message": err.Error(),
				},
			})
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": map[string]interface{}{
				"message":     err.Error(),
				"description": "Unexpected error",
			},
		})

	}
}
