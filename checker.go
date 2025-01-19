package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func CheckAppointments(country, city string) {
	resp, err := http.Get(APIURL)
	if err != nil {
		log.Printf(APIErrorMsg, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	var appointments []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&appointments); err != nil {
		log.Printf(JSONErrorMsg, err)
		return
	}

	availableAppointments := make([]map[string]interface{}, 0)
	for _, appointment := range appointments {
		appointmentDate, ok := appointment[AppointmentDate].(string)
		if !ok || appointmentDate == "" {
			continue
		}

		if appointment[SourceCountry] == TurkeyText &&
			strings.EqualFold(appointment[MissionCountry].(string), country) &&
			strings.Contains(strings.ToLower(appointment[CenterName].(string)), strings.ToLower(city)) {
			availableAppointments = append(availableAppointments, appointment)
		}
	}

	if len(availableAppointments) == 0 {
		fmt.Printf(NoAppointmentsMsg, country, city)
		return
	}

	for _, appointment := range availableAppointments {
		country := GetCountry(appointment[MissionCountry].(string))
		formattedDate := FormatDate(appointment[AppointmentDate].(string))

		message := fmt.Sprintf(AppointmentFoundMsg, country)
		message += fmt.Sprintf(CenterMsg, appointment[CenterName])
		message += fmt.Sprintf(DateMsg, formattedDate)
		message += fmt.Sprintf(CategoryMsg, appointment[VisaCategory])
		if subcategory, ok := appointment[VisaSubCategory].(string); ok && subcategory != "" {
			message += fmt.Sprintf(SubCategoryMsg, subcategory)
		}
		message += fmt.Sprintf(MeetingLinkMsg, appointment[BookNowLink])

		fmt.Println(message)
	}
}
