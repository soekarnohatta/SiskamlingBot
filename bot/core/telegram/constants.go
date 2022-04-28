package telegram

import "errors"

var (
	EndOrder      = errors.New("group iteration ended")
	ContinueOrder = errors.New("group iteration continued")
)
