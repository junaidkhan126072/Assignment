# Assignment

Que-Server is a high-performance server application built specifically for ARM64 architecture on Linux. It efficiently handles requests and serves the **Que-Client**, which can be found in the `client` folder.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)
- [Building from Source](#building-from-source)
- [Client Setup](#client-setup)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Getting Started

These instructions will guide you through setting up and running **Que-Server** on your ARM64-based Linux system.

### Prerequisites

Ensure you have the following software installed on your system:

- [Go](https://golang.org/dl/) (for building from source)
- [Git](https://git-scm.com/)
- ARM64-compatible Linux distribution

## Installation

### Clone repo

```bash
git clone https://github.com/junaidkhan126072/assignment.git

go mod tidy
```

### Run Server

You can quickly set up **Que-Server** using running main.go file :

```bash
go run main.go

```

### Run Client

You can quickly set up **Que-Client** using running client.go file :

```bash
cd client/
go run client.go
