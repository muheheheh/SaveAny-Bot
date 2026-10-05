---
title: "使用帮助"
weight: 10
---

# 使用帮助

这里介绍 Save Any Bot 的一些功能和使用方法, 如果你没有在这里找到你需要的内容, 另请参阅 [配置说明](../deployment/configuration) 或前往 Github [Issues](https://github.com/muheheheh/SaveAny-Bot/issues) 提问.

## 转存文件

要使用 Bot 的转存 Telegram 文件功能, 需要向 Bot 发送或转发以下类型的消息.

1. 文件或媒体消息, 如图片, 视频, 文档等
2. Telegram 消息链接, 例如: `https://t.me/channel_username/123`. 启用 Telegram 存储的 `reuse_media` 后可直接回传媒体；具体可用性取决于源消息访问权限和 Telegram 对引用的接受情况。
3. Telegra.ph 的文章链接, Bot 将下载其中的所有图片