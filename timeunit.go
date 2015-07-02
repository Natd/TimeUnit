package timeunit

type TimeUnit struct {
	Milliseconds milliseconds
	Seconds seconds
	Minutes minutes
	Hours hours
	Days days
}


func New() *TimeUnit {
	return new(TimeUnit)
}