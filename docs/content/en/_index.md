---
title: Introduction
---

# SaveAny-Bot

SaveAny-Bot supports Telegram media relay, website downloads, and transfers to multiple storage backends.

- Enable `reuse_media = true` on Telegram storage to send single files and albums using existing references, without downloading or caching media files.
- Album order and source captions are retained; captions are sent as plain text.
- Rejected references fail the task without falling back to downloading.
- Website downloads and other storage operations may use temporary files.

## Documentation

- [Supported features and quick start](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md)
- [Usage](./usage/_index.md)
- [Installation and updates](./deployment/installation.md)
- [VPS deployment](https://github.com/muheheheh/SaveAny-Bot/blob/main/deploy/vps/README.md)
- [Configuration](./deployment/configuration/_index.md)
- [Storage configuration](./deployment/configuration/storages.md)
- [HTTP API](./usage/api.md) and [CLI](./usage/cli.md)

## Origin and license

Based on [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot), with thanks to the [upstream contributors](https://github.com/krau/SaveAny-Bot/graphs/contributors). The AGPL-3.0 license and upstream notices are retained.
