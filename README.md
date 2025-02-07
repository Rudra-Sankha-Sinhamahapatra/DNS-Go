
# DNS Server in Go

This project implements a simple DNS server in Go that listens for DNS queries and resolves domain names to IP addresses. The server can either resolve known records from a local mapping or perform external DNS lookups for unknown domains.

## Features
- Listen for DNS queries on a specified UDP port.
- Resolve domain names to IP addresses.
- Cache resolved domain names for future requests.
- Perform external DNS resolution for unknown domains.
- Configurable logging and port settings via `.env` file.

---

## Table of Contents
1. [Installation](#installation)
2. [Configuration](#configuration)
3. [Usage](#usage)
4. [Testing](#testing)
5. [Project Structure](#project-structure)
6. [License](#license)

---

## Installation

To run the DNS server locally, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/dns-server-go.git
   cd dns-server-go
   ```

2. Ensure you have [Go](https://golang.org/dl/) installed on your machine.

3. Install the required dependencies:
   ```bash
   go mod tidy
   ```

---

## Configuration

The server uses environment variables for configuration. Create a `.env` file in the root directory and define the following variables:

```bash
DNS_SERVER_PORT=9000
DNS_LOG_FILE=dns_server.log
```

- `DNS_SERVER_PORT`: The port on which the DNS server will listen.
- `DNS_LOG_FILE`: The file where logs will be stored.

Make sure to load the environment variables before running the server:
```bash
source .env
```

---

## Usage

1. **Start the DNS Server:**
   To start the DNS server, simply run the following command:
   ```bash
   go run ./src
   ```
   The server will start listening on `127.0.0.1:9000` (or the port specified in `.env`).

2. **Test the Server with `dig`:**
   Use the `dig` command to test your server's DNS resolution. For example:
   ```bash
   dig @127.0.0.1 -p 9000 google.com
   ```

   The server will respond with an IP address if the domain is resolved in the local records. If the domain is not found, it will attempt an external resolution and cache the result for future queries.

3. **External Domain Resolution:**
   If the server cannot find a domain in its local records, it will perform an external DNS query to resolve the IP. For example, if `google.com` is not in the local records, the server will resolve it via an external DNS service.

---

## Testing

You can test the DNS server via `dig`, as shown in the usage section. Here’s how to test with `dig`:

1. Open a terminal or command prompt.
2. Run:
   ```bash
   dig @127.0.0.1 -p 9000 rudrasankha.in
   ```

   If the domain is found in the server's records, you will get the resolved IP. If not, the server will attempt to resolve it externally and return the IP.

---

## Project Structure

The project is organized as follows:

```
dns-server-go/
│── src/
│   ├── main.go        # Entry point of the DNS server
│   ├── records.go     # Manages local DNS records
│   ├── resolve.go     # Handles external DNS resolution
│   ├── sendDns.go     # Constructs and sends DNS responses
│   ├── serveDns.go    # Handles incoming DNS queries and routing
│   ├── utils/
│       ├── config.go  # Loads environment variables from .env
│       ├── logger.go  # Configures logging
│
├── .env                # Configuration file for server settings
├── go.mod              # Go module definition
└── README.md           # Project documentation
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

