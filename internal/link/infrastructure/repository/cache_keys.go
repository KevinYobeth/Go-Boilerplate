package repository

import "fmt"

func RedirectLinkKey(slug string) string {
	return fmt.Sprintf("link:redirect:%s", slug)
}
