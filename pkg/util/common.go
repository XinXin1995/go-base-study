package util

import (
	"github.com/astaxie/beego/validation"
	"log"
)

func LoopLog(errors []*validation.Error) {
	for _, err := range errors {
		log.Println(err.Key, err.Message)
	}
}
