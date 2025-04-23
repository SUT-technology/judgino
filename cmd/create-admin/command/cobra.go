package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (

	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long: `Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			fmt.Println("Creating admin with username:", username)
			if username == "" || password == "" {
				fmt.Println("Username and password are required")
				return
			}
			err := Run()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Println("Admin created successfully")
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("username", "u", "", "Admin's username")
	rootCmd.PersistentFlags().StringP("password", "p", "", "Admin's password")
}

