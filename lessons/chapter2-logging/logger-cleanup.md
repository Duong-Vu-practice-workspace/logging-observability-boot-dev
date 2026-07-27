## Logger Cleanup

The [buffered writer](https://pkg.go.dev/bufio#Writer) is *faster*, but we added a bug! It *must* be flushed (written to disk) before the program exits, or any pending log messages will be lost!

It's also common to log messages across the network in some scenarios, and those kinds of loggers will also need to be flushed before exit – so we should take that into account in our implementation.

## Assignment

**Clean up logger resources before exit.**

1.
2. ```
	type closeFunc func() error
	func initializeLogger(logFile string) (*log.Logger, closeFunc, error)
	```
3.
4. Once again, we resort to writing directly to `os.Stderr` – the logger isn't in a usable state at this point.
5.

```sh
LINKO_LOG_FILE=linko.access.log go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Master of Mondays, can assist... *for a price*.

Copy/paste one of the following commands into your terminal:

Run

bootdev run 4cf32521-5877-4ea5-bde0-0e3469e72b64

Submit

bootdev run 4cf32521-5877-4ea5-bde0-0e3469e72b64 -s

- Default Base URL: http://localhost:8899
- Optionally configure your CLI to override the default base URL by running `bootdev config base_url <url>`

Run the CLI commands to test your solution.

1. GET /
    - 1.
    Expecting status code: 200
2. POST /admin/shutdown
    - 1.
    Expecting status code: 200
3. cat linko.access.log

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>