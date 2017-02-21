package main

import (
	"fmt"
	"regexp"
)

func checkOkMbid(mbid string) (ok bool, err error) {

	if len(mbid) != 36 {
		ok = false
		err = fmt.Errorf("Submitted mbid has the wrong format")
		return
	}

	//checks that the submitted mbid is in a valid uuid format
	ok, err = regexp.MatchString("([0-9a-z]{8})-([0-9a-z]{4})-([0-9a-z]{4})-([0-9a-z]{4})-([0-9a-z]{12})$", mbid)
	return

}
