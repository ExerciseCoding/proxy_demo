package load_balance

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

/**
加权轮询负载均衡
 */

type WeightNode struct {
	addr string
	wight int //权重值
	currentWeight int // 当前权重值
	effectiveWeight int // 有效权重
}

type WeightRoundRobin struct {
	rss []*WeightNode
	curIndex int
}

func (r *WeightRoundRobin) Add(params ...string)error{
	if len(params) != 2 {
		return errors.New("param len need 2")
	}
	parInt, err := strconv.ParseInt(params[1],10,64)
	if err != nil{
		return err
	}
	node := &WeightNode{addr: params[0], wight: int(parInt)}
	node.effectiveWeight = node.wight
	r.rss = append(r.rss,node)
	return nil
}

func (r *WeightRoundRobin) Next()string{
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]
		// 1. 统计所以有效权重之和
		total += w.effectiveWeight
		// 2. 变更节点临时权重为节点临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight
		// 3. 有效权重默认与权重相同，通讯异常时-1, 通讯成功+1,直到恢复到weight大小
		if w.effectiveWeight < w.wight{
			w.effectiveWeight++
		}
		// 4. 选择最大临时权重点节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}

	}
	if best == nil{
		return ""
	}

	// 5.变更临时权重为临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr

}

func (r *WeightRoundRobin) Get(key string)(string,error){
	return r.Next(),nil
}

func TestWeightRoundRobin(t *testing.T){
	rb := &WeightRoundRobin{}
	rb.Add("127.0.0.1","4")
	rb.Add("127.0.0.2","3")
	rb.Add("127.0.0.3","2")

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())



}