package load_banlance

type LoadBanlance interface {
	Add(params ...string) error
	Get(key string)(string, error)
}


