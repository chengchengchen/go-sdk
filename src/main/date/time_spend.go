package date

import (
	"fmt"
	"time"
)

func timeSpend(op func(res *int)) int {
	start := time.Now()
	res := 0
	op(&res)
	fmt.Printf("op spend time: %v second", time.Since(start).Seconds())
	return res
}
