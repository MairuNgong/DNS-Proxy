# DNS Proxy with Caching and Domain Blocking

## Overview

This project is developed for the purpose of practicing the Go programming language, particularly focusing on the `net` package. It serves as a learning exercise and a functional implementation of a DNS proxy server.

## Project Scope

The DNS proxy performs the following key functions:

- **Proxy DNS Queries:** Acts as an intermediary between clients and DNS servers by forwarding DNS queries and returning the responses.
- **DNS Caching:** Maintains a local cache of DNS query results to improve response time and reduce external DNS traffic.
- **Domain Blocking:** Allows the specification of domains to be blocked, effectively preventing resolution of those domains.

## Features

- Written in Go using the `net` package.
- Simple and lightweight design for easy deployment and understanding.
- Configurable cache mechanism to reduce DNS lookup latency.
- Customizable list of blocked domains.

## Use Cases

- Educational purposes for learning network programming in Go.
- Basic local DNS filtering and acceleration.
- Small-scale DNS traffic monitoring or control.

## Requirements

- Go 1.18 or higher
- Parse DNS quiries with [github.com/miekg/dns](https://github.com/miekg/dns) library
