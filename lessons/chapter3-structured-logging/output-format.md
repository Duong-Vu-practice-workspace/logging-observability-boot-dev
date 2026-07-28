---
title: "Learn Logging and Observability in Go: Output Formats"
source: "https://www.boot.dev/lessons/54fdbc45-64f8-4641-97ad-3a582ae192c5"
author:
published: 2026-07-29
created: 2026-07-28
description: "Add the logs, metrics, traces, and alerts you'll wish you had when production catches fire."
tags:
  - "clippings"
---
## Output Formats

As you know, [`log/slog`](https://pkg.go.dev/log/slog) supports [JSON](https://www.json.org/json-en.html) output, as well as arbitrary custom formats, but so far we've only used text... let's change that.

JSON is great for production logging because it's easy for log aggregation and analysis tools to parse. Most log ingestion systems, such as the [ELK Stack](https://www.elastic.co/elastic-stack/) and [Loki](https://grafana.com/oss/loki/), accept JSON logs – even when they also support their own proprietary formats.

I prefer JSON logs for storage and filtering in production, but text logs for debugging and development.

## Assignment

**Switch your file logs to JSON while keeping local terminal logs readable.**

Restart your server with `LINKO_LOG_FILE=linko.access.log` set:

```sh
LINKO_LOG_FILE=linko.access.log go run .
```

**Run and submit** the CLI tests from the root of the Linko repo.

Add bookmark

Reset lesson

Report Issue with Lesson

Next tab: ctrl+g

Prev tab: ctrl+shift+g

![Boots](https://www.boot.dev/_nuxt/new_boots_profile.DriFHGho.webp)

**Need help?** I, Boots the Sleepy Spellcaster, can assist... *for a price*.

Personal Instructions

Copy/paste one of the following commands into your terminal:

Run

bootdev run 54fdbc45-64f8-4641-97ad-3a582ae192c5

Submit

bootdev run 54fdbc45-64f8-4641-97ad-3a582ae192c5 -s

- Default Base URL: http://localhost:8899
- Optionally configure your CLI to override the default base URL by running `bootdev config base_url <url>`

Run the CLI commands to test your solution.

1. GET /
   If this step fails, you won't lose armor or your Sharpshooter progress.
    - 1.
    Expecting status code: 200
2. GET /api/stats
    - 1.
    Expecting status code: 401
3. POST /admin/shutdown
    - 1.
    Expecting status code: 200
4. cat linko.access.log

View Solution

\-100% XP

Using the Bootdev CLI

The Bootdev CLI is the only way to submit your solution for this type of lesson. We need to be able to run commands in your environment to verify your solution.

You can [install it here](https://github.com/bootdotdev/bootdev). It's a Go program hosted on GitHub, so you'll need Go installed as well. Instructions are on the GitHub page.

<iframe allow="clipboard-write; web-share" src="chrome-extension://cnjifjpddelmedmihgijeibhnjfabmlf/side-panel.html?context=iframe"></iframe>