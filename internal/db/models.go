// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Shop struct {
	ShopID         uuid.UUID
	OwnerID        uuid.UUID
	Title          string
	DefaultDomain  string
	FaviconUrl     pgtype.Text
	CurrencyCode   string
	Status         string
	About          pgtype.Text
	SeoDescription pgtype.Text
	SeoKeywords    pgtype.Text
	SeoTitle       pgtype.Text
	UpdatedAt      pgtype.Timestamptz
	CreatedAt      pgtype.Timestamptz
}

type User struct {
	UserID            uuid.UUID
	Auth0Sub          pgtype.Text
	Email             string
	Name              pgtype.Text
	ProfilePictureUrl pgtype.Text
	CreatedAt         pgtype.Timestamp
	LastLogin         pgtype.Timestamp
}
