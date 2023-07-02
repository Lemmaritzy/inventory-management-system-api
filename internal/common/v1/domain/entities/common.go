package entities

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenClaim struct {
	Id         int64  `json:"id"`
	Lang       string `json:"lang,omitempty"`
	UserType   int    `json:"user_type"`
	ParentId   int64  `json:"parent_id"`
	ParentName string `json:"parent_name"`
	UserTitle  string `json:"user_title"`
	UserName   string `json:"user_name"`
	UserMail   string `json:"user_mail"`
	DeviceType int    `json:"device_type"`
	IsAct      int8   `json:"is_act"`
	Exp        int64  `json:"exp"`
	Iat        int64  `json:"iat"`
	jwt.RegisteredClaims
}

type LogErrors struct {
	Id           uint64    `json:"id"`
	FlagType     uint8     `json:"flag_type"`
	RequestSrc   string    `json:"request_src"`
	RequestPath  string    `json:"request_path,omitempty"`
	RequestFile  string    `json:"request_file,omitempty"`
	RequestLine  int16     `json:"request_line"`
	RequestFunc  string    `json:"request_func,omitempty"`
	RequestType  string    `json:"request_type,omitempty"`
	RequestRaw   string    `json:"request_raw,omitempty"`
	InstanceId   string    `json:"instance_id"`
	ContentError string    `json:"content_error"`
	ContentRaw   string    `json:"content_raw"`
	CorrId       string    `json:"corr_id,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	SentAt       time.Time `json:"sent_at,omitempty"`
	IsSolved     uint8     `json:"is_solved"`
	IsNotified   uint8     `json:"is_notified"`
	IsAssigned   uint8     `json:"is_assigned"`
}

type LogSms struct {
	Id             uint64    `json:"id"`
	RequestSrc     string    `json:"request_src"`
	UserId         uint64    `json:"user_id"`
	IoType         uint8     `json:"io_type"`
	NumFrom        string    `json:"num_from"`
	NumTo          string    `json:"num_to"`
	MessageSize    uint8     `json:"message_size"`
	MessageContent string    `json:"message_content"`
	ErrData        string    `json:"err_data"`
	CorrId         string    `json:"corr_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	SentAt         time.Time `json:"sent_at,omitempty"`
	IsSent         uint8     `json:"is_sent"`
}

type LogMail struct {
	Id             uint64    `json:"id"`
	RequestSrc     string    `json:"request_src"`
	UserId         uint64    `json:"user_id"`
	IoType         uint8     `json:"io_type"`
	NumFrom        string    `json:"num_from"`
	NumTo          string    `json:"num_to"`
	MessageSize    uint8     `json:"message_size"`
	MessageContent string    `json:"message_content"`
	ErrData        string    `json:"err_data"`
	CorrId         string    `json:"corr_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	SentAt         time.Time `json:"sent_at,omitempty"`
	IsSent         uint8     `json:"is_sent"`
}

type LogPush struct {
	Id         uint64    `json:"id"`
	RequestSrc string    `json:"request_src"`
	UserId     uint64    `json:"user_id"`
	PushType   uint8     `json:"push_type"`
	DeviceType uint8     `json:"device_type"`
	DeviceIdTo string    `json:"device_id_to"`
	RawContent string    `json:"raw_content"`
	CorrId     string    `json:"corr_id"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	SentAt     time.Time `json:"sent_at,omitempty"`
	IsSent     uint8     `json:"is_sent"`
}

type LogInApp struct {
	Id         uint64    `json:"id"`
	RequestSrc string    `json:"request_src"`
	UserId     uint64    `json:"user_id"`
	PushType   uint8     `json:"push_type"`
	DeviceType uint8     `json:"device_type"`
	DeviceIdTo string    `json:"device_id_to"`
	RawContent string    `json:"raw_content"`
	CorrId     string    `json:"corr_id"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	SentAt     time.Time `json:"sent_at,omitempty"`
	IsSent     uint8     `json:"is_sent"`
}

type LogSignIn struct {
	AuthId    uint64    `json:"auth_id"`
	Src       uint8     `json:"src"`
	Ip        string    `json:"ip"`
	TimeStamp string    `json:"time_stamp"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserData struct {
	Id            int64  `json:"id"`
	Lang          string `json:"lang,omitempty"`
	UserType      int    `json:"user_type,omitempty"`
	UserFirstName string `json:"user_first_name,omitempty"`
	UserLastName  string `json:"user_last_name,omitempty"`
	UserName      string `json:"user_name,omitempty"`
	UserMail      string `json:"user_mail,omitempty"`
	UserPhone     string `json:"user_phone,omitempty"`
	DeviceType    int    `json:"device_type,omitempty"`
	IsAct         int8   `json:"is_act,omitempty"`
}
