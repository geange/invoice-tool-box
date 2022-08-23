/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/geange/invoice-toolbox/pkg"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		processor := pkg.NewProcessor(&pkg.ProcessorConfig{
			Dir:          *printDir,
			InvoiceField: "",
		})
		processor.Load()

		switch *printFormat {
		case "csv":
			processor.DefaultCSV(*printFileName)
		}
	},
}

var printDir *string
var printFormat *string
var printFileName *string

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	printDir = printCmd.PersistentFlags().String("dir", "./", "workspace")
	printFormat = printCmd.PersistentFlags().String("format", "csv", "output file format")
	printFileName = printCmd.PersistentFlags().String("file", "out.csv", "output file name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// printCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
