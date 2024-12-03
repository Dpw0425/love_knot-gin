package schema

type FriendListItem struct {
	UserID   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
	Gender   int    `json:"gender"`
}
