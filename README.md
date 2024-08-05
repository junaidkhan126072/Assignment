# Assignment

Que-Server is a high-performance server application built specifically for ARM64 architecture on Linux. It efficiently handles requests and serves the **Que-Client**, which can be found in the `client` folder.

## Table of Contents

- [Getting Started](#getting-started)
- [Installation](#installation)
- [Usage](#usage)


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
```

## Usage

### Run Client

You can quickly set up **Que-Client** using running client.go file it take parameters in two way either using file or flags :

```bash
go run client.go -file=commands.json 
```
command.json file format
```bash 

[
    {"action": "addItem", "key": "junaid", "value": "khan"},
    {"action": "addItem", "key": "1", "value": "1"},
    {"action": "addItem", "key": "1", "value": "1"},
    {"action": "addItem", "key": "junaid2", "value": "khan2"},
    {"action": "addItem", "key": "junaid3", "value": "khan3"},
    {"action": "deleteItem", "key": "junaid2"},
    {"action": "getItem", "key": "junaid2"},
    {"action": "getAllItems"}
]

```

You can quickly set up **Que-Client** using running client.go file it take parameters in two way either using file or flags :

```bash
go run client.go -action=addItem -key=junaid -value=khan
go run client.go -action=deleteItem -key=junaid
go run client.go -action=getItem 
go run client.go -action=getAllItems 

```