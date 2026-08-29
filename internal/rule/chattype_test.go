package rule_test

import (
	"testing"

	"github.com/mefoxtrot/gramnfbot/internal/rule"
)

func TestIsDirect(t *testing.T) {
	tests := []struct {
		name     string
		chatType string
		want     bool
	}{
		{
			name:     "private",
			chatType: "private",
			want:     true,
		},
		{
			name:     "group",
			chatType: "group",
			want:     false,
		},
		{
			name:     "supergroup",
			chatType: "supergroup",
			want:     false,
		},
		{
			name:     "channel",
			chatType: "channel",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule.IsDirect(tt.chatType)

			if got != tt.want {
				t.Errorf(
					"IsDirect(%q) = %v; want %v",
					tt.chatType,
					got,
					tt.want,
				)
			}
		})
	}
}
