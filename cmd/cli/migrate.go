package cli

import (
	"leap_store/app/configs"
	"leap_store/app/models"

	"github.com/spf13/cobra"
)

var migrateCMD = &cobra.Command{
	Use: "migrate",
	Short: "Auto Migrate Model",
	Run: func (cmd *cobra.Command, args []string){

		configs.ConnectDB()

		err := configs.DB.AutoMigrate(&models.Cart{})

		if err != nil {
			panic("Error to migrate")
		}
		
	},
}