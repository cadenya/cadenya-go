// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/cadenya/cadenya-go"
	"github.com/cadenya/cadenya-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := cadenya.NewClient(
		option.WithWebhookKey("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
	)
	payload := []byte(`{"data":{"agent":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"agentVariation":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"objective":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"objectiveEvent":{"data":{"assistantMessage":{"content":"content","toolCalls":[{"arguments":"arguments","functionName":"functionName","tool":{"agent":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"cadenyaProvidedTool":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"tool":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}}}}]},"contextWindowCompacted":{"messagesCompacted":0,"newContextWindow":{"completionTokens":0,"objectiveId":"objectiveId","previousWindowContinueInstructions":"previousWindowContinueInstructions","promptTokens":0,"sequence":0},"strategies":["string"],"summary":"summary"},"error":{"message":"message","type":"type"},"memoryRead":{"memoryEntryId":"memoryEntryId","memoryLayerId":"memoryLayerId","message":"message"},"subObjectiveCreated":{"metadata":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}}},"toolApprovalRequested":{"toolCallId":"toolCallId"},"toolApproved":{"toolCallId":"toolCallId"},"toolCalled":{"toolCallId":"toolCallId"},"toolDenied":{"memo":"memo","toolCallId":"toolCallId"},"toolError":{"message":"message","toolCallId":"toolCallId"},"toolResult":{"content":"content","toolCallId":"toolCallId"},"type":"type","userMessage":{"content":"content"}},"metadata":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}},"contextWindowId":"contextWindowId","info":{"createdBy":{"metadata":{"id":"id","accountId":"accountId","name":"name","profileId":"profileId","externalId":"externalId","labels":{"foo":"string"}},"spec":{"type":"PROFILE_TYPE_USER","email":"email","name":"name"}},"objective":{"id":"id","accountId":"accountId","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profileId","workspaceId":"workspaceId","externalId":"externalId","labels":{"foo":"string"}}}}},"timestamp":"2019-12-27T18:11:19.117Z","type":"type"}`)
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Fatal("Failed to sign test webhook message", err)
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Fatal("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Fatal("Failed to unwrap webhook:", err)
	}
}
