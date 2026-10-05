---
title: "使用说明"
weight: 10
---

# 使用说明

## 首次使用

1. 使用 `config.docker.example.toml` 模板，在配置中填写 Bot Token 和允许使用机器人的 Telegram 数字用户 ID。
2. 按 [Docker 部署](../deployment/installation.md) 启动机器人。
3. 打开机器人私聊，点击 **Start** 或发送 `/start`，然后发送或转发媒体。

## Telegram 媒体回传

在私聊中发送以下内容，机器人会把其中的图片、视频或文件发回给你。

| 输入方式 | 操作方法 |
| --- | --- |
| 单张图片、视频或文档 | 直接发送或转发给机器人 |
| 相册 | 将整组图片或视频发送或转发给机器人 |
| Telegram 消息链接 | 发送类似 `https://t.me/channel_username/123` 的链接，机器人读取其中的媒体 |

## 保存与任务管理

先按 [存储配置](../deployment/configuration/_index.md#文件保存) 添加保存位置。发送文件后，选择要保存到的存储。

使用 `/storage` 设置默认存储，再发送 `/silent` 开启自动保存；再次发送 `/silent` 可关闭。

| 命令 | 作用 |
| --- | --- |
| `/storage` | 设置默认存储 |
| `/silent` | 开启或关闭自动保存 |
| `/task` | 列出正在执行的任务 |
| `/task queued` | 列出排队中的任务 |
| `/cancel <任务ID>` | 取消指定任务，也可点击进度消息上的取消按钮 |
| `/help` | 查看机器人帮助 |

回复一条媒体消息并发送 `/save`，可以保存这条消息中的文件。批量保存历史消息使用：

```text
/save <聊天ID或用户名> <起始消息ID>-<结束消息ID> [文本正则]
```

## 下载和文件管理

| 场景 | 操作或前提 | 详细说明 |
| --- | --- | --- |
| HTTP/HTTPS 文件直链 | `/dl <链接1> [链接2]` | [直接下载链接](./directlinks.md) |
| 网站图片和媒体 | 发送 Telegraph 文章链接，或发送 Twitter/X、Kemono 等已启用解析器支持的链接 | [网站解析](./parsers.md) |
| 网站视频与音频 | `/ytdlp <链接> [参数]`，需要安装 yt-dlp，部分处理需要 FFmpeg | [yt-dlp 使用](./ytdlp.md) |
| Aria2 下载 | `/aria2dl <链接>`，需要启用并配置 Aria2 RPC | [Aria2 使用](./aria2.md) |
| 自动选择存储或目录 | 使用 `/rule` 配置规则，使用 `/dir` 管理目录 | [存储规则](./rules.md) |
| 文件命名与重名处理 | `/config`、`/fnametmpl` | [文件命名](./config.md) |
| 存储间传输 | `/transfer`，源存储需支持文件列举和读取 | [存储间传输](./transfer.md) |
| 监听聊天 | `/watch`、`/unwatch`、`/lswatch`，需要启用 UserBot | [监听与通知](./watch.md) |
| 本地文件上传与目录监听 | 在服务器运行 `saveany-bot upload` 或 `saveany-bot watch` | [命令行使用](./cli.md) |
| 程序调用与回调 | 配置 HTTP API，可创建、查询、取消任务并接收 Webhook | [HTTP API](./api.md) |
| 扩展网站支持 | 启用并加载 JavaScript 解析器插件 | [插件开发](https://github.com/muheheheh/SaveAny-Bot/blob/main/plugins/README.md) |

## 常见问题

| 现象 | 检查方法 |
| --- | --- |
| 提示没有权限 | 检查 `[[users]].id` 是否为发消息的用户数字 ID；修改配置后重启程序 |
| 提示没有可用存储 | 检查存储的 `enable` 和用户的 `storages` 配置 |
| 无法读取消息链接 | 检查链接是否有效，以及机器人或 UserBot 账号能否访问源消息 |

更多配置见 [配置说明](../deployment/configuration/_index.md)。反馈问题时，请在 [Issues](https://github.com/muheheheh/SaveAny-Bot/issues) 提供复现步骤和已隐藏 Token 等敏感内容的日志。
