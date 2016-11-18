/*
 *   Copyright 2016 Rémy MATHIEU
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package content

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RedditContent struct {
	Url string
}

var userAgents []string = []string{
	"Mozilla/5.0 (X11; Linux x86_64; rv:42.0) Gecko/20100101 Firefox/42.0",
	"Mozilla/5.0 (X11; Linux x86_64; rv:43.0) Gecko/20100101 Firefox/43.0",
	"Mozilla/5.0 (Linux; Android 5.1; XT1039 Build/LPB23.13-17.6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.95 Mobile Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.103 Safari/537.36",
}

func (reddit RedditContent) callAPI() ([]Content, error) {

	resp, err := reddit.request()
	if err != nil {
		fmt.Println("Error while calling API")
		return nil, err
	}

	doc, err := goquery.NewDocumentFromResponse(resp)

	rv := make([]Content, 0)

	doc.Find(".link").Each(func(i int, selec *goquery.Selection) {
		// ignore sticky posts
		if selec.HasClass("stickied") {
			return
		}

		if len(rv) > 20 {
			return
		}

		l := selec.Find("p.title a.title")
		title := l.First().Text()
		externalLink, _ := l.Attr("href")

		// ignore self posts
		if strings.HasPrefix(externalLink, "/r/") {
			return
		}

		rv = append(rv, Content{
			Text: title,
			Url:  externalLink,
		})
	})

	return rv, nil
}

func (r RedditContent) request() (*http.Response, error) {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", randomUseragent())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("while querying", r.Url, ":", resp.Status)
	}

	return resp, nil
}

func randomUseragent() string {
	return userAgents[rand.Int()%len(userAgents)]
}
