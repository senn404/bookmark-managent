# 🔖 Bookmark Management API

![Go Version](https://img.shields.io/badge/go-1.18+-00ADD8.svg?style=flat-square&logo=go)
![Gin Framework](https://img.shields.io/badge/gin-web_framework-00ADD8.svg?style=flat-square)
![Swagger](https://img.shields.io/badge/swagger-api_docs-85EA2D.svg?style=flat-square&logo=swagger)
![Mockery](https://img.shields.io/badge/test-mockery-blue.svg?style=flat-square)

Một RESTful API Service được viết bằng Golang (Gin framework) tuân thủ kiến trúc **Clean Architecture**. Hiện tại dự án cung cấp các endpoint cơ bản như kiểm tra trạng thái hoạt động (Health Check) và khởi tạo công cụ (Password Generator), nhằm mục đích thiết lập phần khung (skeleton) vững chắc, sẵn sàng phát triển tính năng quản lý bookmark.

---

## 📑 Bảng mục lục (Table of Contents)
- [✨ Tính năng nổi bật](#-tính-năng-nổi-bật)
- [🛠 Công nghệ sử dụng](#-công-nghệ-sử-dụng)
- [📁 Cấu trúc dự án](#-cấu-trúc-dự-án)
- [🚀 Hướng dẫn cài đặt & khởi chạy](#-hướng-dẫn-cài-đặt--khởi-chạy)
- [🧪 Kiểm thử (Testing)](#-kiểm-thử-testing)
- [📚 Tóm tắt API (Swagger)](#-tóm-tắt-api-swagger)

---

## ✨ Tính năng nổi bật

- Cấu trúc thư mục chuẩn Go (Go standard project layout).
- Cấu hình Environment mượt mà sử dụng `kelseyhightower/envconfig` và `joho/godotenv`.
- API endpoints sinh mã tài liệu tự động qua Swagger (`swaggo`).
- Áp dụng Interface, Dependency Injection (DI) dễ dàng Unit testing / Mocking.
- Integration test và Unit test bao phủ các luồng quan trọng của ứng dụng.

---

## 🛠 Công nghệ sử dụng

- **Ngôn ngữ:** [Go (Golang)](https://golang.org/)
- **Web Framework:** [Gin-Gonic](https://github.com/gin-gonic/gin)
- **Tài liệu API:** [Swaggo / gin-swagger](https://github.com/swaggo/gin-swagger)
- **Quản lý cấu hình:** [envconfig](https://github.com/kelseyhightower/envconfig) + [godotenv](https://github.com/joho/godotenv)
- **Unit Testing:** [Mockery](https://github.com/vektra/mockery) & custom assert mechanism.

---

## 📁 Cấu trúc dự án

Dự án áp dụng nguyên lý tách biệt mối quan tâm (Separation of Concerns):

```text
.
├── cmd/
│   └── api/
│       └── main.go       # Entry point chạy dịch vụ API.
├── internal/
│   ├── config/           # Định nghĩa và parsing cấu hình từ `.env`.
│   ├── api/              # Định nghĩa Router và khởi tạo tầng trình diễn.
│   ├── handler/          # HTTP Transport Layer — xử lý request/response JSON.
│   ├── service/          # Business Logic Layer — chứa xử lý nghiệp vụ chính.
│   │   └── mocks/        # Các mock services (Tự động tạo ra bằng Mockery).
│   └── test/
│       └── endpoint/     # Thư mục lưu trữ API Integration Tests.
├── docs/                 # Swagger specs (tự động tạo ra bởi swag cli).
└── Makefile              # Tổng hợp các terminal scripts.
```

---

## 🚀 Hướng dẫn cài đặt & khởi chạy

### 1. Yêu cầu hệ thống
- Go (khuyến nghị phiên bản 1.18+)
- Make (nếu bạn muốn tận dụng các lệnh `Makefile`)
- [Swag CLI](https://github.com/swaggo/swag) (Cho việc generate swagger comments)

### 2. Cấu hình biến môi trường
Mặc định hệ thống tải các file cấu hình tại hệ thống. Bạn có thể tạo nội dung file `.env` ở thư mục gốc (nếu dùng local):

```env
APP_PORT=8080
SERVICE_NAME=bookmark_service
INSTANCE_ID=
```

### 3. Khởi chạy dự án
Bạn có thể dùng `make` để chạy trực tiếp máy chủ.

```bash
# Tải các module gói phụ thuộc
go mod tidy

# Khởi chạy server
make run
```
> Khi bắt đầu, bạn sẽ thấy thiết bị lắng nghe trên `http://localhost:8080` (hoặc cổng được định nghĩa trong biến môi trường).

---

## 🧪 Kiểm thử (Testing)

Dự án được viết test khá hoàn chỉnh ở cả tầng nghiệp vụ (unit test) lẫn integration (kiểm thử endpoint). Test coverage xuất trực tiếp qua tiện ích HTML.

```bash
# Lệnh chạy toàn bộ test & sinh html coverage
make test
```
*(Kết quả coverage file dưới định dạng `cover.html` sẽ được tạo ra tại thư mục gốc của project).*

---

## 📚 Tóm tắt API (Swagger)

Nếu bạn thay đổi các Go Docstrings (ở tầng `handler` hay `main.go`), hãy regenerate Swagger thông qua dòng lệnh:

```bash
make swagger
```

**Truy cập GUI API Docs:** (Khi server đang khởi chạy)
👉 `http://localhost:8080/swagger/index.html`

### Danh sách Endpoint Demo:
- **`GET /health-check`**: Kiểm tra trạng thái hệ thống. (Trả về trạng thái, tên service & UUID instance sinh ngẫu nhiên).
- **`GET /gen-pass`**: Tạo một mật khẩu an toàn theo random byte với độ dài quy chuẩn.
