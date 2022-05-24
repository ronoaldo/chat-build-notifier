package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ronoaldo/chat-build-notifier/chat"
)

const (
	EnvChatWebhook = "CHAT_WEBHOOK"

	BitmojiLamp    = "https://sdk.bitmoji.com/render/panel/82a33c72-3399-441f-8a2e-4132b1f8ce51-4833b118-6c9e-41a4-9557-987f80d99f00-v1.png?transparent=1&palette=1"
	BitmojiShout   = "https://sdk.bitmoji.com/render/panel/44f0e4d9-7f1b-4130-951a-3e49c549add4-4833b118-6c9e-41a4-9557-987f80d99f00-v1.png?transparent=1&palette=1"
	BitmojiYes     = "https://sdk.bitmoji.com/render/panel/987cfa1b-cbd3-49c6-b4de-5d62bdd31bf8-4833b118-6c9e-41a4-9557-987f80d99f00-v1.png?transparent=1&palette=1"
	BitmojiWarning = "https://sdk.bitmoji.com/render/panel/c96df4f1-0e63-4dd3-816a-f3d4d431c4d4-4833b118-6c9e-41a4-9557-987f80d99f00-v1.png?transparent=1&palette=1"
	BitmojiError   = "https://sdk.bitmoji.com/render/panel/4a737c47-e262-44e8-b96e-66d11bb30cab-4833b118-6c9e-41a4-9557-987f80d99f00-v1.png?transparent=1&palette=1"
)

var (
	message     string
	messageType string

	webhook string

	actionLink string
	actionText string
)

var MessageURLs = map[string]string{
	"yes":     BitmojiYes,
	"info":    BitmojiShout,
	"warning": BitmojiWarning,
	"error":   BitmojiError,
}

func init() {
	flag.StringVar(&message, "message", "", "The `MESSAGE` to send via webhook")
	flag.StringVar(&messageType, "type", "info", "The `TYPE` of the message to send: [yes, info, error, warning]")
	flag.StringVar(&webhook, "webhook", "", "The webtook `URL` to send the message to. If not provided, will try the environment var "+EnvChatWebhook)
	flag.StringVar(&actionLink, "link", "", "An optional link `URL` for the user to click")
	flag.StringVar(&actionText, "link-name", "View details", "The action link `NAME`.")
}

func main() {
	flag.Parse()

	imageUrl, ok := MessageURLs[messageType]
	if !ok {
		log.Fatalf("Invalid message type: %v", messageType)
	}

	log.Printf("Building message from type='%v', message='%v'", message, messageType)
	m := chat.Message{
		Cards: []chat.Card{
			{
				Header: &chat.CardHeader{
					Title:    strings.ToUpper(messageType),
					Subtitle: "Sent using the CLI",
					ImageURL: imageUrl,
				},

				Sections: []chat.CardSection{
					{
						Widgets: []chat.Widget{
							{
								TextParagraph: &chat.TextWidget{
									Text: message,
								},
							},
						},
					},
				},
			},
		},
	}

	if actionLink != "" {
		actionButton := chat.Button{
			TextButton: &chat.TextButton{
				Text: actionText,
				OnClick: &chat.ClickEvent{
					OpenLink: &chat.OpenLinkAction{URL: actionLink},
				},
			},
		}
		m.Cards[0].Sections[0].Widgets[0].Buttons = append(m.Cards[0].Sections[0].Widgets[0].Buttons, actionButton)
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Fatalf("Error encoding message: %v", err)
	}
	log.Println("Sending message with body", string(b))

	if webhook == "" {
		log.Printf("No webhook provided, looking up for env %v", EnvChatWebhook)
		webhook = os.Getenv(EnvChatWebhook)
	}
	req, err := http.NewRequest("POST", webhook, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	hc := &http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		log.Fatalf("Error making HTTP call: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status returned: %v %v", resp.StatusCode, resp.Status)
	}

	log.Println("Message sent")
}
