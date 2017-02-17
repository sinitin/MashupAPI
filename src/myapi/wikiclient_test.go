package main

import "testing"

func TestGetWikipediaDescription(t *testing.T) {

	type TestCase struct {
		Expectation bool
		Input       []Relation
	}

	var attributeValues AttributeValues

	url1 := URL{Resource: "https://en.wikipedia.org/wiki/John_Williams", ID: "bce5b81d-8305-4a92-9424-d23e38cede33"}
	rel1 := Relation{Ended: false, TargetType: "url", Direction: "forward", Attributes: nil, URL: url1, AttributeValues: attributeValues, End: nil, Begin: nil, Type: "wikipedia", TypeID: "29651736-fa6d-48e4-aadc-a557c6add1cb", TargetCredit: "", SourceCredit: ""}
	var rellist1 []Relation
	rellist1 = append(rellist1, rel1)
	testcase1 := TestCase{Expectation: true, Input: rellist1}

	url2 := URL{Resource: "", ID: "bce5b81d-8305-4a92-9424-d23e38cede33"}
	rel2 := Relation{Ended: false, TargetType: "url", Direction: "forward", Attributes: nil, URL: url2, AttributeValues: attributeValues, End: nil, Begin: nil, Type: "wikipedia", TypeID: "29651736-fa6d-48e4-aadc-a557c6add1cb", TargetCredit: "", SourceCredit: ""}
	var rellist2 []Relation
	rellist2 = append(rellist2, rel2)
	testcase2 := TestCase{Expectation: false, Input: rellist2}

	var testlist []TestCase
	testlist = append(testlist, testcase1)
	testlist = append(testlist, testcase2)

	for _, value := range testlist {
		_, err := getWikiDescription(value.Input)
		if err != nil && (value.Expectation == true) {
			t.Errorf("Something went wrong, testcase did not expect error: %v", err)
		} else if err == nil && (value.Expectation == false) {
			t.Errorf("Something went wrong, testcase expected an error but did not get one: %v", err)
		}
	}

}
