// Copyright (c) 2026 Lark Technologies Pte. Ltd.
// SPDX-License-Identifier: MIT

package card

import "testing"

func TestKeysIncludesCardActionTrigger(t *testing.T) {
	keys := Keys()
	if len(keys) != 1 {
		t.Fatalf("Keys() len = %d, want 1", len(keys))
	}
	def := keys[0]
	if def.Key != "card.action.trigger" {
		t.Fatalf("Key = %q, want card.action.trigger", def.Key)
	}
	if def.EventType != "card.action.trigger" {
		t.Errorf("EventType = %q, want card.action.trigger", def.EventType)
	}
	if def.Schema.Custom == nil {
		t.Fatal("Schema.Custom must be set for card action callbacks")
	}
	if def.Schema.Native != nil {
		t.Fatal("Schema.Native must not be set for card action callbacks")
	}
	if def.Process != nil {
		t.Fatal("card.action.trigger should stream the raw callback envelope without Process")
	}
	if len(def.AuthTypes) != 1 || def.AuthTypes[0] != "bot" {
		t.Fatalf("AuthTypes = %v, want [bot]", def.AuthTypes)
	}
	if len(def.RequiredConsoleEvents) != 1 || def.RequiredConsoleEvents[0] != "card.action.trigger" {
		t.Fatalf("RequiredConsoleEvents = %v, want [card.action.trigger]", def.RequiredConsoleEvents)
	}
}
