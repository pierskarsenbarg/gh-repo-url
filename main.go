package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/repository"
)

func main() {
	repo, err := repository.Current()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("https://%s/%s/%s", repo.Host, repo.Owner, repo.Name)
	}
}
