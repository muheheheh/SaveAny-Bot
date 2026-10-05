# SaveAny-Bot

SaveAny-Bot 是一个 Telegram 文件机器人，支持媒体回传、文件下载和存储管理。

[使用说明](./docs/content/zh/usage/_index.md) · [Docker 部署](./docs/content/zh/deployment/installation.md) · [配置说明](./docs/content/zh/deployment/configuration/_index.md) · [反馈问题](https://github.com/muheheheh/SaveAny-Bot/issues)

## 支持的功能

| 功能 | 说明 |
| --- | --- |
| Telegram 媒体回传 | 发送或转发图片、视频、文件、相册，机器人会发回给你 |
| 相册与说明文字 | 保留媒体顺序和各条消息的说明文字；说明文字按纯文本发送 |
| Telegram 消息链接 | 发送消息链接，获取其中的图片、视频或文件 |
| 批量处理 | 支持相册和按消息 ID 范围批量保存，可按消息文本过滤 |
| 自动保存 | 配置存储后，用 `/storage` 设置默认存储，用 `/silent` 开启自动保存 |
| 多种存储 | 支持 Telegram、本地磁盘、S3、MinIO、WebDAV、AList 和 Rclone |
| 文件直链下载 | 用 `/dl` 下载一个或多个 HTTP/HTTPS 文件链接 |
| 网站内容解析 | 支持 Telegraph 文章图片、内置 Twitter/X 与 Kemono 解析器，可用 JavaScript 插件扩展 |
| 视频与音频下载 | 通过 `/ytdlp` 调用 yt-dlp，支持其可解析的网站及自定义下载参数 |
| Aria2 下载 | 配置 Aria2 RPC 后，通过 `/aria2dl` 提交下载任务 |
| 存储规则与文件管理 | 按规则选择存储和目录，设置文件命名、重名策略，进行存储间传输 |
| 多用户访问 | 按 Telegram 数字用户 ID 授权，并分别指定用户可用的存储 |
| 任务管理 | 查看执行中和排队中的任务，查看进度并取消任务 |
| 聊天监听 | 配置 UserBot 后监听可访问的聊天，支持过滤条件和任务通知 |
| 命令行与 HTTP API | 支持本地文件上传、目录监听，以及通过 API 创建、查询、取消任务和接收 Webhook 回调 |

## Docker 部署

### 1. 准备环境

- 一台能连接 Telegram 的 Linux 主机，已安装 Docker 和 Docker Compose。
- 在 Telegram 的 `@BotFather` 创建机器人，取得 Bot Token。
- 准备自己的 Telegram **数字用户 ID**，它与用户名、手机号和机器人 ID 不同。

### 2. 下载代码并配置

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp config.docker.example.toml config.toml
chmod 600 config.toml
```

编辑 `config.toml`，填写以下内容：

| 配置项 | 填写方式 |
| --- | --- |
| `[telegram].token` | BotFather 提供的 Bot Token |
| `[[users]].id` | 允许使用机器人的 Telegram 数字用户 ID |
| `lang` | 保持 `"zh-Hans"`，使用简体中文消息 |

多用户和文件保存的配置见 [配置说明](./docs/content/zh/deployment/configuration/_index.md#基础配置)。

### 3. 启动机器人

在源码目录执行：

```sh
docker compose up -d --build
docker compose logs --tail=100 -f
```

按 `Ctrl+C` 退出日志查看，容器会继续运行。

### 4. 在 Telegram 中使用

1. 打开自己的机器人私聊，点击 **Start** 或发送 `/start`。
2. 发送或转发图片、视频、文档、整组相册，或发送机器人能够读取的 Telegram 媒体消息链接。
3. 机器人会把图片、视频或文件发回给你。

文件保存、网站下载等操作见 [使用说明](./docs/content/zh/usage/_index.md)。

## 常用命令

| 命令 | 用途 |
| --- | --- |
| `/start`、`/help` | 查看帮助 |
| `/storage` | 选择默认存储 |
| `/silent` | 开启或关闭自动保存 |
| `/save` | 回复一条媒体消息时保存它；也支持按聊天和消息 ID 范围批量保存 |
| `/task`、`/task queued` | 查看执行中的任务、查看排队任务 |
| `/cancel <任务ID>` | 取消指定任务 |
| `/dl <链接>` | 下载 HTTP/HTTPS 文件直链 |
| `/ytdlp <链接>` | 通过 yt-dlp 下载视频或音频 |
| `/aria2dl <链接>` | 通过已配置的 Aria2 下载 |
| `/rule`、`/dir` | 管理存储规则和目录 |
| `/config`、`/fnametmpl` | 设置文件命名和重名处理方式 |
| `/transfer` | 查看存储间传输的用法 |
| `/parser` | 查看解析器信息 |
| `/watch`、`/unwatch`、`/lswatch` | 管理聊天监听，需要配置 UserBot |

## 更新

在项目目录执行：

```sh
git pull --ff-only
docker compose up -d --build
```

配置和数据保存在挂载目录中。

## 文档

- [Docker 部署](./docs/content/zh/deployment/installation.md)
- [使用说明](./docs/content/zh/usage/_index.md)
- [配置说明](./docs/content/zh/deployment/configuration/_index.md)
- [存储配置](./docs/content/zh/deployment/configuration/storages.md)
- [存储规则](./docs/content/zh/usage/rules.md) · [存储间传输](./docs/content/zh/usage/transfer.md)
- [HTTP API](./docs/content/zh/usage/api.md) · [命令行使用](./docs/content/zh/usage/cli.md)
- [解析器插件开发](./plugins/README.md)

## 来源与许可证

本项目基于 [krau](https://github.com/krau) 和 [上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors) 开发的 SaveAny-Bot，保留 [AGPL-3.0 许可证](./LICENSE) 及原有版权声明。

感谢 [gotd](https://github.com/gotd/td)、[gotgproto](https://github.com/celestix/gotgproto)、[TG-FileStreamBot](https://github.com/EverythingSuckz/TG-FileStreamBot)、[tdl](https://github.com/iyear/tdl) 及所有上游依赖和贡献者。
