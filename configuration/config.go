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
	listIDKey       = "list-id"
	keyKey          = "key"
	tokenKey        = "token"
	outputFolderKey = "output-folder"
)

// GetListID from the configuration file
func GetListID() string {
	return viper.GetString(listIDKey)
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
		Label:    "List ID",
		Validate: validateEmpty,
	}
	listID, err := prompt.Run()
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
%s: %s`,
		listIDKey, listID,
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
