---
title: Introduction
---

# SaveAny-Bot

[muheheheh/SaveAny-Bot](https://github.com/muheheheh/SaveAny-Bot) adds direct Telegram media relay to the upstream project.

- Enable `reuse_media = true` on Telegram storage to send single files and albums using existing references, without downloading or caching media files.
- Album order and source captions are retained; captions are sent as plain text.
- Rejected references fail the task without falling back to downloading.
- Existing website downloads, storage backends, rules, and parsers remain available with their normal transfer and caching behavior.

Start with [installation and updates](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/en/deployment/installation.md), or see [Telegram storage settings](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/en/deployment/configuration/storages.md#telegram). This repository contains the documentation for this fork.

## Documentation

- [Supported features and quick start](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md)
- [Usage, commands, and troubleshooting](./usage/_index.md)
- [VPS deployment, updates, and rollback](https://github.com/muheheheh/SaveAny-Bot/blob/main/deploy/vps/README.md)
- [Configuration](./deployment/configuration/_index.md)
- [HTTP API](./usage/api.md) and [CLI](./usage/cli.md)

## Origin and license

Based on [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot), with thanks to the [upstream contributors](https://github.com/krau/SaveAny-Bot/graphs/contributors). The AGPL-3.0 license and upstream notices are retained.
