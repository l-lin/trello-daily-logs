package printer

import (
	"fmt"
	"io"
	"sort"
	"strings"
	"time"

	"github.com/l-lin/trello-daily-logs/trello"
)

var now = time.Now()

// Printer prints cards
type Printer interface {
	Print(writer io.Writer, cards []trello.Card) error
}

// MarkdownPrinter prints in markdown
type MarkdownPrinter struct {
}

// Print the cards in markdown format
func (p MarkdownPrinter) Print(writer io.Writer, cards []trello.Card) error {
	if len(cards) == 0 {
		return nil
	}
	var b strings.Builder
	labelsMap := map[string]bool{}
	labels := []string{}
	for _, card := range cards {
		for _, label := range card.Labels {
			if !labelsMap[label.Name] {
				labelsMap[label.Name] = true
				labels = append(labels, label.Name)
			}
		}
	}
	sort.Strings(labels)

	b.WriteString(fmt.Sprintf("## %s %02d\n\n", now.Weekday(), now.Day()))

	for _, label := range labels {
		b.WriteString(fmt.Sprintf("- %s\n", label))
		for _, card := range cards {
			if card.ContainLabel(label) {
				b.WriteString(fmt.Sprintf("  - %s\n", card.Name))
			}
		}
	}
	b.WriteString("\n")

	if _, err := writer.Write([]byte(b.String())); err != nil {
		return err
	}

	return nil
}
