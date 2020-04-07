package printer

import (
	"io"
	"text/template"
	"time"

	"github.com/l-lin/trello-daily-logs/trello"
)

const (
	cardsTpl = `{{define "cards"}}{{range $k, $v := .}}
**{{$k}}**
{{range $v}}{{if .Desc}}
<details>
<summary>{{.Name}}</summary>

{{.Desc}}

</details>
{{else}}
&nbsp;&nbsp;&nbsp; {{.Name}}
{{end}}{{end}}{{end}}{{end}}`

	allCardsTpl = `## {{.T.Weekday}} {{printf "%02d" .T.Day}}
{{if .DoneMap}}{{template "cards" .DoneMap}}{{if .TodoMap}}
---
{{end}}{{end}}
{{if .TodoMap}}<details>
<summary>UNFINISHED</summary>
{{template "cards" .TodoMap}}
</details>

{{end}}`
)

// Printer prints cards
type Printer interface {
	Print(writer io.Writer, cardsDone, cardsTodo []trello.Card) error
}

// MarkdownPrinter prints in markdown
type MarkdownPrinter struct {
	t time.Time
}

// NewMarkdownPrinter creates a new trello cards printer that prints in markdown format
func NewMarkdownPrinter(t time.Time) Printer {
	return MarkdownPrinter{t: t}
}

// Print the cards in markdown format
func (p MarkdownPrinter) Print(writer io.Writer, doneCards, todoCards []trello.Card) error {
	if len(doneCards) == 0 && len(todoCards) == 0 {
		return nil
	}

	tpl := template.Must(template.New("all-cards").Parse(allCardsTpl))
	tpl = template.Must(tpl.Parse(cardsTpl))
	tplParams := struct {
		T       time.Time
		DoneMap map[string][]trello.Card
		TodoMap map[string][]trello.Card
	}{
		T:       p.t,
		DoneMap: toMap(doneCards),
		TodoMap: toMap(todoCards),
	}
	if err := tpl.Execute(writer, tplParams); err != nil {
		return err
	}

	return nil
}

func toMap(cards []trello.Card) map[string][]trello.Card {
	m := map[string][]trello.Card{}
	for _, card := range cards {
		for _, label := range card.Labels {
			if _, ok := m[label.Name]; !ok {
				m[label.Name] = []trello.Card{}
			}
			m[label.Name] = append(m[label.Name], card)
		}
	}
	return m
}
