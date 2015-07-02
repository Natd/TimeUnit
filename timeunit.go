package timeunit

type TimeUnit struct {
	Milliseconds Milliseconds
	Seconds      Seconds
	Minutes      Minutes
	Hours        Hours
	Days         Days

}


func New() *TimeUnit {
	return new(TimeUnit)
}