package cmd

import (
	"errors"
	"fmt"
	"os"

	t "github.com/AirHelp/filler/templates"
	"github.com/spf13/cobra"
)

const (
	toScanArgument        = "src"
	fileExtensionArgument = "ext"
)

var rootCmd = &cobra.Command{
	Use:   "filler",
	Short: "Filler - fill templates with environment variables",
	Long:  "Filler - fill templates with environment variables",
	RunE:  template,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String(toScanArgument, "", "directory where app will search for templates or single file template")
	rootCmd.PersistentFlags().String(fileExtensionArgument, "tpl", "template file extension")
}

func template(cmd *cobra.Command, args []string) error {

	toScan, err := cmd.Flags().GetString(toScanArgument)
	if err != nil {
		return err
	}
	if toScan == "" {
		return errors.New("directory path is missing")
	}

	fileExt, err := cmd.Flags().GetString(fileExtensionArgument)
	if err != nil {
		return err
	}
	if fileExt == "" {
		return errors.New("File extension is missing")
	}

	if err := t.SearchAndFill(toScan, fileExt); err != nil {
		return err
	}
	return nil
}
