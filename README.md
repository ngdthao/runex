# Runex

CLI tool to detect and analyze runtime errors in real-time.

## Features

- **Real-time Error Detection**: Monitors stdout/stderr while streaming command output
- **Multi-language Support**: Detects errors from Go, Python, Node.js, Java, Ruby, and Rust
- **Stack Trace Extraction**: Automatically extracts stack traces from error output
- **AI-Powered Analysis**: Optional AI analysis for deeper error understanding
- **Colored Output**: Terminal-friendly colored error messages

## Introduction

Runex is a command-line interface (CLI) tool that helps you monitor and detect runtime errors automatically when running programs. Instead of having to read through the entire output to find errors, Runex monitors stdout/stderr and immediately notifies you when errors are detected along with the full stack trace.

## Installation

### From source (using Go)

```bash
go install github.com/runex/runex@latest
```

### From GitHub Releases

When releases are available, you can download the appropriate binary for your operating system:

```bash
# macOS (Apple Silicon)
curl -L -o runex https://github.com/runex/runex/releases/latest/download/runex-darwin-arm64
chmod +x runex
sudo mv runex /usr/local/bin/

# macOS (Intel)
curl -L -o runex https://github.com/runex/runex/releases/latest/download/runex-darwin-amd64
chmod +x runex
sudo mv runex /usr/local/bin/

# Linux
curl -L -o runex https://github.com/runex/runex/releases/latest/download/runex-linux-amd64
chmod +x runex
sudo mv runex /usr/local/bin/

# Windows
# Download the .exe file from the Releases page and add to PATH
```

## Quick Start

```bash
# Run any command with error detection
runex go run main.go

runex python script.py

runex node app.js
```

## Detailed Usage

### Flags

| Flag | Description | Example |
|------|-------------|---------|
| `-v` | Enable verbose mode, shows additional debug info | `runex -v go run main.go` |
| `--no-color` | Disable colored output | `runex --no-color python script.py` |
| `--ai` | Enable AI-powered analysis for deeper error understanding | `runex --ai node app.js` |
| `--language` | Manually specify language (go, python, node, java, ruby, rust) | `runex --language go run main.go` |

### Language Examples

#### Go

```bash
# Run Go file
runex go run main.go

# Run tests
runex go test ./...

# Build
runex go build -o myapp .
```

#### Python

```bash
# Run Python script
runex python script.py

# Run module
runex python -m mymodule

# Using pipenv
runex pipenv run python main.py
```

#### Node.js

```bash
# Run JavaScript file
runex node app.js

# Run with npm
runex npm start

# Run with yarn
runex yarn start
```

#### Java

```bash
# Run class file
runex java Main

# Run with Maven
runex mvn exec:java

# Run with Gradle
runex gradle run
```

#### Ruby

```bash
# Run Ruby file
runex ruby script.rb

# Run with Bundler
runex bundle exec ruby script.rb

# Run Rails
runex rails server
```

#### Rust

```bash
# Run Rust project
runex cargo run

# Run tests
runex cargo test

# Build
runex cargo build
```

## Configuration

### Configuration File

Create `~/.runex/config.yaml` to store default configuration:

```yaml
verbose: false
noColor: false
ai: false
language: ""
```

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `RUNEX_VERBOSE` | Enable verbose mode | `false` |
| `RUNEX_NO_COLOR` | Disable colors | `false` |
| `RUNEX_AI` | Enable AI analysis | `false` |
| `RUNEX_LANGUAGE` | Default language | `""` |

### Configuration Priority

Priority order (highest to lowest):
1. Command line flags
2. Environment variables
3. Configuration file (`~/.runex/config.yaml`)

## How It Works

1. **Startup**: Runex receives the command from the user and executes it
2. **Monitoring**: Simultaneously reads stdout and stderr from the running command
3. **Detection**: Analyzes output to find known error patterns for each language
4. **Notification**: When an error is detected, displays the notification with stack trace
5. **Completion**: Command finishes, Runex returns the corresponding exit code

### Supported Languages

| Language | Detected Error Types |
|----------|---------------------|
| Go       | panic, runtime error |
| Python   | exceptions, errors |
| Node.js  | Error, TypeError, ReferenceError, UnhandledPromiseRejection |
| Java     | exceptions |
| Ruby     | errors |
| Rust     | panic, compile errors |

## Real-world Examples

### Example 1: Debug Go Application

```bash
# Assuming you have a main.go with an error
runex -v go run main.go
```

Output:
```
[ERROR] Go error detected
panic: index out of range [1] with length 0

goroutine 1 [running]:
main.main()
    /path/to/main.go:10 +0x...
```

### Example 2: Debug Python Script

```bash
# Run Python script with verbose
runex -v python myscript.py
```

### Example 3: Debug Node.js with AI

```bash
# Analyze Node.js errors using AI
runex --ai node app.js
```

### Example 4: Disable Colors for CI/CD

```bash
# Run in CI environment
runex --no-color go test ./...
```

### Example 5: Force Language Detection

```bash
# Force detection for Go
runex --language go run main.go
```

## License

MIT
