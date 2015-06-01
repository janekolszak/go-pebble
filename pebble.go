package pebble

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
