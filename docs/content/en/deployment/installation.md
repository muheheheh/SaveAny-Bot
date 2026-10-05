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
| `[[users]].id` | Your numeric Telegram user ID, authorized to use the bot |

See [configuration](./configuration/_index.md) for all settings.

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

Open a private chat with the bot and send `/start`. See [usage](../usage/_index.md) for instructions.

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

## Updates

Run from the repository root:

```sh
git pull --ff-only
docker compose up -d --build
```

Configuration and data persist in the mounted directories and are reused by updated containers. After changing `config.toml`, run `docker compose restart` to apply it.
