package models

import (
	"time"
)

type StatusApi string
type CodeType string
type Privacy string

const (
	Enabled  StatusApi = "enabled"
	Disabled StatusApi = "disabled"
)

const (
	Public    Privacy = "public"
	Private   Privacy = "private"
	Unlisted  Privacy = "unlisted"
	Protected Privacy = "protected"
)

const (
	Reset  CodeType = "reset"
	Verify CodeType = "verify"
)

type Codes struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Type      string    `gorm:"not null;type:text"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	Code      string    `gorm:"type:text;uniqueIndex"`
	UserId    string    `gorm:"type:uuid;index"`
}

type Users struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	Name      string    `gorm:"type:text"`
	Email     string    `gorm:"type:text;uniqueIndex"`
	Public    bool      `gorm:"default:true"`
	Plan      string    `gorm:"type:text;not null;default:'free'"`
	Suspended bool      `gorm:"default:false"`
	Password  string    `gorm:"type:text;not null"`
	Actived   bool      `gorm:"default:false"`
	Token     string    `gorm:"type:text;not null"`
}

type APIKeys struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Slug      string    `gorm:"unique"`
	UserId    string    `gorm:"type:uuid;index"`
	Read      bool      `gorm:"not null;default:true"`
	Create    bool      `gorm:"column:create;not null;default:true"`
	Delete    bool      `gorm:"not null;default:true"`
	ApiKey    string    `gorm:"unique"`
	StatusApi StatusApi `gorm:"type:text;default:'enabled'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

type Profiles struct {
	Id         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	PublicName string    `gorm:"type:text"`
	Avatar     string    `gorm:"type:text"`
	Bio        string    `gorm:"type:text"`
	Contact    string    `gorm:"type:text"`
	Website    string    `gorm:"type:text"`
	UserId     string    `gorm:"type:uuid;not null;index"`
	Twitter    string    `gorm:"type:text"`
	Github     string    `gorm:"type:text"`
	Location   string    `gorm:"type:text"`
	Company    string    `gorm:"type:text"`
	User       Users     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:UserId"`
	Username   string    `gorm:"type:text;not null;uniqueIndex"`
	Public     bool      `gorm:"default:true"`
}

type Subscriptions struct {
	Id             string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Slug           string    `gorm:"unique"`
	UserId         string    `gorm:"type:uuid;index"`
	Plan           string    `gorm:"type:text;default:'free'"`
	StartDate      time.Time `gorm:"not null"`
	EndDate        time.Time `gorm:"not null"`
	AutoRenew      bool      `gorm:"default:true"`
	CreatedAt      time.Time `gorm:"autoCreateTime;default:now()"`
	SubscriptionId string    `gorm:"type:text;uniqueIndex"`
	CustomerId     string    `gorm:"type:text;uniqueIndex"`
}

type Gemini struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	UserId    string    `gorm:"type:uuid;not null"`
	ModelName string    `gorm:"not null;size:100"`
	ApiKey    string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:now()"`
}

type GeminiHistory struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Slug      string    `gorm:"unique"`
	UserId    string    `gorm:"type:uuid;not null"`
	FromCache bool      `gorm:"default:false"`
	Action    string    `gorm:"type:text;not null"`
	Processed bool      `gorm:"default:false"`
	Error     bool      `gorm:"default:false"`
	Model     string    `gorm:"type:text;not null"`
	Message   string    `gorm:"type:text;null"`
	Prompt    string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:now()"`
}

type Avatars struct {
	Id        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Slug      string    `gorm:"unique"`
	UserId    string    `gorm:"type:uuid;not null;index"`
	Avatar    string    `gorm:"type:text;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime;default:now()"`
}

func (Codes) TableName() string          { return "ds_codes" }
func (Profiles) TableName() string       { return "ds_profiles" }
func (APIKeys) TableName() string        { return "ds_api_keys" }
func (Users) TableName() string          { return "ds_users" }
func (Subscriptions) TableName() string  { return "ds_subscriptions" }
func (Gemini) TableName() string         { return "ds_gemini" }
func (GeminiHistory) TableName() string  { return "ds_gemini_history" }
func (Avatars) TableName() string        { return "ds_avatars" }
