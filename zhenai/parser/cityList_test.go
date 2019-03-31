package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("city_test_data.html")

	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 470

	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but %d got",
			resultSize, len(result.Requests))
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items, but %d got",
			resultSize, len(result.Items))
	}

	exceptedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	exceptedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}

	for i, url := range exceptedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("excepted url #%d: %s, but got %s",
				i, url, result.Requests[i].Url)
		}
	}

	for i, city := range exceptedCities {
		if result.Items[i] != city {
			t.Errorf("excepted city #%d: %s, but got %s",
				i, city, result.Items[i])
		}
	}
}
