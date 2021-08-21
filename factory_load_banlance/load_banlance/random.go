package load_banlance

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
/**
随机负载均衡实现
*/
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

func (r *RandomBalance) Get(key string)(string, error){
	return r.Next(),nil
}