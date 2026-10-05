# SaveAny-Bot（Telegram 快速回传分支）

[English](./README.md) | **简体中文**

基于 [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot)，在
[muheheheh/SaveAny-Bot](https://github.com/muheheheh/SaveAny-Bot) 独立维护。
新增 Telegram 媒体快速回传：复用 Telegram 已有的媒体引用，直接返回图片、视频、文档和相册，无需服务器下载或重新上传媒体文件。

## 本分支新增功能

- Telegram 存储支持 `reuse_media = true`，覆盖单文件和相册。
- 保留相册顺序及每条消息的说明文字，说明文字按纯文本发送。
- 快速回传不创建媒体缓存；引用失败直接报错，不自动下载重传。
- 忽略机器人自身发出的消息，避免回传消息再次被处理。

原项目的存储后端、网站解析、yt-dlp、Aria2、存储规则等功能仍可使用。
网站下载及保存到其他存储端仍走原有传输流程，可能使用临时缓存。
复用 Telegram 媒体时保留原文件名和媒体类型，不执行重命名、格式转换或分卷。
机器人需要能够读取源消息，且 Telegram 必须接受该媒体引用。

## 快速开始

从本仓库构建，确保包含快速回传功能：

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
cp deploy/vps/config.example.toml config.toml
chmod 600 config.toml
```

编辑 `config.toml`：填写 BotFather 提供的 Token，并把 `chat_id = 0` 和
`id = 0` 都改为你自己的 Telegram 数字用户 ID。模板已开启 `reuse_media`，且仅允许此用户访问。

```sh
docker compose up -d --build
```

在 Telegram 中启动机器人，用 `/storage` 选择回传存储，再用 `/silent` 开启自动处理。
随后发送 Telegram 消息链接或媒体消息即可。

小内存 VPS 可按 [VPS 部署说明](./deploy/vps/README.md) 在本地编译后上传。
上游镜像本身不包含本分支改动；VPS 模板仅使用固定版本的上游镜像提供运行依赖，并用本分支编译的程序替换其中的可执行文件。

## 文档与更新

- [安装与更新](./docs/content/zh/deployment/installation.md)
- [配置说明](./docs/content/zh/deployment/configuration/_index.md)
- [Telegram 存储与快速回传配置](./docs/content/zh/deployment/configuration/storages.md#telegram)
- [使用说明](./docs/content/zh/usage/_index.md)
- [在本分支提交问题](https://github.com/muheheheh/SaveAny-Bot/issues)

文档以本仓库内容为准。源码部署通过拉取本分支并重新构建来更新。
`/update` 和命令行更新器只检查 `muheheheh/SaveAny-Bot` 的发布版本；本仓库发布 Release 前，请从源码更新。

## 来源与许可证

本项目基于 [krau](https://github.com/krau) 和
[上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors) 开发的 SaveAny-Bot，
保留 [AGPL-3.0 许可证](./LICENSE) 及原有版权声明。
Go 模块名和导入路径保留 `github.com/krau/SaveAny-Bot` 以维持源码兼容，部署和更新来源为本分支。

感谢 [gotd](https://github.com/gotd/td)、[gotgproto](https://github.com/celestix/gotgproto)、
[TG-FileStreamBot](https://github.com/EverythingSuckz/TG-FileStreamBot)、
[tdl](https://github.com/iyear/tdl) 及所有上游依赖和贡献者。
