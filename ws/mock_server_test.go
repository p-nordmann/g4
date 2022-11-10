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
