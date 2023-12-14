package frontend

import (
	"context"
	"fmt"

	"github.com/3JoB/anthropic-sdk-go/v2"
	"github.com/3JoB/anthropic-sdk-go/v2/data"
	"github.com/3JoB/anthropic-sdk-go/v2/resp"
)

type service struct{}

func (s *service) CallClaude(ctx context.Context, req interface{}) (interface{}, error) {
	claudeClient, claudeClientErr := anthropic.New(&anthropic.Config{
		Key:          "generate the keys",
		DefaultModel: data.ModelFullClaude,
	})
	if claudeClientErr != nil {
		return nil, fmt.Errorf("failed to create client with claude,err: %w", claudeClientErr)
	}

	botResp, botRespErr := claudeClient.Send(&anthropic.Sender{
		Message: data.MessageModule{
			Human: "take this as an input of the rpc",
		},
		// todo(Harleen Singh): read the doc and populate accordingaly
		Sender: &resp.Sender{},
	})
	if botRespErr != nil {
		return nil, fmt.Errorf("failed to get the response from the bot,err: %w", botRespErr)
	}

	return botResp.Response.String(), nil
}
