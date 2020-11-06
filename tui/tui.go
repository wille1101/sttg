package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wille1101/sttg/page"

	input "github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	page      string
	pagenr    int
	textInput input.Model
	viewport  viewport.Model
	vpReady   bool
}

type pageMsg string

func initModel() model {
	inputModel := input.NewModel()
	inputModel.Placeholder = "100"
	inputModel.CharLimit = 3

	return model{
		textInput: inputModel,
		page:      "",
		pagenr:    100,
	}
}

//NewProgram - Kallas av main och startar upp UIet.
func NewProgram() *tea.Program {
	p := tea.NewProgram(initModel())
	p.EnterAltScreen()
	return p
}

func getPageWrap(m model) tea.Cmd {
	return func() tea.Msg {
		p, err := page.GetPage(m.pagenr)
		if err != nil {
			p = "Error: \n Kommer inte åt SVTs hemsida.\n Har du en nätverksanslutning?"
		}
		return pageMsg(p)
	}
}

func getHelpPageWrap(m model) tea.Cmd {
	return func() tea.Msg {
		p := page.GetHelpPage()
		return pageMsg(p)
	}
}

func (m model) Init() tea.Cmd {
	return getPageWrap(m)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.textInput.Focused() {
		return inputUpdate(msg, m)
	}
	return pageUpdate(msg, m)
}

func pageUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case pageMsg:
		m.page = string(msg)
		m.viewport.SetContent(m.page)

	case tea.WindowSizeMsg:
		if !m.vpReady {
			m.viewport = viewport.Model{Width: msg.Width, Height: msg.Height - 3}
			m.vpReady = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - 3
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "right", "l":
			if m.pagenr < 999 {
				m.pagenr++
				return m, getPageWrap(m)
			}
		case "left", "h":
			if m.pagenr > 100 {
				m.pagenr--
				return m, getPageWrap(m)
			}
		case "down", "j":
			m.viewport.LineDown(1)
			return m, nil
		case "up", "k":
			m.viewport.LineUp(1)
			return m, nil
		case ":", "i":
			m.textInput.Reset()
			m.textInput.Focus()
			return inputUpdate(nil, m)
		case "1":
			m.pagenr = 100
			return m, getPageWrap(m)
		case "2":
			m.pagenr = 200
			return m, getPageWrap(m)
		case "3":
			m.pagenr = 300
			return m, getPageWrap(m)
		case "4":
			m.pagenr = 400
			return m, getPageWrap(m)
		case "esc":
			return m, getPageWrap(m)
		case "H":
			return m, getHelpPageWrap(m)
		}
	}
	return m, nil
}

func inputUpdate(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			m.textInput.Blur()
			return pageUpdate(msg, m)

		case "enter":
			m.viewport.GotoTop()
			tempInt, _ := strconv.ParseInt(m.textInput.Value(), 10, 64)
			m.pagenr = int(tempInt)
			if m.pagenr < 100 {
				m.pagenr = 100
			}
			m.textInput.Blur()
			return m, getPageWrap(m)
		}
	}

	return m, cmd
}

func (m model) View() string {
	iw := len(m.textInput.Value())
	shl := ""
	shl += fmt.Sprintf("%s", m.textInput.View())
	if len(m.textInput.Value()) == 0 {
		iw = 2
	}

	shr := fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100)
	shm := strings.Repeat(" ", 34-iw)
	shb, sf := strings.Repeat("-", 41), strings.Repeat("-", 41)
	sh := shl + shm + shr + "\n" + shb

	sm := fmt.Sprintf("%s", viewport.View(m.viewport))

	return fmt.Sprintf("%s\n%s\n%s", sh, sm, sf)
}
