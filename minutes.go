package timeunit

type minutes struct {

}

func (m minutes) ToMilliseconds(value int64) int64 {
	return value * 1000 * 60
}

func (m minutes) ToSeconds(value int64) int64 {
	return value * 60
}

func (m minutes) ToMinutes(value int64) int64 {
	return value
}

func (m minutes) ToHours(value int64) int64 {
	return value / 60
}

func (m minutes) ToDays(value int64) int64 {
	return value / (60 * 24)
}