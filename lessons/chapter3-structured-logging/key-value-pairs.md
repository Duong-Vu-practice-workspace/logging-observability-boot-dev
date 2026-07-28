## Key-Value Pairs

Okay, time to finally put the *structure* in structured logging!

Some logs, like startup and shutdown messages, don't need additional metadata... but some do. Common examples in web services are:

- HTTP response code
- Size of request/response bodies
- Duration of the request
- User ID making the request
- Stack traces in the event of failure

All we need to do is add key-value pairs to the log line, for example:

```
logger.Info("Someone is loose in the server room",
    slog.String("name", "Boots"),
    slog.Int("id", 80045),
)
// 2024-01-15T10:30:45.123Z INFO msg="Someone is loose in the server room" name=Boots id=80045
```

You can use type-specific helpers like [`slog.String`](https://pkg.go.dev/log/slog#String) and [`slog.Int`](https://pkg.go.dev/log/slog#Int), or log values directly. Both are valid, but helpers can be more consistent, and sometimes more performant:

```
logger.Info("Someone is loose in the server room",
    "name", "Boots",
    "id", 80045,
)
// 2024-01-15T10:30:45.123Z INFO msg="Someone is loose in the server room" name=Boots id=80045
```

Structured logs are more readable, but more importantly, their interface can output [JSON](https://www.json.org/json-en.html) (or another format) just by changing the *handler*. Imagine if your app had thousands of log lines and you had to rewrite *all* of them just to ship a new format or destination... **no fun**.

## Assignment

**Add structured fields to your most important logs.**

Restart your server with `LINKO_LOG_FILE=linko.access.log` set:

```sh
LINKO_LOG_FILE=linko.access.log go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

Create similar challenge

Add bookmark

Reset lesson

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Incredibly Fluffy, can assist... *for a price*.

Personal Instructions

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>