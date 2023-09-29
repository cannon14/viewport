package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/machinebox/graphql"
	"github.com/revel/revel"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type accessToken struct {
	Token string `json:"access_token"`
}

func CallGraphql[T any](graphqlRequestString string, variableJson string) T {
	clientId := revel.Config.StringDefault("wiz.client.id", "")
	clientSecret := revel.Config.StringDefault("wiz.client.secret", "")

	authData := url.Values{}
	authData.Set("grant_type", "client_credentials")
	authData.Set("audience", "wiz-api")
	authData.Set("client_id", clientId)
	authData.Set("client_secret", clientSecret)

	client := &http.Client{}

	r, err := http.NewRequest(http.MethodPost, "https://auth.app.wiz.io/oauth/token", strings.NewReader(authData.Encode()))
	r.Header.Add("Encoding", "UTF-8")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	fmt.Println("Getting token.")
	resp, err := client.Do(r)
	fmt.Println("Authentication response: " + resp.Status)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error authenticating.")
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body: %w", resp.StatusCode)
	}

	at := accessToken{}
	jsonErr := json.Unmarshal(bodyBytes, &at)

	if jsonErr != nil {
		log.Fatalf("Failed parsing JSON body: %w", jsonErr)
	}

	graphqlClient := graphql.NewClient("https://api.us7.app.wiz.io/graphql")

	graphqlRequest := graphql.NewRequest(graphqlRequestString)

	var variables map[string]interface{}

	if err := json.Unmarshal([]byte(variableJson), &variables); err != nil {
		log.Fatalf("Failed parsing JSON body: %w", err)
	}

	for k, v := range variables {
		graphqlRequest.Var(k, v)
	}

	graphqlRequest.Header.Set("Authorization", "Bearer "+at.Token)

	var respData T
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &respData); err != nil {
		log.Fatal(err)
	}

	return respData
}
