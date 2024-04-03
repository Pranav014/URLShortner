
# URL Shortener using Go Lang with Redis

This project is a URL shortener implemented in Go, leveraging Redis as the backend data store. It provides a simple and efficient way to generate shortened URLs for long web addresses, facilitating easy sharing while conserving space.

The project utilizes hashing algorithms such as SHA-1 to generate unique short URLs for long addresses, which are then stored in Redis for quick retrieval. It also includes endpoints for accessing statistics on the usage of shortened URLs, as well as health checks to ensure the application is running smoothly.

The application is built using the following technologies:
- Go: A statically typed, compiled programming language designed for simplicity and efficiency.
- Redis: An open-source, in-memory data structure store used as a database, cache, and message broker.
- Gin: A web framework written in Go that provides routing, middleware, and other functionalities for building web applications.


## Features
Endpoints:
- `/shorten`: Accepts a POST request with a JSON payload containing the long URL to be shortened. It generates a unique short URL and stores the mapping in Redis.
- `/{shortURL}`: Redirects to the original long URL associated with the provided short URL.
- `/stats/{shortURL}`: Retrieves statistics for the provided short URL, including the number of times the URL has been accessed and the timestamp of the last access.
- `/health`: Returns the health status of the application.
- `/{shortURL}` : DELETE request to delete the short URL from the database.
- `/Qr/{shortURL}` : GET request to get the QR code of the short URL.
- `/list` : GET request to get the list of all the short URLs.
- `/list/{shortURL}` : GET request to get the details of the short URL.



## Installation

### Prerequisites

Before getting started, ensure you have the following prerequisites installed on your system:

- Go: [Installation Guide](https://golang.org/doc/install)
- Redis: [Installation Guide](https://redis.io/download)


### Steps

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   ```

2. **Install Dependencies**

   ```bash
   go mod tidy
   ```

3. **Start Redis Server**

   Ensure that Redis server is running. Start it by executing the following command in your terminal:

   ```bash
   redis-server
   ```

4. **Verify Redis Installation**

   Open a new terminal window and execute the following command to ensure Redis server is running:

   ```bash
   redis-cli ping
   ```

   You should receive a response of `PONG`, indicating that Redis is up and running.

5. **Run the Application**

   ```bash
   go run main.go
   ```

6. **Access the Application**

   The application should now be running locally. You can access it by navigating to `http://localhost:8080` in your web browser.

## Usage

Once the application is up and running, you can use it to shorten URLs by sending POST requests to the `/shorten` endpoint. You can then access the shortened URLs by navigating to the generated URLs in your browser.

[//]: # (## Contributing)

[//]: # ()
[//]: # (Contributions are welcome! Whether you're fixing a bug, implementing a new feature, or improving documentation, your contributions help make this project better for everyone. Please see the [CONTRIBUTING.md]&#40;CONTRIBUTING.md&#41; file for guidelines on how to contribute.)

## Maintainer

- [Pranav Khismatrao](https://github.com/Pranav014)

---
