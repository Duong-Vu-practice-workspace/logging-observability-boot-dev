## Buffered Logging

Our current logger, especially when writing to a file, is relatively slow. Every time we log a message, it writes to disk, no matter how large or small the message is. That's potentially a *lot* of [disk I/O](https://en.wikipedia.org/wiki/Input/output), and it can really slow down our entire application because many small writes are much slower than a few large writes.

Observability being the reason our app is slow is, frankly, embarrassing.

One solution is to use a buffered writer like [bufio.Writer](https://pkg.go.dev/bufio#Writer) around the file. This allows us to write log messages to an in-memory buffer, and then that buffer is only written to disk when it's full.

```
bufferedFile := bufio.NewWriterSize(file, 1024)
```

## Assignment

**Buffer file logging writes.**

```sh
LINKO_LOG_FILE=linko.access.log go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

If you create the buffered logger the way I did, it will introduce a subtle bug... but don't worry we'll fix it later!

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Hump Day Holdout, can assist... *for a price*.

Copy/paste one of the following commands into your terminal:

Run

bootdev run 0c6aed59-776c-461d-896a-c7ca218dbb6d

Submit

bootdev run 0c6aed59-776c-461d-896a-c7ca218dbb6d -s

- Default Base URL: http://localhost:8899
- Optionally configure your CLI to override the default base URL by running `bootdev config base_url <url>`

Run the CLI commands to test your solution.

1. GET /
    - 1.
    Expecting status code: 200
2. grep -R -n --include='\*.go' 'bufio'.
    - Expecting exit code: 0
3. grep -R -n --include='\*.go' 'NewWriterSize'.
4. POST /admin/shutdown
    - 1.
    Expecting status code: 200

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>