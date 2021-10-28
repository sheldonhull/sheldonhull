package main

import (
	"os"

	"github.com/pterm/pterm"
)

const (
	exitFail    = 1
	exitSuccess = 0
)

func main() {
	pterm.Success.Println("Foo")
	hs := pterm.NewStyle(pterm.BgLightWhite, pterm.BgBlack, pterm.Bold)

	t, err := pterm.DefaultTable.WithStyle(hs).WithData(pterm.TableData{
		{"Name", "Sheldon Hull"},
		{"Website", "sheldonhull.com"},
		{"Github", "github.com/sheldonhull"},
		{"Areas Of Focus", "AWS, Go, PowerShell, Databases, Automation"},
	}).Srender()
	if err != nil {
		pterm.Error.Printf("render defaulttable: %v\n", err)
		os.Exit(exitFail)
	}

	pterm.DefaultBox.WithTitle("Sheldon Hull").Print(t)
	os.Exit(exitSuccess)
}
