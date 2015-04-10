package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gosexy/exif"
)

func expandPath(path string) string {
	home := os.Getenv("HOME")
	return strings.Replace(path, "~", home, 1)
}

func scanImageFiles(path string) ([]string, error) {
	results := make([]string, 0)
	basePath := expandPath(path)

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return results, err
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := filepath.Ext(file.Name())
			if ext == ".JPG" {
				results = append(results, basePath+"/"+file.Name())
			}
		}
	}

	return results, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide a path")
		os.Exit(1)
	}

	path := expandPath(os.Args[1])
	out, err := exec.Command("find", path, "-name", "*.JPG").Output()
	if err != nil {
		fmt.Println("Cant find any images")
		return
	}

	files := strings.Split(strings.TrimSpace(string(out)), "\n")

	exifr := exif.New()

	for _, path := range files {
		err = exifr.Open(path)

		x, _ := strconv.Atoi(exifr.Tags["Pixel X Dimension"])
		y, _ := strconv.Atoi(exifr.Tags["Pixel Y Dimension"])

		if x/y > 2.0 {
			fmt.Println(path)
		}
	}
}
