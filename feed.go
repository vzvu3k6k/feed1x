package feed1x

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
)

type User struct {
	Name string
}

type Photo struct {
	Title    string
	PageURL  string
	ImageURL string
}

func GetFeed(ctx context.Context, userID string) (string, error) {
	lm2, err := getLm2(ctx, userID)
	if err != nil {
		return "", err
	}

	user, photos, err := parseLm2(lm2)
	if err != nil {
		return "", err
	}

	atom, err := buildFeed(user, photos)
	if err != nil {
		return "", err
	}
	return atom, nil
}

func getLm2(_ context.Context, userID string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://1x.com/backend/lm2.php?style=normal&mode=user:%s::&from=0&autoload=&search=&alreadyloaded=", userID))
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func buildFeed(user *User, photos []*Photo) (string, error) {
	now := time.Now()
	feed := &feeds.Feed{
		Title:   fmt.Sprintf("1x: %s", user.Name),
		Created: now,
	}

	for _, photo := range photos {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       photo.Title,
			Id:          photo.PageURL,
			Link:        &feeds.Link{Href: photo.PageURL},
			Description: fmt.Sprintf(`<img src="%s" />`, photo.ImageURL),
		})
	}

	atom, err := feed.ToAtom()
	if err != nil {
		return "", err
	}

	return atom, nil
}
