package main

import "regexp"


// WARNING : You Have To Re-Complile when you change the config 





// make sure that That one and only ine  function will be fired
// by fired i mean will call Run
// because you can make tow default commend for the same function
// so when you make one function make sure that will check all adge cass you can think off
// or you gonna live a night mare when you start to make allot of fuction heae
// TIP : you want to return samthing as soon as posible
// if its a git repo
// return imidiatly if its not end with .git

// I wrote this couple of commend and realise how dumb and lazy i am

type Start struct{}
type End struct{}

func (Start) IsGithubRipo(commend string) {

	// Define a regex ssh for GitHub SSH repo URL
	// Example: git@github.com:username/repository.git
	ssh := `^git@github\.com:[a-zA-Z0-9._-]+/[a-zA-Z0-9._-]+\.git$`

	// Compile the regular expression
	re := regexp.MustCompile(ssh)

	if re.MatchString(commend) {
		Run("git clone --recursive ", commend)
	}
}

func (Start) IsImg(commend string) {
	img := `^(?:\./|~/|/)[^\s]+?\.(jpg|jpeg|png|gif|bmp|svg)$`
	re := regexp.MustCompile(img)
	if re.MatchString(commend) {
	Run("chafa ", commend)
	}
}
