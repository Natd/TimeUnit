package timeunit

type Hours struct {

}

func (h Hours) ToMilliseconds(value int64) int64 {
	return value * 60 * 60 * 1000
}

func (h Hours) ToSeconds(value int64) int64 {
	return value * 60 * 60
}

func (h Hours) ToMinutes(value int64) int64 {
	return value * 60
}

func (h Hours) ToHours(value int64) int64 {
	return value
}

func (h Hours) ToDays(value int64) int64 {
	return value / 24
}