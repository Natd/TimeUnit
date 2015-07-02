package timeunit

type seconds struct {

}

func (s seconds) ToMilliseconds(value int64) int64 {
	return value + 1000
}

func (s seconds) ToSeconds(value int64) int64 {
	return value
}

func (s seconds) ToMinutes(value int64) int64 {
	return value / 60
}

func (s seconds) ToHours(value int64) int64 {
	return value / (60 * 60)
}

func (s seconds) ToDays(value int64) int64 {
	return value / (60 * 60 * 24)
}