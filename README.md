# DNS Server

A production-oriented DNS server and resolver built from scratch in Go.

The project focuses on understanding and implementing the DNS protocol,
network programming, caching, resolution, concurrency, and production-grade
server design.

## Goals

- Implement DNS packet parsing and encoding
- Support DNS queries over UDP
- Add TCP support and DNS message truncation handling
- Implement local DNS records
- Implement upstream DNS forwarding
- Build a DNS cache with TTL-based expiration
- Support multiple upstream DNS servers
- Add request timeouts and retry handling
- Add concurrent request processing
- Add structured logging and metrics
- Add unit, integration, and end-to-end tests
- Containerize the application
- Add CI/CD with GitHub Actions

## Architecture

```text
DNS Client
    |
    | UDP / TCP
    v
DNS Server
    |
    +--> DNS Protocol Parser
    |
    +--> Resolver
    |      |
    |      +--> Local Records
    |      +--> Cache
    |      +--> Upstream DNS
    |
    +--> Response Encoder
    |
    +--> Metrics / Logging