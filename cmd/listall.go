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
	"fmt"
	"github.com/faculerena/bugtracker/internal"
	"os"

	"github.com/spf13/cobra"
)

// listallCmd represents the listall command
var listallCmd = &cobra.Command{
	Use:   "listall ",
	Short: "List all bugs",
	Long:  `Use 'tracker list' to retrieve ALL bugs saved on the tracker`,
	Run: func(cmd *cobra.Command, args []string) {
		t := &tracker.Bugs{}
		err := t.Load(tracker.File)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		t.ListAll()
	},
}

func init() {
	rootCmd.AddCommand(listallCmd)
}
