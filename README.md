# SFDL
Server File Definition Language

## Usage

```bash
go build -o sfdl main.go
./sfdl example.sfdl
```

## Config Format (HCL)

```hcl
name = "my-server"
port = 8080

server "example" {
  host = "localhost"
  timeout = 30
}
```