package main

const (
	WelcomeMsg          = "Schengen Visa Checker"
	DashLine            = "====================================="
	SelectCountryMsg    = "\nSelect your country:"
	SelectCityMsg       = "\nSelect your city:"
	CountryChoiceMsg    = "\nYour choice (1-8): "
	CityChoiceMsg       = "\nYour choice (1-5): "
	ContinueCheckMsg    = "\nWould you like to do another check? (Y/N): "
	ExitingMsg          = "Exiting the program..."
	NoAppointmentsMsg   = "No suitable appointments found for: %s - %s"
	AppointmentFoundMsg = "🎉 Appointment found for %s!\n\n"
	CenterMsg           = "🏢 Center: %s\n"
	DateMsg             = "📅 Date: %s\n"
	CategoryMsg         = "📋 Category: %s\n"
	SubCategoryMsg      = "📝 Sub Category: %s\n"
	MeetingLinkMsg      = "\n🔗 Meeting Link:\n%s"
	APIErrorMsg         = "API didn't respond: %v"
	JSONErrorMsg        = "JSON decode error: %v"
	Yes                 = "Y"
	InvalidCountryMsg   = "Invalid country selection!"
	InvalidCityMsg      = "Invalid city selection!"

	APIURL = "https://api.schengenvisaappointments.com/api/visa-list/?format=json"

	MissionCountry  = "mission_country"
	AppointmentDate = "appointment_date"
	SourceCountry   = "source_country"
	CenterName      = "center_name"
	VisaCategory    = "visa_category"
	VisaSubCategory = "visa_subcategory"
	BookNowLink     = "book_now_link"
	TurkeyText      = "Turkiye"
)

var countries = map[string]string{
	"1": "France",
	"2": "Netherlands",
	"3": "Ireland",
	"4": "Sweden",
	"5": "Czechia",
	"6": "Finland",
	"7": "Denmark",
	"8": "Luxembourg",
}

var citiesTR = map[string]string{
	"1": "Ankara",
	"2": "Istanbul",
	"3": "Izmir",
	"4": "Mugla",
	"5": "Eskisehir",
}

var months = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}
