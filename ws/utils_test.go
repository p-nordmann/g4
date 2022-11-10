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
