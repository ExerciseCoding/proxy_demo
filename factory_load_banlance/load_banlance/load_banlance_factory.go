package load_banlance

type LbType int
const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)
func LoadBanlanceFactory(lbType LbType) LoadBanlance {
	switch lbType {
	case LbRandom:
		return &RandomBalance{}
	case LbRoundRobin:
		return &RoundRobinBalance{}
	case LbWeightRoundRobin:
		return &WeightRoundRobin{}
	case LbConsistentHash:
		return NewConsistantHashBanlance(10,nil)
	default:
		return &RandomBalance{}

	}

}