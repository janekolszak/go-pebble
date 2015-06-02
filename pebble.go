package pebble

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "strings"
)

type Layout struct {
    Type            string   `json:"type"`
    Title           string   `json:"title"`
    TinyIcon        string   `json:"tinyIcon"`
    Body            string   `json:"body,omitempty"`
    PrimaryColor    string   `json:"primaryColor,omitempty"`
    SecondaryColor  string   `json:"secondaryColor,omitempty"`
    BackgroundColor string   `json:"backgroundColor,omitempty"`
    Headings        []string `json:"headings,omitempty"`
    Paragraphs      []string `json:"paragraphs,omitempty"`
    LastUpdated     string   `json:"lastUpdated,omitempty"`
}

type Notification struct {
    Layout Layout `json:"layout,omitempty"`
    Time   string `json:"time,omitempty"`
}

type Reminder struct {
    Time   string `json:"time"`
    Layout Layout `json:"layout"`
}

type Reminders []Reminder

type Action struct {
    Title      string `json:"title"`
    Type       string `json:"type"`
    LaunchCode int    `json:"launchCode"`
}

type Actions []Action

type Pin struct {
    Id                 string       `json:"id"`
    Time               string       `json:"time"`
    Layout             Layout       `json:"layout"`
    Duration           int          `json:"duration,omitempty"`
    CreateNotification Notification `json:"createNotification,omitempty"`
    UpdateNotification Notification `json:"updateNotification,omitempty"`
    Reminders          Reminders    `json:"reminders,omitempty"`
    Actions            []Action     `json:"actions,omitempty"`
}

type Error struct {
    ErrorCode string `json:"errorCode"`
}

func (pin *Pin) String() string {
    buf, _ := json.Marshal(pin)
    return string(buf)
}

func (pin *Pin) Put(client *http.Client, token string) error {
    address := "https://timeline-api.getpebble.com/v1/user/pins/" + pin.Id

    req, err := http.NewRequest("PUT", address, strings.NewReader(pin.String()))
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-User-Token", token)
    resp, err := client.Do(req)
    defer resp.Body.Close()

    if err != nil {
        return err
    }

    // Check the response:
    if resp.StatusCode != 200 {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return err
        }

        // What's the error sent from Pebble:
        var e Error
        err = json.NewDecoder(strings.NewReader(string(body))).Decode(&e)
        if err != nil {
            return err
        }

        err = errors.New("Error from Pebble Server: " + e.ErrorCode)
        return err
    }

    return nil
}
