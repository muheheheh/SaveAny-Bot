# SaveAny-Bot

支持 Telegram 媒体回传、网站下载和多存储转存的机器人。Telegram 媒体可通过已有引用直接发送，无需服务器下载或重新上传。

[使用说明](./docs/content/zh/usage/_index.md) · [VPS 部署](./deploy/vps/README.md) · [配置说明](./docs/content/zh/deployment/configuration/_index.md) · [反馈问题](https://github.com/muheheheh/SaveAny-Bot/issues)

## 支持的功能

| 功能 | 说明 |
| --- | --- |
| Telegram 媒体快速回传 | 启用 `reuse_media = true` 后，直接发送已有媒体引用，支持单文件和相册 |
| 相册与说明文字 | 保留媒体顺序和各条消息的说明文字；说明文字按纯文本发送 |
| Telegram 消息链接 | 发送消息链接，读取其中的媒体并回传或转存；需要具备源消息访问权限 |
| 批量处理 | 支持相册和按消息 ID 范围批量保存，可按消息文本过滤 |
| 自动处理 | 用 `/storage` 设置默认存储，用 `/silent` 开启收到媒体后自动处理 |
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

网站解析取决于对应解析器、登录状态和目标网站的可访问性；yt-dlp、Aria2、Rclone 等功能需要相应工具或服务。

## 文件处理

| 使用场景 | 文件处理方式 |
| --- | --- |
| Telegram 媒体 → 启用 `reuse_media` 的 Telegram 存储 | 只复用媒体引用，不下载、不缓存、不重新上传媒体文件 |
| 网站、直链、yt-dlp、Aria2 → Telegram 或其他存储 | 按原有下载和上传流程处理，可能产生临时文件 |
| Telegram 媒体 → 本地磁盘或其他存储 | 按对应存储的下载、流式传输或上传流程处理 |

快速回传保留原文件名和媒体类型，不执行重命名、格式转换、强制文件发送或分卷。引用失效、权限不足或 Telegram 拒绝发送时，任务直接报错，不会自动切换为下载重传。

机器人仍需在服务器保存配置、会话和用户设置等运行数据。纯文本聊天记录、投票等不属于媒体回传范围；受访问权限或内容保护限制的消息不保证可处理。

## 快速开始

### 1. 准备环境

- 一台能连接 Telegram 的服务器。
- Docker 和 Docker Compose；小内存 VPS 可按 [VPS 部署说明](./deploy/vps/README.md) 在本地编译后上传。
- 在 Telegram 的 `@BotFather` 创建机器人，取得 Bot Token。
- 准备自己的 Telegram **数字用户 ID**，它与用户名、手机号和机器人 ID 不同。

### 2. 下载代码并配置

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp deploy/vps/config.example.toml config.toml
chmod 600 config.toml
```

编辑 `config.toml`，填写以下内容：

| 配置项 | 填写方式 |
| --- | --- |
| `[telegram].token` | BotFather 提供的 Bot Token |
| `[[storages]].chat_id` | 你自己的 Telegram 数字用户 ID，机器人将媒体返回到这个私聊 |
| `[[users]].id` | 同一个数字用户 ID，仅允许该用户使用机器人 |
| `[[storages]].reuse_media` | 保持 `true`，启用 Telegram 快速回传 |
| `lang` | 保持 `"zh-Hans"`，使用简体中文消息 |

模板已配置好名为“回传到聊天”的存储和用户白名单。`chat_id` 与 `id` 填写数字，不要填写 `@用户名`。

### 3. 启动机器人

在源码目录执行：

```sh
docker compose up -d --build
docker compose logs --tail=100 -f
```

按 `Ctrl+C` 退出日志查看，容器会继续运行。

### 4. 在 Telegram 中使用

1. 打开自己的机器人私聊，点击 **Start** 或发送 `/start`。
2. 发送 `/storage`，选择 **回传到聊天**，设置为默认存储。
3. 发送一次 `/silent`，确认机器人提示已开启静默模式。这个命令是开关，再发一次会关闭。
4. 发送或转发图片、视频、文档或整组相册，也可以发送机器人能够读取的 Telegram 消息链接。
5. 机器人会把媒体发送到配置的目标聊天，并反馈处理结果。

关闭静默模式时，机器人会先让你选择目标存储。完整操作和更多场景见 [使用说明](./docs/content/zh/usage/_index.md)。

## 常用命令

| 命令 | 用途 |
| --- | --- |
| `/start`、`/help` | 查看帮助 |
| `/storage` | 选择默认存储 |
| `/silent` | 开启或关闭自动处理 |
| `/save` | 回复一条媒体消息时保存它；也支持按聊天和消息 ID 范围批量保存 |
| `/task`、`/task queued` | 查看执行中的任务、查看排队任务 |
| `/cancel <任务ID>` | 取消指定任务 |
| `/dl <链接>` | 下载 HTTP/HTTPS 文件直链 |
| `/ytdlp <链接>` | 通过 yt-dlp 下载视频或音频 |
| `/aria2dl <链接>` | 通过已配置的 Aria2 下载 |
| `/rule`、`/dir` | 管理存储规则和目录 |
| `/config`、`/fnametmpl` | 管理用户文件命名和重名策略；快速回传仍保留原媒体文件名 |
| `/transfer` | 查看存储间传输的用法 |
| `/parser` | 查看解析器信息 |
| `/watch`、`/unwatch`、`/lswatch` | 管理聊天监听，需要配置 UserBot |

## 部署与更新

从源码构建的 Docker 部署可在源码目录执行：

```sh
git pull --ff-only
docker compose up -d --build
```

小内存 VPS 使用 [本地编译后部署](./deploy/vps/README.md) 的更新步骤。更新前保留配置、`data` 目录和当前使用的程序或镜像。

`/update` 和 `./saveany-bot up` 需要可用的 GitHub Release。Docker 部署使用上述重建步骤更新。

## 文档

- [安装与更新](./docs/content/zh/deployment/installation.md)
- [VPS 部署](./deploy/vps/README.md)
- [使用说明](./docs/content/zh/usage/_index.md)
- [配置说明](./docs/content/zh/deployment/configuration/_index.md)
- [存储配置](./docs/content/zh/deployment/configuration/storages.md)
- [存储规则](./docs/content/zh/usage/rules.md) · [存储间传输](./docs/content/zh/usage/transfer.md)
- [HTTP API](./docs/content/zh/usage/api.md) · [命令行使用](./docs/content/zh/usage/cli.md)
- [解析器插件开发](./plugins/README.md)

## 来源与许可证

本项目基于 [krau](https://github.com/krau) 和 [上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors) 开发的 SaveAny-Bot，保留 [AGPL-3.0 许可证](./LICENSE) 及原有版权声明。

感谢 [gotd](https://github.com/gotd/td)、[gotgproto](https://github.com/celestix/gotgproto)、[TG-FileStreamBot](https://github.com/EverythingSuckz/TG-FileStreamBot)、[tdl](https://github.com/iyear/tdl) 及所有上游依赖和贡献者。
