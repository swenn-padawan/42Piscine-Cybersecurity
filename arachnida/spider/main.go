package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"spider/utils"
	"strconv"
	"strings"
)

type ParseErrors int

const (
	ParsingOK ParseErrors = iota
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
	return ParsingOK, "", p
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	if !strings.Contains(string(body), "<img") {
		fmt.Printf("%s[WARN]: Spider: %.20s... Don't contain images\n%s", utils.YELLOW, p.url, utils.RESET)
		return ParsingOK
	} else {
		//TODO Parse Body response (strings.Index, strings.Contain, etc...)
		//TODO if <img> read the raw bytes of the file, and copy in the path,
		err := os.Mkdir(p.path, os.ModeDir)
		if err != nil {
			return FailedToCreatePath
		}
	}
	return ParsingOK
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s[Error]: Spider: Not enough arguments\n%s", utils.RED, utils.RESET)
	}
	err, strerr, p := parse_args(os.Args)
	if err > ParsingOK {
		switch err {
		case LevelWithoutRecurs:
			fmt.Printf("%s[Error]: Spider: Try to add -r parameter first\n%s", utils.RED, utils.RESET)
		case FailedAtoi:
			fmt.Printf("%s[Error]: Spider: %s not valid parameter, spider recurse flag need number\n%s", utils.RED, strerr, utils.RESET)
		}
		return
	}
	err = scrapping(p)
	if err > ParsingOK {
		switch err {
		case FailedToCreatePath:
			fmt.Printf("%s[Error]: Spider: Failed to create: %s\n%s", utils.RED, p.path, utils.RESET)
		}
	}
}
