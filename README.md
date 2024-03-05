# Crypto Server

Crypto Server to get real time data for any crypto

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)

## Installation

To install the Crypto Server, use the following command:

```bash
go get -u github.com/saifaligts/crypto-server
```

## Usage

Run the application via Go:

```bash
go mod tidy
go run main.go
```

Run the application via Docker:

```bash
cd crypto-server
docker build -t crypto-server .
docker run -p 8080:8080 crypto-server
```