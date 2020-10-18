## A simple example of how to cancel/terminate goroutine

- Creation of a cobra CLI
- Simple core with goroutines
- Use of context and signal to catch if master process is terminate

In one terminal :

```bash
go run tests/main/server.go
```

In another terminal :

```bash
go run main.go test
```