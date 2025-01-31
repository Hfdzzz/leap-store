package cli

import (
	"fmt"
	
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use: "leap_store",
	Short: "Leap Store CLI",
	Run: func(cmd *cobra.Command, args []string){
		if len(args) == 0 {
			fmt.Println("Welcome to the leap store cli")
			fmt.Println("Use 'go run main.go serve' for running the server")
			
			

		}else{
			fmt.Printf("uknown command %s", args[0])
			os.Exit(1)
		}
	},
	SilenceUsage: true,
}

func Execute(){
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init(){
	rootCMD.AddCommand(serveCMD)
	rootCMD.AddCommand(cobaCMD)
	rootCMD.AddCommand(migrateCMD)
}