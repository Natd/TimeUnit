package timeunit

type days struct {

}

func (d days) ToMilliseconds(value int64) int64 {
	return value * 60 * 60 * 1000 * 24
}

func (d days) ToSeconds(value int64) int64 {
	return value * 60 * 60 * 24
}

func (d days) ToMinutes(value int64) int64 {
	return value * 60 * 24
}

func (d days) ToHours(value int64) int64 {
	return value * 24
}

func (d days) ToDays(value int64) int64 {
	return value
}