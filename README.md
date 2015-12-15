# go-pebble [![Build Status](https://travis-ci.org/janekolszak/go-pebble.svg)](https://travis-ci.org/janekolszak/go-pebble)
Interacting with the Pebble's timeline in Go.

### Setup
<pre>go get -u github.com/janekolszak/go-pebble</pre>

### Test
<pre>go test github.com/janekolszak/go-pebble</pre>

### Example
```go
package main

import (
    "fmt"
    "github.com/janekolszak/go-pebble"
    "net/http"
)

func main() {
    layout := pebble.Layout{
        Type:     "genericPin",
        Title:    "Title",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
        Body:     "Body",
    }

    creationLayout := pebble.Layout{
        Type:     "genericPin",
        Title:    "Creation Title",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
        Body:     "Creation Body",
    }

    creationNotification := pebble.Notification{
        Layout: &creationLayout,
    }

    updateLayout := pebble.Layout{
        Type:     "genericPin",
        Title:    "Update Title",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
        Body:     "Update Body",
    }

    updateNotification := pebble.Notification{
        Layout: &updateLayout,
        Time:   time.Now().Format(time.RFC3339),
    }

    reminderLayout := pebble.Layout{
        Type:     "genericReminder",
        Title:    "Reminder Title",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
    }

    reminder := pebble.Reminder{
        Time:   time.Now().Format(time.RFC3339),
        Layout: &reminderLayout,
    }

    pin := pebble.Pin{
        Id:                 "UNIQUE ID",
        Time:               time.Now().Format(time.RFC3339),
        Layout:             &layout,
        CreateNotification: &creationNotification,
        UpdateNotification: &updateNotification,
        Reminders:          &pebble.Reminders{reminder},
    }

    uPin.userPin = pebble.UserPin{
        Pin:   pin,
        Token: "TOKEN",
    }

    fmt.Println(pin.String())


    client := &http.Client{
        CheckRedirect: redirectPolicyFunc,
    }
    uPin.Put(client)
    uPin.Delete(client)
}
```
