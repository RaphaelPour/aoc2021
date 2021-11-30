// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/fatih/color"

	"github.com/MarkusFreitag/advent-of-code/util"
)

const (
	YEAR = 2015
	DAY  = 3
)

func main() {
	fmt.Println("HELLO")
	//day := time.Now().Day()
	day := DAY

	/* create day folder */
	folder := fmt.Sprintf("%02d", day)
	path := filepath.Join("days", folder)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		color.Red("%s already exists", path)
		return
	}

	/* create folder */
	if err := os.MkdirAll(path, 0777); err != nil {
		color.Red("error creating path %s: %s", path, err)
		return
	}

	/* create part.go from template */
	templates, err := template.ParseGlob("templates/*.tpl")
	if err != nil {
		color.Red("error initializing templates: %s", err)
		return
	}

	fd, err := os.OpenFile(
		filepath.Join(path, "part.go"),
		os.O_RDWR|os.O_CREATE,
		0777,
	)
	if err != nil {
		color.Red(
			"error creating part file %s: %s",
			filepath.Join(path, "part.go"),
			err,
		)
		return
	}

	if err := templates.ExecuteTemplate(fd, "part.tpl", nil); err != nil {
		color.Red("error executing part template: %s", err)
		return
	}

	if err := fd.Close(); err != nil {
		color.Red("error writing part file: %s", err)
		return
	}

	/* download input */
	input, err := util.InputFromURL(YEAR, day)
	if err != nil {
		color.Red("error getting input: %s", err)
		return
	}

	err = ioutil.WriteFile(filepath.Join(path, "input"), []byte(input), 0777)
	if err != nil {
		color.Red("error writing input to file: %s", err)
		return
	}

	color.Green("Day %d is ready to solve :)", day)
}
