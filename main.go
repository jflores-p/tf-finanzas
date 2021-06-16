package main

import (
	"flag"
)

var (
	flagAction     string
	flagUser       string
	flagRepo       string
	flagCreateRepo bool
	flagDeleteRepo bool
)

func init() {
	flag.StringVar(&flagAction, "a", "", "Specify the option\nremote: get html/ssh repo url")
	flag.StringVar(&flagUser, "u", "jflores-p", "User of github")
	flag.StringVar(&flagRepo, "r", "", "Repository name")
	flag.BoolVar(&flagCreateRepo, "c", false, "Triggers create repo")
	flag.BoolVar(&flagDeleteRepo, "d", false, "Triggers delete repo")

	flag.Parse()
}

func main() {

}
