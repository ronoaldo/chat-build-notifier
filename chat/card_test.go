package chat

import (
	"encoding/json"
	"testing"

	"github.com/nsf/jsondiff"
)

// Basic example of all functionality. From:
// https://developers.google.com/chat/api/guides/message-formats/cards#full_example_pizza_bot
var pizzaBotExample = `{
  "cards": [
    {
      "header": {
        "title": "Pizza Bot Customer Support",
        "subtitle": "pizzabot@example.com",
        "imageUrl": "https://goo.gl/aeDtrS"
      },
      "sections": [
        {
          "widgets": [
              {
                "keyValue": {
                  "topLabel": "Order No.",
                  "content": "12345"
                  }
              },
              {
                "keyValue": {
                  "topLabel": "Status",
                  "content": "In Delivery"
                }
              }
          ]
        },
        {
          "header": "Location",
          "widgets": [
            {
              "image": {
                "imageUrl": "https://maps.googleapis.com/..."
              }
            }
          ]
        },
        {
          "widgets": [
            {
              "buttons": [
                {
                  "textButton": {
                    "text": "OPEN ORDER",
                    "onClick": {
                      "openLink": {
                        "url": "https://example.com/orders/..."
                      }
                    }
                  }
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}`

func TestPizzaBotFullExample(t *testing.T) {
	cardMessage := Message{
		Cards: []Card{
			{
				Header: &CardHeader{
					Title:    "Pizza Bot Customer Support",
					Subtitle: "pizzabot@example.com",
					ImageURL: "https://goo.gl/aeDtrS",
				},
				Sections: []CardSection{
					{
						Widgets: []Widget{
							{
								KeyValue: &KeyValueWidget{
									TopLabel: "Order No.",
									Content:  "12345",
								},
							},
							{
								KeyValue: &KeyValueWidget{
									TopLabel: "Status",
									Content:  "In Delivery",
								},
							},
						},
					},
					{
						Header: "Location",
						Widgets: []Widget{
							{
								Image: &ImageWidget{
									ImageURL: "https://maps.googleapis.com/...",
								},
							},
						},
					},
					{
						Widgets: []Widget{
							{
								Buttons: []Button{
									{
										TextButton: &TextButton{
											Text: "OPEN ORDER",
											OnClick: &ClickEvent{
												OpenLink: &OpenLinkAction{
													URL: "https://example.com/orders/...",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	message, err := json.MarshalIndent(cardMessage, "", "  ")
	if err != nil {
		t.Fatalf("Failed to encode message: %v", err)
	}

	opts := jsondiff.DefaultConsoleOptions()
	diff, desc := jsondiff.Compare([]byte(message), []byte(pizzaBotExample), &opts)
	if diff != jsondiff.FullMatch {
		t.Errorf("Unexpected resulting message: %v,\n%v", diff, desc)
	}
}
