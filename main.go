//go:generate go install -v github.com/kevinburke/go-bindata/go-bindata
//go:generate go-bindata -pkg bindata -o app/bindata/bindata.go app/settings.json data/... app.conf
//go:generate go install -v github.com/josephspurrier/goversioninfo/cmd/goversioninfo
//go:generate goversioninfo -icon=app.ico -manifest=app.manifest

package main

import (
	_ "github.com/josephspurrier/goversioninfo"
	_ "github.com/kevinburke/go-bindata"

	"fmt"

	"github.com/crazy-max/WindowsSpyBlocker/app/cmds/dev"
	"github.com/crazy-max/WindowsSpyBlocker/app/cmds/telemetry"
	"github.com/crazy-max/WindowsSpyBlocker/app/menu"
	"github.com/crazy-max/WindowsSpyBlocker/app/utils/app"
	"github.com/crazy-max/WindowsSpyBlocker/app/utils/config"
	"github.com/crazy-max/WindowsSpyBlocker/app/utils/windows"
	"github.com/fatih/color"
	"github.com/mcuadros/go-version"
)

func init() {
	windows.SetConsoleTitle(fmt.Sprintf("%s %s", config.AppName, config.AppVersion))
}

func main() {
	color.New(color.FgHiWhite).Println(config.AppName + " " + config.AppVersion)
	color.New(color.FgHiWhite).Println(config.AppUrl)

	latestVersion, err := app.GetLatestVersion()
	if err != nil {
		color.New(color.FgRed).Printf("\n%s can't contact the update server: %s", config.AppName, err.Error())
	} else if version.Compare(config.AppVersion, latestVersion, "<") {
		color.New(color.FgHiGreen).Print("\nA new release is available : ")
		color.New(color.FgHiGreen, color.Bold).Print(latestVersion)
		color.New(color.FgHiGreen).Print("\nDownload : ")
		color.New(color.FgHiGreen, color.Bold).Print(config.AppUrl + "/releases/latest\n")
	}

	menuCommands := []menu.CommandOption{
		{
			Description: "> Telemetry : Block telemetry and data collection",
			Color:       color.FgYellow,
			Function:    telemetry.Menu,
		},
		{
			Description: "> Dev       : Several tools used by WindowsSpyBlocker",
			Color:       color.FgYellow,
			Function:    dev.Menu,
		},
	}

	menuOptions := menu.NewOptions("Main", "'menu' for help [main]> ", 0, "")

	menuN := menu.NewMenu(menuCommands, menuOptions)
	menuN.Start()
}
