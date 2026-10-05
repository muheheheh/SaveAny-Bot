---
title: 介绍
---

# SaveAny-Bot

SaveAny-Bot 支持 Telegram 媒体回传、网站下载和多存储转存。

- 私聊媒体和 Telegram 消息链接默认回传到当前聊天；只回传时不下载或缓存媒体文件。
- 用户有可用存储时，回传后继续保存；没有可用存储时，成功后不保留进度提示。
- 相册顺序和原消息说明文字会保留，说明文字按纯文本发送。
- 引用发送失败时保留错误，不自动下载重传。
- 网站下载及其他存储操作可能使用临时文件。

## 文档

- [支持的功能](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md#支持的功能)
- [使用说明](./usage/_index.md)
- [Docker 部署](./deployment/installation.md)
- [配置说明](./deployment/configuration/_index.md)
- [存储配置](./deployment/configuration/storages.md)
- [HTTP API](./usage/api.md) 与 [命令行使用](./usage/cli.md)

## 来源与许可证

基于 [krau/SaveAny-Bot](https://github.com/krau/SaveAny-Bot)，感谢 [上游贡献者](https://github.com/krau/SaveAny-Bot/graphs/contributors)。保留 AGPL-3.0 许可证和原有版权声明。
