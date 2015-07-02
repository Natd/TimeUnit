# TimeUnit

```go
package main

import (
	"github.com/natd/timeunit"
	"fmt"
)

func main() {
	fmt.Println(timeunit.New().Milliseconds.ToSeconds(1000))
	fmt.Println(timeunit.New().Seconds.ToMinutes(60))
	fmt.Println(timeunit.New().Minutes.ToSeconds(1))
	fmt.Println(timeunit.New().Hours.ToMinutes(24))
	fmt.Println(timeunit.New().Days.ToHours(1))
}
```