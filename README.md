# go-pebble
Interacting with the Pebble's timeline in Go

### Setup
<pre>go get github.com/janekolszak/go-pebble</pre>

### Example
```go
package main

import (
    "encoding/json"
    "fmt"
    "github.com/janekolszak/go-pebble"
)

func main() {
    pin := pebble.Pin{
        Id:                 id,
        Time:               time,
        Layout:             layout,
        Duration:           duration,
        CreateNotification: creationNotification,
        UpdateNotification: updateNotification,
        Reminders:          pebble.Reminders{reminder},
        Actions:            pebble.Actions{action}
    }
    pinJSON, _ := json.Marshal(pin)
    fmt.Println(pinJSON)
}
```