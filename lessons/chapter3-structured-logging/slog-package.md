## Slog Package

In Go 1.21 (June 2023), the [`log/slog` package](https://pkg.go.dev/log/slog) was added to the standard library to support *structured* logging. There was much rejoicing.

There are several third-party logging libraries, like:

- [logrus](https://github.com/sirupsen/logrus)
- [zerolog](https://github.com/rs/zerolog)
- [zap](https://github.com/uber-go/zap)

Most predate `log/slog`, and unless you have very specific needs, `log/slog` is probably all you need these days. Compare a standard [`log`](https://pkg.go.dev/log) message:

```
2023/10/01 12:00:00 This is a log message
```

To a structured `log/slog` message:

```
2024-01-15T10:30:45.123Z INFO msg="user login successful" user_id=12345 username=john_doe ip_address=192.168.1.100 duration_ms=245
```

## Handlers

The `log/slog` package introduces *handlers*, which accept arbitrary key-value pairs and format them into a log record. Two built-in handlers are:

- [`log/slog.NewTextHandler`](https://pkg.go.dev/log/slog#NewTextHandler): Formats logs as plain text
- [`log/slog.NewJSONHandler`](https://pkg.go.dev/log/slog#NewJSONHandler): Formats logs as JSON.

## Initialization

The [`log/slog.New`](https://pkg.go.dev/log/slog#New) function takes a handler as an argument, and returns a logger instance that can be used to log messages:

```
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
```

Then we can use it like this:

```
logger.Info("This is an info message")
```

The structured logger doesn't support formatting methods like `Infof`, so use [`fmt.Sprintf`](https://pkg.go.dev/fmt#Sprintf) when needed:

```
logger.Info(fmt.Sprintf("Failed to open file %s: %s", filename, err))
```

## Assignment

**Switch Linko to structured logging with `slog`.**

1.
2. Use `fmt.Sprintf` to format strings as needed. `slog` doesn't have `Infof` or similar methods.
3.

Restart your server:

```sh
go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

Add bookmark

Reset lesson

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Gormless Glutton, can assist... *for a price*.

Personal Instructions

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>