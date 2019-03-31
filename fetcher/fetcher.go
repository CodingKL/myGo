package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
  input a url, fetch the contents of that page
*/
var rateLimit = time.Tick(10 * time.Millisecond)
func Fetch(url string) (contents []byte, err error) {
	//<-rateLimit
	request, _ := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.67 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if err != nil {
			return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
		}
	}
	return ioutil.ReadAll(resp.Body)
}
