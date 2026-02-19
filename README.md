# Runex

CLI tool to detect and analyze runtime errors in real-time.

## Features

- **Real-time Error Detection**: Monitors stdout/stderr while streaming command output
- **Multi-language Support**: Detects errors from Go, Python, Node.js, Java, Ruby, and Rust
- **Stack Trace Extraction**: Automatically extracts stack traces from error output
- **AI-Powered Analysis**: Optional AI analysis for deeper error understanding
- **Colored Output**: Terminal-friendly colored error messages

## Installation

```bash
go install github.com/runex/runex@latest
```

## Usage

```bash
# Run any command with error detection
runex go run main.go

runex python script.py

runex node app.js

# Enable verbose mode
runex -v go run main.go

# Disable colored output
runex --no-color go run main.go

# Enable AI-powered analysis
runex --ai go run main.go

# Force specific language
runex --language go run main.go
```

## Supported Languages

| Language | Error Types |
|----------|-------------|
| Go       | panic, runtime error |
| Python   | exceptions, errors |
| Node.js  | Error, TypeError, ReferenceError, UnhandledPromiseRejection |
| Java     | exceptions |
| Ruby     | errors |
| Rust     | panic, compile errors |

## Configuration

Create `~/.runex.yaml` for persistent configuration:

```yaml
verbose: false
noColor: false
ai: false
language: ""
```

## License

MIT
