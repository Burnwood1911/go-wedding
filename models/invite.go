package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"gorm.io/gorm"
)

type inviteType string

type InviteStatus string

const (
	Single inviteType = "SINGLE"
	Double inviteType = "DOUBLE"
)

const (
	Used InviteStatus = "USED"

	Unused InviteStatus = "UNUSED"
)

type Invite struct {
	Id uint `json:"id" gorm:"primaryKey"`

	Secret string `json:"secret" gorm:"unique;not null"`

	Name string `json:"name" gorm:"unique;not null"`

	MaxUsage int `json:"maxUsage" gorm:"not null"`

	Table int `json:"table" gorm:"not null"`

	Uses int `json:"uses" gorm:"not null"`

	InviteType inviteType `json:"inviteType" gorm:"not null"`

	InviteStatus InviteStatus `json:"inviteStatus" gorm:"not null"`
}

type InviteRepo struct {
	DB *gorm.DB
}

func NewInviteRepo(DB *gorm.DB) *InviteRepo {
	return &InviteRepo{DB}
}

func (a *Invite) Validate() error {
	return validation.ValidateStruct(a,
		validation.Field(&a.Secret, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&a.MaxUsage, validation.Required, validation.Min(1), validation.Max(2)),
		validation.Field(&a.Uses, validation.NotNil, validation.Min(0), validation.Max(2)),
		validation.Field(&a.InviteType, validation.Required),
		validation.Field(&a.InviteStatus, validation.Required),
		validation.Field(&a.Table, validation.Required),
	)
}

func (r *InviteRepo) CreateAlias(a *Invite) error {

	if result := r.DB.Create(a); result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *InviteRepo) GetInviteBySecret(secret string) (*Invite, error) {

	invite := new(Invite)

	if result := r.DB.Where("secret = ?", secret).First(invite); result.Error != nil {
		return nil, result.Error
	}

	return invite, nil
}

func (r *InviteRepo) UpdateInvite(invite *Invite) error {

	if result := r.DB.Save(invite); result.Error != nil {
		return result.Error
	}

	return nil
}
