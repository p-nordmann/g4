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

package ws

import (
	"testing"
)

func TestAreUrlEqual(t *testing.T) {
	t.Run("identical url are equal", func(t *testing.T) {
		if !areUrlEqual("127.0.0.1", "127.0.0.1") {
			t.Error("got false but wanted true")
		}
	})
	t.Run("identical url with different ports are equal", func(t *testing.T) {
		if !areUrlEqual("127.0.0.1:1234", "127.0.0.1:8080") {
			t.Error("got false but wanted true")
		}
	})
	t.Run("different addresses are not equal", func(t *testing.T) {
		if areUrlEqual("127.0.0.1", "127.0.0.2") {
			t.Error("got true but wanted false")
		}
	})
	t.Run("localhost and 127.0.0.1 are equal", func(t *testing.T) {
		if !areUrlEqual("localhost", "127.0.0.1") {
			t.Error("got false but wanted true")
		}
	})
	t.Run("complete addresses should be handled correctly", func(t *testing.T) {
		if !areUrlEqual("http://127.0.0.1", "127.0.0.1") {
			t.Error("got false but wanted true")
		}
	})
}
