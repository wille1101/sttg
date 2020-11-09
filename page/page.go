package page

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/wille1101/sttg/config"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
)

// GetPage - Visar sidan
func GetPage(pagenr int) (string, error) {
	svtURLslice := []string{"https://www.svt.se/svttext/tv/pages/", strconv.Itoa(pagenr), ".html"}
	svtURL := strings.Join(svtURLslice, "")
	resp, err := http.Get(svtURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	sida := ""
	doc.Find("pre.root").Contents().Each(func(i int, s *goquery.Selection) {
		st := s.Text()
		switch {
		case s.HasClass("Y"):
			sida += color.Yellow.Render(st)
		case s.HasClass("C"):
			sida += color.Cyan.Render(st)
		case s.HasClass("B"):
			sida += color.Blue.Render(st)
		default:
			sida += st
		}

	})

	return sida, nil
}

// GetHelpPage - Visar hjälpsidan
func GetHelpPage() string {
	sida := ""
	sida += fmt.Sprintf("%s / %s : Gå en sida åt vänster \n", config.Keymap["Left"][0], config.Keymap["Left"][1])
	sida += fmt.Sprintf("%s / %s : Gå en sida åt höger\n", config.Keymap["Right"][0], config.Keymap["Right"][1])
	sida += fmt.Sprintf("%s / %s : Skrolla ner på sidan\n", config.Keymap["Down"][0], config.Keymap["Down"][1])
	sida += fmt.Sprintf("%s / %s : Skrolla upp på sidan\n", config.Keymap["Up"][0], config.Keymap["Up"][1])
	sida += fmt.Sprintf("%s / %s : Gå direkt till en sida\n", config.Keymap["SetPage"][0], config.Keymap["SetPage"][1])
	sida += fmt.Sprintf("%s / %s : Stäng programmet\n", config.Keymap["Quit"][0], config.Keymap["Quit"][1])
	sida += fmt.Sprintf("%s / %s : Visar denna hjälpsida\n", config.Keymap["GetHelp"][0], config.Keymap["GetHelp"][1])
	sida += "\n"
	sida += "esc stänger denna sida"
	return sida
}
