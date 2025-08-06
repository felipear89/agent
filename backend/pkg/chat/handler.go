package chat

import (
	"errors"
	"fmt"
	"github.com/felipear89/agent/pkg/server/apperror"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
}

func NewHandler(api *gin.RouterGroup) *Handler {
	h := &Handler{}
	h.RegisterRoutes(api)
	return h
}

func (h *Handler) RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/chat")
	{
		group.POST("", h.Completation)
	}
}

func (h *Handler) Completation(ctx *gin.Context) {
	//authorization := "eyJhbGciOiJFUzI1NiIsImtpZCI6Imdlbi1haV8xNzIyMzY1MDAwIiwidHlwIjoiSldUIn0.eyJhdWQiOiI0NmFhMjVkYi03MmJmLTExZjAtYjA3NC00ZTAxM2UyZGRkZTQiLCJleHAiOjE3NTQ1MDQ1NjEsImp0aSI6IjU2NWU1Yjc1LTQ5NjgtNDRiMi1hYTFmLTZiNjNmMTI2MWU5NSIsImlhdCI6MTc1NDUwNDI2MSwiaXNzIjoiZ2VuLWFpLWFwaSIsIm5iZiI6MTc1NDUwNDI2MSwic3ViIjoiMTc5LjIwOS4xNDAuMTgwIiwidGVhbV9pZCI6MjQ0NjQ5ODQsImxlYWZfYWdlbnQiOnsidXVpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCJ9LCJtb2RlbCI6eyJ1dWlkIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIn19.mki_fvLybTfufmWrGA9m5_bGtTwRAuC8VqfM9sXifvToJnnSR3COea8cRdyrg3SFtS_clPi0amZaVeETpkDnmw"
	// POST https://ot7q6l7i6k57sxp2ikoi72vx.agents.do-ai.run/api/v1/chat/completions with content-type text/event-stream;

	//ch := make(chan string)

}

type EventStreamRequest struct {
	Message string `form:"message" json:"message" binding:"required,max=100"`
}

func HandleEventStreamPost(c *gin.Context, ch chan string) {
	var request EventStreamRequest
	if err := c.ShouldBind(&request); err != nil {
		errorMessage := fmt.Sprintf("request validation error: %s", err.Error())
		apperror.BadRequestResponse(c, errors.New(errorMessage))

		return
	}

	ch <- request.Message

	//CreatedResponse(c, &request.Message)

	return
}

func request(token string) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		// handle error
	}

	// Add headers
	req.Header.Add("Authorization", token)
	req.Header.Add("Content-Type", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
}

type Payload struct {
	Messages            []interface{} `json:"messages"`
	Temperature         int           `json:"temperature"`
	TopP                int           `json:"top_p"`
	MaxTokens           int           `json:"max_tokens"`
	MaxCompletionTokens int           `json:"max_completion_tokens"`
	Stream              bool          `json:"stream"`
	K                   int           `json:"k"`
	RetrievalMethod     string        `json:"retrieval_method"`
	FrequencyPenalty    int           `json:"frequency_penalty"`
	PresencePenalty     int           `json:"presence_penalty"`
	Stop                string        `json:"stop"`
	StreamOptions       struct {
		IncludeUsage bool `json:"include_usage"`
	} `json:"stream_options"`
	KbFilters []struct {
		Index string `json:"index"`
		Path  string `json:"path,omitempty"`
	} `json:"kb_filters"`
	FilterKbContentByQueryMetadata bool   `json:"filter_kb_content_by_query_metadata"`
	InstructionOverride            string `json:"instruction_override"`
	IncludeFunctionsInfo           bool   `json:"include_functions_info"`
	IncludeRetrievalInfo           bool   `json:"include_retrieval_info"`
	IncludeGuardrailsInfo          bool   `json:"include_guardrails_info"`
	ProvideCitations               bool   `json:"provide_citations"`
}

type Response struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Delta struct {
			Id       string `json:"id"`
			Role     string `json:"role"`
			Content  string `json:"content"`
			SentTime string `json:"sentTime"`
		} `json:"delta"`
		Index        int    `json:"index"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Guardrails struct {
		TriggeredGuardrails []struct {
			Message  string `json:"message"`
			RuleName string `json:"rule_name"`
		} `json:"triggered_guardrails"`
	} `json:"guardrails"`
	Functions struct {
		CalledFunctions []string `json:"called_functions"`
	} `json:"functions"`
	Retrieval struct {
		RetrievedData []struct {
			Id    string `json:"id"`
			Index string `json:"index"`
			Score int    `json:"score"`
		} `json:"retrieved_data"`
	} `json:"retrieval"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
