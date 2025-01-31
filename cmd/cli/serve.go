package cli

import (
	"fmt"
	"leap_store/app/configs"
	

	//"leap_store/app/controllers"
	"leap_store/routes"

	"github.com/spf13/cobra"
)

var serveCMD = &cobra.Command{
	Use: "serve",
	Short: "Server running on :8000",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("Server running on :8000")

		

		
		configs.ConnectDB()
		//configs.InsertDummyData(configs.DB, 2000000)
		routes.RegisterRoutes()

	},
}

