package linter

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LintFiles(dir string) error {
	var err error
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	files, err := listFiles(dir)
	if err != nil {
		return errors.New(fmt.Sprintf("Listing files failed: ", err.Error()))
	}

	for _, f := range files {
		err = os.Rename(filepath.Join(dir, f), filepath.Join(dir, lint(f)))
		if err != nil {
			return err
		}
	}
	return nil
}

func lint(name string) string {
	lintedName := strings.Replace(name, " ", "-", -1)
	return lintedName
}

func listFiles(dir string) ([]string, error) {
	var result []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return result, err
	}

	for _, f := range files {
		result = append(result, f.Name())
	}

	return result, nil
}
