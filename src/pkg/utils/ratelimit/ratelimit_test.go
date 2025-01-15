package ratelimit

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	testCases := []struct {
		name        string
		config      *LimitConfig
		expectedErr error
	}{
		{
			name:        "invalid burst and qps",
			config:      &LimitConfig{},
			expectedErr: errors.New("LimitConfig Burst and QPS cannot be empty"),
		},
		{
			name:        "qps grat than burst",
			config:      &LimitConfig{QPS: 10, Burst: 1},
			expectedErr: errors.New("LimitConfig QPS(10) cannot exceed Burst(1)"),
		},
		{
			name:        "set cache size",
			config:      &LimitConfig{QPS: 2, Burst: 4, CacheSize: 100},
			expectedErr: nil,
		},
		{
			name:        "config success",
			config:      &LimitConfig{QPS: 2, Burst: 4},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cacheSize := tc.config.CacheSize
			err := tc.config.Validate()
			assert.Equal(t, tc.expectedErr, err)

			if cacheSize == 0 && tc.expectedErr == nil {
				assert.Equal(t, defaultCacheSize, tc.config.CacheSize)
			}
		})
	}
}
