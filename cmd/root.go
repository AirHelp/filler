package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/AirHelp/filler/log"
	"github.com/AirHelp/filler/templates"
	"github.com/spf13/cobra"
)

const (
	toScanArgument        = "src"
	fileExtensionArgument = "ext"
	deleteTemplateFile    = "delete"
	FailIfMissing = "fail-if-missing"
)

var (
	leftDelimiter  string
	rightDelimiter string
	inPlace        bool
	verbose        bool
	failIfMissing  bool
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
	rootCmd.PersistentFlags().Bool(deleteTemplateFile, false, "delete template file after filling")
	rootCmd.PersistentFlags().StringVar(&leftDelimiter, "left-delimiter", "", "left delimiter")
	rootCmd.PersistentFlags().StringVar(&rightDelimiter, "right-delimiter", "", "right delimiter")
	rootCmd.PersistentFlags().BoolVar(&inPlace, "in-place", false, "template file without creating new one")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "turn on debug logging")
	rootCmd.PersistentFlags().BoolVar(&failIfMissing, FailIfMissing, false, "will return an error if any variable is missing")
}

func template(cmd *cobra.Command, args []string) error {
	if verbose {
		log.InitLogger("debug")
	}

	if failIfMissing {
		templates.SetFailIfMissing()
	}

	if err := setDelimiters(); err != nil {
		return err
	}

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
		return errors.New("file extension is missing")
	}

	deleteFile, err := cmd.Flags().GetBool(deleteTemplateFile)
	if err != nil {
		return err
	}

	if err := templates.SearchAndFill(toScan, fileExt, deleteFile, inPlace); err != nil {
		return err
	}
	return nil
}

func setDelimiters() error {
	if (rightDelimiter != "" && leftDelimiter == "") || (rightDelimiter == "" && leftDelimiter != "") {
		return errors.New("both delimiters must be set or none")
	}

	if rightDelimiter != "" && leftDelimiter != "" {
		templates.SetDelimiters(leftDelimiter, rightDelimiter)
	}

	return nil
}
