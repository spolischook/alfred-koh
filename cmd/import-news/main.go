package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Title     string
	Alias     string
	Introtext string
	Fulltext  string
	Created   string
}

func main() {
	// Connect to the MySQL database
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3406)/alfred_koh")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT title, alias, introtext, `fulltext`, created FROM ako_content WHERE catid=1")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()

	var articles []Article

	// Fetch the results
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Title, &article.Alias, &article.Introtext, &article.Fulltext, &article.Created); err != nil {
			fmt.Println("Error:", err)
			return
		}
		articles = append(articles, article)
	}

	// Iterate over the articles
	for _, article := range articles {
		// Create a new directory for each article
		newDir := filepath.Join("content/news-archive", article.Alias)
		os.MkdirAll(newDir, os.ModePerm)

		// Create a new file in the new directory
		f, err := os.Create(filepath.Join(newDir, "index.md"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// Parse the date
		t, err := time.Parse("2006-01-02 15:04:05", article.Created)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		intro := strings.ReplaceAll(article.Introtext, `src="images`, `src="/images`)
		full := strings.ReplaceAll(article.Fulltext, `src="images`, `src="/images`)
		intro = strings.ReplaceAll(intro, `/ z_main.gif`, `/z_main.gif`)
		full = strings.ReplaceAll(full, `/ z_main.gif`, `/z_main.gif`)
		intro = strings.ReplaceAll(intro, `http://lh6.ggpht.com/_Jv_-CaL0c2w/S46_i_icNQI/AAAAAAAAAd4/XgYqipNuAAQ/s400/Kharkov_bULLON.jpg`, `/images/Kharkov_bULLON.jpg`)
		full = strings.ReplaceAll(full, `http://lh6.ggpht.com/_Jv_-CaL0c2w/S46_i_icNQI/AAAAAAAAAd4/XgYqipNuAAQ/s400/Kharkov_bULLON.jpg`, `/images/Kharkov_bULLON.jpg`)

		// Write the required content into the file
		_, err = f.WriteString(fmt.Sprintf(`---
title: "%s"
date: "%s"
---
{{< rawhtml >}}
%s
%s
{{< /rawhtml >}}
`, strings.ReplaceAll(article.Title, `"`, `\"`), t.Format("2006-01-02 15:04:05"), intro, full))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		f.Close()
	}
}
