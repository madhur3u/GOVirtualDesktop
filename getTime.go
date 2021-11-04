package main

// THIS IS MADE USING LINUX (UBUNTU 20.04) AND MIGHT NOT WORK WELL WITH WINDOWS

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getTime() string {

	weekday := [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	response, _ := http.Get("http://worldtimeapi.org/api/timezone/Asia/Kolkata")
	body, _ := ioutil.ReadAll(response.Body)
	time, _ := UnmarshalWelcome(body)

	s1 := time.Datetime
	s2 := time.DayOfWeek

	return weekday[s2] + "    " + s1[:10] + "    " + s1[11:19]
}

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}
func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Abbreviation string      `json:"abbreviation"`
	ClientIP     string      `json:"client_ip"`
	Datetime     string      `json:"datetime"`
	DayOfWeek    int64       `json:"day_of_week"`
	DayOfYear    int64       `json:"day_of_year"`
	Dst          bool        `json:"dst"`
	DstFrom      interface{} `json:"dst_from"`
	DstOffset    int64       `json:"dst_offset"`
	DstUntil     interface{} `json:"dst_until"`
	RawOffset    int64       `json:"raw_offset"`
	Timezone     string      `json:"timezone"`
	Unixtime     int64       `json:"unixtime"`
	UTCDatetime  string      `json:"utc_datetime"`
	UTCOffset    string      `json:"utc_offset"`
	WeekNumber   int64       `json:"week_number"`
}
