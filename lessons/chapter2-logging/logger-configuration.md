## Logger Configuration

It's more standard to decide what goes to `STDERR` and what goes to a file based on the [environment](https://en.wikipedia.org/wiki/Deployment_environment) your application is running in, rather than on separate loggers. Common environments are:

- "development" (local development, like when you're running the application on your own machine)
- "staging" (a pre-production environment that mimics production)
- "production" (the live environment that users interact with)

## Multiwriter Configuration

There's no reason a logger can't write to both `STDERR` and a file at the same time! The [`io.MultiWriter`](https://pkg.go.dev/io#MultiWriter) function takes multiple `io.Writer` objects and returns a single `io.Writer` that writes to all of them. For example:

```
file, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
if err != nil {
    log.Fatalf("failed to open log file: %v", err)
}
multiWriter := io.MultiWriter(os.Stderr, file)
logger := log.New(multiWriter, "INFO: ", log.LstdFlags)
```

## Assignment

**Use one logger that changes output based on `LINKO_LOG_FILE`.**

Assume that in production, Linko has a `LINKO_LOG_FILE` environment variable set. In local development and staging, it is not set.

If `LINKO_LOG_FILE` is set, the logger should write to *both* the file and `STDERR`. Otherwise, it should only write to `STDERR`.

Restart your server, setting the `LINKO_LOG_FILE` environment variable so the tests can verify the file is created:

```sh
LINKO_LOG_FILE=linko.access.log go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Fearless Friday Deployer, can assist... *for a price*.

Copy/paste one of the following commands into your terminal:

Run

bootdev run a85039cb-42ac-4f37-98af-6a05b7aa6f30

Submit

bootdev run a85039cb-42ac-4f37-98af-6a05b7aa6f30 -s

- Default Base URL: http://localhost:8899
- Optionally configure your CLI to override the default base URL by running `bootdev config base_url <url>`

Run the CLI commands to test your solution.

1. GET /
    - 1.
    Expecting status code: 200
2. GET /api/stats
    - 1.
    Expecting status code: 401
3. POST /admin/shutdown
    - 1.
    Expecting status code: 200
4. cat linko.access.log
    - Expecting exit code: 0
        - Expecting stdout to contain all of:
        - Served request: GET /
          - Served request: GET /api/stats

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>