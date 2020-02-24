package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var (
	count int
	all   bool
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Reverse the migration",
	Long: `Reverse the migration.

You may set --count flag to set how many migrations to be reversed.
You may also set the --all flag to reverse all of the migrations.
Otherwise, it will reverse only to one migration before.
However, if --all flag is set, it will ignore --count flag
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if all {
			if err := sch.ReverseAll(); err != nil {
				log.Println("Error occured when reversing migrations: ", err)
				return
			}
			log.Println("Migrations reversed successfully")
			return
		}
		if count > 0 {
			if err := sch.ReverseN(count); err != nil {
				log.Println("Error occured when reversing migrations: ", err)
				return
			}
			log.Println("Migrations reversed successfully")
			return
		}

		if err := sch.ReverseLast(); err != nil {
			log.Println("Error occured when reversing migrations: ", err)
			return
		}
		log.Println("Migrations reversed successfully")
	},
}

func init() {
	rootCmd.AddCommand(downCmd)

	downCmd.Flags().IntVarP(&count, "count", "c", 0, "reverse n migrations")
	downCmd.Flags().BoolVarP(&all, "all", "a", false, "reverse all migrations")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
