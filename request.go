package groq

// CompletionCreateParams struct to handle API parameters
type CompletionCreateParams struct {
	Messages         []Message      `json:"messages"`
	Model            string         `json:"model"`
	FrequencyPenalty float32        `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	Logprobs         bool           `json:"logprobs,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	N                int            `json:"n,omitempty"`
	PresencePenalty  float32        `json:"presence_penalty,omitempty"`
	ResponseFormat   ResponseFormat `json:"response_format,omitempty"`
	Seed             int            `json:"seed,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	Temperature      float32        `json:"temperature,omitempty"`
	ToolChoice       ToolChoice     `json:"tool_choice,omitempty"`
	Tools            []Tool         `json:"tools,omitempty"`
	TopLogprobs      int            `json:"top_logprobs,omitempty"`
	TopP             float32        `json:"top_p,omitempty"`
	User             string         `json:"user,omitempty"`
}

// MessageToolCallFunction struct to handle function details
type MessageToolCallFunction struct {
	Arguments string `json:"arguments,omitempty"`
	Name      string `json:"name,omitempty"`
}

// MessageToolCall struct to handle tool calls in messages
type MessageToolCall struct {
	ID       string                  `json:"id,omitempty"`
	Function MessageToolCallFunction `json:"function,omitempty"`
	Type     string                  `json:"type,omitempty"`
}

// Message struct to handle messages
type Message struct {
	Content    string            `json:"content"` // Required fields, not omitting in JSON
	Role       string            `json:"role"`    // Required fields, not omitting in JSON
	Name       string            `json:"name,omitempty"`
	ToolCallID string            `json:"tool_call_id,omitempty"`
	ToolCalls  []MessageToolCall `json:"tool_calls,omitempty"`
}

// ResponseFormat struct to handle response formatting
type ResponseFormat struct {
	Type string `json:"type,omitempty"`
}

// ToolChoiceToolChoiceFunction struct to handle tool choice functions
type ToolChoiceToolChoiceFunction struct {
	Name string `json:"name,omitempty"`
}

// ToolChoiceToolChoice struct to handle nested tool choices
type ToolChoiceToolChoice struct {
	Function ToolChoiceToolChoiceFunction `json:"function,omitempty"`
	Type     string                       `json:"type,omitempty"`
}

type ToolChoice string

const (
	ToolChoiceAuto ToolChoice = "auto"
	ToolChoiceNone ToolChoice = "none"
)

// ToolFunction struct to handle tool functions
type ToolFunction struct {
	Description string                 `json:"description,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
}

// Tool struct to handle tools
type Tool struct {
	Function ToolFunction `json:"function,omitempty"`
	Type     string       `json:"type,omitempty"`
}
