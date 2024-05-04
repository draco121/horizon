package constants

type Action string

const (
	Read   Action = "read"
	Write  Action = "write"
	Delete Action = "delete"
	All    Action = "all"
)
