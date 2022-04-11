package cmd

import (
	"fmt"
	"github.com/leddzip/back-finder"
	"os"

	"github.com/spf13/cobra"
)

var From string
var BackTo string
var NoFileAsErrorFlag bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "back-find",
	Short: "Find a specific file in the parent hierarchy.",
	Long: `Will try to find the exact file name in one of the upper directory.
If multiple candidate exist, it will always stop at the closer to the base 
directory (default to the current directory)`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		filePresence, err := back_finder.FindFileBetween(args[0], From, BackTo)
		if err != nil {
			os.Exit(1)
		}

		if filePresence.IsFilePresent {
			fmt.Print(filePresence.FilePathIfExist)
		} else {
			fmt.Print("")
			if NoFileAsErrorFlag {
				os.Exit(2)
			}
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&From, "from", "f", ".", "From where do you want to search")
	rootCmd.Flags().StringVarP(&BackTo, "back-to", "b", "/", "Where do you want to stop the search")
	rootCmd.Flags().BoolVarP(&NoFileAsErrorFlag, "error", "e", false, "If flag enable, the command return an error (exit code 2) if no file is present")
}
