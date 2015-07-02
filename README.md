# TimeUnit

```go
package main

import (
	"github.com/natd/timeunit"
	"fmt"
)

func main() {
	fmt.Println(timeunit.Milliseconds{}.ToSeconds(1000))
	fmt.Println(timeunit.Seconds{}.ToMinutes(60))
	fmt.Println(timeunit.Minutes{}.ToSeconds(1))
	fmt.Println(timeunit.Hours{}.ToMinutes(24))
	fmt.Println(timeunit.Days{}.ToHours(1))

	TimeLeft(timeunit.Milliseconds{})
	TimeLeft(timeunit.Seconds{})
	TimeLeft(timeunit.Minutes{})
	TimeLeft(timeunit.Hours{})
	TimeLeft(timeunit.Days{})

}


func TimeLeft(timeUnit timeunit.TimeUnit){
	fmt.Println(timeUnit.ToMilliseconds(1))
}
```