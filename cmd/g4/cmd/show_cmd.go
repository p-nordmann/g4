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

import "github.com/spf13/cobra"

// showCmd provides the show command allowing to display information about G4.
//
// TODO
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display specific information about G4.",
	Long:  ``,
	Args:  cobra.NoArgs,
}

// helpCmd displays help page.
//
// TODO
var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show help.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// rulesCmd displays rules page.
//
// TODO
var rulesCmd = &cobra.Command{
	Use:   "rules",
	Short: "Show rules of the game.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// readmeCmd displays README.md.
//
// TODO
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Show README.md.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// licenseCmd displays project's license.
//
// TODO
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Show COPYING file.",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	showCmd.AddCommand(helpCmd)
	showCmd.AddCommand(rulesCmd)
	showCmd.AddCommand(readmeCmd)
	showCmd.AddCommand(licenseCmd)
}
