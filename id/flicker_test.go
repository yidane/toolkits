package id

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestFlicker_NewID(t *testing.T) {
	db, err := sql.Open("mysql", "root:sasasasasa@/test")
	if err != nil {
		t.Fatal(err)
	}
	flicker := NewFlicker(db)
	err = flicker.Ready()
	if err != nil {
		t.Fatal(err)
	}

	id, err := flicker.NewID()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}
