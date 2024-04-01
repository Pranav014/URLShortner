
# URL Shortener using Go Lang with Redis

This project is a URL shortener implemented in Go, leveraging Redis as the backend data store. It provides a simple and efficient way to generate shortened URLs for long web addresses, facilitating easy sharing while conserving space.

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
