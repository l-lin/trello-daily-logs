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
						Name: "career: find good path",
						Labels: []trello.Label{
							trello.Label{Name: "PERSO"},
						},
					},
					trello.Card{
						Name: "api@gravitee: install in prod environment",
						Labels: []trello.Label{
							trello.Label{Name: "WORK"},
						},
						Desc: "# Gravitee in prod\n## Getting started\n\n- get up from bed\n- brush teeth",
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
						Desc: "# Foo\n\n> Foobar",
					},
				},
			},
			expected: `## Sunday 09

**ABANDONED**

&nbsp;&nbsp;&nbsp; api@spec: write OpenAPI specifications

**PERSO**

&nbsp;&nbsp;&nbsp; career@interview: find good questions for interview

&nbsp;&nbsp;&nbsp; career: find good path

**WORK**

<details>
<summary>api@gravitee: install in prod environment</summary>

# Gravitee in prod
## Getting started

- get up from bed
- brush teeth

</details>

&nbsp;&nbsp;&nbsp; api@spec: write OpenAPI specifications

---

<details>
<summary>UNFINISHED</summary>

**PERSO**

&nbsp;&nbsp;&nbsp; shopping: buy milk

**WORK**

&nbsp;&nbsp;&nbsp; projectA@taskA: study solutions

<details>
<summary>projectB@taskB: implement solution</summary>

# Foo

> Foobar

</details>

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

&nbsp;&nbsp;&nbsp; career@interview: find good questions for interview

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

&nbsp;&nbsp;&nbsp; career@interview: find good questions for interview

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
