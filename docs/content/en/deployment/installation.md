---
title: "Docker Deployment"
---

# Docker Deployment

Use a Linux host with Docker and Docker Compose installed and network access to Telegram.

## 1. Download the source

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp config.docker.example.toml config.toml
chmod 600 config.toml
```

## 2. Configure the bot

Edit `config.toml`:

| Setting | Value |
| --- | --- |
| `[telegram].token` | The bot token provided by BotFather |
| `[[storages]].chat_id` | Your numeric Telegram user ID, which receives relayed media |
| `[[users]].id` | The same numeric user ID, authorized to use the bot |

The template enables `reuse_media = true` and sends media to your private chat with the bot. See [configuration](./configuration/_index.md) for other storage backends and settings.

## 3. Start

Run from the repository root:

```sh
docker compose up -d --build
```

Docker builds the image and starts the container. The runtime image includes FFmpeg and yt-dlp.

Check status and logs:

```sh
docker compose ps
docker compose logs --tail=100 -f
```

Press `Ctrl+C` to stop following logs; the container keeps running.

Open the bot's private chat, send `/start`, select the relay storage with `/storage`, then enable automatic processing with `/silent`. See [usage](../usage/_index.md) for details.

## Common operations

| Operation | Command |
| --- | --- |
| Show status | `docker compose ps` |
| Follow logs | `docker compose logs --tail=100 -f` |
| Restart | `docker compose restart` |
| Stop | `docker compose stop` |
| Start stopped containers | `docker compose start` |

## Data directories

Configuration and data are mounted from the project directory:

| File or directory | Purpose |
| --- | --- |
| `config.toml` | Bot, user, and storage configuration |
| `data/` | Sessions, database, and user settings |
| `downloads/` | Default destination for local storage |
| `cache/` | Temporary files for download tasks |

Telegram media relay uses existing references without writing media cache files. Website downloads and other storage tasks use their respective transfer paths.

## Updates

Run from the repository root:

```sh
git pull --ff-only
docker compose up -d --build
```

Configuration and data persist in the mounted directories and are reused by updated containers. After changing `config.toml`, run `docker compose restart` to apply it.
