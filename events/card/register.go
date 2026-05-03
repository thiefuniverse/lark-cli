// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

// Package card registers CardKit callback EventKeys.
package card

import (
	"reflect"

	"github.com/larksuite/cli/internal/event"
)

// Keys returns all card callback EventKey definitions.
func Keys() []event.KeyDefinition {
	return []event.KeyDefinition{
		{
			Key:         "card.action.trigger",
			DisplayName: "Card action trigger",
			Description: "Triggered when a user interacts with an interactive card component",
			EventType:   "card.action.trigger",
			Schema: event.SchemaDef{
				Custom: &event.SchemaSpec{Type: reflect.TypeOf(CardActionTriggerOutput{})},
			},
			AuthTypes:             []string{"bot"},
			RequiredConsoleEvents: []string{"card.action.trigger"},
		},
	}
}

// CardActionTriggerOutput is the raw V2 callback envelope emitted by `card.action.trigger`.
type CardActionTriggerOutput struct {
	Schema string                  `json:"schema,omitempty" desc:"Feishu event protocol version" enum:"2.0"`
	Header CardActionTriggerHeader `json:"header,omitempty" desc:"Callback header"`
	Event  CardActionTriggerEvent  `json:"event,omitempty"  desc:"Card action callback payload"`
}

type CardActionTriggerHeader struct {
	AppID      string `json:"app_id,omitempty"      desc:"App ID receiving the callback"`
	CreateTime string `json:"create_time,omitempty" desc:"Callback creation time, in millisecond timestamp string" kind:"timestamp_ms"`
	EventID    string `json:"event_id,omitempty"    desc:"Globally unique callback ID; safe for deduplication"`
	EventType  string `json:"event_type,omitempty"  desc:"Callback type; always card.action.trigger" enum:"card.action.trigger"`
	TenantKey  string `json:"tenant_key,omitempty"  desc:"Tenant key"`
	Token      string `json:"token,omitempty"       desc:"Verification token from the app callback configuration"`
}

type CardActionTriggerEvent struct {
	Operator     CardActionOperator `json:"operator,omitempty"      desc:"User who triggered the card action"`
	Token        string             `json:"token,omitempty"         desc:"Card update token from the callback; use it for delayed card updates"`
	Action       CardAction         `json:"action,omitempty"        desc:"Interactive component action details"`
	Host         string             `json:"host,omitempty"          desc:"Card host, such as im_message or im_top_notice"`
	DeliveryType string             `json:"delivery_type,omitempty" desc:"Card delivery channel"`
	Context      CardActionContext  `json:"context,omitempty"       desc:"Conversation and message context for the card"`
}

type CardActionOperator struct {
	TenantKey string `json:"tenant_key,omitempty" desc:"Operator tenant key"`
	UserID    string `json:"user_id,omitempty"    desc:"Operator user ID, if available" kind:"user_id"`
	OpenID    string `json:"open_id,omitempty"    desc:"Operator open ID" kind:"open_id"`
	UnionID   string `json:"union_id,omitempty"   desc:"Operator union ID, if available" kind:"union_id"`
}

type CardAction struct {
	Value      map[string]interface{} `json:"value,omitempty"       desc:"Custom callback payload from the card component value or callback behavior"`
	Tag        string                 `json:"tag,omitempty"         desc:"Component tag that triggered the action, such as button"`
	Option     string                 `json:"option,omitempty"      desc:"Selected option for select components"`
	Timezone   string                 `json:"timezone,omitempty"    desc:"User timezone from date/time components"`
	Name       string                 `json:"name,omitempty"        desc:"Component name"`
	FormValue  map[string]interface{} `json:"form_value,omitempty" desc:"Submitted form field values"`
	InputValue string                 `json:"input_value,omitempty" desc:"Input component value"`
	Options    []string               `json:"options,omitempty"     desc:"Selected options for multi-select components"`
	Checked    bool                   `json:"checked,omitempty"     desc:"Checkbox or switch state"`
}

type CardActionContext struct {
	URL           string `json:"url,omitempty"             desc:"URL context for URL preview callbacks"`
	PreviewToken  string `json:"preview_token,omitempty"   desc:"URL preview token, when present"`
	OpenMessageID string `json:"open_message_id,omitempty" desc:"Message ID of the interactive card" kind:"message_id"`
	OpenChatID    string `json:"open_chat_id,omitempty"    desc:"Chat ID containing the interactive card" kind:"chat_id"`
}
