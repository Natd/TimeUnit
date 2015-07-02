package timeunit

type Milliseconds struct {

}

func (m Milliseconds) ToMilliseconds(value int64) int64 {
	return value
}

func (m Milliseconds) ToSeconds(value int64) int64 {
	return value / 1000
}

func (m Milliseconds) ToMinutes(value int64) int64 {
	return value / (60 * 1000)
}

func (m Milliseconds) ToHours(value int64) int64 {
	return value / (60 * 60 * 1000)
}

func (m Milliseconds) ToDays(value int64) int64 {
	return value / (60 * 60 * 24 * 1000)
}