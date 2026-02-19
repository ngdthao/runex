# Runex

Công cụ CLI để phát hiện và phân tích lỗi runtime theo thời gian thực.

## Tính năng

- **Phát hiện lỗi thời gian thực**: Giám sát stdout/stderr khi streaming output của command
- **Hỗ trợ đa ngôn ngữ**: Phát hiện lỗi từ Go, Python, Node.js, Java, Ruby và Rust
- **Trích xuất Stack Trace**: Tự động trích xuất stack trace từ error output
- **Phân tích AI**: Tùy chọn phân tích bằng AI để hiểu sâu hơn về lỗi
- **Màu sắc cho Terminal**: Thông báo lỗi có màu sắc thân thiện với terminal

## Giới thiệu

Runex là một công cụ dòng lệnh (CLI) giúp bạn giám sát và phát hiện lỗi runtime một cách tự động khi chạy các chương trình. Thay vì phải đọc toàn bộ output để tìm lỗi, Runex sẽ theo dõi stdout/stderr và ngay lập tức thông báo khi phát hiện lỗi cùng với stack trace đầy đủ.

## Cài đặt

### Từ source (sử dụng Go)

```bash
go install github.com/runex/runex@latest
```

### Từ GitHub Releases

Khi có bản phát hành, bạn có thể tải binary phù hợp với hệ điều hành của mình:

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
# Tải file .exe từ trang Releases và thêm vào PATH
```

## Bắt đầu nhanh

```bash
# Chạy bất kỳ command nào với phát hiện lỗi
runex go run main.go

runex python script.py

runex node app.js
```

## Sử dụng chi tiết

### Các cờ (flags)

| Cờ | Mô tả | Ví dụ |
|-----|-------|-------|
| `-v` | Bật chế độ verbose, hiển thị thêm thông tin debug | `runex -v go run main.go` |
| `--no-color` | Tắt màu sắc trong output | `runex --no-color python script.py` |
| `--ai` | Bật phân tích AI để hiểu sâu hơn về lỗi | `runex --ai node app.js` |
| `--language` | Chỉ định ngôn ngữ thủ công (go, python, node, java, ruby, rust) | `runex --language go run main.go` |

### Ví dụ theo ngôn ngữ

#### Go

```bash
# Chạy file Go
runex go run main.go

# Chạy test
runex go test ./...

# Build
runex go build -o myapp .
```

#### Python

```bash
# Chạy script Python
runex python script.py

# Chạy module
runex python -m mymodule

# Sử dụng pipenv
runex pipenv run python main.py
```

#### Node.js

```bash
# Chạy file JavaScript
runex node app.js

# Chạy với npm
runex npm start

# Chạy với yarn
runex yarn start
```

#### Java

```bash
# Chạy file class
runex java Main

# Chạy với Maven
runex mvn exec:java

# Chạy với Gradle
runex gradle run
```

#### Ruby

```bash
# Chạy file Ruby
runex ruby script.rb

# Chạy với Bundler
runex bundle exec ruby script.rb

# Chạy Rails
runex rails server
```

#### Rust

```bash
# Chạy project Rust
runex cargo run

# Chạy test
runex cargo test

# Build
runex cargo build
```

## Cấu hình

### File cấu hình

Tạo file `~/.runex/config.yaml` để lưu cấu hình mặc định:

```yaml
verbose: false
noColor: false
ai: false
language: ""
```

### Biến môi trường

| Biến | Mô tả | Giá trị mặc định |
|------|-------|------------------|
| `RUNEX_VERBOSE` | Bật chế độ verbose | `false` |
| `RUNEX_NO_COLOR` | Tắt màu sắc | `false` |
| `RUNEX_AI` | Bật phân tích AI | `false` |
| `RUNEX_LANGUAGE` | Ngôn ngữ mặc định | `""` |

### Ưu tiên cấu hình

Thứ tự ưu tiên (từ cao đến thấp):
1. Cờ dòng lệnh (command line flags)
2. Biến môi trường (environment variables)
3. File cấu hình (`~/.runex/config.yaml`)

## Cách hoạt động

1. **Khởi động**: Runex nhận command từ người dùng và thực thi nó
2. **Giám sát**: Đồng thời đọc stdout và stderr từ command đang chạy
3. **Phát hiện**: Phân tích output để tìm các pattern lỗi đã biết của từng ngôn ngữ
4. **Thông báo**: Khi phát hiện lỗi, hiển thị thông báo với stack trace
5. **Kết thúc**: Command kết thúc, Runex trả về exit code tương ứng

### Ngôn ngữ được hỗ trợ

| Ngôn ngữ | Loại lỗi phát hiện |
|----------|-------------------|
| Go       | panic, runtime error |
| Python   | exceptions, errors |
| Node.js  | Error, TypeError, ReferenceError, UnhandledPromiseRejection |
| Java     | exceptions |
| Ruby     | errors |
| Rust     | panic, compile errors |

## Ví dụ thực tế

### Ví dụ 1: Debug ứng dụng Go

```bash
# Giả sử bạn có file main.go với lỗi
runex -v go run main.go
```

Output:
```
[ERROR] Phát hiện lỗi Go
panic: index out of range [1] with length 0

goroutine 1 [running]:
main.main()
    /path/to/main.go:10 +0x...
```

### Ví dụ 2: Debug script Python

```bash
# Chạy script Python với verbose
runex -v python myscript.py
```

### Ví dụ 3: Debug Node.js với AI

```bash
# Phân tích lỗi Node.js bằng AI
runex --ai node app.js
```

### Ví dụ 4: Tắt màu cho CI/CD

```bash
# Chạy trong CI environment
runex --no-color go test ./...
```

### Ví dụ 5: Chỉ định ngôn ngữ thủ công

```bash
# Force detection cho Go
runex --language go run main.go
```

## License

MIT
