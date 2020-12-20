package models

import (
	"time"
)

type Migration struct {
	Version string
	Date    time.Time
}
