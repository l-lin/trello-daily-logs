package printer

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/l-lin/trello-daily-logs/trello"
)

func TestMarkdownPrinter_Print(t *testing.T) {
	var tests = map[string]struct {
		given    []trello.Card
		expected string
	}{
		"multiple cards": {
			given: []trello.Card{
				trello.Card{
					Name: "career@interview: find good questions for interview",
					Labels: []trello.Label{
						trello.Label{Name: "PERSO"},
					},
				},
				trello.Card{
					Name: "api@gravitee: install in prod environment",
					Labels: []trello.Label{
						trello.Label{Name: "WORK"},
					},
				},
				trello.Card{
					Name: "api@spec: write OpenAPI specifications",
					Labels: []trello.Label{
						trello.Label{Name: "WORK"},
						trello.Label{Name: "ABANDONED"},
					},
				},
			},
			expected: fmt.Sprintf(`## %s %02d

- ABANDONED
  - api@spec: write OpenAPI specifications
- PERSO
  - career@interview: find good questions for interview
- WORK
  - api@gravitee: install in prod environment
  - api@spec: write OpenAPI specifications

`, time.Now().Weekday(), time.Now().Day()),
		},
		"one cards": {
			given: []trello.Card{
				trello.Card{
					Name: "career@interview: find good questions for interview",
					Labels: []trello.Label{
						trello.Label{Name: "PERSO"},
					},
				},
			},
			expected: fmt.Sprintf(`## %s %02d

- PERSO
  - career@interview: find good questions for interview

`, time.Now().Weekday(), time.Now().Day()),
		},
		"no card": {
			given:    []trello.Card{},
			expected: "",
		},
	}
	p := MarkdownPrinter{}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := bytes.NewBufferString("")
			p.Print(actual, tt.given)
			if actual.String() != tt.expected {
				t.Errorf("expected %v, actual %v", tt.expected, actual)
			}
		})
	}
}
