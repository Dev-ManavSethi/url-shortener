package utils

import "log"

func HandleErr(err error, ErrMessage, SuccessMessage string) {
	if err != nil {
		log.Println(ErrMessage)
		log.Fatalln(err)
	} else if SuccessMessage != "" {
		log.Println(SuccessMessage)
	}
}
