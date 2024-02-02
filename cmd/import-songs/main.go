package main

import (
	"encoding/csv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Song struct {
	Title    string
	Alias    string
	Fulltext string
	Created  time.Time
}

func main() {
	file, err := os.Open("cmd/import-songs/songs.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','
	reader.LazyQuotes = true

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		panic(err)
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		created, err := time.Parse("2006-01-02 15:04:05", record[3])
		if err != nil {
			panic(err)
		}

		song := Song{
			Title:    record[0],
			Alias:    record[1],
			Fulltext: record[2],
			Created:  created,
		}

		dirPath := filepath.Join("content", "songs", song.Alias)
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}

		filePath := filepath.Join(dirPath, "index.md")
		f, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString("---\n")
		_, err = f.WriteString("title: \"" + song.Title + "\"\n")
		_, err = f.WriteString("slug: \"" + song.Alias + "\"\n")
		_, err = f.WriteString("date: \"" + song.Created.Format(time.RFC3339) + "\"\n")
		_, err = f.WriteString("---\n")
		_, err = f.WriteString("{{< rawhtml >}}\n")
		_, err = f.WriteString(strings.ReplaceAll(song.Fulltext, "\"", "\\\"") + "\n")
		_, err = f.WriteString("{{< /rawhtml >}}\n")

		if err != nil {
			f.Close()
			panic(err)
		}

		err = f.Close()
		if err != nil {
			panic(err)
		}
	}
}
