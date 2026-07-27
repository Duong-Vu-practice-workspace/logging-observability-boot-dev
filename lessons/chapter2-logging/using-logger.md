## Use the Logger

So [`log.Printf`](https://pkg.go.dev/log#Printf) seems to work well... but *there's a better way*!

A "logger" is an instance of a [`log.Logger`](https://pkg.go.dev/log#Logger) that can be used to produce logs. Generally it's **better to use a logger object** than the `log` package's functions directly, for a few reasons:

- You can easily change where the logs go, all in one place
- You can add prefixes to the logs, again all in one place
- You can change where the logs go at runtime, again... all in one place

## Using STDERR

It's usually best to send logs to [`os.Stderr`](https://pkg.go.dev/os#Stderr) instead of [`os.Stdout`](https://pkg.go.dev/os#Stdout) because `STDOUT` is typically used for the main output of a program, and we don't want to gum that up with logs meant for developers.

When you create a new logger with [`log.New`](https://pkg.go.dev/log#New), you can specify the output destination, and `os.Stderr` is usually the right choice.

```
// create a logger
var logger = log.New(os.Stderr, "MESSAGE: ", log.LstdFlags)

// use a logger
logger.Printf("The Lisan al-Gaib arrived")
// MESSAGE: 2024/06/01 12:00:00 The Lisan al-Gaib arrived
```

- [`os.Stderr`](https://pkg.go.dev/os#Stderr) is the standard error output stream
- The second argument is a prefix for the log messages (here we're using "MESSAGE: ")
- The third argument is the [log flags](https://pkg.go.dev/log#pkg-constants), which can include things like timestamps, file names, and line numbers. `log.LstdFlags` simply includes the date and time.

## Enforcing Loggers

If you find yourself forgetting to use a logger, the [golangci-lint](https://golangci-lint.run/) linter comes with a sublinter called [forbidigo](https://golangci-lint.run/usage/linters/#forbidigo) that can be configured to prohibit the use of these functions:

```yaml
version: "2"

linters:
  settings:
    forbidigo:
      forbid:
        - pattern: ^fmt\.Print.*$
          msg: Use logger instead.
      analyze-types: true
```

This is totally optional of course, but it's nice to know about.

## Assignment

**Move from package-level `log` calls to a shared logger instance.**

```sh
go run . 2>&1 | sh -c 'trap "" INT; tee linko.out.log'
```

**Run and submit** the CLI tests from the root of the Linko repo.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Incredibly Fluffy, can assist... *for a price*.

Copy/paste one of the following commands into your terminal:

Run

bootdev run 8bbb5dae-95b3-4e4e-a67f-37829535a524

Submit

bootdev run 8bbb5dae-95b3-4e4e-a67f-37829535a524 -s

- Default Base URL: http://localhost:8899
- Optionally configure your CLI to override the default base URL by running `bootdev config base_url <url>`

Check the output of your CLI command to see a results breakdown.

1. GET /
    - 1.
    Expecting status code: 200
2. POST /admin/shutdown
    - 1.
    Expecting status code: 200
3. cat linko.out.log

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>