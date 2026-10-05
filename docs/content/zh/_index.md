---
title: 介绍
---

# SaveAny-Bot

SaveAny-Bot 支持 Telegram 媒体回传、网站下载和多存储转存。

- Telegram 存储启用 `reuse_media = true` 后，单文件和相册直接复用媒体引用发送，不下载或缓存媒体文件。
- 相册顺序和原消息说明文字会保留，说明文字按纯文本发送。
- 引用发送失败时直接报错，不自动下载重传。
- 网站下载及其他存储操作可能使用临时文件。

## 文档

- [支持的功能与快速开始](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md)
- [使用说明](./usage/_index.md)
- [安装与更新](./deployment/installation.md)
- [VPS 部署](https://github.com/muheheheh/SaveAny-Bot/blob/main/deploy/vps/README.md)
- [配置说明](./deployment/configuration/_index.md)
- [存储配置](./deployment/configuration/storages.md)
- [HTTP API](./usage/api.md) 与 [命令行使用](./usage/cli.md)

## 来源与许可证

基于 [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot)，感谢 [上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors)。保留 AGPL-3.0 许可证和原有版权声明。
