package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cobaCMD = &cobra.Command{
	Use: "test",
	Short: "Test",
	Run: func (cmd *cobra.Command, args []string)  {

		fmt.Println("ini dari coba")
		
	},
}