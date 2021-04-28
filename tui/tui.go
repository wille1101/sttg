package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wille1101/sttg/config"
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
		pagenr:    config.DefPageNr,
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
			p = fmt.Sprintf("Error: \n %s", err)
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

		case config.Keymap["Quit"][0], config.Keymap["Quit"][1]:
			return m, tea.Quit

		case config.Keymap["Right"][0], config.Keymap["Right"][1]:
			if m.pagenr < 999 {
				m.pagenr++
				return m, getPageWrap(m)
			}
		case config.Keymap["Left"][0], config.Keymap["Left"][1]:
			if m.pagenr > 100 {
				m.pagenr--
				return m, getPageWrap(m)
			}
		case config.Keymap["Up"][0], config.Keymap["Up"][1]:
			m.viewport.LineUp(1)
			return m, nil

		case config.Keymap["Down"][0], config.Keymap["Down"][1]:
			m.viewport.LineDown(1)
			return m, nil

		case config.Keymap["GoTop"][0], config.Keymap["GoTop"][1]:
			m.viewport.GotoTop()
			return m, nil

		case config.Keymap["GoBot"][0], config.Keymap["GoBot"][1]:
			m.viewport.GotoBottom()
			return m, nil

		case config.Keymap["GoViewUp"][0], config.Keymap["GoViewUp"][1]:
			m.viewport.ViewUp()
			return m, nil

		case config.Keymap["GoViewDown"][0], config.Keymap["GoViewDown"][1]:
			m.viewport.ViewDown()
			return m, nil

		case config.Keymap["SetPage"][0], config.Keymap["SetPage"][1]:
			m.textInput.Reset()
			m.textInput.Focus()
			return inputUpdate(nil, m)

		case "1", "2", "3", "4", "5", "6", "7", "8", "9":
			m.textInput.Reset()
			m.textInput.Focus()
			return inputUpdate(msg, m)

		case "esc":
			return m, getPageWrap(m)

		case config.Keymap["GetHelp"][0], config.Keymap["GetHelp"][1]:
			m.viewport.GotoTop()
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
	shl := fmt.Sprintf("%s", m.textInput.View())
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
