package timeunit

type Days struct {

}

func (d Days) ToMilliseconds(value int64) int64 {
	return value * 60 * 60 * 1000 * 24
}

func (d Days) ToSeconds(value int64) int64 {
	return value * 60 * 60 * 24
}

func (d Days) ToMinutes(value int64) int64 {
	return value * 60 * 24
}

func (d Days) ToHours(value int64) int64 {
	return value * 24
}

func (d Days) ToDays(value int64) int64 {
	return value
}