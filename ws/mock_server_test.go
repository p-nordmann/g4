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
	"context"
	"errors"
	"net/http"
)

type MockServer struct {
	server                             *http.Server
	mustCrash                          bool
	cancel                             chan context.Context
	countListenAndServe, countShutdown int
}

func (mockServer *MockServer) ListenAndServe() error {
	mockServer.countListenAndServe++
	go func() {
		mockServer.server.ListenAndServe()
	}()
	if mockServer.mustCrash {
		mockServer.server.Close()
		return errors.New("server crashed")
	}
	ctx := <-mockServer.cancel
	mockServer.server.Shutdown(ctx)
	return errors.New("server was canceled")
}

func (mockServer *MockServer) Shutdown(ctx context.Context) error {
	mockServer.countShutdown++
	mockServer.cancel <- ctx
	return nil
}
