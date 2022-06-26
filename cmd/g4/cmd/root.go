/*
G4 is an open-source board game inspired by the popular game of connect-4.
Copyright (C) 2022  Pierre-Louis Nordmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base g4 command.
var rootCmd = &cobra.Command{
	Use:   "g4",
	Short: "G4 command-line interface.",
	Long:  `G4 is a game based on connect-4, with added rotation moves.`,
	// PreRun: func(cmd *cobra.Command, args []string) {
	// },
	Args: cobra.NoArgs,
}

func init() {
	// rootCmd.PersistentFlags().StringVarP(&ManifestoFile, "manifesto", "m", "", "manifesto file path")
	// rootCmd.MarkPersistentFlagRequired("manifesto")

	rootCmd.AddCommand(playCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(readCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
