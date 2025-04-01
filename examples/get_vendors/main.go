package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
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

CONDUCTOR_SECRET=YOUR_SECRET END_USER_ID=YOUR_END_USER_ID go run main.go

This will create an end user and generate a session auth url for the created user.
*/
func main() {
	conductorSecret := os.Getenv("CONDUCTOR_SECRET")
	endUserId := os.Getenv("END_USER_ID")
	authCtx := context.WithValue(context.Background(), conductor.ContextAccessToken, conductorSecret)
	configuration := conductor.NewConfiguration()
	apiClient := conductor.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.
		QuickbooksDesktopVendorsGet(authCtx).
		ConductorEndUserId(endUserId).
		Execute()
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Full HTTP response: %+v\n", r)
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.QuickbooksDesktopVendorsGet``: %+v\n", err)
		respBody, _ := io.ReadAll(r.Body)
		if r.StatusCode == 200 {
			var data map[string]any
			err = json.Unmarshal(respBody, &data)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error unmarshalling response: %+v\n", err)
			} else {
				fmt.Printf("Response body: %+v\n", data)
			}
		}
		_ = os.WriteFile("./resp.json", respBody, 0644)
	}
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.QuickbooksDesktopVendorsGet`: %+v\n", resp)
}
