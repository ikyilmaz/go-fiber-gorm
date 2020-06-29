package lib

import "log"

func CheckErr(err error, msg ...string) {
	if err != nil {
		if len(msg) > 0 {
			log.Println(msg[0])
		}
		panic(err)
	}
}
