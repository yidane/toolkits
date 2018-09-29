package id

import (
	"fmt"
	"testing"
)

func TestSnowflake_NewID(t *testing.T) {
	snowflake, err := NewSnowflake(1, 1, 0)
	if err != nil {
		t.Error(err)
	}

	id, err := snowflake.NewID()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(id)
}

func BenchmarkSnowflake_NewID(b *testing.B) {
	snowflake, err := NewSnowflake(1, 1, 0)
	if err != nil {
		b.Error(err)
	}

	ids := map[int64]int{}

	for i := 0; i < b.N; i++ {
		id, err := snowflake.NewID()
		if err != nil {
			b.Fatal(err)
		}

		if _, ok := ids[id]; !ok {
			ids[id] = 1
		} else {
			b.Fatal(id)
		}
	}

	fmt.Println("ids:", len(ids))
}
