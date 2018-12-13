package main

import (
	"bufio"
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

func writeScript(input string, path string) {
	d1 := []byte(input)
	err := ioutil.WriteFile(path, d1, 0644)
	check(err)
}

func askName() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("echo name script: ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(text, "\n")
}

func writeScriptValidation(arg []string) bool {
	return arg[0] == "new"
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
	for strings.Index(script, "%"+strconv.Itoa(n)) > 0 {
		n++
	}
	return n - 1
}

func normalizeDollarSign(script string) string {
	n := 1
	for strings.Index(script, "#"+strconv.Itoa(n)) > 0 {
		script = strings.Replace(script, "#"+strconv.Itoa(n), "$"+strconv.Itoa(n), -1)
		n++
	}
	return script
}

func genScript(script string, argv []string) string {
	script = normalizeDollarSign(script)
	reqArgsLen := countArgFromScript(script)
	if reqArgsLen > len(argv) {
		return "echo not enaugh arguments , " + strconv.Itoa(reqArgsLen)
	}
	for index, arg := range argv {
		script = strings.Replace(script, "%"+strconv.Itoa(index+1), arg, -1)
	}
	return script
}

func main() {
	basePath := preparation()
	arg := os.Args[1:]

	if len(arg) == 0 {
		fmt.Print("echo Hellow A")
		return
	}

	_, secondLayerParams := arg[0], arg[1:]
	inputShiftJoin := strings.Join(secondLayerParams, " ")
	found, scipt := loadIfExist(basePath + "/" + arg[0])
	if found {
		fmt.Print(genScript(scipt, secondLayerParams))
		return
	}

	if !writeScriptValidation(arg) {
		fmt.Print("echo script not found!")
		return
	}
	name := askName()
	writeScript(inputShiftJoin, basePath+"/"+name)
	fmt.Print(name + " created!")

}
