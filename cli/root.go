package cli

import (
	"GoEncrypt/pkg/keys"
	"GoEncrypt/pkg/utils"
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goencrypt",
	SilenceUsage: true,
	Short:        "A simple file encryption tool",
	Long:         "This tool provides functionalities to encrypt files on your system.",
}

var MainContext context.Context

// starting app.
func Execute() {
	//intial check.
	init, err := InitialConfig()
	if err != nil || !init {
		fmt.Println("\n[ERROR] loading initial config")
		keys.CleanKeys()
		os.Exit(1)
	}

	var cancel context.CancelFunc
	MainContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	go func() {
		select {
		case <-signalChan:
			fmt.Println("\n[ERROR] Keyboard interrupt detected, terminating...")
			cancel()
		case <-MainContext.Done():
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// Leaving this in results in the same error appearing twice
		// Once before and once after the help output. Not sure if
		// this is going to be needed to output other errors that
		// aren't automatically outputted.
		// fmt.Println(err)
		os.Exit(1)
	}
}

func InitialConfig() (bool, error) {
	containsKeys, err := utils.ContainsKeys()
	if err != nil {
		return false, errors.New("error append")
	}

	if !containsKeys {
		fmt.Println("It seems, to be your first time using GoEncrypt ----> Generating new pair of keys.")

		//generating keys.
		rsaKeys, err := keys.GenerateRSAKeys(2048)
		if err != nil {
			return false, err
		}

		//saving keys.
		errSaveKeys := keys.AddKeysToFile(rsaKeys, &rsaKeys.PublicKey)
		if errSaveKeys != nil {
			return false, errSaveKeys
		}

		fmt.Println("\n------------------------------------------")
		fmt.Println("Keys successfuly generated in /data/keys !")

		return true, nil

	} else {
		return true, nil
	}
}
