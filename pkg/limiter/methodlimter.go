package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func NewMethodLimiter() LimiterIface {
	l := &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)}
	return MethodLimiter{
		Limiter: l,
	}
}

// Key 根据RequestURI切割出核心路由作为键值对名称
func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}
	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.limiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rule ...LimiterBucketRule) LimiterIface {
	for _, rule := range rule {
		if _, ok := l.limiterBuckets[rule.Key]; !ok {
			bucket := ratelimit.NewBucketWithQuantum(
				rule.FillInertval,
				rule.Capacity,
				rule.Quantum,
			)
			l.limiterBuckets[rule.Key] = bucket
		}
	}

	return l
}
