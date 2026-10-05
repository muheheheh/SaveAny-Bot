# SaveAny-Bot (Telegram relay fork)

**English** | [简体中文](./README_zh.md)

A fork of [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot), maintained in
[muheheheh/SaveAny-Bot](https://github.com/muheheheh/SaveAny-Bot), with direct
Telegram media relay. Existing photos, videos, documents, and albums can be
sent back using Telegram media references, without downloading or uploading
the media bytes.

## What this fork adds

- Opt-in `reuse_media = true` for Telegram storage, covering single files and albums.
- Original album order and plain-text captions are preserved.
- No media cache files or automatic download fallback on the relay path.
- Outgoing bot messages are ignored to prevent replies being processed again.

The upstream storage backends, website parsers, yt-dlp, Aria2, rules, and other
features remain available. Website downloads and saves to other backends use
their normal transfer paths and may create temporary files. Reusing Telegram
media preserves its original filename and type; renaming, conversion, and
splitting do not apply. The bot must be able to read the source and Telegram
must accept its media reference.

## Quick start

Build from this repository to include the relay changes:

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp deploy/vps/config.example.toml config.toml
chmod 600 config.toml
```

Edit `config.toml`: set the BotFather token, and replace both `chat_id = 0`
and `id = 0` with your own numeric Telegram user ID. The template enables
`reuse_media` and restricts access to that user.

```sh
docker compose up -d --build
```

Start the bot in Telegram, use `/storage` to select the relay destination,
then `/silent` to enable automatic processing. Send a Telegram message link
or a media message to use the relay.

For a small VPS, [build the binary locally and deploy it](./deploy/vps/README.md).
The original upstream image does not include this fork's changes; the VPS
template uses a pinned upstream image only for runtime dependencies and
replaces its executable with the binary built from this fork.

## Documentation and updates

- [Installation and updates](./docs/content/en/deployment/installation.md)
- [Configuration](./docs/content/en/deployment/configuration/_index.md)
- [Telegram storage and relay settings](./docs/content/en/deployment/configuration/storages.md#telegram)
- [Usage](./docs/content/en/usage/_index.md)
- [Report an issue in this fork](https://github.com/muheheheh/SaveAny-Bot/issues)

Documentation is maintained in this repository. For source-built deployments,
pull this fork and rebuild. `/update` and the CLI updater check releases in
`muheheheh/SaveAny-Bot`; until releases are published here, update from source.

## Origin and license

Based on SaveAny-Bot by [krau](https://github.com/krau) and the
[upstream contributors](https://github.com/krau/SaveAny-Bot/graphs/contributors).
This fork retains the [AGPL-3.0 license](./LICENSE) and upstream notices.
The Go module/import path remains `github.com/krau/SaveAny-Bot` for source
compatibility; it is not the deployment or update source.

Thanks to [gotd](https://github.com/gotd/td),
[gotgproto](https://github.com/celestix/gotgproto),
[TG-FileStreamBot](https://github.com/EverythingSuckz/TG-FileStreamBot),
[tdl](https://github.com/iyear/tdl), and all upstream dependencies and contributors.
