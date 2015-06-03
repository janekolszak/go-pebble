package pebble

import (
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "strings"
)

const (
    userPinURL   = "https://timeline-api.getpebble.com/v1/user/pins/"
    sharedPinURL = "https://timeline-api.getpebble.com/v1/shared/pins/"
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

type UserPin struct {
    Pin
    Token string
}

type SharedPin struct {
    Pin
    APIKey string
    Topics string
}

func (pin *Pin) String() string {
    buf, _ := json.Marshal(pin)
    return string(buf)
}

func (pin *Pin) doRequest(client *http.Client, req *http.Request) error {
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

func (uPin *UserPin) address() string {
    return userPinURL + uPin.Id
}

func (sPin *SharedPin) address() string {
    return sharedPinURL + sPin.Id
}

func (uPin *UserPin) Put(client *http.Client) error {

    req, err := http.NewRequest("PUT", uPin.address(), strings.NewReader(uPin.String()))
    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-User-Token", uPin.Token)

    return uPin.doRequest(client, req)
}

func (uPin *UserPin) Delete(client *http.Client) error {

    // TODO: Is body required here?
    req, err := http.NewRequest("DELETE", uPin.address(), strings.NewReader(uPin.String()))
    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-User-Token", uPin.Token)

    return uPin.doRequest(client, req)
}

func (sPin *SharedPin) Put(client *http.Client) error {

    req, err := http.NewRequest("PUT", sPin.address(), strings.NewReader(sPin.String()))
    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-API-Key", sPin.APIKey)
    req.Header.Add("X-Pin-Topics", sPin.Topics)

    return sPin.doRequest(client, req)
}

func (sPin *SharedPin) Delete(client *http.Client) error {

    // TODO: Is body required here?
    req, err := http.NewRequest("DELETE", sPin.address(), strings.NewReader("sPin.String()"))
    if err != nil {
        return err
    }
    req.Header.Add("Content-Type", "application/json")
    req.Header.Add("X-API-Key", sPin.APIKey)

    return sPin.doRequest(client, req)
}
