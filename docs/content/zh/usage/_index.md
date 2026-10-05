---
title: "使用说明"
weight: 10
---

# 使用说明

SaveAny-Bot 可以将 Telegram 媒体回传到指定聊天，也可以将 Telegram 或网站上的文件保存到其他存储。完整功能列表见 [项目首页](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md#支持的功能)。

## 首次使用

1. 按 [Docker 部署](../deployment/installation.md) 启动机器人，使用 `config.docker.example.toml` 配置模板。
2. 在配置中填写 Bot Token，将 Telegram 存储的 `chat_id` 和允许用户的 `id` 都设为你自己的数字用户 ID，并开启该存储的 `reuse_media = true`。
3. 打开机器人私聊，点击 **Start** 或发送 `/start`。
4. 发送 `/storage`，选择“回传到聊天”作为默认存储。
5. 发送 `/silent`，确认机器人提示静默模式已开启，然后发送或转发媒体。

默认存储和静默模式按用户保存。`/silent` 是切换开关，再发一次会关闭；关闭时，机器人会先询问本次保存的目标存储。

## Telegram 媒体回传

| 输入方式 | 操作方法 |
| --- | --- |
| 单张图片、视频或文档 | 直接发送或转发给机器人 |
| 相册或媒体组 | 将整组媒体发送或转发给机器人，尽量保留原有分组 |
| Telegram 消息链接 | 发送类似 `https://t.me/channel_username/123` 的链接，机器人读取其中的媒体 |
| 已有媒体消息 | 回复该媒体消息并发送 `/save` |
| 多条历史消息 | 使用 `/save <聊天ID或用户名> <起始消息ID>-<结束消息ID> [文本正则]` 按范围保存 |

选择启用 `reuse_media` 的 Telegram 存储时，上述 Telegram 媒体任务通过已有媒体引用发送。相册保留顺序，每条消息的说明文字按纯文本保留。目标由存储配置决定；使用个人回传模板时，就是你与机器人的私聊。

消息链接、批量历史消息和私密聊天都需要相应访问权限。启用 UserBot 需要另行登录用户账号，也需要该账号本身能够访问源消息；快速回传不会绕过 Telegram 的权限或内容保护。

## 自动处理与任务管理

| 命令 | 作用 |
| --- | --- |
| `/storage` | 设置默认存储 |
| `/silent` | 开启或关闭自动处理 |
| `/task` | 列出正在执行的任务 |
| `/task queued` | 列出排队中的任务 |
| `/cancel <任务ID>` | 取消指定任务，也可点击进度消息上的取消按钮 |
| `/help` | 查看机器人帮助 |

开启静默模式后，支持的媒体和链接会按默认存储及已配置的规则处理。网站链接使用对应的解析器或下载器。

## 其他支持的功能

| 场景 | 操作或前提 | 详细说明 |
| --- | --- | --- |
| 保存到其他存储 | 配置本地磁盘、S3、MinIO、WebDAV、AList、Rclone 或 Telegram 存储 | [存储端配置](../deployment/configuration/storages.md) |
| HTTP/HTTPS 文件直链 | `/dl <链接1> [链接2]` | [直接下载链接](./directlinks.md) |
| 网站图片和媒体 | 发送 Telegraph 文章链接，或发送 Twitter/X、Kemono 等已启用解析器支持的链接 | [网站解析](./parsers.md) |
| 网站视频与音频 | `/ytdlp <链接> [参数]`，需要安装 yt-dlp，部分处理需要 FFmpeg | [yt-dlp 使用](./ytdlp.md) |
| Aria2 下载 | `/aria2dl <链接>`，需要启用并配置 Aria2 RPC | [Aria2 使用](./aria2.md) |
| 自动选择存储或目录 | 使用 `/rule` 配置规则，使用 `/dir` 管理目录 | [存储规则](./rules.md) |
| 文件命名与重名处理 | `/config`、`/fnametmpl`；快速回传保留原媒体文件名 | [文件命名](./config.md) |
| 存储间传输 | `/transfer`，源存储需支持文件列举和读取 | [存储间传输](./transfer.md) |
| 监听聊天 | `/watch`、`/unwatch`、`/lswatch`，需要启用 UserBot | [监听与通知](./watch.md) |
| 本地文件上传与目录监听 | 在服务器运行 `saveany-bot upload` 或 `saveany-bot watch` | [命令行使用](./cli.md) |
| 程序调用与回调 | 配置 HTTP API，可创建、查询、取消任务并接收 Webhook | [HTTP API](./api.md) |
| 扩展网站支持 | 启用并加载 JavaScript 解析器插件 | [插件开发](https://github.com/muheheheh/SaveAny-Bot/blob/main/plugins/README.md) |

## 快速回传的范围

- 仅 Telegram 来源的媒体发送到启用 `reuse_media` 的 Telegram 存储时，不下载、不缓存、不重新上传媒体文件。
- 网站下载和其他存储仍按原有流程传输文件，可能使用服务器临时目录。
- 保留原媒体类型和文件名，不应用强制文件发送、重命名、转码或分卷；说明文字中的富文本格式不会重建。
- 无法获取或使用媒体引用时，任务报错，不会自动下载重传。
- 纯文本聊天记录、投票等不在媒体回传范围内。
- 配置、会话、用户设置和日志仍属于服务器运行数据。

## 常见问题

| 现象 | 检查方法 |
| --- | --- |
| 提示没有权限 | 检查 `[[users]].id` 是否为发消息的用户数字 ID；修改配置后重启程序 |
| 提示没有可用存储 | 检查存储已启用，且用户的存储过滤规则允许访问它 |
| 每次都要求选择存储 | 用 `/storage` 设置默认存储，再用 `/silent` 开启静默模式 |
| 无法读取消息链接 | 检查链接和访问权限；普通用户能看到的消息不一定对机器人账号可见 |
| 引用发送失败 | 查看错误和日志；确认源媒体仍可访问后重新提交。此模式不会自动下载重试 |
| 服务器出现缓存文件 | 检查任务是否来自网站、使用其他存储，或选中了未开启 `reuse_media` 的 Telegram 存储 |

更多配置见 [配置说明](../deployment/configuration/_index.md)。反馈问题时，请在 [Issues](https://github.com/muheheheh/SaveAny-Bot/issues) 提供复现步骤和已隐藏 Token 等敏感内容的日志。
