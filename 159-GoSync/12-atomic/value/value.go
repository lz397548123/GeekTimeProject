package value

import "math/rand"

type Config struct {
	NodeName string
	Addr     string
	Count    int32
}

func loadNewConfig() Config {
	return Config{
		NodeName: "深圳",
		Addr:     "1.1.1.1",
		Count:    rand.Int31(),
	}
}
