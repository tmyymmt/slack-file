package main

import (
	"github.com/slack-go/slack"
	"log"
	"sync"
)

type SlackFile struct {
	ID    string
	Title string
}

type SlackFileDeleter struct {
	files     []SlackFile
	waitGroup sync.WaitGroup
}

func newInstance() SlackFileDeleter {
	return SlackFileDeleter{
		waitGroup: sync.WaitGroup{},
	}
}

func (s *SlackFileDeleter) registerFile(file SlackFile) {
	s.files = append(s.files, file)
}

func (s *SlackFileDeleter) delete(api *slack.Client) {
	for _, targetFile := range s.files {
		s.waitGroup.Add(1)
		go s.deleteImpl(api, targetFile)
	}
	s.waitGroup.Wait()
}

func (s *SlackFileDeleter) deleteImpl(api *slack.Client, targetFile SlackFile) {
	defer s.waitGroup.Done()
	if err := api.DeleteFile(targetFile.ID); err == nil {
		log.Printf("Deleted %s %s at Slack.\n", targetFile.ID, targetFile.Title)
	} else {
		log.Println(err)
	}
}
