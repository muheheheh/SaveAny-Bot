---
title: Introduction
---

# SaveAny-Bot (Telegram relay fork)

[muheheheh/SaveAny-Bot](https://github.com/muheheheh/SaveAny-Bot) adds direct Telegram media relay to the upstream project.

- Enable `reuse_media = true` on Telegram storage to send single files and albums using existing references, without downloading or caching media files.
- Album order and source captions are retained; captions are sent as plain text.
- Rejected references fail the task without falling back to downloading.
- Existing website downloads, storage backends, rules, and parsers remain available with their normal transfer and caching behavior.

Start with [installation and updates](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/en/deployment/installation.md), or see [Telegram storage settings](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/en/deployment/configuration/storages.md#telegram). This repository contains the documentation for this fork.

## Origin and license

Based on [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot), with thanks to the [upstream contributors](https://github.com/krau/SaveAny-Bot/graphs/contributors). The AGPL-3.0 license and upstream notices are retained.
