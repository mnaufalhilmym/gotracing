# gotracing

Application level tracing for Go

## Usage

- Initialization

```go
import "github.com/mnaufalhilmym/gotracing"

logLevel := "debug" // off, error, warn, info, debug, trace

printLevel := gotracing.LevelFilterError
if level := gotracing.NewLevelFilter(logLevel); level.IsOk() { // if you want to parse from a string
    printLevel = level.Unwrap()
}
gotracing.SetMinConsolePrintLevel(printLevel)
gotracing.SetMinStackTrace(gotracing.LevelFilterError)
gotracing.SetMaxProgramCounters(10)

// Add this if you want to save logs to storage
import "github.com/mnaufalhilmym/gotracingstorageinmemory" // anything, but it must implement the storage interface

gotracing.SetMinStoreLevel(gotracing.LevelFilterTrace)
storage = gotracingstorageinmemory.New(100)
gotracing.SetStorage(&storage)
```

- Print log

```go
gotracing.Info("Example of an info log") // 2024-09-25T16:20:51+07:00  INFO: Example of an info log
gotracing.Error("Example of an error log") // 2024-09-25T16:20:51+07:00  ERROR: Example of an error log (with stacktrace)
```

> [!TIP]
> Other usage examples can be seen in the gotracing_test.go file.
