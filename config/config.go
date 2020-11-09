package config

import (
	"os"
	"fmt"
	v "github.com/spf13/viper"
)

var (
	//Keymap är en map av all handligar samt tangenterna som förknippas med dom
	Keymap = map[string][]string{
		"Up":      {"", ""},
		"Down":    {"", ""},
		"Left":    {"", ""},
		"Right":   {"", ""},
		"SetPage": {"", ""},
		"GetHelp": {"", ""},
		"Quit":    {"", ""},
	}

	//DefPageNr är standardsidan som visas när programmet startar
	DefPageNr int
	
	winPath = fmt.Sprintf("%s%s", os.Getenv("APPDATA"), "\\sttg\\")
)

//LoadCon laddar in configfilen om den finns och sätter till standardvärden om den inte skulle finnas eller om något inte är definerat i den
func LoadCon() error {
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("$HOME/.config/sttg/")
	v.AddConfigPath(winPath)
	v.AddConfigPath(".")

	v.SetDefault("Keys.Up", []string{"k", "up"})
	v.SetDefault("Keys.Down", []string{"j", "down"})
	v.SetDefault("Keys.Left", []string{"h", "left"})
	v.SetDefault("Keys.Right", []string{"l", "right"})
	v.SetDefault("Keys.SetPage", []string{"i", ":"})
	v.SetDefault("Keys.GetHelp", []string{"H", ""})
	v.SetDefault("Keys.Quit", []string{"q", "ctrl + q"})

	v.SetDefault("Page.DefPageNr", 100)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigFileNotFoundError); !ok {
			//Fel när filen läses in,
			return err
		}
	}

	Keymap["Up"] = v.GetStringSlice("Keys.Up")
	Keymap["Down"] = v.GetStringSlice("Keys.Down")
	Keymap["Left"] = v.GetStringSlice("Keys.Left")
	Keymap["Right"] = v.GetStringSlice("Keys.Right")
	Keymap["SetPage"] = v.GetStringSlice("Keys.SetPage")
	Keymap["GetHelp"] = v.GetStringSlice("Keys.GetHelp")
	Keymap["Quit"] = v.GetStringSlice("Keys.Quit")

	DefPageNr = v.GetInt("Page.DefPageNr")

	return nil
}
