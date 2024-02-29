# Structured Application Logging with Context

This is a small reference project to demonstrate various ways to use the `log/slog` package to achieve structured logging, including passing a logger object through the application using the `context` package.

Running the script is as simple as the following command:

```bash
go run ./cmd/logging/main.go
```

I recommend piping the output of the application to `jq` or a similar JSON parser tool to make the output easier to read.

```bash
go run ./cmd/logging/main.go | jq
```
