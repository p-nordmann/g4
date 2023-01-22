package main

import (
	"context"
	"errors"
	"g4"
	"g4/bitsim"

	tea "github.com/charmbracelet/bubbletea"
)

type SelectedModel int

const (
	selectedBoard SelectedModel = iota
	selectedPreview
	selectedSelector
	selectedError
)

type Size struct {
	Height, Width int
}

type AppModel struct {
	board         BoardModel
	preview       BoardModel
	selector      SelectorModel
	errorModel    ErrorModel
	selectedModel SelectedModel

	statusBar StatusBarModel

	game bitsim.Game

	width, height int
	descr         string
}

func (app AppModel) Init() tea.Cmd {
	cmd, err := p2pService.connect(context.Background(), app.descr)
	if err != nil {
		return handleError(err)
	}
	return tea.Batch(
		app.board.Init(),
		app.preview.Init(),
		app.selector.Init(),
		app.statusBar.Init(),
		cmd,
	)
}

func (app AppModel) updateChildren(msg tea.Msg) (tea.Model, tea.Cmd) {
	board, cmd1 := app.board.Update(msg)
	preview, cmd2 := app.preview.Update(msg)
	selector, cmd3 := app.selector.Update(msg)
	errorModel, cmd4 := app.errorModel.Update(msg)
	statusBar, cmd5 := app.statusBar.Update(msg)
	app.board = board.(BoardModel)
	app.preview = preview.(BoardModel)
	app.selector = selector.(SelectorModel)
	app.errorModel = errorModel.(ErrorModel)
	app.statusBar = statusBar.(StatusBarModel)
	return app, tea.Batch(cmd1, cmd2, cmd3, cmd4, cmd5)
}

func (app AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Quit messages.
	if isQuitMessage(msg) {
		return ErrorModel{
			err: errors.New("SIGTERM"),
		}, tea.Quit
	}

	// In case of error, we can only quit.
	if app.selectedModel == selectedError {
		return app, nil
	}

	// Error message.
	if err, ok := msg.(error); ok {
		app.errorModel = ErrorModel{err}
		app.selectedModel = selectedError
	}

	// Regular messages.
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		app.height = msg.Height
		app.width = msg.Width
		return app.updateChildren(Size{
			Height: app.height - 1, // Keep 1 line free for the status bar.
			Width:  app.width,
		})

	case SentMove:
		game, err := app.game.Apply(g4.Move(msg))
		if err != nil {
			return app, handleError(err)
		}
		app.game = game
		app.board.game = game
		app.preview.game = game
		app.selectedModel = selectedBoard

	case ReceivedMove:
		game, err := app.game.Apply(g4.Move(msg))
		if err != nil {
			return app, handleError(err)
		}
		app.game = game
		app.board.game = game
		app.preview.game = game
		app.selectedModel = selectedBoard

	case SelectedMove:
		game, err := app.preview.game.Apply(g4.Move(msg))
		if err != nil {
			return app, handleError(err)
		}
		app.preview.game = game
		app.selectedModel = selectedPreview

	default:
		return app.updateChildren(msg)
	}

	return app, nil
}

func (app AppModel) View() string {
	s := ""
	switch app.selectedModel {
	case selectedBoard:
		s += app.board.View()
	case selectedPreview:
		s += app.preview.View()
	case selectedSelector:
		s += app.selector.View()
	case selectedError:
		s += app.errorModel.View()
	}
	return s + "\n" + app.statusBar.View()
}
