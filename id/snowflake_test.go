package id

import (
	"fmt"
	"testing"
)

func TestSnowflake_NewID(t *testing.T) {
	snowflake, err := NewSnowflake(2, 3, 0)
	if err != nil {
		t.Error(err)
	}

	id, err := snowflake.NewID()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(id)
}

func TestSnowflake_NewID1(t *testing.T) {
	snowflake, err := NewSnowflake(1, 1, 0)
	if err != nil {
		t.Error(err)
	}

	ids := map[int64]int{}

	total := 1 << 25
	for i := 0; i < total; i++ {
		id, err := snowflake.NewID()
		if err != nil {
			t.Fatal(err)
		}

		if _, ok := ids[id]; !ok {
			ids[id] = 1
		} else {
			t.Fatal(id)
		}
	}

	fmt.Println("ids:", len(ids), "=", total)
}

func BenchmarkSnowflake_NewID(b *testing.B) {
	snowflake, err := NewSnowflake(1, 1, 0)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		id, err := snowflake.NewID()
		if err != nil {
			b.Fatal(err)
		}

		if id < 0 {
			b.Fatal(id)
		}
	}
}
