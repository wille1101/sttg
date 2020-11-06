package page

import (
	"net/http"
	"strconv"
	"strings"

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
	sida += "h / vänster piltanget: Gå en sida åt vänster \n"
	sida += "l / höger piltanget  : Gå en sida åt höger\n"
	sida += ": / i                : Gå direkt till en sida\n"
	sida += "q / Ctrl + c         : Stäng programmet\n"
	sida += "1                    : Gå direkt till Nyheter\n"
	sida += "2                    : Gå direkt till Ekonomi\n"
	sida += "3                    : Gå direkt till Sport\n"
	sida += "4                    : Gå direkt till Väder\n"
	sida += "H                    : Visar denna hjälpsida\n"
	sida += "\n"
	return sida
}
