package timeunit

type milliseconds struct {

}

func (m milliseconds) ToMilliseconds(value int64) int64 {
	return value
}

func (m milliseconds) ToSeconds(value int64) int64 {
	return value / 1000
}

func (m milliseconds) ToMinutes(value int64) int64 {
	return value / (60 * 1000)
}

func (m milliseconds) ToHours(value int64) int64 {
	return value / (60 * 60 * 1000)
}

func (m milliseconds) ToDays(value int64) int64 {
	return value / (60 * 60 * 24 * 1000)
}