package groq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var doneEvent = []byte("[DONE]")

type Client struct {
	apiKey  string
	baseUrl string
	debug   bool
}

type ClientOption func(*Client)

func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

func WithDebug(debug bool) ClientOption {
	return func(c *Client) {
		c.debug = debug
	}
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{
		apiKey:  os.Getenv("GROQ_API_KEY"),
		baseUrl: "https://api.groq.com",
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

func (c *Client) CreateChatCompletion(params CompletionCreateParams) (*ChatCompletion, error) {
	if !params.Stream {
		return c.syncChatCompletion(params)
	}

	return c.streamChatCompletion(params)
}

func (c *Client) syncChatCompletion(params CompletionCreateParams) (*ChatCompletion, error) {
	req, err := c.newRequest("POST", "/openai/v1/chat/completions", params)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errResp ErrorResponse

		if c.debug {
			for k, v := range resp.Header {
				fmt.Printf("%s: %s\n", k, v)
			}
		}

		err = json.Unmarshal(body, &errResp)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
		}

		return nil, errResp.Error
	}

	var result ChatCompletion

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Client) streamChatCompletion(params CompletionCreateParams) (*ChatCompletion, error) {
	req, err := c.newRequest("POST", "/openai/v1/chat/completions", params)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// TODO: manage channel size
	result := ChatCompletion{Stream: make(chan *ChatChunkCompletion, 100)}

	go c.startStream(resp, &result)

	return &result, nil
}

func (c *Client) startStream(resp *http.Response, result *ChatCompletion) {
	reader := NewStreamReader(resp.Body)
	for {
		event, err := reader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
			return
		}

		if c.debug {
			log.Println(string(event.Data))
		}

		if bytes.HasPrefix(event.Data, doneEvent) {
			break
		}

		var delta ChatChunkCompletion
		err = json.Unmarshal(event.Data, &delta)
		if err != nil {
			log.Println(err)
			continue
		}

		result.Stream <- &delta
	}

	// Close the channel to signal that no more data will be sent
	close(result.Stream)
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseUrl+path, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
		reqBody, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		if c.debug {
			log.Println(string(reqBody))
		}

		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
	}

	return req, nil
}
