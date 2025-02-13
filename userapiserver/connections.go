package userapiserver

import (
	"encoding/json"
	"net/http"

	models "github.com/ZolaraProject/user-api/models"
)

var (
	PkiVaultServiceHost string
	PkiVaultServicePort string
)

func writeStandardResponse(r *http.Request, w http.ResponseWriter, grpcToken string, message string) {
	responseObj := &models.Response{
		Token:   grpcToken,
		Message: message,
	}

	response, _ := json.Marshal(responseObj)
	w.Write(response)
}
