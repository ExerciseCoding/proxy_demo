package load_balance

import (
	"errors"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type RandomBalance struct {
	curIndex int
	rss []string
}

func (r *RandomBalance) Add(params ...string)error{
	if len(params) == 0{
		return errors.New("param len 1 at least")
	}
	addr := params[0]
	r.rss = append(r.rss,addr)
	return nil

}

func (r *RandomBalance) Next()string{
	if len(r.rss) == 0{
		return ""
	}
	rand.Seed(time.Now().UnixNano())
	r.curIndex = rand.Intn(len(r.rss))
	fmt.Println(r.curIndex)
	return r.rss[r.curIndex]
}

func TestRandomBalance(t *testing.T){
	ipaddr := "127.0.0.1"
	ipaddr1 := "168.172.1.3"
	rss := []string{ipaddr,ipaddr1}

	b := &RandomBalance{
		curIndex: 0,
		rss: rss,
	}
	index := b.Next()
	fmt.Println(index)
}