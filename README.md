# Runex

Công cụ CLI để phát hiện và phân tích lỗi runtime theo thời gian thực.

## Tính năng

- **Phát hiện lỗi thời gian thực**: Giám sát stdout/stderr trong khi truyền tải đầu ra của lệnh
- **Hỗ trợ đa ngôn ngữ**: Phát hiện lỗi từ Go, Python, Node.js, Java, Ruby và Rust
- **Trích xuất Stack Trace**: Tự động trích xuất stack trace từ đầu ra lỗi
- **Phân tích AI**: Tùy chọn phân tích AI để hiểu sâu hơn về lỗi
- **Màu sắc đầu ra**: Thông báo lỗi có màu sắc thân thiện với terminal

---

## Giới thiệu

Runex là công cụ giao diện dòng lệnh (CLI) giúp bạn giám sát và phát hiện lỗi runtime một cách tự động khi chạy chương trình. Thay vì phải đọc toàn bộ đầu ra để tìm lỗi, Runex giám sát stdout/stderr và ngay lập tức thông báo cho bạn khi phát hiện lỗi cùng với full stack trace.

---

## Cài đặt

### Từ source code (sử dụng Go)

```bash
go install github.com/runex/runex@latest
```

### Từ GitHub Releases

Khi các bản phát hành có sẵn, bạn có thể tải xuống file nhị phân phù hợp cho hệ điều hành của mình:

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

---

## Bắt đầu nhanh

```bash
# Chạy bất kỳ lệnh nào với phát hiện lỗi
runex go run main.go

runex python script.py

runex node app.js
```

---

## Sử dụng chi tiết

### Các cờ (Flags)

| Cờ | Mô tả | Ví dụ |
|------|-------------|---------|
| `-v` | Bật chế độ verbose, hiển thị thông tin debug bổ sung | `runex -v go run main.go` |
| `--no-color` | Tắt màu sắc đầu ra | `runex --no-color python script.py` |
| `--ai` | Bật phân tích AI để hiểu sâu hơn về lỗi | `runex --ai node app.js` |
| `--language` | Chỉ định ngôn ngữ thủ công (go, python, node, java, ruby, rust) | `runex --language go run main.go` |

### Ví dụ theo ngôn ngữ

#### Go

```bash
# Chạy file Go
runex go run main.go

# Chạy tests
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

# Chạy tests
runex cargo test

# Build
runex cargo build
```

---

## Cấu hình

### File cấu hình

Tạo `~/.runex/config.yaml` để lưu cấu hình mặc định:

```yaml
verbose: false
noColor: false
ai: false
language: ""
```

### Biến môi trường

| Biến | Mô tả | Mặc định |
|----------|-------------|---------|
| `RUNEX_VERBOSE` | Bật chế độ verbose | `false` |
| `RUNEX_NO_COLOR` | Tắt màu sắc | `false` |
| `RUNEX_AI` | Bật phân tích AI | `false` |
| `RUNEX_LANGUAGE` | Ngôn ngữ mặc định | `""` |

### Ưu tiên cấu hình

Thứ tự ưu tiên (cao nhất đến thấp nhất):
1. Cờ dòng lệnh
2. Biến môi trường
3. File cấu hình (`~/.runex/config.yaml`)

---

## Cách hoạt động

1. **Khởi động**: Runex nhận lệnh từ người dùng và thực thi
2. **Giám sát**: Đồng thời đọc stdout và stderr từ lệnh đang chạy
3. **Phát hiện**: Phân tích đầu ra để tìm các mẫu lỗi known cho từng ngôn ngữ
4. **Thông báo**: Khi phát hiện lỗi, hiển thị thông báo kèm stack trace
5. **Hoàn thành**: Lệnh kết thúc, Runex trả về mã thoát tương ứng

### Ngôn ngữ được hỗ trợ

| Ngôn ngữ | Loại lỗi được phát hiện |
|----------|---------------------|
| Go       | panic, runtime error |
| Python   | exceptions, errors |
| Node.js  | Error, TypeError, ReferenceError, UnhandledPromiseRejection |
| Java     | exceptions |
| Ruby     | errors |
| Rust     | panic, compile errors |

---

## Ví dụ thực tế

### Ví dụ 1: Debug ứng dụng Go

```bash
# Giả sử bạn có main.go với một lỗi
runex -v go run main.go
```

Đầu ra:
```
[ERROR] Phát hiện lỗi Go
panic: index out of range [1] with length 0

goroutine 1 [running]:
main.main()
    /path/to/main.go:10 +0x...
```

### Ví dụ 2: Debug Python Script

```bash
# Chạy Python script với verbose
runex -v python myscript.py
```

### Ví dụ 3: Debug Node.js với AI

```bash
# Phân tích lỗi Node.js bằng AI
runex --ai node app.js
```

### Ví dụ 4: Tắt màu cho CI/CD

```bash
# Chạy trong môi trường CI
runex --no-color go test ./...
```

### Ví dụ 5: Chỉ định ngôn ngữ thủ công

```bash
# Chỉ định phát hiện cho Go
runex --language go run main.go
```

---

## Giấy phép

MIT
