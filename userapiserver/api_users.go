package userapiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/grpc"

	grpctoken "github.com/ZolaraProject/library/grpctoken"
	logger "github.com/ZolaraProject/library/logger"
	pkiVaultService "github.com/ZolaraProject/pki-vault-service/pkivaultrpc"
	models "github.com/ZolaraProject/user-api/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx, grpcToken := grpctoken.CreateContextFromHeader(r)

	// Create gRPC client
	conn, err := grpc.Dial(fmt.Sprintf("%v:%v", PkiVaultServiceHost, PkiVaultServicePort), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeStandardResponse(r, w, grpcToken, fmt.Sprintf("CreateAbstractClass could not establish gRPC connection: %v", err))
		return
	}
	defer conn.Close()
	client := pkiVaultService.NewPkiVaultClient(conn)

	users, err := client.GetUsers(ctx, &pkiVaultService.UserRequest{})
	if err != nil {
		logger.Err("failed to get user: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		writeStandardResponse(r, w, grpcToken, "failed to get users")
		return
	}

	userList := []models.UserInList{}
	for _, user := range users.GetUsers() {
		logger.Info("User: %v", user)
		userList = append(userList, models.UserInList{
			Id:       user.GetId(),
			Username: user.GetUsername(),
			Email:    user.GetEmail(),
			Role:     user.GetRole(),
		})
	}

	response, err := json.Marshal(&models.UserList{
		Users: userList,
		Total: users.GetTotal(),
	})
	if err != nil {
		logger.Err("failed to marshal response: %s", err)

		w.WriteHeader(http.StatusInternalServerError)
		writeStandardResponse(r, w, grpcToken, "failed to marshal response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
