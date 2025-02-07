# Basic Information API

## Overview

The Basic Information API is a simple RESTful API that provides essential information, including an email address, the current date and time in ISO 8601 format (UTC), and a GitHub URL.

## Base URL

```
https://basic-information-api-du6s.onrender.com/

```

## Endpoints

1. Get Basic Information

Request

- Method: `GET`

- Endpoint: `/`

- Headers: `None`

- Query Parameters: `None`

Response

- Status Code: `200 OK`

- Content-Type: `application/json`

Example Response

```json
{
  "email": "berylatieno30@gmail.com",
  "current_datetime": "2024-02-07T12:34:56Z",
  "github_url": "https://github.com/BerylCAtieno/basic-information-api"
}
```

## Setup and Installation

### Prerequisites

- Go (latest version recommended)

- Git

### Clone the Repository

```bash
git clone https://github.com/BerylCAtieno/basic-information-api.git
cd basic-information-api
```

### Install Dependencies

```bash
go mod tidy
```

### Run the Server

```bash
go run main.go
```

The API will be running at `https://basic-information-api-du6s.onrender.com/`

## Testing the API

You can test the API using:

#### cURL Command

```bash
curl -X GET http://localhost:8080/
```

#### Postman

- Open Postman.

- Create a new GET request.

- Enter `https://basic-information-api-du6s.onrender.com/` as the request URL.

- Click Send.

You should receive the JSON response.


## Author

`Beryl Atieno`

GitHub: (BerylCAtieno)[https://github.com/BerylCAtieno]

