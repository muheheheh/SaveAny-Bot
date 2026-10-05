---
title: "Installation and Updates"
---

# Installation and Updates

## Build this fork from source

Install Go 1.25 or newer, then build from this repository:

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
CGO_ENABLED=0 go build -trimpath -o saveany-bot .
```

Build from source until this fork publishes releases. Upstream release binaries do not include this fork's relay changes.

Copy `deploy/vps/config.example.toml` to `config.toml`, then fill in the bot token, destination chat ID, and allowed user ID. The template enables media reuse. See the [Configuration Guide](../configuration) for other settings.

Run:

```bash
chmod +x saveany-bot
./saveany-bot
```

### Daemon

{{< tabs "daemon" >}}
{{< tab "systemd (Regular Linux)" >}}

Create a file <code>/etc/systemd/system/saveany-bot.service</code> and write the following content:

{{< codeblock >}}
[Unit]
Description=SaveAnyBot
After=systemd-user-sessions.service

[Service]
Type=simple
WorkingDirectory=/yourpath/
ExecStart=/yourpath/saveany-bot
Restart=always

[Install]
WantedBy=multi-user.target
{{< /codeblock >}}

Enable startup on boot and start the service:

{{< codeblock >}}
systemctl enable --now saveany-bot
{{< /codeblock >}}

{{< /tab >}}

{{< tab "procd (OpenWrt)" >}}

<h4>Add Boot Autostart Service</h4>

Create a file <code>/etc/init.d/saveanybot</code>, refer to <a href="https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/confs/wrt_init" target="_blank">wrt_init</a> and modify as needed:

{{< codeblock >}}
#!/bin/sh /etc/rc.common

#This is the OpenWRT init.d script for SaveAnyBot

START=99 
STOP=10
description="SaveAnyBot"

WORKING_DIR="/mnt/mmc1-1/SaveAnyBot"
EXEC_PATH="$WORKING_DIR/saveany-bot"
start() {
    echo "Starting SaveAnyBot..."
    cd $WORKING_DIR
    $EXEC_PATH &
}
stop() {
    echo "Stopping SaveAnyBot..."
    killall saveany-bot
}
reload() {
    stop
    start
}

{{< /codeblock >}}

Set permissions:

{{< codeblock >}}
chmod +x /etc/init.d/saveanybot
{{< /codeblock >}}

Then copy the file to <code>/etc/rc.d</code> and rename it to <code>S99saveanybot</code>, also set permissions:

{{< codeblock >}}
chmod +x /etc/rc.d/S99saveanybot
{{< /codeblock >}}

<h4>Add Shortcut Commands</h4>

Create a file <code>/usr/bin/sabot</code>, refer to <a href="https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/confs/wrt_bin" target="_blank">wrt_bin</a> and modify as needed. Note that the file encoding here only supports ANSI 936.

Then set permissions:

{{< codeblock >}}
chmod +x /usr/bin/sabot
{{< /codeblock >}}

Usage: <code>sudo sabot start|stop|restart|status|enable|disable</code>

{{< /tab >}}
{{< /tabs >}}


## Deploy using Docker

From a full checkout of this repository, prepare `config.toml` as above and run:

```sh
docker compose up -d --build
```

The root Compose file builds this fork instead of pulling the upstream application image.
For a small VPS, [build locally and deploy the binary](https://github.com/muheheheh/SaveAny-Bot/blob/main/deploy/vps/README.md).

Slim variants can be built from this fork's `Dockerfile.micro` or `Dockerfile.pico`. Choose the variant according to the dependencies and build tags in its Dockerfile.

## Updates

Pull this fork, then rebuild and recreate the container:

```sh
git pull --ff-only
docker compose up -d --build
```

For the VPS template, rebuild locally, upload the binary, and run `docker compose up -d --build` in the server deployment directory.

For native binary deployments, repeat the Go build command, replace the executable, and restart the service.
`/update` and `./saveany-bot up` only query this fork's GitHub releases and cannot update until releases are published here. Docker deployments are updated by rebuilding their images.
