package main

import "fmt"
import "bytes"
import "strconv"
import "github.com/PuerkitoBio/goquery"

type vote struct {
	author  string
	party   string
	yes     bool
	no      bool
	present bool
}

func numPages(posts int) (pages int) {
	fullPages := (posts / 20) + 1
	return fullPages
}

func newLink(baseLink string, offset int) (newLink string) {
	var buffer bytes.Buffer
	buffer.WriteString(baseLink)
	buffer.WriteString("&start=")
	buffer.WriteString(strconv.Itoa(offset))
	return buffer.String()
}

func runForTopic(posts int, baseLink string) {
	fullPages := numPages(posts)
	if fullPages > 1 {
		firstLink := newLink(baseLink, 0)
		exampleScrape(firstLink)

		// Now do the rest
		for i := 1; i < fullPages; i++ {
			offset := (i * 20) + 1
			thisLink := newLink(baseLink, offset)
			exampleScrape(thisLink)
		}
	} else {
		exampleScrape(baseLink)
	}

}

func exampleScrape(link string) {
	doc, err := goquery.NewDocument(link)
	if err != nil {
		fmt.Printf("Something went wrong")
	}

	doc.Find(".block-start .tablebg tr").Each(func(i int, s *goquery.Selection) {
		var skip = false
		// Find Author
		author := s.Find(".postauthor").Text()
		// Find Party
		party := s.Find(".posterrank").Text()
		// Find Post Test
		postText := s.Find(".postbody").First().Text()

		// See if anything is empty
		if author == "" || party == "" || postText == "" || len(postText) > 75 {
			skip = true
		}

		if skip == false {
			fmt.Printf("%s, %s, %s\n", author, party, postText)
		}
	})
}

func main() {
	//exampleScrape("http://usgovsim.net/USG/viewtopic.php?f=471&t=26168")
	runForTopic(38, "http://usgovsim.net/USG/viewtopic.php?f=471&t=26168")
}
