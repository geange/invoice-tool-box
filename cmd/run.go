/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/geange/invoice-toolbox/pkg"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		processor := pkg.NewProcessor(&pkg.ProcessorConfig{
			Dir:          *dir,
			InvoiceField: *field,
		})
		err := processor.Load()
		if err != nil {
			//log.Println(err)
			return
		}
		err = processor.Run()
		if err != nil {
			//log.Println(err)
			return
		}
		err = processor.Close()
		if err != nil {
			//log.Println(err)
			return
		}
	},
}

var (
	dir   *string
	field *string
)

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	dir = runCmd.PersistentFlags().String("dir", "./", "set the work Directory")
	field = runCmd.PersistentFlags().String("field", "InvoiceNum", "set the new file name")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
