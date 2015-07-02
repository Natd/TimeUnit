package timeunit

type Minutes struct {

}

func (m Minutes) ToMilliseconds(value int64) int64 {
	return value * 1000 * 60
}

func (m Minutes) ToSeconds(value int64) int64 {
	return value * 60
}

func (m Minutes) ToMinutes(value int64) int64 {
	return value
}

func (m Minutes) ToHours(value int64) int64 {
	return value / 60
}

func (m Minutes) ToDays(value int64) int64 {
	return value / (60 * 24)
}