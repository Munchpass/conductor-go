package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/Munchpass/conductor-go"
)

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes), nil
}

/*
Run by:

CONDUCTOR_SECRET=YOUR_SECRET CONDUCTOR_PUBLISHABLE_KEY=YOUR_PUBLISHABLE_KEY go run main.go

This will create an end user and generate a session auth url for the created user.
*/
func main() {
	conductorSecret := os.Getenv("CONDUCTOR_SECRET")
	conductorPublishableKey := os.Getenv("CONDUCTOR_PUBLISHABLE_KEY")
	authCtx := context.WithValue(context.Background(), conductor.ContextAccessToken, conductorSecret)
	sourceId, _ := generateRandomString(32)
	endUsersPostRequest := *conductor.NewEndUsersPostRequest("Acme Inc.", sourceId, "alice@acme.com") // EndUsersPostRequest |  (optional)

	configuration := conductor.NewConfiguration()
	apiClient := conductor.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EndUsersPost(authCtx).EndUsersPostRequest(endUsersPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EndUsersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EndUsersPost`: EndUser
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EndUsersPost`: %v\n", resp)

	authSessReq := *conductor.NewAuthSessionsPostRequest(conductorPublishableKey, resp.Id)
	authResp, authR, authErr := apiClient.DefaultAPI.AuthSessionsPost(authCtx).AuthSessionsPostRequest(authSessReq).Execute()
	if authErr != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AuthSessionsPost`: %v\n", authErr)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", authR)
	}
	// response from `AuthSessionsPost`: AuthSession
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AuthSessionsPost`: %v\n", authResp)
	fmt.Printf("Auth URL: %s\n", authResp.AuthFlowUrl)
}
