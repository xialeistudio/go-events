# golang事件处理模块
基于NodeJs events.EventEmitter设计

## 使用

```go
package main

import (
	"fmt"
	"github.com/xialeistudio/go-events/events"
	"time"
)

func main() {
	emitter := events.NewEventEmitter()
	emitter.On("click", func(args ...interface{}) {
		fmt.Println(args...)
	})
	emitter.Emit("click", 1, 2, 3)
	time.Sleep(time.Second)
}
```