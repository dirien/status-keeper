package statuspage

import "strings"

type StatusPages []StatusPage

func (t StatusPages) Len() int { return len(t) }

func (t StatusPages) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func (t StatusPages) Less(i, j int) bool {
	var ti = t[i]
	var tj = t[j]
	var tiNameLower = strings.ToLower(ti.Name)
	var tjNameLower = strings.ToLower(tj.Name)
	if tiNameLower == tjNameLower {
		return ti.Name < tj.Name
	}
	return tiNameLower < tjNameLower
}

func GetStatusPages() StatusPages {

	statusPages := []StatusPage{}
	statusPages = append(statusPages,

		StatusPage{
			Name:    "STACKIT",
			Url:     "status.stackit.cloud",
			Version: "v2",
		},
		StatusPage{
			Name:    "jFrog",
			Url:     "status.jfrog.io",
			Version: "v2",
		},
		StatusPage{
			Name:    "github",
			Url:     "www.githubstatus.com",
			Version: "v2",
		},
		StatusPage{
			Name:    "digitalocean",
			Url:     "status.digitalocean.com",
			Version: "v2",
		},
		StatusPage{
			Name:    "dropbox",
			Url:     "status.dropbox.com",
			Version: "v2",
		},
		StatusPage{
			Name:    "reddit",
			Url:     "www.redditstatus.com",
			Version: "v2",
		},
		StatusPage{
			Name:    "scaleway",
			Url:     "status.scaleway.com",
			Version: "v2",
		},
	)

	return statusPages
}
