package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

const (
	listIDsKey      = "list-ids"
	keyKey          = "key"
	tokenKey        = "token"
	outputFolderKey = "output-folder"
)

var (
	listDoneIDKey = listIDsKey + ".done"
	listTodoIDKey = listIDsKey + ".todo"
)

// GetListDoneID from the configuration file
func GetListDoneID() string {
	return viper.GetString(listDoneIDKey)
}

// GetListTodoID from the configuration file
func GetListTodoID() string {
	return viper.GetString(listTodoIDKey)
}

// GetKey from the configuration file
func GetKey() string {
	return viper.GetString(keyKey)
}

// GetToken from the configuration file
func GetToken() string {
	return viper.GetString(tokenKey)
}

// GetOutputFolder from the configuration file
func GetOutputFolder() string {
	return viper.GetString(outputFolderKey)
}

// InitConfig initialize the cli configuration
func InitConfig(cfgFile string) {
	prompt := promptui.Prompt{
		Label:    "Your DONE list ID",
		Validate: validateEmpty,
	}
	listDoneID, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label:    "Your TODO list ID",
		Validate: validateEmpty,
	}
	listTodoID, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label:    "Key",
		Validate: validateEmpty,
		Mask:     '*',
	}
	key, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label:    "Token",
		Validate: validateEmpty,
		Mask:     '*',
	}
	token, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}
	prompt = promptui.Prompt{
		Label: "Output folder",
	}
	outputFolder, err := prompt.Run()
	if err != nil {
		log.Fatal(err)
	}

	cfgContent := []byte(fmt.Sprintf(`%s: %s
%s: %s
%s: %s
%s: %s
%s: %s`,
		listDoneIDKey, listDoneID,
		listTodoIDKey, listTodoID,
		keyKey, key,
		tokenKey, token,
		outputFolderKey, outputFolder,
	))
	err = ioutil.WriteFile(cfgFile, cfgContent, 0644)

	if err != nil {
		log.Fatalf("%v", err)
	}
}

func validateEmpty(input string) error {
	if strings.Trim(input, " ") == "" {
		return errors.New("Content must not be empty")
	}
	return nil
}
