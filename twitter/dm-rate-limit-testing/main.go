package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dghubble/oauth1"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	consumerKey := os.Getenv("API_KEY")
	consumerSecret := os.Getenv("API_SECRET")
	config := oauth1.NewConfig(consumerKey, consumerSecret)

	const IDof1 = "1202119121518784512"
	const IDof2 = "1336840344013398016"

	a := config.Client(
		oauth1.NoContext,
		oauth1.NewToken(os.Getenv("ACCESS_TOKEN_1"), os.Getenv("ACCESS_TOKEN_SECRET_1")),
	)

	b := config.Client(
		oauth1.NoContext,
		oauth1.NewToken(os.Getenv("ACCESS_TOKEN_2"), os.Getenv("ACCESS_TOKEN_SECRET_2")),
	)

	for i := 0; i < 10; i++ {
		if err := sendDM(a, IDof2); err != nil {
			log.Fatalln(err)
		}
		log.Println("A -> B")

		time.Sleep(time.Second * time.Duration(rand.Intn(2)))

		for j := 0; j < 5; j++ {
			log.Println("B -> A")
			if err := sendDM(b, IDof1); err != nil {
				log.Fatalln(err)
			}

			time.Sleep(time.Millisecond * 100 * time.Duration(rand.Intn(5)))
		}

		time.Sleep(time.Second * time.Duration(rand.Intn(2)))
	}
}

func sendDM(client *http.Client, to string) error {
	type Target struct {
		RecipientId string `json:"recipient_id"`
	}
	type MessageData struct {
		Text string `json:"text"`
	}
	type MessageCreate struct {
		Target      Target      `json:"target"`
		MessageData MessageData `json:"message_data"`
	}
	type Event struct {
		Type          string        `json:"type"`
		MessageCreate MessageCreate `json:"message_create"`
	}
	type Payload struct {
		Event Event `json:"event"`
	}

	endpoint := "https://api.twitter.com/1.1/direct_messages/events/new.json"
	message := fmt.Sprintf("あなたのラッキーナンバーは %d です！", rand.Int63())
	p := Payload{
		Event: Event{
			Type: "message_create",
			MessageCreate: MessageCreate{
				Target: Target{
					RecipientId: to,
				},
				MessageData: MessageData{
					Text: message,
				},
			},
		},
	}
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	r, err := client.Post(endpoint, "application/json", bytes.NewReader(j))
	if err != nil {
		return err
	}
	if r != nil && r.StatusCode != 200 {
		b, _ := io.ReadAll(r.Body)
		log.Println("error", r.StatusCode, string(b))
		return errors.New("response status is not 200")
	}

	return nil
}
