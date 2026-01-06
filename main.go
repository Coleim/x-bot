package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dghubble/oauth1"
)

// TweetRequest represents the request body for creating a tweet
type TweetRequest struct {
	Text string `json:"text"`
}

// TweetResponse represents the response from Twitter API v2
type TweetResponse struct {
	Data struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
}

// UserResponse represents the user information from Twitter API v2
type UserResponse struct {
	Data struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	} `json:"data"`
}

// UserTweetsResponse represents recent tweets from the user
type UserTweetsResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	} `json:"data"`
	Meta struct {
		ResultCount int `json:"result_count"`
	} `json:"meta"`
}

// OpenAI API structures
type OpenAIRequest struct {
	Model     string          `json:"model"`
	Messages  []OpenAIMessage `json:"messages"`
	MaxTokens int             `json:"max_tokens,omitempty"`
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OpenAIResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
}

func postTweet(client *http.Client, message string) (*TweetResponse, error) {
	tweetReq := TweetRequest{Text: message}
	body, err := json.Marshal(tweetReq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal tweet request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "TwitterBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var tweetResp TweetResponse
	if err := json.Unmarshal(respBody, &tweetResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &tweetResp, nil
}

func getAuthenticatedUser(client *http.Client) (*UserResponse, error) {
	req, err := http.NewRequest("GET", "https://api.twitter.com/2/users/me", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "TwitterBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var userResp UserResponse
	if err := json.Unmarshal(respBody, &userResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return &userResp, nil
}

func getRecentTweets(client *http.Client, userID string, count int) ([]string, error) {
	url := fmt.Sprintf("https://api.twitter.com/2/users/%s/tweets?max_results=%d", userID, count)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("User-Agent", "TwitterBot/1.0")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var tweetsResp UserTweetsResponse
	if err := json.Unmarshal(respBody, &tweetsResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	var recentTweets []string
	for _, tweet := range tweetsResp.Data {
		recentTweets = append(recentTweets, tweet.Text)
	}

	return recentTweets, nil
}

func generateTweetWithAI(apiKey, provider string, recentTweets []string) (string, error) {
	systemPrompt := `You are a witty and insightful Twitter bot for software engineers, developers, and game developers.

Your tweets should:
- Be engaging and spark conversation
- Mix humor with technical insights
- Stay under 280 characters
- Alternate between humorous and serious/thoughtful content
- Cover topics like: coding, debugging, dev life, tech trends, game dev, programming languages, DevOps, architecture, best practices

Tone variations:
- Humorous: Funny observations, memes in text form, relatable developer struggles
- Serious: Technical insights, career advice, industry trends, best practices

Generate ONE tweet only. No quotes around it. Just the tweet text.

IMPORTANT: The user's recent tweets will be provided. You MUST generate something completely different in topic, style, or perspective. Avoid repeating similar jokes, themes, or observations.`

	userPrompt := "Generate an engaging tweet for developers. Make it either humorous OR serious/insightful."

	// Add recent tweets context if available
	if len(recentTweets) > 0 {
		userPrompt += "\n\nMy recent tweets (DO NOT repeat similar content):\n"
		for i, tweet := range recentTweets {
			userPrompt += fmt.Sprintf("%d. %s\n", i+1, tweet)
		}
		userPrompt += "\nGenerate something COMPLETELY DIFFERENT from the above tweets."
	}

	// Determine API endpoint and model based on provider
	var apiURL, model string
	switch provider {
	case "groq":
		apiURL = "https://api.groq.com/openai/v1/chat/completions"
		model = "llama-3.3-70b-versatile"
	case "openai":
		apiURL = "https://api.openai.com/v1/chat/completions"
		model = "gpt-4o-mini"
	default:
		apiURL = "https://api.groq.com/openai/v1/chat/completions"
		model = "llama-3.3-70b-versatile"
	}

	reqBody := OpenAIRequest{
		Model: model,
		Messages: []OpenAIMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
		MaxTokens: 100,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal AI request: %v", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create AI request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send AI request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read AI response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("AI API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var aiResp OpenAIResponse
	if err := json.Unmarshal(respBody, &aiResp); err != nil {
		return "", fmt.Errorf("failed to unmarshal AI response: %v", err)
	}

	if len(aiResp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI")
	}

	tweet := strings.TrimSpace(aiResp.Choices[0].Message.Content)
	tweet = strings.Trim(tweet, "\"") // Remove quotes if AI added them

	return tweet, nil
}

func main() {
	log.Println("Starting Twitter Bot (API v2)...")

	// Get credentials from environment variables
	consumerKey := os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	accessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	accessSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		log.Fatal("Missing required environment variables. Please check KEY.MD for setup instructions.")
	}

	// Configure OAuth1
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// Verify credentials by getting user info
	user, err := getAuthenticatedUser(httpClient)
	if err != nil {
		log.Fatalf("Failed to verify credentials: %v", err)
	}

	log.Printf("✅ Authenticated as: @%s (%s)", user.Data.Username, user.Data.Name)

	// Get AI credentials
	aiAPIKey := os.Getenv("AI_API_KEY")
	aiProvider := os.Getenv("AI_PROVIDER")

	// Backward compatibility with OPENAI_API_KEY
	if aiAPIKey == "" {
		aiAPIKey = os.Getenv("OPENAI_API_KEY")
		if aiAPIKey != "" {
			aiProvider = "openai"
		}
	}

	if aiProvider == "" {
		aiProvider = "groq" // Default to free Groq
	}

	// Check if AI API key is provided
	if aiAPIKey == "" {
		log.Fatal("❌ AI_API_KEY is required. Get a FREE Groq API key at https://console.groq.com")
	}

	// Fetch recent tweets to avoid duplication
	log.Println("📥 Fetching recent tweets to ensure uniqueness...")
	recentTweets, err := getRecentTweets(httpClient, user.Data.ID, 10)
	if err != nil {
		log.Printf("⚠️  Could not fetch recent tweets: %v. Will generate without context.", err)
		recentTweets = []string{}
	} else {
		log.Printf("✅ Found %d recent tweets to avoid duplicating", len(recentTweets))
	}

	// Generate tweet with AI
	log.Printf("🤖 Generating tweet with AI (%s)...", aiProvider)
	message, err := generateTweetWithAI(aiAPIKey, aiProvider, recentTweets)
	if err != nil {
		log.Fatalf("❌ AI generation failed: %v. No tweet will be posted.", err)
	}

	log.Printf("✨ AI generated: %s", message)

	tweetResp, err := postTweet(httpClient, message)
	if err != nil {
		log.Fatalf("Failed to post tweet: %v", err)
	}

	log.Printf("✅ Tweet posted successfully!")
	log.Printf("Tweet ID: %s", tweetResp.Data.ID)
	log.Printf("Tweet URL: https://twitter.com/%s/status/%s", user.Data.Username, tweetResp.Data.ID)
	log.Printf("Message: %s", tweetResp.Data.Text)

	log.Println("✅ Bot completed successfully!")
}
