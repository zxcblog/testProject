// 使用redis记录id， 随机增长，保持id的唯一性
package cache

import (
	"context"
	"new-project/global"
	"new-project/pkg/util"
	"time"
)

// UniqueSN 通过redis中记录的数值进行唯一id的获取
func UniqueSN(ctx context.Context, key string) string {
	// 记录每个小时内的随机增长id
	key += time.Now().Format("200601021504")
	global.Redis.IncrBy(ctx, key, int64(util.RandomInt(100)))

	val, _ := global.Redis.Get(ctx, key).Result()
	return util.RandomStr(16) + val
}
