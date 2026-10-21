// Package main validates that every cheat YAML file still parses and reports
// how many titles/descs contain CJK characters (a proxy for "translated").
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"gopkg.in/yaml.v3"

	"github.com/halilkirazkaya/arsenal-ng/internal/model"
)

func main() {
	root := "internal/loader/cheat-files"
	totalFiles := 0
	totalActions := 0
	totalErrors := 0
	cjkDesc := 0
	cjkTitle := 0

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".yaml" && ext != ".yml" {
			return nil
		}
		totalFiles++

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("READ ERROR %s: %v\n", path, err)
			totalErrors++
			return nil
		}

		var cf model.CheatFile
		if err := yaml.Unmarshal(data, &cf); err != nil {
			fmt.Printf("PARSE ERROR %s: %v\n", path, err)
			totalErrors++
			return nil
		}

		for _, a := range cf.Actions {
			totalActions++
			if containsCJK(a.Desc) {
				cjkDesc++
			}
			if containsCJK(a.Title) {
				cjkTitle++
			}
			if strings.TrimSpace(a.Command) == "" {
				fmt.Printf("EMPTY COMMAND %s (title: %s)\n", path, a.Title)
				totalErrors++
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("WALK ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("files=%d actions=%d parse_errors=%d cjk_title=%d cjk_desc=%d\n",
		totalFiles, totalActions, totalErrors, cjkTitle, cjkDesc)
	if totalErrors > 0 {
		os.Exit(1)
	}
}

func containsCJK(s string) bool {
	for _, r := range s {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}
