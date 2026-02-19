RUNEX – Runtime Error Explorer for Developers

Runex là một công cụ dòng lệnh (CLI) giúp lập trình viên bắt và phân tích lỗi xảy ra trong quá trình chương trình đang chạy (runtime), một cách độc lập với IDE, ngôn ngữ lập trình và framework.

Mục tiêu của Runex là giải quyết vấn đề phổ biến trong workflow lập trình: khi chương trình chạy lỗi, developer phải copy stacktrace, log lỗi, dán vào công cụ khác (Google, AI chat) và thường bị mất context quan trọng. Runex giữ lại toàn bộ context runtime và đưa ra thông tin có cấu trúc ngay tại terminal.

⸻

TECH STACK

Ngôn ngữ chính:
	•	Go (Golang)

Lý do chọn Go:
	•	Build được single static binary
	•	Hiệu năng tốt, memory thấp
	•	Quản lý process, stdout/stderr stream rất mạnh
	•	Phù hợp cho CLI / dev tooling / infra tools
	•	Cross-platform (macOS, Linux, Windows)

Kiến trúc runtime:
	•	OS process wrapper (os/exec)
	•	Context-based process control
	•	Real-time stdout / stderr streaming
	•	Exit code & signal handling

CLI & UX:
	•	Native Go CLI (flag / cobra có thể dùng sau)
	•	Terminal-first UX
	•	Không phụ thuộc IDE
	•	Không yêu cầu daemon chạy nền

Error detection:
	•	Heuristic-based detection (panic, exception, traceback)
	•	Rule-based language identification
	•	Raw stacktrace extraction
	•	Không dùng static analysis
	•	Không can thiệp source code

Ngôn ngữ & hệ sinh thái hỗ trợ (giai đoạn đầu):
	•	Java (Spring Boot, Maven, Gradle)
	•	Python
	•	Node.js
	•	Go

AI integration (optional, pluggable):
	•	Provider-agnostic design
	•	AIProvider interface
	•	Hỗ trợ OpenAI / Anthropic / các LLM khác
	•	API key do user tự cấu hình (env / config)
	•	Core không phụ thuộc AI

Cấu hình & bảo mật:
	•	Environment variables
	•	Local config file (~/.config/runex)
	•	Không lưu API key trong source code
	•	Phù hợp open-source

Phân phối:
	•	Single binary
	•	go install
	•	Homebrew (macOS)
	•	Manual download cho Linux / Windows

Testing & chất lượng:
	•	Manual runtime error scenarios
	•	Cross-language error simulation
	•	Focus vào behavior thay vì unit test phức tạp
	•	Ưu tiên độ ổn định của CLI

⸻

Cách hoạt động

Runex hoạt động như một wrapper cho các lệnh chạy chương trình hiện có.

Ví dụ:
	•	runex mvn spring-boot:run
	•	runex npm run dev
	•	runex python app.py
	•	runex go run main.go

Về mặt kỹ thuật, Runex:
	•	Spawn một process con để chạy lệnh gốc
	•	Stream stdout và stderr theo thời gian thực
	•	Theo dõi exit code và tín hiệu lỗi
	•	Thu thập stderr để phát hiện runtime error
	•	Phân tích lỗi sau khi process kết thúc

⸻

Những gì Runex làm
	•	Phát hiện runtime error (panic, exception, traceback, crash)
	•	Nhận diện ngôn ngữ và loại lỗi bằng heuristic
	•	Trích xuất stacktrace thô
	•	Chuẩn hóa thông tin lỗi thành context thống nhất
	•	Hiển thị lỗi rõ ràng trong terminal
	•	(Tuỳ chọn) Giải thích lỗi bằng AI

⸻

Những gì Runex KHÔNG làm
	•	Không phải debugger
	•	Không phải linter
	•	Không phải formatter
	•	Không tự động sửa code
	•	Không yêu cầu IDE
	•	Không bắt buộc dùng AI

⸻

Giá trị cốt lõi
	•	Runtime truth: dựa trên lỗi thực tế khi chương trình chạy
	•	Context preservation: không mất thông tin quan trọng
	•	IDE-agnostic: chạy ở mọi môi trường terminal
	•	CLI-first: phù hợp backend, DevOps, infra

⸻

Triết lý thiết kế
	•	Fail-safe: lỗi phân tích không làm crash chương trình gốc
	•	Không phá workflow hiện tại
	•	Ưu tiên đơn giản và ổn định
	•	Open-source friendly

⸻

Định hướng phát triển
	•	Mở rộng rule & parser cho nhiều ngôn ngữ
	•	Hỗ trợ chế độ CI
	•	Team error knowledge base
	•	Phiên bản cloud / enterprise (tuỳ chọn)