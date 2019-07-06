package util

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func LoopLog(errors []*validation.Error) {
	for _, err := range errors {
		log.Fatalln(err.Key, err.Message)
	}
}
