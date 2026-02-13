# Quantum-Secure Chat Go

High-performance chat infrastructure resistant against quantum computing attacks.

## Features

- **Post-Quantum Cryptography**: Lattice-based KEM integration
- **Perfect Forward Secrecy**: Ephemeral per-session keys
- **Hybrid Encryption**: ECC + PQC combined security
- **High Concurrency**: Goroutine-per-client model
- **Scalable Architecture**: Horizontal scaling support

## Architecture

```
quantum-secure-chat-go/
├── cmd/              # Server and client applications
├── internal/         # Internal packages
├── pkg/
│   ├── crypto/       # Cryptographic implementations
│   ├── kem/          # Key encapsulation mechanisms
│   └── chat/         # Chat protocol
├── README.md         # This file
└── .github/
    └── workflows/
        └── ci.yml   # CI/CD pipeline
```

## Usage

### Build

```bash
go build -o server ./cmd/server
go build -o client ./cmd/client
```

### Run Server

```bash
./server
```

### Run Client

```bash
./client
```

## Security

This implementation uses hybrid cryptography:
- **Traditional ECC**: X25519 key exchange
- **Post-Quantum**: Lattice-based KEM (simulated)
- **Symmetric**: AES-256-GCM encryption

## Requirements

- Go 1.21+
- Cryptographic libraries

## Disclaimer

**Educational Use Only**: This tool is for security learning and authorized testing only.

## License

MIT
