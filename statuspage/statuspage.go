package statuspage

import "fmt"

type StatusPage struct {
	Name    string
	Url     string
	Version string
}

func GetStatusPageComponentURL(statusPage *StatusPage) string {
	return fmt.Sprintf(
		"https:/%s/api/%s/components.json",
		statusPage.Url, statusPage.Version)
}
