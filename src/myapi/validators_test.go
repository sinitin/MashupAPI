package main

import "testing"

func TestCheckMbid(t *testing.T) {

	testmap := make(map[string]bool)

	testmap["02f7d891-5e88-4c3a-a5fd-eba599904e61"] = true
	testmap["02f7d891-5e88-4c3a-a5fd-eba599904e6"] = false
	testmap[""] = false
	testmap["02f7d891-5e88-4c3a-a5fd-eba599904e61-"] = false
	testmap["02f7d891-5e88-4c3a-a5fd-Eba599904e61"] = false
	testmap["02f7d891-5e88-4c3a-a5fdeba599904e61"] = false
	testmap["f-5-4-a"] = false

	for uuid, expected := range testmap {
		v, _ := checkOkMbid(uuid)
		if v != expected {
			t.Error(
				"For", uuid,
				"expected", expected,
				"got", v,
			)
		}
	}
}
