package bommodel

import (
	"encoding/xml"
	"time"
)

type Product struct {
	XMLName  xml.Name `xml:"product"`
	AMOC     Amoc     `xml:"amoc"`
	Forecast Forecast `xml:"forecast"`
}

type Amoc struct {
	Source                    Source `xml:"source"`
	Identifier                string `xml:"identifier"`
	IssueTimeUTC              string `xml:"issue-time-utc"`
	IssueTimeLocal            Time   `xml:"issue-time-local"`
	SentTime                  string `xml:"sent-time"`
	ExpiryTime                string `xml:"expiry-time"`
	ValidityBgnTimeLocal      Time   `xml:"validity-bgn-time-local"`
	ValidityEndTimeLocal      Time   `xml:"validity-end-time-local"`
	NextRoutineIssueTimeUTC   string `xml:"next-routine-issue-time-utc"`
	NextRoutineIssueTimeLocal Time   `xml:"next-routine-issue-time-local"`
	Status                    string `xml:"status"`
	Service                   string `xml:"service"`
	SubService                string `xml:"sub-service"`
	ProductType               string `xml:"product-type"`
	Phase                     string `xml:"phase"`
}

// Source represents the <source/> element within <amoc/>
type Source struct {
	Sender     string `xml:"sender"`
	Region     string `xml:"region"`
	Office     string `xml:"office"`
	Copyright  string `xml:"copyright"`
	Disclaimer string `xml:"disclaimer"`
}

// Forecast represents the <forecast/> element
type Forecast struct {
	Area []Area `xml:"area"`
}

// Area represents the <area/> element within <forecast/>
type Area struct {
	AAC             string           `xml:"aac,attr"`
	Description     string           `xml:"description,attr"`
	Type            string           `xml:"type,attr"`
	ParentAAC       string           `xml:"parent-aac,attr,omitempty"`
	ForecastPeriods []ForecastPeriod `xml:"forecast-period"`
}

// ForecastPeriod represents the <forecast-period/> element within <area/>
type ForecastPeriod struct {
	Index          string    `xml:"index,attr,omitempty"`
	StartTimeLocal Time      `xml:"start-time-local,attr"`
	EndTimeLocal   Time      `xml:"end-time-local,attr"`
	StartTimeUTC   string    `xml:"start-time-utc,attr"`
	EndTimeUTC     string    `xml:"end-time-utc,attr"`
	Text           []Text    `xml:"text"`
	Elements       []Element `xml:"element"`
}

// Text represents the <text/> element within <forecast-period/>
type Text struct {
	Type string `xml:"type,attr"`
	Text string `xml:",chardata"`
}

// Element represents the <element/> element within <forecast-period/>
type Element struct {
	Type  string `xml:"type,attr"`
	Units string `xml:"units,attr,omitempty"`
	Text  string `xml:",chardata"`
}

// Time represents a time element with timezone attribute
type Time struct {
	time.Time
	Tz string `xml:"tz,attr,omitempty"`
}

// UnmarshalXMLAttr customizes the unmarshalling process for the Time type
func (t *Time) UnmarshalXMLAttr(attr xml.Attr) error {
	// Assuming your time format is "2006-01-02T15:04:05Z"
	parsedTime, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	t.Time = parsedTime
	return nil
}
