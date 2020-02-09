package printer

import (
	"bytes"
	"testing"
	"time"

	"github.com/l-lin/trello-daily-logs/trello"
)

func TestMarkdownPrinter_Print(t *testing.T) {
	ti, err := time.Parse(time.RFC3339, "2020-02-09T11:12:13Z")
	if err != nil {
		t.Fatal(err)
	}
	type given struct {
		doneCards []trello.Card
		todoCards []trello.Card
	}
	var tests = map[string]struct {
		given    given
		expected string
	}{
		"multiple cards": {
			given: given{
				doneCards: []trello.Card{
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
				todoCards: []trello.Card{
					trello.Card{
						Name: "projectA@taskA: study solutions",
						Labels: []trello.Label{
							trello.Label{Name: "WORK"},
						},
					},
					trello.Card{
						Name: "shopping: buy milk",
						Labels: []trello.Label{
							trello.Label{Name: "PERSO"},
						},
					},
					trello.Card{
						Name: "projectB@taskB: implement solution",
						Labels: []trello.Label{
							trello.Label{Name: "WORK"},
						},
					},
				},
			},
			expected: `## Sunday 09

**ABANDONED**

- api@spec: write OpenAPI specifications

**PERSO**

- career@interview: find good questions for interview

**WORK**

- api@gravitee: install in prod environment
- api@spec: write OpenAPI specifications

<details>
<summary>UNFINISHED</summary>

**PERSO**

- shopping: buy milk

**WORK**

- projectA@taskA: study solutions
- projectB@taskB: implement solution

</details>

`,
		},
		"one done card": {
			given: given{
				doneCards: []trello.Card{
					trello.Card{
						Name: "career@interview: find good questions for interview",
						Labels: []trello.Label{
							trello.Label{Name: "PERSO"},
						},
					},
				},
				todoCards: []trello.Card{},
			},
			expected: `## Sunday 09

**PERSO**

- career@interview: find good questions for interview

`,
		},
		"one todo card": {
			given: given{
				doneCards: []trello.Card{},
				todoCards: []trello.Card{
					trello.Card{
						Name: "career@interview: find good questions for interview",
						Labels: []trello.Label{
							trello.Label{Name: "PERSO"},
						},
					},
				},
			},
			expected: `## Sunday 09

<details>
<summary>UNFINISHED</summary>

**PERSO**

- career@interview: find good questions for interview

</details>

`,
		},
		"no card": {
			given: given{
				doneCards: []trello.Card{},
				todoCards: []trello.Card{},
			},
			expected: "",
		},
	}
	p := MarkdownPrinter{t: ti}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := bytes.NewBufferString("")
			err := p.Print(actual, tt.given.doneCards, tt.given.todoCards)
			if err != nil {
				t.Fatalf("err: %v", err)
			}
			if actual.String() != tt.expected {
				t.Errorf("expected:\n%v\nactual:\n%v", tt.expected, actual)
			}
		})
	}
}
