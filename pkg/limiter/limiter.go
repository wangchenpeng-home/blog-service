package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

type LimiterIface interface {
	Key(c *gin.Context) string                         // 获取对应的限流器的键值对名称
	GetBucket(key string) (*ratelimit.Bucket, bool)    // 获取令牌桶
	AddBuckets(rules ...LimiterBucketRule) LimiterIface // 新增多个令牌桶
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string        // 自定义的键值对名称
	FillInertval time.Duration // 间隔多久时间放N个令牌
	Capacity     int64         // 令牌桶的容量
	Quantum      int64         // 令牌桶的每次放入数量
}
