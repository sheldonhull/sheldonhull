package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/glamour"
)

const (
	exitFail    = 1
	exitSuccess = 0
)

func main() {

in := `# The Epic Personal Go Card Of Sheldon Hull

Hi, I'm Sheldon!

You must love Go too.

My personal slice of the internet is:

- [sheldonhull](https://www.sheldonhull.com)


## For My Fellow Devs:

- _Preferred Editor:_ Visual Studio Code, and I do enjoy Jetbrains Goland.
- _Light or Dark Theme:_ Dark, of course. Currently using Night Owl.
- _Spaces or Tabs:_ whatever gitattributes defines.
- Font: Jet Brains Mono & FuraCode Nerd Font Mono (ligature based font)
- _Source Control:_ Git
- _Automation:_ Azure DevOps, GitHub Actions (GitLab also seems promising)
- _Preferred Computer Setup:_ Ultrawide, mechanical keyboard, and latte...
- _Blog:_ As time permits, sheldonhull.com, though most projects I've done are not public
- _Terminal:_ Warp (Rust Based) terminal for macOS, VSCode Terminal for almost everything else. I've got a slick zinit zsh setup synced through chezmoi.
- _If I was an emoji I'd be:_ 🌮🌮🌮

`

r, _ := glamour.NewTermRenderer(
    // detect background color and pick either the default dark or light theme
    glamour.WithAutoStyle(),
    // wrap output at specific width
    glamour.WithWordWrap(100),

)
out, err := r.Render(in)
if err != nil {
	fmt.Printf("I have failed you and apologize. Quite embarassing. Can I just say, \"Worked on my machine?\"\n")
os.Exit(1)
}
fmt.Print(out)

	// pterm.Success.Println("The Epic Personal Go Card Of Sheldon Hull")
	// hs := pterm.NewStyle(pterm.BgLightWhite, pterm.BgBlack, pterm.Bold)

	// t, err := pterm.DefaultTable.WithStyle(hs).WithData(pterm.TableData{
	// 	{"Name", "Sheldon Hull"},
	// 	{"Website", "sheldonhull.com"},
	// 	{"Github", "github.com/sheldonhull"},
	// 	{"Areas Of Focus", "AWS, Go, PowerShell, Databases, Automation"},
	// }).Srender()
	// if err != nil {
	// 	pterm.Error.Printf("render defaulttable: %v\n", err)
	// 	os.Exit(exitFail)
	// }

	// pterm.DefaultBox.WithTitle("Sheldon Hull").Print(t)
	os.Exit(exitSuccess)
}
