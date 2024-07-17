package main

import (
	"context"
	"github.com/go-git/go-git/v5"
	"golang.design/x/clipboard"
	"log"
	"os"
	"time"
)

var repoPath = "./repo"
var decryptedPath = "./repo"

func main() {
	github_url := os.Getenv("GITHUB_URL")

	if github_url == "" {
		log.Fatalln("please enter all required ENVs")
		return
	}
	var repo *git.Repository
	repo, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:      "",
		Progress: os.Stdout,
	})
	if err != nil {
		if err == git.ErrRepositoryAlreadyExists {
			// pull updates
			repo, err = git.PlainOpen(repoPath)
			if err != nil {
				panic(err)
			}
			w, err := repo.Worktree()
			if err != nil {
				panic(err)
			}
			err = w.Pull(&git.PullOptions{RemoteName: "origin"})
		} else {
			panic(err)
		}
	}

	err = clipboard.Init()
	if err != nil {
		panic(err)
	}

	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	for data := range ch {
		// print out clipboard data whenever it is changed
		println(string(data))
		println(time.Now().Format("[2006-01-02T15:04:05Z]"))
	}
}
