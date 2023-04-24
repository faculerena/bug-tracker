/*
Copyright © 2023 Facundo Lerena  <contacto@faculerena.com.ar>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"
	tracker "github.com/faculerena/bugtracker/internal"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates the tracker",
	Long:  `Create the .tracker.json file and start tracking!`,
	Run: func(cmd *cobra.Command, args []string) {
		
		_, err := os.Open(tracker.File)
		if errors.Is(err, os.ErrNotExist) != true {
			fmt.Println("File already exists. If you want to clear the tracker use 'tracker clear'")
			os.Exit(1)
		}

		fmt.Println("Initializing files")

		_, err = os.Create(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = os.Create(".id")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.WriteFile(".id", []byte(strconv.Itoa(1)), 0644)

		fmt.Println("You can start tracking bugs!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}
