package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func viewModal(content string, hover bool) string {
	// See lipgloss example:
	// https://github.com/charmbracelet/lipgloss/blob/master/examples/layout/main.go

	dialogBoxStyle := lipgloss.NewStyle().
		Padding(1, 0).
		Border(lipgloss.RoundedBorder()).BorderForeground(pink).
		BorderTop(true).BorderLeft(true).BorderRight(true).BorderBottom(true)

	contentStyle := lipgloss.NewStyle().
		PaddingLeft(3).PaddingRight(3)

	buttonStyle := lipgloss.NewStyle().
		Foreground(light).Background(pink).
		Padding(0, 3).MarginTop(1)
	if hover {
		buttonStyle = buttonStyle.
			Foreground(light).Background(pinker)
	}

	return dialogBoxStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			contentStyle.Render(content),
			buttonStyle.Render("OK"),
		),
	)
}

// TODO handle errors here too
func handleModalEvents(app AppModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.MouseMsg:
		btnX, btnY := computeBtnCoords(app)
		switch msg.Type {

		case tea.MouseMotion:
			if isInsideBtn(msg.X, msg.Y, btnX, btnY) {
				app.modalHover = true
			} else {
				app.modalHover = false
			}

		case tea.MouseLeft:
			if isInsideBtn(msg.X, msg.Y, btnX, btnY) {
				app = closeModal(app)
			}

		}

	case tea.WindowSizeMsg:
		app.height = msg.Height
		app.width = msg.Width

	case tea.KeyMsg:
		combo := app.keyHandler.handle(msg.String())
		app.debug = combo
		switch combo {

		case "enter":
			app = closeModal(app)

		case "quit":
			if p2pService.ch != nil {
				p2pService.ch.Close()
			}
			return closeModal(app), tea.Quit

		}
	}

	return app, nil
}

func closeModal(app AppModel) AppModel {
	app.modalContent = ""
	app.modalHover = false
	return app
}

func computeBtnCoords(app AppModel) (int, int) {
	const btnW = 8

	modalView := viewModal(app.modalContent, false)

	// Dimensions of the modal.
	modalW := lipgloss.Width(modalView)
	modalH := lipgloss.Height(modalView)

	// Position of the button inside the modal.
	btnX := (modalW - btnW + 1) / 2 // lipgloss seems to round up so we add 1
	btnY := (modalH - 3)

	// Modal position in the terminal.
	modalX := (app.width - modalW + 1) / 2
	modalY := (app.height - 1 - modalH) / 2

	return modalX + btnX, modalY + btnY
}

func isInsideBtn(x, y, btnX, btnY int) bool {
	if y != btnY {
		return false
	}
	if x < btnX || btnX+8 <= x {
		return false
	}
	return true
}
