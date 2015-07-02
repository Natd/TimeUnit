package timeunit

type hours struct {

}

func (h hours) ToMilliseconds(value int64) int64 {
	return value * 60 * 60 * 1000
}

func (h hours) ToSeconds(value int64) int64 {
	return value * 60 * 60
}

func (h hours) ToMinutes(value int64) int64 {
	return value * 60
}

func (h hours) ToHours(value int64) int64 {
	return value
}

func (h hours) ToDays(value int64) int64 {
	return value / 24
}