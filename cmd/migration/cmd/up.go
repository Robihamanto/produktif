package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run all migrations",
	Long:  `Run all migrations.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := sch.Migrate(); err != nil {
			log.Println("Error occured when running migration: ", err)
			return
		}
		log.Println("Migration run successfully")
	},
}

func init() {
	rootCmd.AddCommand(upCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// upCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// upCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
