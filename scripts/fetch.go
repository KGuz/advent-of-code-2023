package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

//go:embed .env
var fs embed.FS

func getenv(name string) string {
	file, err := fs.Open(".env")
	if err != nil {
		log.Fatalln("have you tried turning it off and on again?")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, name) {
			_, val, _ := strings.Cut(line, "=")
			return strings.Trim(val, "\"")
		}
	}
	return ""
}

func fetch(url string, session string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("cookie", session)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func example(resp string) string {
	s := strings.Index(resp, "<pre><code>") + len("<pre><code>")
	e := strings.Index(resp, "</code></pre>")
	return resp[s:e]
}

func description(resp string) string {
	article := `<article class="day-desc">`
	s := strings.Index(resp, article) + len(article)
	e := strings.Index(resp, "</article>")
	desc := resp[s:e]

	resp = resp[e+len("</article>"):]
	s = strings.Index(resp, article) + len(article)
	e = strings.Index(resp, "</article>")
	desc = desc + "\n" + resp[s:e]

	desc = strings.ReplaceAll(desc, "<p>", "<p>\n")
	return stripHtmlTags(desc)
}

func stripHtmlTags(s string) string {
	var builder strings.Builder
	builder.Grow(len(s) + utf8.UTFMax)

	in := false
	start := 0
	end := 0

	for i, c := range s {
		if (i+1) == len(s) && end >= start {
			builder.WriteString(s[end:])
		}

		if c != '<' && c != '>' {
			continue
		}

		if c == '<' {
			if !in {
				start = i
			}
			in = true

			builder.WriteString(s[end:start])
			continue
		}
		in = false
		end = i + 1
	}
	s = builder.String()
	return s
}

func main() {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d", 2022, 1)
	input := fetch(url, getenv("AOC_TOKEN"))

	println(example(input))
	println(description(input))
}
