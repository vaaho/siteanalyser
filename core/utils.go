package core

import "log"

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
