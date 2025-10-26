package models

type Perm int

const (
	ADMIN Perm = iota + 1
	MOD
	USER
)
