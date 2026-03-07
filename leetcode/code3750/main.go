package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	token, err := getAccessToken()
	if err != nil {
		fmt.Println("here", err)
		log.Fatal(err)
	}

	emails, err := fetchUnreadEmails(token)
	if err != nil {
		fmt.Println("here", err)
		log.Fatal(err)
	}

	for _, m := range emails {
		log.Println("---- EMAIL ----")
		log.Println("Subject:", m.Subject)
		log.Println("Received:", m.ReceivedDateTime)
		log.Println("ID", m.ID)
		// log.Println("Body:", truncate(m.Body.Content, 300))
	}
	fmt.Println("len(),", len(emails))
}

func getAccessToken() (string, error) {
	cfg := clientcredentials.Config{
		ClientID:     "6e61609c-0d09-4388-a36a-8d1f46614066",
		ClientSecret: "REMOVED_AZURE_SECRET",
		TokenURL: fmt.Sprintf(
			"https://login.microsoftonline.com/%s/oauth2/v2.0/token",
			"bc9fa0fc-6ced-4a94-87bd-a41ab913a1be",
		),
		Scopes: []string{"https://graph.microsoft.com/.default"},
	}

	token, err := cfg.Token(context.Background())
	if err != nil {
		fmt.Println("err", err)
		return "", err
	}

	return token.AccessToken, nil
}

type GraphResponse struct {
	Value []Message `json:"value"`
}

type Message struct {
	ID               string `json:"id"`
	Subject          string `json:"subject"`
	ReceivedDateTime string `json:"receivedDateTime"`
	Body             Body   `json:"body"`
}

type Body struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

func fetchUnreadEmails(token string) ([]Message, error) {
	baseURL := fmt.Sprintf(
		"https://graph.microsoft.com/v1.0/users/%s/mailFolders/inbox/messages",
		"notifications@easyto.travel",
	)

	q := url.Values{}
	q.Set("$filter", "isRead eq false")
	q.Set("$orderby", "receivedDateTime desc")
	q.Set("$top", "10")

	req, _ := http.NewRequest("GET", baseURL+"?"+q.Encode(), nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("graph error: %s", b)
	}

	var gr GraphResponse
	if err := json.NewDecoder(resp.Body).Decode(&gr); err != nil {
		return nil, err
	}

	return gr.Value, nil
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max] + "..."
}

func minimumFlips(n int) int {
	nBinary := strconv.FormatInt(int64(n), 2)
	reversed := string(reverseSlice([]rune(nBinary)))

	var count int
	for i := 0; i < len(nBinary); i++ {
		if nBinary[i] != reversed[i] {
			count++
		}
	}

	return count
}

func reverseSlice(slice []rune) []rune {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
