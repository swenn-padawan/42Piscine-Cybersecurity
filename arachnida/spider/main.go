package main

import (
	"spider/utils"
	"fmt"
	"os"
)

func parse_args(argv []string)(bool, string){
	for i, s:= range argv{
		if i == 0 {continue}
		fmt.Printf("%s\n", s)
	}
	return false, ""
}

func main(){
	if len(os.Args) < 2 {
		fmt.Printf("%s[Error]: Spider: Not enough arguments\n", utils.RED)
	}
	err, strerr := parse_args(os.Args)
	if err {
		fmt.Printf("%s[Error]: Spider: %s wrong parameter\n", utils.RED, strerr)
	}
}
