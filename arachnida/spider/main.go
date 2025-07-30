package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"spider/utils"
	"strconv"
	"strings"
	"regexp"
	"path/filepath"
)

type ParseErrors int

const (
	OK ParseErrors = iota
	LevelWithoutRecurs
	FailedAtoi
	FailedToCreatePath
)

type params struct {
	is_recurs bool
	r_lvl     int
	path      string
	url       string
}

func parse_args(argv []string) (ParseErrors, string, params) {
	p := params{is_recurs: false, r_lvl: 5, path: "./data"}
	var err error

	for i, s := range argv {
		if i == 0 {
			continue
		}
		if s == "-l" && !p.is_recurs {
			return LevelWithoutRecurs, "", p
		}
		if s == "-l" {
			if i+1 < len(argv) {
				p.r_lvl, err = strconv.Atoi(argv[i+1])
				if err != nil {
					return FailedAtoi, argv[i+1], p
				}
			}
		}
		if s == "-r" {
			p.is_recurs = true
		}
		if s == "-p" {
			if i+1 < len(argv) {
				p.path = argv[i+1]
				i += 1
			}
		}
	}
	p.url = argv[len(argv)-1]
	fmt.Printf("%s[LOG]: Spider: Parsing OK\n%s", utils.GREEN, utils.RESET)
	return OK, "", p
}

func scrapping(p params) ParseErrors {
	client := &http.Client{
		Transport: &http.Transport{},
	}
	req, err := http.NewRequest("GET", p.url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	body := string(bodyBytes)
	re := regexp.MustCompile(`(?i)https?://[^\s'"]+\.(jpg|jpeg|png|gif|bmp)`)
	matches := re.FindAllString(body, -1)

	seen := make(map[string]bool)
	for _, imgURL := range matches{
		if seen[imgURL] {
			continue
		}
		seen[imgURL] = true
		fmt.Println("Downloading:", imgURL)

		resp, err := http.Get(imgURL)
		if err != nil{
			fmt.Println("Failed:", err)
			continue;
		}
		defer resp.Body.Close()
		err = os.Mkdir(p.path, 0775)
		if err != nil {
			return FailedToCreatePath
		}
		filename := imgURL[strings.LastIndex(imgURL, "/")+1:]
		filename = filepath.Join(p.path, filename)
		out, err := os.Create(filename)
		if err != nil{
			fmt.Println("Cannot create file:", err)
			continue;
		}
		io.Copy(out, resp.Body)
		out.Close()
	}
	return OK
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s[Error]: Spider: Not enough arguments\n%s", utils.RED, utils.RESET)
	}
	err, strerr, p := parse_args(os.Args)
	if err > OK {
		switch err {
		case LevelWithoutRecurs:
			fmt.Printf("%s[Error]: Spider: Try to add -r parameter first\n%s", utils.RED, utils.RESET)
		case FailedAtoi:
			fmt.Printf("%s[Error]: Spider: %s not valid parameter, spider recurse flag need number\n%s", utils.RED, strerr, utils.RESET)
		}
		return
	}
	err = scrapping(p)
	if err > OK {
		switch err {
		case FailedToCreatePath:
			fmt.Printf("%s[Error]: Spider: Failed to create: %s\n%s", utils.RED, p.path, utils.RESET)
		}
	}
}
