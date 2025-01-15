package ratelimit

import (
	"errors"
	"fmt"
	"net/http"

	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/time/rate"
)

const (
	defaultCacheSize           = 2048
	ServerLimitType  LimitType = "server"
	IPLimitType      LimitType = "ip"
)

type LimitType string

type LimitConfig struct {
	LimitType LimitType `yaml:limitType`
	Burst     int       `yaml:burst`
	QPS       int       `yaml:qps`
	CacheSize int       `yaml:cacheSize`
}

func (c *LimitConfig) Validate() error {
	if c.QPS == 0 || c.Burst == 0 {
		return fmt.Errorf("LimitConfig Burst and QPS cannot be empty")
	}

	if c.QPS > c.Burst {
		return fmt.Errorf("LimitConfig QPS(%d) cannot exceed Burst(%d)", c.QPS, c.Burst)
	}

	if c.CacheSize == 0 {
		c.CacheSize = defaultCacheSize
	}

	return nil
}

type RateLimiter struct {
	limitType          LimitType
	keyFunc            func(r *http.Request) string
	cache              *lru.Cache[string, *rate.Limiter]
	rateLimiterFactory func() *rate.Limiter
}

func NewRateLimiter(conf *LimitConfig) (*RateLimiter, error) {
	if conf == nil {
		return nil, errors.New("invalid config")
	}

	if err := conf.Validate(); err != nil {
		return nil, err
	}

	var keyFunc func(r *http.Request) string
	switch conf.LimitType {
	case ServerLimitType:
		keyFunc = func(r *http.Request) string {
			return ""
		}
	case IPLimitType:
		keyFunc = func(r *http.Request) string {
			return r.RemoteAddr
		}
	default:
		return nil, fmt.Errorf("unknown LimitType: %s", conf.LimitType)
	}

	c, err := lru.New[string, *rate.Limiter](conf.CacheSize)
	if err != nil {
		return nil, err
	}

	rateLimiterFactory := func() *rate.Limiter {
		return rate.NewLimiter(rate.Limit(conf.QPS), conf.Burst)
	}

	return &RateLimiter{
		limitType:          conf.LimitType,
		keyFunc:            keyFunc,
		cache:              c,
		rateLimiterFactory: rateLimiterFactory,
	}, nil
}

func (rl *RateLimiter) get(key string) *rate.Limiter {
	value, found := rl.cache.Get(key)
	if !found {
		new := rl.rateLimiterFactory()
		rl.cache.Add(key, new)
		return new
	}
	return value
}

func (rl *RateLimiter) Accept(r *http.Request) error {
	key := rl.keyFunc(r)
	limiter := rl.get(key)

	if !limiter.Allow() {
		return fmt.Errorf("limit reached on %s for key %v", rl.limitType, key)
	}

	return nil
}
