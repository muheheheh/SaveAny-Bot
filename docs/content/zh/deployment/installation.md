---
title: "Docker 部署"
---

# Docker 部署

需要一台已安装 Docker 和 Docker Compose 的 Linux 主机，并能连接 Telegram。

## 1. 下载代码

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp config.docker.example.toml config.toml
chmod 600 config.toml
```

## 2. 配置机器人

编辑 `config.toml`：

| 配置项 | 内容 |
| --- | --- |
| `[telegram].token` | BotFather 提供的 Bot Token |
| `[[storages]].chat_id` | 接收回传媒体的 Telegram 数字用户 ID |
| `[[users]].id` | 允许使用机器人的 Telegram 数字用户 ID，与上面填写相同值 |

模板已开启 `reuse_media = true`，默认将媒体回传到你与机器人的私聊。配置层级、完整回传示例和多用户配置见 [配置文件结构](./configuration/_index.md#配置文件结构)。

## 3. 启动

在项目根目录执行：

```sh
docker compose up -d --build
```

Docker 会构建镜像并启动容器，运行镜像包含 FFmpeg 和 yt-dlp。

查看状态和日志：

```sh
docker compose ps
docker compose logs --tail=100 -f
```

按 `Ctrl+C` 退出日志查看，容器继续运行。

打开机器人私聊，发送 `/start`，通过 `/storage` 选择“回传到聊天”，再用 `/silent` 开启自动处理。详细操作见 [使用说明](../usage/_index.md)。

## 常用操作

| 操作 | 命令 |
| --- | --- |
| 查看状态 | `docker compose ps` |
| 查看日志 | `docker compose logs --tail=100 -f` |
| 重启 | `docker compose restart` |
| 停止 | `docker compose stop` |
| 启动已停止的容器 | `docker compose start` |

## 数据目录

配置和数据通过项目目录挂载到容器：

| 文件或目录 | 用途 |
| --- | --- |
| `config.toml` | 机器人、用户和存储配置 |
| `data/` | 会话、数据库和用户设置 |
| `downloads/` | 本地存储的默认保存目录 |
| `cache/` | 下载任务的临时目录 |

Telegram 媒体回传使用已有媒体引用，不写入媒体缓存。网站下载和其他存储任务按各自流程处理文件。

## 更新

在项目根目录执行：

```sh
git pull --ff-only
docker compose up -d --build
```

配置和数据保存在挂载目录中，更新容器时会继续使用。修改 `config.toml` 后，执行 `docker compose restart` 使配置生效。
