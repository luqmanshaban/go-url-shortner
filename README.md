
# URL Shortener

A simple URL shortener built with Go, Gin, and MongoDB. This project allows users to convert long URLs into shorter, easily shareable links. The project is live and deployed on [url.tanelt.com](http://url.tanelt.com).

## Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
- [Environment Variables](#environment-variables)
- [Endpoints](#endpoints)
- [Project Structure](#project-structure)
- [License](#license)

## Features

- Generate short URLs for any valid long URL
- Retrieve and redirect to the original URL using the short URL
- Prevent duplicate entries for the same long URL
- Simple and responsive HTML form for URL submission

## Tech Stack

- **Go**: Backend language
- **Gin**: Web framework for routing and request handling
- **MongoDB**: Database for storing URLs and short codes


## Getting Started

### Prerequisites

- Go 1.17+ installed
- MongoDB (local or cloud instance)
- Git (for cloning the repository)
- [Docker](https://www.docker.com/get-started) (optional, for containerization)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/luqmanshaban/go-url-shortener.git
   cd go-url-shortener
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file in the root directory to store environment variables (see [Environment Variables](#environment-variables) section).

4. Start the server:
   ```bash
   go run main.go
   ```

The server should now be running on `http://localhost:4000`.

## Environment Variables

To run this project, create a `.env` file in the root directory and add the following:

```plaintext
MONGODB_URI=mongodb://your_mongodb_uri_here
```

## Endpoints

### HTML Form

- **GET /** - Renders the HTML form for submitting a long URL.

### API Endpoints

- **POST /submit**
  - **Description**: Accepts a long URL, generates a short URL, and saves it to the database.
  - **Body**: `url=<long_url>`
  - **Response**:
    ```json
    {
      "ShortUrl": "http://url.tanelt.com/<short_url_code>"
    }
    ```

- **GET /url/:shortUrl**
  - **Description**: Redirects to the original URL based on the provided `shortUrl`.
  - **Params**: `shortUrl` - The generated short code.
  - **Response**: Redirects to the original URL if found, otherwise returns a `404` error.

## Project Structure

```plaintext
.
├── main.go               # Main application file
├── go.mod                # Go module dependencies
├── static/
│   └── form.html         # HTML form for URL input
└── README.md             # Project documentation
```



## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

