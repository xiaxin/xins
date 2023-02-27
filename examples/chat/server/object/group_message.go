package object

type GroupMessage struct {
	GID     string `json:"gid"`     // 组ID
	UID     string `json:"uid"`     // 用户ID
	Content string `json:"content"` // 内容
}

func NewGroupMessage(gid, uid, content string) *GroupMessage {
	return &GroupMessage{gid, uid, content}
}
