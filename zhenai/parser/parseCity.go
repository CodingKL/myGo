package parser

import (
	"crawler/engine"
	"regexp"
)

// input the web page contents, output the city urls and city name
// ParserResult.Requests -> []Request, which stores city url
// ParserResult.Items -> interface{}, which stores city name
func ParseCity(contents []byte) (parseResult engine.ParserResult) {
	re := regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[\d]+)"[^>]*>([^>]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		name := string(m[2])
		//fmt.Printf("Get user %s, url: %s\n", m[2], m[1])

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			// Closure to input name
			// Note that 'name' has to be a copy to m[2]
			ParserFunc: func(contents []byte) engine.ParserResult {
				return parseProfile(contents, name)
			},
		})
	}

	return result
}


