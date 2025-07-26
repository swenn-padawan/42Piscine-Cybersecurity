package main

import (
	"spider/utils"
	"strconv"
	"fmt"
	"net/http"
	"os"
)

type params struct{
	is_recurs	bool
	r_lvl   	int
	path  		string
	url			string
}

func parse_args(argv []string)(bool, string, params){
	p := params{is_recurs: false, r_lvl: 5, path: "./data"}
	var err error

	for i, s:= range argv{
		if i == 0 {continue}
		if s == "-l" {
			if i+1 < len(argv){
				p.r_lvl, err = strconv.Atoi(argv[i + 1])
				if err != nil {return true, argv[i + 1], p}
			}
		}
		if s == "-r" {p.is_recurs = true}
		if s == "-p" {
			if i+1 < len(argv){
				p.path = argv[i + 1]
				i += 1
			}
		}
	}
	p.url = argv[len(argv) - 1]
	fmt.Printf("%s[LOG]: Spider: Parsing OK\n%s", utils.GREEN, utils.RESET)
	return false, "", p
}

func scrapping(p params)(bool){
	fmt.Println(p)
	resp, err := http.Get(p.url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response Status:", resp.Status)
	return true
}

func main(){
	if len(os.Args) < 2 {
		fmt.Printf("%s[Error]: Spider: Not enough arguments\n%s", utils.RED, utils.RESET)
	}
	err, strerr, p := parse_args(os.Args)
	if err {
		fmt.Printf("%s[Error]: Spider: %s wrong parameter\n%s", utils.RED, strerr, utils.RESET)
	}
	err = scrapping(p)
	if err {
		fmt.Printf("%s[Error]: Spider: Something went wrong with the scrapping\n%s", utils.RED, utils.RESET)
	}
}
