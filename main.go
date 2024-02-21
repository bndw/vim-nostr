package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"
	"strings"
	"time"

	"github.com/nbd-wtf/go-nostr"
)

func main() {
	var (
		post = flag.Bool("post", false, "post content from stdin")
	)
	flag.Parse()

	if !*post {
		log.Println("USAGE: 'echo test | vim-nostr -post'")
		os.Exit(1)
	}

	var config Config
	if err := loadConfig(&config); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	if config.PrivateKey == "" {
		log.Println("must set 'privatekey' in config")
		os.Exit(1)
	}
	if len(config.WriteRelays()) == 0 {
		log.Println("must set atleast 1 write relay in config")
		os.Exit(1)
	}

	var content string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		content += line + "\n"
	}
	content = strings.TrimSuffix(content, "\n")

	event, err := createEvent(config.PrivateKey, content)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	publish(config.WriteRelays(), event)
}

func createEvent(sk, content string) (*nostr.Event, error) {
	pk, err := nostr.GetPublicKey(sk)
	if err != nil {
		return nil, err
	}

	event := &nostr.Event{
		Kind:      1,
		PubKey:    pk,
		Content:   content,
		CreatedAt: nostr.Now(),
	}

	err = event.Sign(sk)

	return event, err
}

func publish(relays []string, event *nostr.Event) {
	for _, url := range relays {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		relay, err := nostr.RelayConnect(ctx, url)
		if err != nil {
			log.Printf("err: relay connect %q: %v", url, err)
			continue
		}
		defer relay.Close()

		err = relay.Publish(ctx, *event)
		if err != nil {
			log.Printf("err: relay publish %q: %v", url, err)
			continue
		}

		log.Printf("published to %s", url)
	}
}
