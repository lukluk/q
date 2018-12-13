package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func preparation() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	HOME := usr.HomeDir + "/.a"
	if _, err := os.Stat(HOME); os.IsNotExist(err) {
		os.Mkdir(HOME, os.ModePerm)
	}
	return HOME
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func loadIfExist(path string) (bool, string) {
	if exists(path) {
		dat, err := ioutil.ReadFile(path)
		check(err)
		return true, string(dat)
	}
	return false, ""
}

func countArgFromScript(script string) int {
	n := 1
	for strings.Index(script, "$"+strconv.Itoa(n)) > 0 {
		n++
	}
	return n - 1
}

func validate(path string, script string, argv []string) string {
	reqArgsLen := countArgFromScript(script)
	if reqArgsLen > len(argv) {
		return "echo not enaugh arguments ,  " + strconv.Itoa(reqArgsLen)
	}
	return "bash " + path + " " + strings.Join(argv, " ")
}

func main() {
	basePath := preparation()
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Print("echo Hellow A")
		return
	}

	command, secondLayerParams := arg[0], arg[1:]

	if command == "new" && arg[1] != "" {
		fmt.Println("vim " + basePath + "/" + arg[1])
		return
	}
	found, scipt := loadIfExist(basePath + "/" + arg[0])
	if found {
		fmt.Print(validate(basePath+"/"+arg[0], scipt, secondLayerParams))
		return
	}
	fmt.Print("echo script not found!")

}
