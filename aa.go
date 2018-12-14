package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"
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

func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str, end)
	if e < 0 {
		return ""
	}
	return str[s:e]
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

func parseHumanParam(str string) (lenParam int, params []string) {
	for found := "init"; found != ""; found = GetStringInBetween(str, "<", ">") {
		if found == "init" {
			continue
		}
		lenParam++
		params = append(params, found)
		str = strings.Replace(str, "<"+found+">", "", -1)
	}
	return
}

func validate(name string, path string, script string, argv []string) string {
	reqArgsLen := countArgFromScript(script)
	lenHParams, hParams := parseHumanParam(script)
	if reqArgsLen > len(argv) {
		return "echo not enaugh arguments ,  " + strconv.Itoa(reqArgsLen)
	} else if lenHParams > len(argv) {
		return "echo q " + name + " " + strings.Join(hParams, "  ")
	} else if lenHParams == len(argv) {
		for i, param := range hParams {
			script = strings.Replace(script, "<"+param+">", "$"+strconv.Itoa(i+1), -1)
		}
		d1 := []byte(script)
		compiledPath := path + "/." + name
		err := ioutil.WriteFile(compiledPath, d1, 0644)
		check(err)
		return "bash " + compiledPath + " " + strings.Join(argv, " ")
	}

	return "bash " + path + "/" + name + " " + strings.Join(argv, " ")
}

func getArgs(index int) string {
	if index < len(os.Args[1:]) {
		return os.Args[1:][index]
	} else {
		return ""
	}
}

func main() {
	basePath := preparation()
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Print("echo 🥒 && ls -t " + basePath)
		return
	}

	command, secondLayerParams := arg[0], arg[1:]
	if command == "repo" {
		if getArgs(1) == "" {
			fmt.Println("echo q repo {git uri}")
			return
		}
		fmt.Println("cd " + basePath + " && git init && git remote add origin " + getArgs(1))
		return
	}
	if command == "pull" {
		if !exists(basePath + "/.git") {
			fmt.Println("echo q repo {git uri}")
			return
		}
		fmt.Println("cd " + basePath + " && git pull origin master")
	}
	if command == "push" {
		if !exists(basePath + "/.git") {
			fmt.Println("echo q repo {git uri}")
			return
		}
		message := getArgs(1)
		if message == "" {
			t := time.Now()
			message = "update at " + t.Format(time.RFC3339)
		}
		fmt.Println("cd " + basePath + " && git add --all && git commit -m \"" + message + "\" && git pull -r origin master && git push git push --set-upstream origin master")
	}
	if command == "new" {
		if getArgs(1) == "" {
			fmt.Println("echo q new {script-name}")
			return
		}
		fmt.Println("vim " + basePath + "/" + getArgs(1))
		return
	}
	found, scipt := loadIfExist(basePath + "/" + arg[0])
	if found {
		fmt.Print(validate(arg[0], basePath, scipt, secondLayerParams))
		return
	}
	fmt.Print("echo script not found!")

}
