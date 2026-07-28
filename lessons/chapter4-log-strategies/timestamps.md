This lesson's interactive features are locked, please upgrade to keep using them

## Timestamps

This might seem obvious (and most default loggers do this), but **always include [timestamps](https://en.wikipedia.org/wiki/Timestamp) in your logs**.

Even if your logs are *complete jank*, timestamps at least let us do brute-force investigation. Take a look:

Each log entry alone isn't very useful, but the timestamps allow us to deduce that they're *probably related*, and that the "File not found" error likely relates to opening Alice's profile configuration file.

One exception is in *automated tests*. You may want to remove or overwrite timestamps for deterministic output.

## Assignment

**Write a test for `requestLogger` that verifies timestamped output.**

1. ```
	func Test_requestLogger(t *testing.T) {
	    logBuffer := &bytes.Buffer{}
	    logger := slog.New(slog.NewTextHandler(logBuffer, &slog.HandlerOptions{
	        ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
	            if a.Key == slog.TimeKey {
	                return slog.Time(slog.TimeKey, time.Date(2023, 10, 1, 12, 34, 57, 0, time.UTC))
	            }
	            return a
	        },
	    }))
	    requestLoggerMiddleware := requestLogger(logger)
	    dummyHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	    loggedHandler := requestLoggerMiddleware(dummyHandler)
	    req := httptest.NewRequest("GET", "http://lin.ko/api/stats", nil)
	    rr := httptest.NewRecorder()
	    loggedHandler.ServeHTTP(rr, req)
	    const expectedLogString = \`time=2023-10-01T12:34:57.000Z level=INFO msg="Served request" method=GET path=/api/stats client_ip=192.0.2.1:1234\` + "\n"
	    const expectedStatusCode = http.StatusOK
	    // replace the .Skip() call with two checks to verify the log string and status code here
	    // If either doesn't match, use t.Errorf to report the failure with a helpful message.
	    t.Skip()
	}
	```
   Notice that we're using the [`httptest`](https://pkg.go.dev/net/http/httptest) package to create a dummy HTTP request and response recorder. This is a cool way to "end-to-end" test an individual HTTP handler.
2. - Compare `logBuffer.String()` to the expected log string.
    - Compare `rr.Code` to the expected status code.

**Run and submit** the CLI tests from the root of the Linko repo.

Add bookmark

Become a member to Reset

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Bear with a Back-End, can assist... *for a price*.

Personal Instructions

Copy/paste one of the following commands into your terminal:

Run

bootdev run dc065478-2849-4069-9a14-6fa40b211426

Submit

bootdev run dc065478-2849-4069-9a14-6fa40b211426 -s

To run and submit the tests for this lesson, you must have an active Boot.dev membership

Become a member to view solution

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>