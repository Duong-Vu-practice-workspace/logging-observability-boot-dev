# Logger Failure — Giải thích
## Logger Failure

Here's the code I used to create Linko's logger:

```
func initializeLogger(logFile string) (*log.Logger, error) {
    if logFile != "" {
        file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
        if err != nil {
            return nil, fmt.Errorf("failed to open log file: %w", err)
        }
        multiWriter := io.MultiWriter(os.Stderr, file)
        return log.New(multiWriter, "", log.LstdFlags), nil
    }
    return log.New(os.Stderr, "", log.LstdFlags), nil
}

func run(ctx context.Context, httpPort int, dataDir string) int {
    logger, err := initializeLogger(os.Getenv("LINKO_LOG_FILE"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
        return 1
    }
    // ...
}
```

Notice that if an error occurs when opening the file, I `return` an `error` from `initializeLogger`, and then in `run()` I write a message to [`os.Stderr`](https://pkg.go.dev/os#Stderr) and return a non-zero exit code. If *you* used [`log.Fatal`](https://pkg.go.dev/log#Fatal) or [`log.Panic`](https://pkg.go.dev/log#Panic) instead, you might have a couple of problems in your code that would:

- Make it impossible (or very difficult) to test that behavior in a unit test.
- Prevent the program from running any deferred functions or doing other cleanup.

If you ask me, [`log.Fatal`](https://pkg.go.dev/log#Fatal) and [`log.Panic`](https://pkg.go.dev/log#Panic) should be avoided... I don't even like that they're *in* the standard library, because they couple logging with control flow – but that's a different discussion.

Instead, I prefer to let the caller of the `initializeLogger` function decide how to behave in the event of a failure! Then, when it's time to *handle* the error (in the `run` function), this is one of the few times it's okay to log without a logger (by using [`fmt.Fprintf`](https://pkg.go.dev/fmt#Fprintf)) because it was the logger itself that failed to initialize!

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Magnificent, can assist... *for a price*.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>

## Vấn đề với `log.Fatal` và `log.Panic`

Khi bạn dùng `log.Fatal` hoặc `log.Panic`, chúng làm 2 việc cùng lúc:
1. Ghi log
2. **Kết thúc chương trình** (exit) hoặc **panic**

Điều này được gọi là "coupling logging với control flow" — nghĩa là ghi log và quyết định thoát chương trình bị trộn vào nhau, rất khó kiểm tra (testing).

Ví dụ code thường thấy:
```go
f, err := os.OpenFile(logFile, ...)
if err != nil {
    log.Fatalf("failed to open log file: %v", err)  // vừa log, vừa exit
}
multiWriter = io.MultiWriter(os.Stderr, f)  // dead code, không bao giờ chạy
```

## Cách tốt hơn: tách riêng "tạo logger" và "xử lý lỗi"

```go
// Chỉ tạo logger, trả về error nếu thất bại — KHÔNG tự ý thoát
func initializeLogger(logFile string) (*log.Logger, error) {
    if logFile != "" {
        file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
        if err != nil {
            return nil, fmt.Errorf("failed to open log file: %w", err)
        }
        multiWriter := io.MultiWriter(os.Stderr, file)
        return log.New(multiWriter, "", log.LstdFlags), nil
    }
    return log.New(os.Stderr, "", log.LstdFlags), nil
}

// Caller quyết định cách xử lý lỗi
func run(...) int {
    logger, err := initializeLogger(os.Getenv("LINKO_LOG_FILE"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "failed to initialize logger: %v\n", err)
        return 1
    }
    // ... dùng logger bình thường
}
```

## Lợi ích

1. **Test được** — bạn có thể gọi `initializeLogger` với file sai trong unit test và kiểm tra nó trả về error, thay vì phải catch `os.Exit` hoặc `panic`.

2. **Defer vẫn chạy** — `log.Fatal` gọi `os.Exit` bên trong, **bỏ qua tất cả `defer`**. Nếu bạn có cleanup (đóng DB, ghi file, v.v.), chúng sẽ không bao giờ chạy.

3. **Phân tách trách nhiệm** — hàm `initializeLogger` chỉ lo *tạo* logger, không lo *quyết định số phận* của chương trình. Đó là việc của `run()` (caller).

## Áp dụng cho code hiện tại

Thay vì:
```go
if err != nil {
    log.Fatalf("failed to open log file: %v", err)
}
multiWriter = io.MultiWriter(os.Stderr, f)  // dead code
```

Nên:
```go
if err != nil {
    return fmt.Errorf("failed to open log file: %w", err)
}
multiWriter = io.MultiWriter(os.Stderr, f)
```

## Tóm lại

**Không dùng `log.Fatal`/`log.Panic`** trong hàm khởi tạo. Trả về `error` để caller tự quyết định. Đó là DI (dependency injection) áp dụng cho chính việc xử lý lỗi.
