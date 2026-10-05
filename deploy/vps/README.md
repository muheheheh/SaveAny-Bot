# VPS deployment

Build the Linux amd64 binary from this fork on your development machine:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath \
  -ldflags '-s -w -X github.com/krau/SaveAny-Bot/config.Version=0.62.1-fast-relay -X github.com/krau/SaveAny-Bot/config.Docker=true' \
  -o deploy/vps/saveany-bot .
```

Copy `saveany-bot`, `Dockerfile`, `.dockerignore`, and `compose.yaml` into the
server deployment directory. Create `config.toml` from `config.example.toml`,
fill in the bot token and your numeric Telegram user ID, and set its mode to
`600`. Use the same ID for the Telegram storage destination and allowed user.

```sh
mkdir -p data cache
chmod 700 data cache
docker compose build
docker compose up -d
```

The image retains the official release's FFmpeg and yt-dlp runtime tools.
`reuse_media = true` sends existing Telegram photos, videos, documents, and
albums without downloading or uploading media bytes. Website downloads still
use the original transfer path and temporary cache. Only `data` (settings and
session) and `cache` are mounted; the bot has no published network ports.

Keep a copy of the previous binary, compose file, and configuration before
upgrading. To roll back, restore those files, rebuild the local image, and run
`docker compose up -d` again. Avoid the bot's upstream `/update` command when
running this fork, since it installs an official binary without these changes.
