## Logging Requests

It's *very* common to log requests in a web service. One of the cleaner ways to implement this is with a [middleware](https://en.wikipedia.org/wiki/Middleware) function that logs the request after it's been served:

```
func requestLogger(logger *log.Logger) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            next.ServeHTTP(w, r)
            logger.Printf("Wake up babe, a new %s request to %s just dropped", r.Method, r.URL.Path)
        })
    }
}
```

This one simply logs the request method and path after the request has been served. Notice that it takes a [`*log.Logger`](https://pkg.go.dev/log#Logger) as an argument, allowing you to use any logger you want on a per-handler basis. So, instead of declaring a handler that we want to log like this:

```
mux.HandleFunc("POST /api/shorten", apiCfg.handlerShortenURL)
```

We can use middleware:

```
mux.Handle("/api/shorten", requestLogger(logger)(http.HandlerFunc(apiCfg.handlerShortenURL)))
```

Alternatively, we can wrap the entire `mux` with the middleware, so that *all* requests are logged:

```
srv = &http.Server{
    Addr:    fmt.Sprintf(":%d", port),
    Handler: requestLogger(logger)(mux),
}
```

## Assignment

**Log each served request with middleware.**

1. ```
	Served request: METHOD Path
	```
   Where `METHOD` is the HTTP method of the request, and `Path` is the path of the request. For example:
   ```
   Served request: GET /
   ```
2. ```sh
	go run . 2>&1 | sh -c 'trap "" INT; tee linko.out.log'
	```
3.

**Run and submit** the CLI tests from the root of the Linko repo.

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Gormless Glutton, can assist... *for a price*.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>