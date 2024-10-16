package cli

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goencrypt",
	SilenceUsage: true,
}

var MainContext context.Context

// starting app.
func Execute() {
	var cancel context.CancelFunc
	MainContext, cancel = context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()

	//base root check => first use ?

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
