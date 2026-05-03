// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package source

import (
	"context"
	"testing"

	larkevent "github.com/larksuite/oapi-sdk-go/v3/event"
	"github.com/larksuite/oapi-sdk-go/v3/event/dispatcher/callback"

	eventlib "github.com/larksuite/cli/internal/event"
)

func TestCardActionTriggerHandlerEmitsRawEventAndAcks(t *testing.T) {
	payload := []byte(`{
		"schema": "2.0",
		"header": {
			"event_id": "evt_card_1",
			"event_type": "card.action.trigger",
			"create_time": "1776409469274",
			"app_id": "cli_test"
		},
		"event": {
			"operator": {"open_id": "ou_user"},
			"token": "c-card-token",
			"action": {
				"tag": "button",
				"value": {"action": "continue", "task_id": "task_1"}
			},
			"context": {
				"open_message_id": "om_card",
				"open_chat_id": "oc_chat"
			}
		}
	}`)

	var captured *eventlib.RawEvent
	s := &FeishuSource{}
	handler := s.buildCardActionTriggerHandler(func(raw *eventlib.RawEvent) {
		captured = raw
	})

	resp, err := handler(context.Background(), &callback.CardActionTriggerEvent{
		EventReq: &larkevent.EventReq{Body: payload},
	})
	if err != nil {
		t.Fatalf("handler returned error: %v", err)
	}
	if resp == nil {
		t.Fatal("handler must acknowledge card callbacks with an empty response")
	}
	if captured == nil {
		t.Fatal("expected raw event to be emitted")
	}
	if captured.EventID != "evt_card_1" {
		t.Errorf("EventID = %q, want evt_card_1", captured.EventID)
	}
	if captured.EventType != "card.action.trigger" {
		t.Errorf("EventType = %q, want card.action.trigger", captured.EventType)
	}
	if captured.SourceTime != "1776409469274" {
		t.Errorf("SourceTime = %q, want 1776409469274", captured.SourceTime)
	}
	if string(captured.Payload) != string(payload) {
		t.Error("Payload should preserve the original callback body")
	}
}
