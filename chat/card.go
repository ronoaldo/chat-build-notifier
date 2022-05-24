package chat

// Type Message represents a Google Chat message object that can be sent via
// Webhook. It is expected that you only pass either Text or Cards value, but
// not both.
type Message struct {
	Text string `json:"text,omitempty"`

	Cards        []Card `json:"cards,omitempty"`
	FallbackText string `json:"fallbackText,omitempty"`
}

// Type Card defines a card that can be used to compose a more elaborated
// message object.
type Card struct {
	// Name of the card
	Name string `json:"name,string"`

	// Header is shown at the top of the card.
	Header *CardHeader `json:"header,omitempty"`

	// The card must have at least one section, which is composed by several
	// widgets.
	Sections []CardSection `json:"sections,omitempty"`
}

// Type CardHeader has the text details of a header shown in a card.
type CardHeader struct {
	Title      string `json:"title,omitempty"`
	Subtitle   string `json:"subtitle,omitempty"`
	ImageURL   string `json:"imageUrl,omitempty"`
	ImageStyle string `json:"imageStyle,omitempty"`
}

// Type CardSection is one element of the card, composed by an optional header
// and one type of a Widget.
type CardSection struct {
	Header  string   `json:"header,omitempty"`
	Widgets []Widget `json:"widgets,omitempty"`
}

// Type Widget represents a single widget inside the card section.
//
// The widget type is expected to have only one of the available fields filled:
// you can set either TextParagraph, KeyValue or Image fields, but not more than
// one at the same time.
//
// The widget can also have one or more buttons.
type Widget struct {
	// Your widget will have only one of TextParagraph, KeyValue or Image options.
	TextParagraph *TextWidget     `json:"textParagraph,omitempty"`
	KeyValue      *KeyValueWidget `json:"keyValue,omitempty"`
	Image         *ImageWidget    `json:"image,omitempty"`

	// Your widget can have several buttons.
	Buttons []Button `json:"buttons,omitempty"`
}

// Type TextWidget represents very simple widget with just a set of text on it.
type TextWidget struct {
	Text string `json:"text,omitempty"`
}

// Type KeyValueWidget displays a lable and a value.
type KeyValueWidget struct {
	TopLabel         string `json:"topLabel,omitempty"`
	Content          string `json:"content,omitempty"`
	ContentMultiline bool   `json:"contentMultiline,omitempty"`
	BottomLabel      string `json:"bottomLabel,omitempty"`

	Icon    string      `json:"icon,omitempty"`
	IconURL string      `json:"iconUrl,omitempty"`
	Button  *Button     `json:"button,omitempty"`
	OnClick *ClickEvent `json:"onClick,omitempty"`
}

// Type ImageWidget displays an image.
type ImageWidget struct {
	ImageURL string      `json:"imageUrl,omitempty"`
	OnClick  *ClickEvent `json:"onClick,omitempty"`
}

// Type ClickEvent allows you to add actions to several elements.  You are
// expected to fill either the OpenLink action, or define a custom function
// action. Custom function actions may not work as expected if you are only
// using a webhook call. Check the full documentation on the Chat API on
// https://developers.google.com/chat/api/ to learn more.
type ClickEvent struct {
	OpenLink *OpenLinkAction `json:"openLink,omitempty"`
	Action   *FormAction     `json:"action,omitempty"`
}

// Type OpenLinkAction allows the user to open the specified URL when the element is
// clicked.
type OpenLinkAction struct {
	URL string `json:"url,omitempty"`
}

type FormAction struct {
	MethodName string            `json:"actionMethodName,omitempty"`
	Parameters map[string]string `json:"parameters,omitempty"`
}

// Type Button allows messages to have either Text or Image buttons. Like the
// Widget, it is expected that you specify either TextButton or ImageButton.
type Button struct {
	TextButton *TextButton `json:"textButton,omitempty"`
}

// Type TextButton is a simple button with text.
type TextButton struct {
	Text    string      `json:"text,omitempty"`
	OnClick *ClickEvent `json:"onClick,omitempty"`
}

// Type ImageButton is a more fancy button with an icon image. Either specify
// one of the built-in icons or an icon via URL.
type ImageButton struct {
	Icon    string `json:"icon,omitempty"`
	IconURL string `json:"iconUrl,omitempty"`

	OnClick *ClickEvent `json:"onClick,omitempty"`
}
