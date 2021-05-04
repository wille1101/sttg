package page

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gookit/color"
	"github.com/wille1101/sttg/config"
)

//Page - är en typ som representerar en sida med dess information.
type Page struct {
	Content       []string `json:"content"`
	ContentParsed string
	NextPage      string `json:"next_page"`
	PrevPage      string `json:"prev_page"`
}

//Parr - är en array av sid-slices
var Parr [999][1]Page

func downloadPage(pagenr int) error {
	svtURL := fmt.Sprintf("http://api.texttv.nu/api/get/%d?app=svttexttvtgo", pagenr)
	req, err := http.NewRequest("GET", svtURL, nil)
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&Parr[pagenr]); err != nil {
		return err
	}

	if err := parsePage(pagenr); err != nil {
		return err
	}

	return nil
}

func parsePage(pagenr int) error {
	var sida strings.Builder
	for i := range Parr[pagenr][0].Content {
		doc, err := goquery.NewDocumentFromReader(strings.NewReader(Parr[pagenr][0].Content[i]))
		if err != nil {
			return err
		}
		doc.Find("div.root").Contents().Each(func(i int, s *goquery.Selection) {
			switch {
			case s.Children().HasClass("Y"):
				sida.WriteString(color.Yellow.Render(s.Text()))
			case s.Children().HasClass("C"):
				sida.WriteString(color.Cyan.Render(s.Text()))
			case s.Children().HasClass("B"):
				sida.WriteString(color.Blue.Render(s.Text()))
			default:
				sida.WriteString(s.Text())
			}
		})
		sida.WriteString("\n\n")
		Parr[pagenr][0].ContentParsed = sida.String()
	}
	return nil
}

// GetPage - Visar sidan
func GetPage(pagenr int) (string, error) {
	if Parr[pagenr][0].ContentParsed == "" {
		if err := downloadPage(pagenr); err != nil {
			return "", err
		}
	}

	return Parr[pagenr][0].ContentParsed, nil
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
