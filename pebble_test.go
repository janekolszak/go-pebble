package pebble

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

const (
    minimalPin = `{"id":"example-pin-generic-1","time":"2015-03-19T18:00:00Z","layout":{"type":"genericPin","title":"News at 6 o'clock","tinyIcon":"system://images/NOTIFICATION_FLAG"}}`
    maximalPin = `{"id":"meeting-453923","time":"2015-03-19T15:00:00Z","layout":{"type":"genericPin","title":"Client Meeting","tinyIcon":"system://images/TIMELINE_CALENDAR","body":"Meeting in Kepler at 4:00pm. Topic: discuss pizza toppings for party."},"duration":60,"createNotification":{"layout":{"type":"genericNotification","title":"New Item","tinyIcon":"system://images/NOTIFICATION_FLAG","body":"A new appointment has been added to your calendar at 4pm."}},"updateNotification":{"layout":{"type":"genericNotification","title":"Reminder","tinyIcon":"system://images/NOTIFICATION_FLAG","body":"The meeting has been rescheduled to 4pm."},"time":"2015-03-19T16:00:00Z"},"reminders":[{"time":"2015-03-19T14:45:00Z","layout":{"type":"genericReminder","title":"Meeting in 15 minutes","tinyIcon":"system://images/TIMELINE_CALENDAR"}},{"time":"2015-03-19T14:55:00Z","layout":{"type":"genericReminder","title":"Meeting in 5 minutes","tinyIcon":"system://images/TIMELINE_CALENDAR"}}],"actions":[{"title":"View Schedule","type":"openWatchApp","launchCode":15},{"title":"Show Directions","type":"openWatchApp","launchCode":22}]}`
)

func TestStringMinimal(t *testing.T) {
    assert.True(t, true, "True is true!")
    layout := Layout{
        Type:     "genericPin",
        Title:    "News at 6 o'clock",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
    }

    pin := Pin{
        Id:     "example-pin-generic-1",
        Time:   "2015-03-19T18:00:00Z",
        Layout: &layout,
    }

    assert.Equal(t, pin.String(), minimalPin)
}

func TestStringMaximal(t *testing.T) {
    layout := Layout{
        Type:     "genericPin",
        Title:    "Client Meeting",
        TinyIcon: "system://images/TIMELINE_CALENDAR",
        Body:     "Meeting in Kepler at 4:00pm. Topic: discuss pizza toppings for party.",
    }

    createLayout := Layout{
        Type:     "genericNotification",
        Title:    "New Item",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
        Body:     "A new appointment has been added to your calendar at 4pm.",
    }

    createNotification := Notification{
        Layout: &createLayout,
    }

    updateLayout := Layout{
        Type:     "genericNotification",
        Title:    "Reminder",
        TinyIcon: "system://images/NOTIFICATION_FLAG",
        Body:     "The meeting has been rescheduled to 4pm.",
    }

    updateNotification := Notification{
        Time:   "2015-03-19T16:00:00Z",
        Layout: &updateLayout,
    }

    r1Layout := Layout{
        Type:     "genericReminder",
        Title:    "Meeting in 15 minutes",
        TinyIcon: "system://images/TIMELINE_CALENDAR",
    }

    r1 := Reminder{
        Time:   "2015-03-19T14:45:00Z",
        Layout: &r1Layout,
    }

    r2Layout := Layout{
        Type:     "genericReminder",
        Title:    "Meeting in 5 minutes",
        TinyIcon: "system://images/TIMELINE_CALENDAR",
    }

    r2 := Reminder{
        Time:   "2015-03-19T14:55:00Z",
        Layout: &r2Layout,
    }

    a1 := Action{
        Title:      "View Schedule",
        Type:       "openWatchApp",
        LaunchCode: 15,
    }

    a2 := Action{
        Title:      "Show Directions",
        Type:       "openWatchApp",
        LaunchCode: 22,
    }

    pin := Pin{
        Id:                 "meeting-453923",
        Time:               "2015-03-19T15:00:00Z",
        Duration:           60,
        Layout:             &layout,
        CreateNotification: &createNotification,
        UpdateNotification: &updateNotification,
        Reminders:          &Reminders{r1, r2},
        Actions:            &Actions{a1, a2},
    }

    assert.Equal(t, pin.String(), maximalPin)
}
