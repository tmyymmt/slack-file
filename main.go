package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/namsral/flag"
	"github.com/slack-go/slack"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const envFile = ".env"

// These variables are set in build step
var (
	Version  = "unset"
	Revision = "unset"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if exists(envFile) {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatal("Error loading .env file. ", err)
		}
	}

	slackToken := flag.String("token", "", "Your slack api token")
	doShow := flag.Bool("show", true, "Show slack files")
	doDownload := flag.Bool("download", false, "Download slack files")
	doDelete := flag.Bool("delete", false, "Download file from slack")
	channelId := flag.String("channel-id", "", "Filter files by channel id")
	excludeChannelIds := flag.String("exclude-channel-ids", "", "Filter files by excluded channel ids")
	fileType := flag.String("type", "all", "Filter files by type")
	beforeTimestamp := flag.Int64("before-timestamp", 0, "Filter files by before the timestamp")
	beforeDays := flag.Int("before-days", 0, "Filter files by more than ? days old")
	beforeEndOfMonth := flag.Int("before-end-of-month", 0, "Filter files by more than the end of ? month(s) ago")
	to := flag.String("to", "downloads", "Download slack files to specified download folder")
	toWidthDate := flag.Bool("to-with-date", false, "Add date info to download folder name")
	toWidthChannels := flag.Bool("to-with-channels", false, "Add channels folder at download folder")
	flag.Parse()

	if *slackToken == "" {
		log.Fatal("Error token didn't set. Please set your slack api token.")
	}

	now := time.Now()
	beforeTimestampResult := now
	if *beforeTimestamp != 0 {
		beforeTimestampResult = time.Unix(*beforeTimestamp, 0)
	}
	beforeDaysResult := now
	if *beforeDays != 0 {
		beforeDaysResult = time.Date(now.Year(), now.Month(), now.Day()-(*beforeDays), 0, 0, 0, 0, time.Local).AddDate(0, 0, 1).Add(time.Duration(1) * time.Nanosecond * -1)
	}
	beforeEndOfMonthResult := now
	if *beforeEndOfMonth != 0 {
		beforeEndOfMonthResult = time.Date(now.Year(), now.Month()-time.Month(*beforeEndOfMonth), 1, 0, 0, 0, 0, time.Local).AddDate(0, 1, 0).Add(time.Duration(1) * time.Nanosecond * -1)
	}
	beforeResult := int64(math.Min(math.Min(float64(beforeTimestampResult.Unix()), float64(beforeDaysResult.Unix())), float64(beforeEndOfMonthResult.Unix())))

	excludeChannelIdsResult := []string{}
	if *excludeChannelIds != "" {
		excludeChannelIdsResult = strings.Split(*excludeChannelIds, ",")
	}

	toResult := *to
	if *toWidthDate {
		toResult = toResult + "_" + now.Format("20060102")
	}
	if *doDownload && !exists(toResult) {
		os.Mkdir(toResult, 0777)
	}

	api := slack.New(*slackToken)

	files, paging, err := getFiles(api, *fileType, beforeResult, *channelId, 1)
	if err != nil {
		log.Fatal(err)
	}

	var slackFileDeleter SlackFileDeleter
	if *doDelete {
		slackFileDeleter = newInstance()
	}
	waitGroup := sync.WaitGroup{}
	for paging.Page <= paging.Pages {
		for _, slackFile := range files {
			exclude := false
			for _, excludeChannelId := range excludeChannelIdsResult {
				if contains(slackFile.Channels, excludeChannelId) {
					exclude = true
					continue
				}
			}
			if exclude {
				continue
			}

			if *doShow {
				fmt.Println(
					slackFile.ID,
					quote(slackFile.Channels),
					quote(slackFile.Title),
					slackFile.Size,
					quote(time.Unix(int64(slackFile.Created), 0).Format("2006-01-02 15:04:05 MST")),
					slackFile.URLPrivateDownload)
			}
			if *doDownload {
				waitGroup.Add(1)
				go download(&waitGroup, slackFile, *slackToken, toResult, *toWidthChannels)
			}
			if *doDelete {
				deleteFile := SlackFile{slackFile.ID, slackFile.Title}
				slackFileDeleter.registerFile(deleteFile)
			}
		}
		files, paging, err = getFiles(api, *fileType, *beforeTimestamp, *channelId, paging.Page+1)
		if err != nil {
			log.Fatal(err)
		}
	}
	waitGroup.Wait()

	if *doDelete {
		slackFileDeleter.delete(api)
	}
}

func quote(text interface{}) string {
	return fmt.Sprintf("\"%v\"", text)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func contains(array []string, value string) bool {
	for _, one := range array {
		if one == value {
			return true
		}
	}
	return false
}

func download(waitGroup *sync.WaitGroup, slackFile slack.File, slackToken string, downloadFolder string, widthChannels bool) {
	defer waitGroup.Done()
	req, err := http.NewRequest("GET", slackFile.URLPrivateDownload, nil)
	req.Header.Set("Authorization", "Bearer "+slackToken)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Println(err, slackFile.URLPrivateDownload)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err, slackFile.URLPrivateDownload)
		return
	}

	downloadFolderResult := downloadFolder
	if widthChannels {
		channelsResult := strings.Join(slackFile.Channels, "_")
		if len(slackFile.Channels) != 0 {
			downloadFolderResult = filepath.Join(downloadFolderResult, channelsResult)
			os.Mkdir(downloadFolderResult, 0777)
		}
	}

	withId := false
	if exists(getFileName(slackFile, withId, downloadFolderResult)) {
		withId = true
	}
	file, err := os.OpenFile(getFileName(slackFile, withId, downloadFolderResult), os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err, slackFile.URLPrivateDownload)
		return
	}
	defer file.Close()
	_, err = file.Write(body)
	if err != nil {
		log.Println(err, slackFile.URLPrivateDownload)
		return
	}
	fmt.Println("Downloaded:", slackFile.ID, quote(slackFile.Title), slackFile.URLPrivateDownload)
}

func getFiles(api *slack.Client, fileType string, beforeTimestamp int64, channel string, page int) ([]slack.File, *slack.Paging, error) {
	return api.GetFiles(slack.GetFilesParameters{
		Types:       fileType,
		Count:       1000,
		Page:        page,
		Channel:     channel,
		TimestampTo: slack.JSONTime(beforeTimestamp),
	})
}

func getFileName(slackFile slack.File, withId bool, downloadFolder string) string {
	if withId {
		return fmt.Sprintf("%s/%s_%s", downloadFolder, slackFile.ID, slackFile.Title)
	} else {
		return fmt.Sprintf("%s/%s", downloadFolder, slackFile.Title)
	}
}
