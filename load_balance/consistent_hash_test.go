package load_balance

import (
	"fmt"
	"testing"
)

func TestConsistentHash(t *testing.T){
	c := NewConsistantHashBanlance(10,nil)
	c.Add("127.0.0.1")
	c.Add("127.0.0.2")
	c.Add("127.0.0.3")
	c.Add("127.0.0.4")
	c.Add("127.0.0.5")

	// url hash
	fmt.Println(c.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(c.Get("http://127.0.0.1:2002/base/error"))
	fmt.Println(c.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(c.Get("http://127.0.0.1:2002/base/changepwd"))

	// ip hash
	fmt.Println(c.Get("192.168.0.1"))
	fmt.Println(c.Get("192.168.0.2"))
	fmt.Println(c.Get("192.168.0.1"))
}