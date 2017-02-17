package main

import (
	"regexp"
	"fmt"
)

func checkOkMbid(mbid string) (ok bool, err error) {

	fmt.Println("     ")
	fmt.Printf("the mbid I got was %s", mbid)
	fmt.Println("     ")

	if len(mbid)!=36{
		return false, fmt.Errorf("Submitted mbid has the wrong format")
	}

	ok, err = regexp.MatchString("([0-9a-z]{8})-([0-9a-z]{4})-([0-9a-z]{4})-([0-9a-z]{4})-([0-9a-z]{12})$", mbid)
	return

}
