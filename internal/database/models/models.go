package models

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Model struct {
	CreatedAt time.Time `json:"created_at" bun:"created_at,nullzero"`
	UpdatedAt time.Time `json:"updated_at" bun:"updated_at,nullzero"`
}
type SoftDeleteModel struct {
	DeletedAt *time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}

type User struct {
	Model
	SoftDeleteModel
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int    `json:"id" bun:"id,pk"`
	Username      string `json:"username" bun:"username,notnull"`
	FirstName     string `json:"first_name" bun:"first_name,notnull"`
	LastName      string `json:"last_name" bun:"last_name"`
	Email         string `json:"email" bun:"email"`
	PasswordHash  string `json:"password_hash" bun:"password_hash"`
	Blocked       bool   `json:"blocked" bun:"blocked"`
	Active        bool   `json:"active" bun:"active"`
	Roles         []Role `bun:"m2m:users_to_roles,join:User=Role"`
}

type Role struct {
	Model
	SoftDeleteModel
	bun.BaseModel `bun:"table:roles,alias:r"`
	ID            int    `json:"id" bun:"id,pk"`
	Name          string `json:"name" bun:"name,notnull"`
	Description   string `json:"description" bun:"description"`
	Type          string `json:"type" bun:"type,notnull"`
	Users         []User `bun:"m2m:users_to_roles,join:Role=User"`
}

type UserToRole struct {
	bun.BaseModel `bun:"table:users_to_roles,alias:utr"`
	UserID        int   `bun:",pk"`
	User          *User `bun:"rel:belongs-to,join:user_id=id"`
	RoleID        int   `bun:",pk"`
	Role          *Role `bun:"rel:belongs-to,join:role_id=id"`
}

type Note struct {
	Model
	SoftDeleteModel
	bun.BaseModel `bun:"table:notes"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	Title         string `json:"title" bun:"title,notnull"`
	Body          string `json:"body" bun:"body,type:text"`
	ImageSrc      string `json:"image_src" bun:"image_src"`
	AuthorID      int    `bun:"author_id"`
	Author        *User  `bun:"rel:belongs-to,join:author_id=id"`
}

type Fact struct {
	Model
	SoftDeleteModel
	bun.BaseModel `bun:"table:facts"`
	ID            int    `json:"id" bun:"id,pk,autoincrement"`
	Question      string `json:"question" bun:"question,notnull"`
	Answer        string `json:"answer" bun:"answer,notnull"`
}

var _ bun.BeforeAppendModelHook = (*Model)(nil)

func (m *Model) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}

	return nil
}
