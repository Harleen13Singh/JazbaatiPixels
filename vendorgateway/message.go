package vendorgateway

import (
	"github.com/3JoB/ulib/litefmt"
	"github.com/sugawarayuuta/sonnet"
)

type VendorRequest struct {
	Message *Message
	// session id is useful to store the context of the chat with the bot
	// it will be unique for every user and will be stored in the system
	SessionId string
	// url of the vendor api where request has to be made
	Url string
	// todo(Harleen Singh): move this behind an interface to accomodate multiple bots
	BotSpecificData *BotSpecificData
}

type Message struct {
	// data returned from the model
	// ***************NOTE***********
	// Do not update the Assistant value
	Assistant string
	// data input by the user
	User string
}

type BotSpecificData struct {
	Prompt        string   `json:"prompt"`                   // (required) The prompt you want Claude to complete. For proper response generation you will most likely want to format your prompt as follows:See [our comments on prompts](https://console.anthropic.com/docs/prompt-design#what-is-a-prompt) for more context.
	Model         string   `json:"model"`                    // (required) As we improve Claude, we develop new versions of it that you can query. This controls which version of Claude answers your request
	StopSequences []string `json:"stop_sequences,omitempty"` // (optional) A list of strings upon which to stop generating. You probably want , as that's the cue for the next turn in the dialog agent. Our client libraries provide a constant for this value (see examples below["\n\nHuman:"])
	Stream        bool     `json:"stream"`                   // (optional) Amount of randomness injected into the response. Ranges from 0 to 1. Use temp closer to 0 for analytical / multiple choice, and temp closer to 1 for creative and generative tasks.
	Temperature   float64  `json:"temperature,omitempty"`    // (required) Amount of randomness injected into the response. Defaults to 1. Ranges from 0 to 1. Use temp closer to 0 for analytical / multiple choice, and closer to 1 for creative and generative tasks.
	TopK          float64  `json:"top_k,omitempty"`          // (optional) Only sample from the top K options for each subsequent token. Used to remove "long tail" low probability responses. Defaults to -1, which disables it.
	TopP          uint     `json:"top_p,omitempty"`          // (optional) Does nucleus sampling, in which we compute the cumulative distribution over all the options for each subsequent token in decreasing probability order and cut it off once it reaches a particular probability specified by . Defaults to -1, which disables it. Note that you should either alter or , but not both.`top_ptemperaturetop_p``
	MaxToken      uint     `json:"max_tokens_to_sample"`     // (required) A maximum number of tokens to generate before stopping.
}

type VendorResponse struct {
	SessionId  string
	Completion string `json:"completion"`          // The resulting completion up to and excluding the stop sequences.
	StopReason string `json:"stop_reason"`         // The reason we stopped sampling, either if we reached one of your provided , or if we exceeded `.stop_sequencestop_sequencesmax_tokensmax_tokens_to_sample`
	Stop       string `json:"stop"`                // If the is , this contains the actual stop sequence (of the list passed-in) that was `seenstop_reasonstop_sequencestop_sequences`
	LogID      string `json:"log_id"`              // The ID of the log that generated the response
	Exception  string `json:"exception,omitempty"` // exception
	Model      string `json:"model"`               // Model
	Truncated  bool   `json:"truncated"`           // truncated
}

func (r *VendorResponse) ToString() string {
	d, _ := sonnet.Marshal(r)
	return string(d)
}

func (s *BotSpecificData) Set(m *Message) error {
	if m.User == "" {
		return ErrEmptyUserInput
	}
	if m.Assistant == "" {
		s.Prompt = litefmt.Sprint("\n\nHuman: ", m.User, "\n\nAssistant:")
	} else {
		s.Prompt = litefmt.Sprint(m.User, m.Assistant)
	}
	return nil
}

func (s *BotSpecificData) Build(next string, m *Message) error {
	if m.User == "" {
		return ErrEmptyUserInput
	}
	if m.Assistant == "" {
		s.Prompt = litefmt.Sprint(next, "\n\nHuman: ", m.User, "\n\nAssistant:")
	} else {
		s.Prompt = litefmt.Sprint(next, "\n\nHuman: ", m.User, "\n\nAssistant:", m.Assistant)
	}
	return nil
}
