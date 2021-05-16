package main

import (
	"fmt"
	"github.com/dirien/status-keeper/statuspage"
	"sort"
)

func main() {
	statusPages := statuspage.GetStatusPages()
	sort.Sort(statusPages)
	name := "STACKIT"
	var status *statuspage.StatusPage

	for _, t := range statusPages {
		if t.Name == name {
			status = &t
			break
		}
	}

	if status == nil {
		fmt.Errorf("cannot get status: %s", name)
	}

	fmt.Printf("Status %s\n", status.Name)

}
