

## More Log Levels

I previously mentioned the conventional 4 log levels:

- Debug
- Info
- Warn
- Error

Many projects use even more levels like `TRACE`, `FATAL`, or `PANIC`, but as I already mentioned... I'd avoid them. Let's go a bit deeper on each level.

## Debug

Use `Debug` for detailed information that's useful for debugging, but not necessary for normal operation. Storing `Debug` logs can be expensive *and* noisy, so it's common to disable them in production and keep them enabled locally.

You can *almost* think of Debug logs as permanent "print debugging" statements.

Sometimes I even throw them in during local development, and *delete* them before committing if they're *super* ad-hoc.

## Info

`Info` level logs are used to record *important events* that are *not errors*. For example:

- "User 'alice' logged in"
- "File 'config.json' loaded successfully"
- "Server started on port 8080"

They're useful for understanding normal system behavior, and can help identify trends over time. In many cases, `Info` logs can be replaced with aggregated [metrics](https://prometheus.io/docs/concepts/metric_types/) – more on that later.

## Warn

Whoops. Did I go out of order? Yes.. but intentionally!

`Warn` is the level between `Info` and `Error`, so when should you use it?

To be blunt: *I think you shouldn't*!

It's a weird gray area between `Info` and `Error`, and in my experience, if it's not an actual error, it should usually be demoted to `Info`. Here's how I think about it:

1. Does the message represent *any* sort of potential bug? `Error`.
2. Does the message need to make its way to the user? Send a `400` and use `Info`.
3. Will the issue resolve itself? Don't even log it.

## Error

`Error` level logs record... errors. Obviously.

They should include enough information to diagnose the problem. Things like:

- The error message
- A stack trace
- Context (user, permissions, external API)

System errors and user errors are not the same. A request for a non-existent resource should return a `404 Not Found` – but the *system* didn't fail.

System errors (in general, that corresponds to 5XX codes in web-speak) should be logged as `Error` level logs. User errors (4XX codes) should be logged as `Debug` or `Info` level logs or not logged at all.

Add bookmark

Reset lesson

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Incredibly Fluffy, can assist... *for a price*.

Personal Instructions

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>