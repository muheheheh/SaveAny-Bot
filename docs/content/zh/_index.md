---
title: 介绍
---

# SaveAny-Bot（Telegram 快速回传分支）

[muheheheh/SaveAny-Bot](https://github.com/muheheheh/SaveAny-Bot) 基于原项目增加了 Telegram 媒体快速回传。

- Telegram 存储启用 `reuse_media = true` 后，单文件和相册直接复用媒体引用发送，不下载或缓存媒体文件。
- 相册顺序和原消息说明文字会保留，说明文字按纯文本发送。
- 引用发送失败时直接报错，不自动下载重传。
- 原有网站下载、存储后端、规则和解析器仍可使用，它们保持原有传输和缓存行为。

从 [安装与更新](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/zh/deployment/installation.md) 开始，或查看 [Telegram 存储配置](https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/content/zh/deployment/configuration/storages.md#telegram)。本分支文档以本仓库为准。

## 来源与许可证

基于 [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot)，感谢 [上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors)。保留 AGPL-3.0 许可证和原有版权声明。
