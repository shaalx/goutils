package goutils

import (
	"fmt"
	"log"
)

func CheckErr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func CheckNoLogErr(err error) bool {
	if nil != err {
		return true
	}
	return false
}

func LogCheckErr(err error) bool {
	if nil != err {
		log.Println(err)
		return true
	}
	return false
}

func Log(v ...interface{}) {
	log.Print(v...)
}

func Print(v ...interface{}) {
	fmt.Print(v...)
}
