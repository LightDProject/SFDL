# SFDL

Server File Definition Language - A configuration language for service packaging, inspired by Terraform.

## Architecture

SFDL is designed as a **library** with LSP support, following the Terraform pattern:

```
sfdl/                    # Core library
├── sfdl.go             # Main API (Parse, Parser)
├── sfdl_test.go        # Tests
├── cmd/
│   ├── lsp/main.go     # LSP server
│   └── cli/main.go     # CLI tool
├── example.sfdl
└── LANGUAGE_SPEC.md
```

## Core Library

```go
import sfdl "github.com/lightDproject/SFDL"

config := sfdl.Config{
    Filename: "config.sfdl",
    Content:  []byte(`name = "test"`),
}
file, err := sfdl.Parse(config)
```

## Commands

### CLI

```bash
go build -o sfdl-cli ./cmd/cli
sfdl-cli -parse config.sfdl
```

### LSP Server

```bash
go build -o sfdl-lsp ./cmd/lsp
sfdl-lsp -addr ":4389"
```

Connect with any LSP-compatible editor (VSCode, Neovim, etc.).

## Features

- **Providers** - Declare and configure external providers
- **Registry** - Configure remote file registries with authentication
- **Functions** - Define executable functions with input/output specs
- **HCL-based** - Built on HashiCorp HCL

## Documentation

See [LANGUAGE_SPEC.md](LANGUAGE_SPEC.md) for language specification.