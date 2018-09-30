package id

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFlicker_NewID(t *testing.T) {
	flicker := NewFlicker(nil, 1000)
	flicker.NewID()
}
