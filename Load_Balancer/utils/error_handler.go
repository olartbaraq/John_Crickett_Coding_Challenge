package utils

import (
	"fmt"
	"log"
)

func OnPanicError(err error, errText string) {
	if err != nil {
		log.Fatalf("%v: \n, %v", errText, err)
		panic(fmt.Sprintf("%v: \n, %v", errText, err))
	}
}
