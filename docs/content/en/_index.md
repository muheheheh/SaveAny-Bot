---
title: Introduction
---

# SaveAny-Bot

SaveAny-Bot supports Telegram media relay, website downloads, and transfers to multiple storage backends.

- Private-chat media and Telegram message links automatically relay to the current chat. Relay alone does not download or cache media.
- Available storage enables saving after relay. Without storage, successful relay leaves no progress messages.
- Album order and source captions are retained; captions are sent as plain text.
- Rejected references leave an error without falling back to downloading.
- Website downloads and other storage operations may use temporary files.

## Documentation

- [Supported features](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md#支持的功能)
- [Usage](./usage/_index.md)
- [Docker deployment](./deployment/installation.md)
- [Configuration](./deployment/configuration/_index.md)
- [Storage configuration](./deployment/configuration/storages.md)
- [HTTP API](./usage/api.md) and [CLI](./usage/cli.md)

## Origin and license

Based on [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot), with thanks to the [upstream contributors](https://github.com/krau/SaveAny-Bot/graphs/contributors). The AGPL-3.0 license and upstream notices are retained.
