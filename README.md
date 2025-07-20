# Nightscout Twilio Alerts

A Go service that monitors your [Nightscout](https://www.nightscout.info/) instance and sends urgent low blood sugar alerts via phone call using [Twilio](https://www.twilio.com/).

## Features

- Periodically fetches the latest blood sugar reading from your Nightscout API.
- If the reading is below 70 mg/dL, triggers a phone call alert to your personal number using Twilio.
- Configurable via environment variables.

## Project Structure

```
.
├── .env
├── .gitignore
├── example.env
├── go.mod
├── go.sum
├── Makefile
└── src
    ├── main.go
    ├── nightscout/
    │   └── nightscout.go
    ├── twilio/
    │   └── twilio.go
    └── utils/
        └── config.go
```

## Setup

### 1. Clone the repository

```sh
git clone https://github.com/brkss/nightscout-twillio-alerts.git
cd nightscout-twillio-alerts
```

### 2. Configure Environment Variables

Copy [`example.env`](example.env) to [`.env`](.env) and fill in your credentials:

```sh
cp example.env .env
```

Edit [`.env`](.env):

```env
ACCOUNT_SID=your_twilio_account_sid
AUTH_TOKEN=your_twilio_auth_token
TWILLIO_NUMBER=your_twilio_phone_number
PERSONAL_NUMBER=your_personal_phone_number
NIGHTSCOUT_URL=https://your-nightscout.herokuapp.com
```

### 3. Install Dependencies

```sh
go mod tidy
```

### 4. Build and Run the Service

You can use the provided `Makefile`:

```sh
make build
make run
```

Or run directly:

```sh
go run ./src/main.go
```

## How it Works

- [`main.go`](src/main.go) loads configuration and starts the monitoring loop.
- [`nightscout/nightscout.go`](src/nightscout/nightscout.go) fetches the latest blood sugar entry and triggers alerts.
- [`twilio/twilio.go`](src/twilio/twilio.go) handles phone call alerts via Twilio.
- [`utils/config.go`](src/utils/config.go) loads environment variables.

## Notes

- Make sure your Twilio account is set up and your phone numbers are verified.
- The [`.env`](.env) file is ignored by git for security.
- The polling interval is 5 minutes by default.

## License

MIT