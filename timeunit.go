package timeunit

type TimeUnit interface {
	ToMilliseconds(value int64) int64
	ToSeconds(value int64) int64
	ToMinutes(value int64) int64
	ToHours(value int64) int64
	ToDays(value int64) int64
}