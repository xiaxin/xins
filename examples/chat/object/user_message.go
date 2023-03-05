package object

type UserMessage struct {
	UID     string `json:"uid"`     // 用户ID
	Content string `json:"content"` // 内容
}

func NewUserMessage(uid, content string) *UserMessage {
	return &UserMessage{uid, content}
}
