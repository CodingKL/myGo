package parser

import (
	"crawler/engine"
	"regexp"
)

// input the web page contents, output the city urls and city name
// ParserResult.Requests -> []Request, which stores city url
// ParserResult.Items -> interface{}, which stores city name
func ParseCityList(contents []byte) (parseResult engine.ParserResult) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-Z]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParserResult{}
	for _, m := range matches {
		//if i >= 5 {
		//	break
		//}
		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})

	}

	return result
}
