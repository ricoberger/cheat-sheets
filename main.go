package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/orderedmap"
)

type CheatSheet []CheatSheetSection

type CheatSheetSection struct {
	Title    string                `json:"title"`
	Commands orderedmap.OrderedMap `json:"commands"`
	Tip      Tip                   `json:"tip"`
}

type Tip struct {
	Title    string                `json:"title"`
	Commands orderedmap.OrderedMap `json:"commands"`
}

// readCheatSheet reads the cheat sheet from the provided file and parse it
// into a `CheatSheet` struct.
func readCheatSheet() CheatSheet {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: ./generator data/vim.json")
	}

	f, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalln("Failed to read cheat sheet", err)
	}

	var cheatSheet CheatSheet
	err = json.Unmarshal([]byte(f), &cheatSheet)
	if err != nil {
		log.Fatalln("Failed to parse cheat sheet", err)
	}

	return cheatSheet
}

// renderCheatSheet renders the cheat sheet into a HTML string.
func renderCheatSheet(cheatSheet CheatSheet) string {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.New("index.html").Funcs(template.FuncMap{
		"unescapeHTML": func(s string) template.HTML {
			return template.HTML(s)
		},
		"getCommand": func(commands orderedmap.OrderedMap, key string) string {
			val, _ := commands.Get(key)
			return val.(string)
		},
		"formatCommand": func(s string) template.HTML {
			return template.HTML(strings.ReplaceAll(strings.ReplaceAll(strings.TrimPrefix(strings.TrimPrefix(s, "!!! "), "||| "), " or ", "</kbd> or <kbd>"), " + ", "</kbd> + <kbd>"))
		},
		"hasPrefix": strings.HasPrefix,
	}).ParseFiles(fp)
	if err != nil {
		log.Fatalln("Failed to parse template", err)
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, cheatSheet); err != nil {
		log.Fatalln("Failed to render template", err)
	}

	return buf.String()
}

func main() {
	cheatSheet := readCheatSheet()
	renderedCheatSheet := renderCheatSheet(cheatSheet)
	fmt.Println(renderedCheatSheet)
}
