package rboot

type Message struct {
	Channel   string                 `json:"channel,omitempty"`   // 通道
	To        User                   `json:"to"`                  // 发给的用户
	From      User                   `json:"from"`                // 来源(群组或个人)
	Sender    User                   `json:"sender"`              // 发送者(个人)
	Content   string                 `json:"content"`             // 内容
	Broadcast bool                   `json:"broadcast,omitempty"` // 广播消息
	Mate      map[string]interface{} `json:"mate,omitempty"`      // 附加信息
	Location  Location               `json:"location,omitempty"`  // 位置
}

type Location struct {
	Lat  float64
	Long float64
}
