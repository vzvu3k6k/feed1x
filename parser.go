package feed1x

import (
	"encoding/xml"
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseLm2(body string) (*User, []*Photo, error) {
	type wrapperXML struct {
		Data string `xml:"data"`
	}

	var x wrapperXML
	err := xml.Unmarshal([]byte(body), &x)
	if err != nil {
		return nil, nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(x.Data))
	if err != nil {
		return nil, nil, err
	}

	user := &User{}
	user.Name = doc.Find(".profile-thumbs-name").First().Text()

	photos := []*Photo{}
	for _, s := range doc.Find(`[id^="imgcontainersecondary-"]`).EachIter() {
		id, ok := s.Attr("id")
		if !ok {
			return nil, nil, errors.New("id is not found")
		}
		imageID := id[22:] // "imgcontainersecondary-123" -> "123"

		src, ok := s.Find(".photos-feed-image-container, .photos-feed-image").First().Attr("src")
		if !ok {
			return nil, nil, errors.New("src is not found")
		}

		photos = append(photos, &Photo{
			Title:    s.Find(".photos-feed-data-title-container, .photos-feed-data-name").First().Text(),
			PageURL:  fmt.Sprintf("https://1x.com/photo/%s", imageID),
			ImageURL: src,
		})
	}

	return user, photos, nil
}
