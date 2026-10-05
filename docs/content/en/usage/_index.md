---
title: "Usage"
weight: 10
---

# Usage

This page introduces some of Save Any Bot's features and basic usage. If you can't find what you need here, please also see the [Configuration Guide](../deployment/configuration) or ask in GitHub [Issues](https://github.com/muheheheh/SaveAny-Bot/issues).

## File Transfer

To use the bot's Telegram file saving feature, you need to send or forward the following types of messages to the bot:

1. File or media messages, such as images, videos, documents, etc.
2. Telegram message links, for example: `https://t.me/channel_username/123`. Enable `reuse_media` on Telegram storage to relay media directly; availability depends on access to the source and Telegram accepting its reference.
3. Telegra.ph article links. The bot will download all images in the article.