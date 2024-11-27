# Weather API Service

A high-performance weather information service built with Go, featuring request deduplication using singleflight pattern and SQLite caching.

## Features

![](./views/weather_diagram.png)

- **Request Deduplication**: Implements the singleflight pattern to prevent duplicate API calls for the same weather data
- **Caching System**: Uses SQLite database to cache weather information
- **Fast Response Times**: Built with the high-performance Fiber web framework
- **Clean Architecture**: Follows clean architecture principles with clear separation of concerns

## Tech Stack

- **Go**: Primary programming language
- **Fiber**: Web framework for handling HTTP requests
- **GORM**: ORM library for database operations
- **SQLite**: Database for caching weather data
- **Singleflight**: For request deduplication

## Project Structure

```
.
├── app/
│   ├── handler/     # HTTP request handlers
│   ├── service/     # Business logic layer
│   ├── store/       # Data access layer
│   └── weather/     # Weather-related components
├── config/          # Configuration management
├── views/           # Template files
├── main.go         # Application entry point
├── go.mod          # Go module file
└── go.sum          # Go module checksums
```

## Prerequisites

- Go 1.16 or higher
- SQLite

## Setup and Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/batuhannoz/paribu-case.git
   cd paribu-case
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the application:
   ```bash
   go run main.go
   ```

The server will start on port 3000 by default.

## API Endpoints

### Get Weather Information
```
GET /weather?q=<location>
```

Returns weather information for the specified parameters.

#### Response:
```json
{
  "location"    : "<location>",
  "temperature" : "<average-temp>"
}
```

## Architecture

The project follows a clean architecture pattern with the following layers:

- **Handlers**: Handle HTTP requests and responses
- **Services**: Contain business logic and implement singleflight pattern
- **Store**: Manages data persistence and caching using SQLite
