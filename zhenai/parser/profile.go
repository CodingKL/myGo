package parser

import (
	"crawler/engine"
	"crawler/model"
	"log"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([0-9]+)岁</div>`)
var workPlaceRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>工作地:([^<]+)</div>`)
var incomeRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>月收入:([^<]+)</div>`)
var weightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)kg</div>`)
var heightRe = regexp.MustCompile(`<div class="m-btn purple"[^>]*>([\d]+)cm</div>`)
//<div class="m-btn purple" data-v-bff6f798="">170cm</div>

func parseProfile(contents []byte, name string) engine.ParserResult {
	profile := model.Profile{}

	profile.Name = name
	profile.Age = extractInt(contents, ageRe)
	profile.Weight = extractInt(contents, weightRe)
	profile.Height = extractInt(contents, heightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.WorkPlace = extractString(contents, workPlaceRe)

	//log.Printf("Got user info: %v", profile)

	return engine.ParserResult{
		Requests: nil,
		Items:    []interface{}{profile},
	}
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	match := re.FindSubmatch(contents)
	//fmt.Printf("%s\n\n\n", contents)
	//fmt.Printf("extractInt: %s\n", match)
	if match != nil && len(match) >= 2 {
		n, err := strconv.Atoi(string(match[1]))
		if err != nil {
			log.Printf("extractInt: convert to int failed: %s", match[1])
			return -1
		}
		return n
	}
	return -1
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	//fmt.Printf("extractString: %s\n", match)
	if match != nil && len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
