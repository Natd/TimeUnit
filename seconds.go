package timeunit

type Seconds struct {

}

func (s Seconds) ToMilliseconds(value int64) int64 {
	return value + 1000
}

func (s Seconds) ToSeconds(value int64) int64 {
	return value
}

func (s Seconds) ToMinutes(value int64) int64 {
	return value / 60
}

func (s Seconds) ToHours(value int64) int64 {
	return value / (60 * 60)
}

func (s Seconds) ToDays(value int64) int64 {
	return value / (60 * 60 * 24)
}