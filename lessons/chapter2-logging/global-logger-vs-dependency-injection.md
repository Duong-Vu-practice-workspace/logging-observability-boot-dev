## Global Logger vs. Dependency Injection

Notice that our `requestLogger` middleware accepts a `*log.Logger` as a parameter:

```
func requestLogger(logger *log.Logger) func(next http.Handler) http.Handler
```

But why not just keep using the *global* logger we already declared everywhere? Wouldn't that be simpler? On the surface, perhaps.

But globals are generally a bad idea because they make testing and debugging more difficult... and a global logger is no exception!

## Using Dependency Injection

[Dependency injection (DI)](https://en.wikipedia.org/wiki/Dependency_injection) is a really fancy term for a really simple idea: pass a function or method's dependencies in as arguments. Generally it makes testing much easier, because you don't need to continuously mutate shared state.

By passing the logger object as an argument to our middleware function, we can avoid these problems. It also means we can pass a distinct logger object for each test, and even run tests in parallel without worrying about global state.

So, non-global loggers are great because:

- It's easy to use *different* loggers in different parts of your app if you'd like
- You can more easily pass [`context`](https://pkg.go.dev/context) information to the logger as needed
- There's no global state to worry about when you're writing unit tests that use the logger

## Assignment

**Replace the global logger with injected loggers.**

Restart your server:

```sh
go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Gormless Glutton, can assist... *for a price*.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>