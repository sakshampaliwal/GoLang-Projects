# URL Shortener Service

This is a URL shortener service built using Go and Gin framework. It provides functionality to shorten long URLs into compact, unique identifiers that users can use to access the original URLs.

## Features
- Shorten long URLs into compact codes
- Redirect users from short URLs to original long URLs
- Simple and easy-to-use API

## Prerequisites
- Go (v1.16 or higher)
- MongoDB or MongoDB container also works

## Building and Running

To run the URL shortener service, follow these steps:

1. **Build the Application:**
   - Run the following command to build the application:
     ```bash
     go build .
     ```
   This command will generate an executable file.

2. **Run the Application:**
   - Execute the generated executable by running:
     ```bash
     ./urlshortner
     ```
   This command will start the URL shortener service, and it will be accessible at the specified port (default port is 8000).

3. **Accessing the Service:**
   Once the service is running, you can access it via HTTP requests using the specified endpoints.

4. **Testing:**
   Test the functionality of the service by sending requests and verifying the responses.

Make sure you have all the prerequisites installed and configured before building and running the application.
Make sure MongoDB is running.

## API Endpoints

### Shorten URL
- **URL:** `/api/shorten`
- **Method:** `POST`
- **Request Body:**
  ```json
  {
    "url": "https://example.com/long-url"
  }


## Setting Up MongoDB with Docker (Optional)

If MongoDB is not installed locally, you can use a Docker container to run MongoDB. Execute the following command to create and run a MongoDB Docker container:

```bash
docker run --name my-mongodb -p 27017:27017 -d mongo
```

## Shortening URL via POST Request

To shorten a URL, send a POST request to the following endpoint:

- **Endpoint:** `POST http://localhost:8000/v1/url/short`
- **Request Body:**
  ```json
  {
      "longUrl": "https://www.example.com/very/long/url/that/needs/shortening"
  }

## Checking Shortened URL

After shortening the URL, you can check the shortened URL by visiting:

- **URL:** `http://localhost:8000/v1/url/<short_code>`

Replace `<short_code>` with the code generated after shortening the URL.

Example:
If the short code generated after shortening is `abc123`, then you can access the original URL by visiting:
```http
http://localhost:8000/v1/url/abc123
```

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to:

- Open an issue to report a problem or suggest new features.
- Submit a pull request to contribute directly to the codebase.

### How to Contribute

1. **Fork the repository:** Click on the "Fork" button at the top right corner of the repository's page on GitHub.
2. **Clone the forked repository:** Clone the repository to your local machine using the following command: