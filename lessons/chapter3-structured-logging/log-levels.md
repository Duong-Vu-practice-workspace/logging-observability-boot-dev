## Log Levels

Log levels aren't specific to structured logging, but [`log/slog`](https://pkg.go.dev/log/slog) provides a built-in way to handle them. They're a convention for labeling each log with a severity. The standard library defines four levels by default, and you can define custom levels if needed:

- `slog.LevelError`: Error messages, indicating failures or issues that need attention.
- `slog.LevelWarn`: Warning messages, indicating potential issues that are not critical.
- `slog.LevelInfo`: Informational messages, typically used for general application events.
- `slog.LevelDebug`: Debug messages, useful for development and debugging.

You'll usually only need these four.

I don't even use `Warn` messages very often if truth be told...

## Filtering

Have you ever added `print` statements while debugging? Sometimes you need that same visibility in production because you can't reproduce the issue locally. You still don't want to flood your normal logs with extra noise, so you use `Debug`. Then if something else goes wrong, you can filter debug logs out or send them elsewhere. For example, you can use different handlers with different minimum levels:

```
// logs DEBUG and above
debugLogger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelDebug,
}))

// logs ERROR and above
errorLogger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelError,
}))
```

## Combining Loggers

A more practical approach is a single logger that routes logs to different destinations by level. For example, everything goes to `STDERR`, but only `INFO` and higher go to a file. As of Go 1.26, this is easy with [`slog.NewMultiHandler`](https://pkg.go.dev/log/slog#NewMultiHandler):

```
debugHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
    Level: slog.LevelDebug,
})

logFile, err := os.OpenFile("linko.access.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
if err != nil {
    return err
}
defer logFile.Close()
infoHandler := slog.NewTextHandler(logFile, &slog.HandlerOptions{
    Level: slog.LevelInfo,
})

logger := slog.New(slog.NewMultiHandler(
    debugHandler,
    infoHandler,
))
```

## Assignment

**Split logs by severity, but keep one app-wide logger.**

Start your server with `LINKO_LOG_FILE=linko.access.log` set:

```sh
LINKO_LOG_FILE=linko.access.log go run . 2>&1 | sh -c 'trap "" INT; tee linko.out.log'
```

**Run and submit** the CLI tests from the root of the Linko repo.

Add bookmark

Reset lesson

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Hump Day Holdout, can assist... *for a price*.

Personal Instructions

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>