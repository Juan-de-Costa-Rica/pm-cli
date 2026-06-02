package cli

import (
	"testing"

	"github.com/bscott/pm-cli/internal/imap"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{"single", "uid", []string{"uid"}, false},
		{"multiple", "uid,subject,from_address", []string{"uid", "subject", "from_address"}, false},
		{"whitespace trimmed", " uid , subject ", []string{"uid", "subject"}, false},
		{"empty segments skipped", "uid,,subject,", []string{"uid", "subject"}, false},
		{"unknown field", "uid,bogus", nil, true},
		{"empty string", "", nil, true},
		{"only commas", ",,,", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFields(tt.input)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error for %q, got none", tt.input)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != len(tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("index %d: got %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestParseFieldsAllKnown(t *testing.T) {
	// Every field in the canonical order must have an accessor.
	for _, f := range listFieldOrder {
		if _, ok := fieldAccessors[f]; !ok {
			t.Errorf("field %q listed in order but has no accessor", f)
		}
	}
}

func TestFilterMessageFields(t *testing.T) {
	messages := []imap.MessageSummary{
		{UID: 1, Subject: "Hello", From: "Alice", FromAddress: "alice@example.com"},
		{UID: 2, Subject: "World", From: "Bob", FromAddress: "bob@example.com"},
	}

	rows := filterMessageFields(messages, []string{"uid", "from_address"})
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2", len(rows))
	}

	if rows[0]["uid"] != uint32(1) {
		t.Errorf("row 0 uid: got %v, want 1", rows[0]["uid"])
	}
	if rows[0]["from_address"] != "alice@example.com" {
		t.Errorf("row 0 from_address: got %v", rows[0]["from_address"])
	}
	// Unselected fields must be absent, not zero-valued.
	if _, ok := rows[0]["subject"]; ok {
		t.Errorf("row 0 should not contain unselected field 'subject'")
	}
	if len(rows[0]) != 2 {
		t.Errorf("row 0 has %d keys, want 2", len(rows[0]))
	}
}
