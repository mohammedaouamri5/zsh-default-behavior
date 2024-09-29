package main

import (
	"fmt"
	"os"
	"reflect"
	"sync"
)


// This Function will run the command and exit 
// It look like kt just print the command 
// but if you look to the shell sqript
// you will see that the Im wraping the command with `eval` 
func Run(args ...interface{}) {
	fmt.Println(args...)
	os.Exit(0)
}

func GetAtStart(commend string) {
	var s Start
	TheValues := reflect.ValueOf(s)
	TheType := reflect.TypeOf(s)
	wg := new(sync.WaitGroup)
	wg.Add(TheType.NumMethod())
	for i := 0; i < TheType.NumMethod(); i++ {
		go func() {
			defer wg.Done()
			method := TheType.Method(i)
			args := []reflect.Value{reflect.ValueOf(commend)}
			method.Func.Call(append([]reflect.Value{TheValues}, args...))
		}()
	}
	wg.Wait()
	// if None Of The Above
	Run(commend)
}

func GetAtEnd(commend string, errcode string) {
	var s End
	TheValues := reflect.ValueOf(s)
	TheType := reflect.TypeOf(s)
	for i := 0; i < TheType.NumMethod(); i++ {
		go func() {
			method := TheType.Method(i)
			args := []reflect.Value{reflect.ValueOf(commend)}
			method.Func.Call(append([]reflect.Value{TheValues}, args...))
		}()
	}
}

func main() {
	// Check if a command was provided as an argument
	if len(os.Args) < 1 {
		fmt.Println("Error in zsh-default-behavior")
		return
	}
	if len(os.Args) == 2 {
		GetAtStart(os.Args[1])
		return
	}
	if len(os.Args) == 3 {
		GetAtEnd(os.Args[1], os.Args[2])
		return
	}
}
