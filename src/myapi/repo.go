package main

import "fmt"

var m map[string]Summary

func init() {
	m = make(map[string]Summary)
}

func RepoCheckMbid(mbid string) error {

	if _, ok := m[mbid]; ok {
		return nil
	}
	return fmt.Errorf("Mbid does not exist in repo")
}

func RepoGetSummary(mbid string) Summary {
	summary, _ := m[mbid]
	return summary
}

func RepoAddSummary(s Summary) {

	if _, ok := m[s.Mbid]; !ok {
		m[s.Mbid] = s
	}
}
