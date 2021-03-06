package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/l-lin/trello-daily-logs/configuration"
	"github.com/l-lin/trello-daily-logs/printer"
	"github.com/l-lin/trello-daily-logs/trello"
	"github.com/l-lin/trello-daily-logs/worklog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	cfgFileName = ".trello-daily-logs"
)

var (
	cfgFile string
	format  string
	rootCmd = &cobra.Command{
		Use:   "trello-daily-logs",
		Short: "Fetch cards from a list of trello (usually DONE), and write the card names in a markdown file",
		Run:   run,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version, buildDate string) {
	rootCmd.Version = func(version, buildDate string) string {
		res, err := json.Marshal(cliBuild{Version: version, BuildDate: buildDate})
		if err != nil {
			log.Fatal(err)
		}
		return string(res)
	}(version, buildDate)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(`{{printf "%s" .Version}}`)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.trello-daily-logs.yaml)")
	rootCmd.PersistentFlags().StringVarP(&format, "format", "f", "console", "output format to display the content (available: console, file)")
}

func run(cmd *cobra.Command, args []string) {
	listDoneID := configuration.GetListDoneID()
	listTodoID := configuration.GetListTodoID()
	key := configuration.GetKey()
	token := configuration.GetToken()
	doneCardsCh := make(chan getCardsResult)
	todoCardsCh := make(chan getCardsResult)
	go getCards(listDoneID, key, token, doneCardsCh)
	go getCards(listTodoID, key, token, todoCardsCh)
	doneCardsResult := <-doneCardsCh
	todoCardsResult := <-todoCardsCh

	if doneCardsResult.err != nil {
		log.Fatal(doneCardsResult.err)
	}
	if todoCardsResult.err != nil {
		log.Fatal(todoCardsResult.err)
	}
	p := printer.NewMarkdownPrinter(time.Now())
	if format == "file" {
		if err := worklog.Write(p, doneCardsResult.cards, todoCardsResult.cards, configuration.GetOutputFolder()); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := p.Print(os.Stdout, doneCardsResult.cards, todoCardsResult.cards); err != nil {
			log.Fatal(err)
		}
	}
}

func getCards(listID, key, token string, cardsResultCh chan getCardsResult) {
	cards, err := trello.GetCards(listID, key, token)
	if err != nil {
		cardsResultCh <- getCardsResult{err: err}
	}
	cardsResultCh <- getCardsResult{cards: cards}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".trello-daily-logs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(cfgFileName)
		cfgFile = fmt.Sprintf("%s/%s.yaml", home, cfgFileName)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		listDoneID := configuration.GetListDoneID()
		if listDoneID == "" {
			log.Printf("Could not read the properties. Initializing them in %s", cfgFile)
			configuration.InitConfig(cfgFile)
			viper.ReadInConfig()
		}
	} else { // Else we create the file
		log.Printf("Could not read the config file '%s'. Creating it.\n", cfgFile)
		configuration.InitConfig(cfgFile)
		viper.ReadInConfig()
	}
}

type cliBuild struct {
	Version   string `json:"version"`
	BuildDate string `json:"buildDate"`
}

type getCardsResult struct {
	cards []trello.Card
	err   error
}
