package worklog

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/l-lin/trello-daily-logs/printer"
	"github.com/l-lin/trello-daily-logs/trello"
)

var now = time.Now()

// Write the cards in the work log file
func Write(cards []trello.Card, outputFolder string) error {
	if !exists(outputFolder) {
		return fmt.Errorf("Folder '%s' does not exists", outputFolder)
	}
	yearFolder := fmt.Sprintf("%s/%d", outputFolder, now.Year())
	if !exists(yearFolder) {
		if err := os.Mkdir(yearFolder, 0755); err != nil {
			return err
		}
	}
	monthFile := fmt.Sprintf("%s/%02d.md", yearFolder, now.Month())
	f, err := os.OpenFile(monthFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Printf("Writing content to '%s'\n", outputFolder)
	p := printer.MarkdownPrinter{}
	if err = p.Print(f, cards); err != nil {
		return err
	}
	return nil
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}
