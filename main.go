package main

import (
	"fmt"
	"log"
	"strings"

	"bufio"

	"github.com/nlopes/slack"
	"os"
)

func fetchEnv(key string) (value string) {
	value, present := os.LookupEnv(key)
	if !present {
		errMsg := fmt.Sprintf("Abort: missing environment variable %s", key)
		log.Fatal(errMsg)
	} else {
		log.Printf("Found environment variable %s : %v", key, value)
	}
	return value
}


func main() {
	version := "0.0.1"
	log.Printf("slack-poster version %s", version)

	token := fetchEnv("SLACK_POSTER_API_TOKEN")
	postTo := fetchEnv("SLACK_POSTER_DESTINATION")

	api := slack.New(token)

	var body string
	var lines []string

	if len(os.Args) >1 {
		body = strings.Join(os.Args[1:], " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		body = strings.Join(lines, "\n")
	}

	log.Println("Posting message {", body, "} to destination {", postTo, "}")

	params := slack.NewPostMessageParameters()
	params.LinkNames = 1

	log.Println("posting to Slack API...")
	_, _, err := api.PostMessage(postTo, slack.MsgOptionText(body, false), slack.MsgOptionPostMessageParameters(params))
	if err != nil {
		log.Fatalf("%s\n", err)
	}
}
