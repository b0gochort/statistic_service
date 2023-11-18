package model

type OnlineItem struct {
	UserId int64 `json:"user_id" reindex:"user_id,hash,pk"`
	Auth   bool  `json:"auth" reindex:"auth,-"`
	Time   int64 `json:"time" reindex:"time,ttl,expire_after=120"`
}
type OnlineLens struct {
	LenAuth   int `json:"len_auth"`
	LenUnAuth int `json:"len_un_auth"`
}

type NewChatItem struct {
	ID            int64       `json:"id" reindex:"id,hash,pk"`
	Time          int64       `json:"time" reindex:"time,tree"`
	Message       MessageType `json:"message"`
	LastHostStaff bool        `json:"last_host_staff" reindex:"last_host_staff,-"`
	UID           int64       `json:"uid" reindex:"uid,hash"`
	IP            string      `json:"ip" reindex:"ip,hash"`
	Category      string      `json:"category" reindex:"category,hash"`
}

type MessageType struct {
	NumberOfUnread int    `json:"number_of_unread" reindex:"number_of_unread,-"`
	LastMessage    string `json:"last_message" reindex:"last_message,-"`
}

type NewChatRes struct {
	ID            int64       `json:"id"`
	Time          int64       `json:"time"`
	Message       MessageType `json:"message"`
	LastHostStaff bool        `json:"last_host_staff"`
	UID           int64       `json:"uid"`
	IP            string      `json:"ip"`
	Category      string      `json:"category"`
}
