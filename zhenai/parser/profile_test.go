package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestProfile(t *testing.T) {
	body, err := ioutil.ReadFile("user_test_data.html")
	if err != nil {
		panic(err)
	}
	result := parseProfile(body, "name")
	fmt.Printf("%v", result.Items)
}
