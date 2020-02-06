package main

import "log"

func errcheck(err error) {
	if err != nil {
		log.Print(err)
	}


func usercheck(s1, s2 string) bool {
		return s1 == s2
	}
	
}
