package msg

import (
	"encoding/json/jsontext"
)

type Message struct {
	Type    string         `json:"type"`
	Payload jsontext.Value `json:"payload"`
}

type MessagePayload struct {
	Type    string `json:"type"`
	Payload any    `json:"payload,omitempty"`
}

type Event struct {
	Type    string         `json:"type"` // ability.drop|ability.accept
	Value   string         `json:"value,omitempty"`
	Service string         `json:"service,omitempty"`
	Channel string         `json:"channel,omitempty"`
	Payload jsontext.Value `json:"payload,omitempty"`
}

type EventPayload struct {
	Event
	Payload any `json:"payload,omitempty"`
}

type SourceChatStatePayload struct {
	State string `json:"state"`
}

type SourceConnected struct {
	Service string `json:"service"`
}

type SourceChatChannelJoinPayload struct { // Команда для Подключения или Отключения от канала
	Channel string `json:"channel"`
}

type SourceChatChannelJoin struct {
	Type    string                       `json:"type"`
	Payload SourceChatChannelJoinPayload `json:"payload"`
}

type Config struct {
	Type    string        `json:"type"`
	Payload ConfigPayload `json:"payload"`
}

type ConfigPayload struct {
	Channels map[string][]string `json:"channels"`
	Rooms    []ConfigRoom        `json:"rooms"`
}

type ConfigRoom struct {
	Id       int             `json:"id"`
	Title    string          `json:"title"`
	FontSize int             `json:"font_size"`
	ThemeId  int             `json:"theme_id"`
	Theme    ConfigRoomTheme `json:"theme"`
}

type ConfigRoomTheme struct {
	Id     int    `json:"id"`
	Key    string `json:"key"`
	Title  string `json:"title"`
	Weight int    `json:"weight"`
}

type ChatMessage struct {
	Type    string             `json:"type"`
	Payload ChatMessagePayload `json:"payload"`
}

type ChatMessagePayload struct {
	Service string             `json:"service"`
	Channel string             `json:"channel"`
	User    ChatMessageUser    `json:"user"`
	Message ChatMessageMessage `json:"message"`
	Tags    map[string]string  `json:"tags"`
}

type ChatMessageUser struct {
	Id     string   `json:"id"`
	Nick   string   `json:"nick"`
	Login  string   `json:"login"`
	Status string   `json:"status"`
	Avatar string   `json:"avatar"`
	Badges []string `json:"badges"`
	Color  string   `json:"color"`
}

type ChatMessageMessage struct {
	Html string `json:"html"`
	Text string `json:"text"`
	Src  string `json:"src"`
}
