package model

type Usuario struct {
	Nome string `json:"nome" db:"nome"`
	RA   string `json:"ra" db:"ra"`
}
