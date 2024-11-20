package dto

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPayload_UserRate(t *testing.T) {
	var p Payload
	p.UserId = ""
	fmt.Println(p.Random())
}

func BenchmarkPayload_UserRate(b *testing.B) {
	var num int64
	for i := 0; i < b.N; i++ {
		var p Payload
		p.AndroidId = random.RandString(32)
		if p.Random() < 50 {
			num++
		}
	}
	fmt.Println("rate", b.N, num)
}

func TestPayload_Random(t *testing.T) {
	body, err := os.ReadFile("android.txt")
	if err != nil {
		panic(err)
	}
	line := strings.Split((string(body)), "\n")
	size := len(line)
	fmt.Println(size)
	var p Payload
	var i int
	for _, v := range line {
		p.AndroidId = v
		if p.Random() < 50 {
			i++
		}
	}
	fmt.Println("小于50%", i)
	fmt.Println("大于等于50%", size-i)

}

func TestPayload_RegisterAfter(t *testing.T) {
	p := Payload{
		RegisterTime: "2024-07-01 16:23:33",
	}
	b := p.RegisterAfter("2024-05-01")
	assert.True(t, b)
	b = p.RegisterAfter("2024-05-01 18:20:33")
	assert.True(t, b)
	b = p.RegisterBefore("2024-05-01")
	assert.False(t, b)
	b = p.RegisterBefore("2024-05-01 18:20:33")
	assert.False(t, b)
	b = p.RegisterBetween("2024-07-01", "2024-08-01")
	assert.True(t, b)
	b = p.RegisterBetween("2024-05-01", "2024-06-01 20:00:00")
	assert.False(t, b)
}
