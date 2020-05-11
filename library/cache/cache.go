package cache

type Cache interface {
	init()
	Set(key string, data string, ttl int) bool
	Get(key string) string
	SAdd(key string, member []byte)
	SCard(key string) int
	SMembers(key string) []string
}

var (
	Client Cache
)

const (
	ExpireNever 	= 0
	ExpireMinute 	= 60
	ExpireHour 		= 3600
	ExpireDay 		= 86400
)

func Init(cache string) {
	Client = clientFactory(cache)
}

func Set(key string, data string, ttl int) bool {
	return Client.Set(key, data, ttl)
}

func Get(key string) string {
	return Client.Get(key)
}

func clientFactory(vendor string) Cache {
	var c Cache

	switch vendor {
	case "redis":
		c = new(RedisClient)
	default:
		panic("cache vendor not found")
	}

	c.init()

	return c
}