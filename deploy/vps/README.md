# VPS 部署说明

适用于在开发机器上编译程序，再上传到 Linux VPS 运行的部署方式，可减少服务器上的编译资源占用。功能介绍和日常操作见 [项目首页](../../README.md) 与 [使用说明](../../docs/content/zh/usage/_index.md)。

## 部署前准备

- 开发机器安装 Go 1.25 或更新版本，并已下载本仓库源码。
- 服务器安装 Docker 和 Docker Compose，能连接 Telegram 和镜像仓库。
- 准备 BotFather 提供的 Bot Token，以及自己的 Telegram 数字用户 ID。

以下示例使用 Linux x86-64 服务器。命令中的 `root@your-server` 需替换为实际 SSH 用户和服务器地址。

## 1. 在本地编译

在本仓库根目录执行：

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags '-s -w -X github.com/krau/SaveAny-Bot/config.Version=0.62.1-fast-relay -X github.com/krau/SaveAny-Bot/config.Docker=true' -o deploy/vps/saveany-bot .
```

ARM64 服务器将 `GOARCH=amd64` 改为 `GOARCH=arm64`，并确认所用基础镜像支持对应架构。

## 2. 准备配置并上传

首次部署时，在本地复制配置模板：

```sh
cp deploy/vps/config.example.toml config.toml
chmod 600 config.toml
```

编辑 `config.toml`，填写 `[telegram].token`，并将 `[[storages]].chat_id` 与 `[[users]].id` 都改为你自己的 Telegram 数字用户 ID。模板默认使用中文、仅允许该用户访问，并开启 `reuse_media = true`。

创建服务器部署目录并上传程序、构建文件和配置：

```sh
ssh root@your-server 'mkdir -p /opt/saveany-bot'
scp deploy/vps/saveany-bot deploy/vps/Dockerfile deploy/vps/.dockerignore deploy/vps/compose.yaml config.toml root@your-server:/opt/saveany-bot/
```

## 3. 在服务器启动

登录服务器，在部署目录执行：

```sh
cd /opt/saveany-bot
chmod 600 config.toml
mkdir -p data cache
chmod 700 data cache
docker compose up -d --build
docker compose ps
docker compose logs --tail=100 -f
```

按 `Ctrl+C` 退出日志查看，容器继续运行。Compose 已设置自动重启。

该模板以固定版本的上游镜像提供 FFmpeg、yt-dlp 等运行依赖，再用本仓库编译的程序替换其中的可执行文件。快速回传功能来自上传的程序。

## 4. 首次使用

1. 在 Telegram 中打开机器人私聊，点击 **Start** 或发送 `/start`。
2. 发送 `/storage`，选择“回传到聊天”。
3. 发送 `/silent`，确认提示静默模式已开启。
4. 发送或转发媒体、相册，或发送机器人有权限读取的 Telegram 消息链接。

再次发送 `/silent` 会关闭自动处理。更多命令见 [使用说明](../../docs/content/zh/usage/_index.md)。

## 服务器上的文件

| 文件或目录 | 用途 |
| --- | --- |
| `config.toml` | Bot Token、用户白名单和存储配置，以只读方式挂载到容器 |
| `data/` | 会话、用户设置和数据库，需要持久保存 |
| `cache/` | 普通下载任务可能使用的临时目录 |

启用 `reuse_media` 的 Telegram 媒体回传不下载、不缓存、不重新上传媒体文件。网站下载、yt-dlp 和其他存储操作仍按原有流程处理，可能占用临时磁盘空间。

模板默认关闭 HTTP API，Compose 不发布网络端口。若要使用其他存储或 API，请按 [配置说明](../../docs/content/zh/deployment/configuration/_index.md) 单独配置。

## 更新

1. 在本地执行 `git pull --ff-only`，按上面的命令重新编译。
2. 在服务器保存当前程序、Compose 文件、配置及 `data/` 的备份。
3. 上传新的 `saveany-bot`；仅在模板有改动时同步构建文件，保留已填写的 `config.toml`。
4. 在服务器的 `/opt/saveany-bot` 目录执行 `docker compose up -d --build`，并检查启动日志。

Docker 部署通过重建镜像和容器更新。`/update` 查询本仓库的 Release，不能代替上述 Docker 更新步骤。

## 回滚

恢复上一个版本的程序、Compose 文件和配置，在部署目录执行 `docker compose up -d --build`，并检查启动日志。保留原有 `data/`；若升级涉及数据库结构变化，应使用与旧程序匹配的数据备份。
