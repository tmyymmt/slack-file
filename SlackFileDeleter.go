package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"log"
	"time"
)

type SlackFile struct {
	ID    string
	Title string
}

type SlackFileDeleter struct {
	files []SlackFile
}

func newInstance() SlackFileDeleter {
	return SlackFileDeleter{}
}

func (s *SlackFileDeleter) registerFile(file SlackFile) {
	s.files = append(s.files, file)
}

func (s *SlackFileDeleter) delete(api *slack.Client) {
	for _, targetFile := range s.files {
		s.deleteImpl(api, targetFile)
		// https://api.slack.com/docs/rate-limits
		// https://api.slack.com/methods/files.delete
		// 50/sec is the limit of slack api 'files.delete'.
		// 20 * time.Millisecond is the maximum.
		// +5 is a margin.
		time.Sleep(25 * time.Millisecond)
	}
}

func (s *SlackFileDeleter) deleteImpl(api *slack.Client, targetFile SlackFile) {
	if err := api.DeleteFile(targetFile.ID); err == nil {
		fmt.Println("Deleted:", targetFile.ID, quote(targetFile.Title))
	} else {
		log.Println(err)
	}
}
