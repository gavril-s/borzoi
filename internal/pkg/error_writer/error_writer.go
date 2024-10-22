package errorwriter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gavril-s/borzoi/pkg/api"
)

func WriteError(w http.ResponseWriter, code int, message string, err error) {
	w.WriteHeader(code)
	resp := api.Error{
		ErrorCode: code,
		Message:   fmt.Sprintf("%s: %v", message, err),
	}
	json.NewEncoder(w).Encode(resp)
}
