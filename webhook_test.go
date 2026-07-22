// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cadenya_test

import (
	"go.cadenya.com/cadenya-go"
	"go.cadenya.com/cadenya-go/option"
	"net/http"
	"strconv"
	"testing"
	"time"

	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := cadenya.NewClient(
		option.WithWebhookKey("whsec_c2VjcmV0Cg=="),
		option.WithAPIKey("My API Key"),
	)
	payload := []byte(`{"data":{"agent":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","workspaceId":"workspace_01HXKD2E5NQM3T9AYWCF133E3Q","externalId":"externalId","labels":{"foo":"string"},"updatedAt":"2019-12-27T18:11:19.117Z"},"agentVariation":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","createdAt":"2019-12-27T18:11:19.117Z","name":"name","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","workspaceId":"workspace_01HXKD2E5NQM3T9AYWCF133E3Q","externalId":"externalId","labels":{"foo":"string"},"updatedAt":"2019-12-27T18:11:19.117Z"},"objective":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","workspaceId":"workspace_01HXKD2E5NQM3T9AYWCF133E3Q","externalId":"externalId","labels":{"foo":"string"}},"objectiveEvent":{"data":{"type":"userMessage","userMessage":{"content":"content"}},"metadata":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","workspaceId":"workspace_01HXKD2E5NQM3T9AYWCF133E3Q","externalId":"externalId","labels":{"foo":"string"}},"contextWindowId":"objwin_01HXKD2E5NQM3T9AYWCFN7BSTR","duration":"-160513s","info":{"createdBy":{"metadata":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","name":"name","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","createdAt":"2019-12-27T18:11:19.117Z","externalId":"externalId","labels":{"foo":"string"}},"spec":{"type":"PROFILE_TYPE_UNSPECIFIED","email":"email","name":"name"}},"objective":{"id":"id","accountId":"account_01HXKD2E5NQM3T9AYWCFTJHJVF","createdAt":"2019-12-27T18:11:19.117Z","profileId":"profile_01HXKD2E5NQM3T9AYWCFS0AP08","workspaceId":"workspace_01HXKD2E5NQM3T9AYWCF133E3Q","externalId":"externalId","labels":{"foo":"string"}}}}},"timestamp":"2019-12-27T18:11:19.117Z","type":"type"}`)
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
