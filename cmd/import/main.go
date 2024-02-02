package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/mozillazg/go-unidecode"
)

type Record struct {
	ImgDate string
	Title   string
	Folder  string
}

func main() {
	// Open the CSV file
	file, err := os.Open("cmd/import/morfeo.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	// Define the template for the index.md file
	tmpl := `---
title: {{.Title}}
date: {{.ImgDate}}
---
{{"{{< gallery match=\"images/*\" sortOrder=\"asc\" rowHeight=\"150\" margins=\"5\" thumbnailResizeOptions=\"600x600 q90 Lanczos\" showExif=true previewType=\"blur\" embedPreview=true loadJQuery=true >}}"}}
`

	// Parse the CSV file
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		record := Record{
			ImgDate: row[0],
			Title:   row[1],
			Folder:  row[2],
		}

		// Slugify the title
		title := unidecode.Unidecode(record.Title)
		slug := strings.ToLower(strings.ReplaceAll(title, " ", "-"))

		// Create a new directory
		newDir := fmt.Sprintf("content/gallery/%s", slug)
		err = os.MkdirAll(newDir, 0755)
		if err != nil {
			panic(err)
		}

		// Create the index.md file
		indexFile, err := os.Create(fmt.Sprintf("%s/index.md", newDir))
		if err != nil {
			panic(err)
		}

		// Write the required content into the index.md file
		t := template.Must(template.New("tmpl").Parse(tmpl))
		err = t.Execute(indexFile, record)
		if err != nil {
			panic(err)
		}
		indexFile.Close()

		// Create a new directory for images
		newImagesDir := fmt.Sprintf("%s/images", newDir)
		err = os.MkdirAll(newImagesDir, 0755)
		if err != nil {
			panic(err)
		}

		// Copy the images
		files, err := ioutil.ReadDir(fmt.Sprintf("alfred-koh/images/morfeoshow/%s/big", record.Folder))
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if !file.IsDir() {
				srcFile := fmt.Sprintf("alfred-koh/images/morfeoshow/%s/big/%s", record.Folder, file.Name())
				destFile := fmt.Sprintf("%s/%s", newImagesDir, file.Name())

				input, err := ioutil.ReadFile(srcFile)
				if err != nil {
					panic(err)
				}

				err = ioutil.WriteFile(destFile, input, 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
