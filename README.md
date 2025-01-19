# Schengen Visa Appointment Checker

## Overview
This is a CLI application to check for available Schengen visa appointments in various countries and cities. The application fetches data from an API and displays the relevant information to the user.

## Features
- Allows users to select a country and city.
- Fetches appointment data from a Schengen visa appointment API.
- Filters appointments based on user input.
- Displays available appointments with detailed information.
- Supports multiple checks in a single session.

## Technologies Used
- **Programming Language**: Go
- **HTTP Client**: To fetch data from the API.
- **JSON Parsing**: To decode API responses.

## Prerequisites
- **Go**: Ensure you have Go installed (version 1.21.13 or newer).

## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/schengen-visa-checker.git
   cd schengen-visa-checker
   ```
2. Build the application:
   ```bash
   go build -o visa-checker
   ```
3. Run the application:
   ```bash
   ./visa-checker
   ```

## How to Use
1. Start the application.
2. Select a country by entering the corresponding number from the list.
3. Select a city by entering the corresponding number from the list.
4. The application will fetch and display available appointments.
5. You can continue checking for other countries and cities or exit the application.

### Example Usage
```
Schengen Visa Checker
=====================================

Select your country:
1. France
2. Netherlands
3. Ireland
...
Your choice (1-8): 1

Select your city:
1. Ankara
2. Istanbul
...
Your choice (1-5): 2

🎉 Appointment found for France!

🏢 Center: VFS Global - Istanbul
📅 Date: 15 January 2025
📋 Category: Tourist Visa
🔗 Meeting Link:
https://example.com/appointment
```

## Configuration
The application uses the following constants for API interaction and messages:

- **APIURL**: The endpoint for fetching visa appointment data.
- **APIErrorMsg**: Error message for API failures.
- **JSONErrorMsg**: Error message for JSON decoding failures.

## File Structure
```
.
├── main.go           # Main application logic
├── constants.go      # Constants for messages and API keys
├── utils.go          # Utility functions (e.g., formatting dates)
├── README.md         # Documentation
```

## Key Functions
### `CheckAppointments(country, city string)`
- Fetches and filters appointments based on the country and city selected by the user.
- Prints available appointments or a message if no appointments are found.

### `GetUserInput() (string, string)`
- Prompts the user to select a country and city.
- Returns the selected country and city names.

### `FormatDate(dateStr string)`
- Converts a date string from `YYYY-MM-DD` format to `DD Month YYYY` format.

## Notes
- Ensure a stable internet connection to fetch data from the API.
- If no appointments are found, the program will notify you.

## License
This project is licensed under the MIT License. See `LICENSE` for details.

