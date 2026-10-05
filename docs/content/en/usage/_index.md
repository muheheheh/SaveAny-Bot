---
title: "Usage"
weight: 10
---

# Usage

SaveAny-Bot can relay Telegram media to a configured chat or save Telegram and website files to other storage. The [repository homepage](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md) provides the complete feature overview.

## First use

1. Follow [Docker deployment](../deployment/installation.md) using the `config.docker.example.toml` template.
2. Set the bot token, use your numeric user ID for both the Telegram storage's `chat_id` and the allowed user's `id`, and enable `reuse_media = true` on that storage.
3. Open a private chat with the bot and press **Start** or send `/start`.
4. Send `/storage` and select the relay storage named “回传到聊天” as the default.
5. Send `/silent`, check that silent mode is enabled, then send or forward media.

The default storage and silent mode are saved per user. `/silent` is a toggle: sending it again disables automatic processing. With silent mode disabled, the bot asks you to select a destination for each operation.

## Telegram media relay

| Input | How to use it |
| --- | --- |
| A photo, video, or document | Send or forward it to the bot |
| An album or media group | Send or forward the group together, preserving its grouping where possible |
| A Telegram message link | Send a link such as `https://t.me/channel_username/123` to read its media |
| An existing media message | Reply to the message with `/save` |
| A range of messages | Use `/save <chat-ID-or-username> <first-message-ID>-<last-message-ID> [text-regex]` |

When the destination is Telegram storage with `reuse_media` enabled, these Telegram media tasks send existing references. Album order and each source caption are retained; captions are plain text. The storage configuration determines the destination. In the personal relay template, it is your private chat with the bot.

Message links, historical ranges, and private chats require access to the source. UserBot requires a separate user login and that account must also have access. Relay does not bypass Telegram permissions or content protection.

## Automatic processing and task management

| Command | Purpose |
| --- | --- |
| `/storage` | Select the default storage |
| `/silent` | Toggle automatic processing |
| `/task` | List running tasks |
| `/task queued` | List queued tasks |
| `/cancel <task-ID>` | Cancel a task; progress messages also provide a cancel button |
| `/help` | Show bot help |

With silent mode enabled, supported media and links use the default storage and configured rules. Website links use the corresponding parser or downloader.

## Other supported features

| Scenario | Command or requirement | Details |
| --- | --- | --- |
| Other storage | Configure Local, S3, MinIO, WebDAV, AList, Rclone, or Telegram storage | [Storage configuration](../deployment/configuration/storages.md) |
| HTTP/HTTPS file links | `/dl <URL1> [URL2]` | [Direct links](./directlinks.md) |
| Website images and media | Send Telegraph articles or URLs supported by enabled parsers, including Twitter/X and Kemono | [Parsers](./parsers.md) |
| Website video and audio | `/ytdlp <URL> [flags]`; requires yt-dlp and FFmpeg for some processing | [yt-dlp](./ytdlp.md) |
| Aria2 downloads | `/aria2dl <URL>`; requires enabled and configured Aria2 RPC | [Aria2](./aria2.md) |
| Storage and directory rules | Use `/rule` for rules and `/dir` for directories | [Rules](./rules.md) |
| Naming and duplicate handling | `/config`, `/fnametmpl`; relay keeps the original media filename | [Naming](./config.md) |
| Cross-storage transfer | `/transfer`; the source must support listing and reading files | [Transfer](./transfer.md) |
| Chat monitoring | `/watch`, `/unwatch`, `/lswatch`; requires UserBot | [Watch and notifications](./watch.md) |
| Local uploads and directory watching | Run `saveany-bot upload` or `saveany-bot watch` on the server | [CLI](./cli.md) |
| Programmatic tasks and callbacks | Configure the HTTP API to create, query, and cancel tasks and receive webhooks | [HTTP API](./api.md) |
| Additional websites | Enable and load JavaScript parser plugins | [Plugin development](https://github.com/muheheheh/SaveAny-Bot/blob/main/plugins/README.md) |

## Relay scope

- Download-free, cache-free media relay applies only to Telegram sources sent to Telegram storage with `reuse_media` enabled.
- Website downloads and other storage paths retain their normal transfer behavior and may use temporary server files.
- Relay keeps the original media type and filename. Forced document sending, renaming, transcoding, and splitting do not apply; caption formatting is not reconstructed.
- Unavailable or rejected media references fail the task without a download fallback.
- Plain-text chat archives, polls, and other non-media content are outside the media relay scope.
- Configuration, sessions, user settings, and logs remain server-side runtime data.

## Troubleshooting

| Symptom | What to check |
| --- | --- |
| Permission denied | Ensure `[[users]].id` is the sender's numeric user ID; restart after editing configuration |
| No available storage | Ensure the storage is enabled and allowed by the user's storage filters |
| Repeated destination prompts | Set a default with `/storage`, then enable `/silent` |
| Cannot read a message link | Check the link and source access; a message visible to your user account may not be visible to the bot |
| Reference sending fails | Inspect the error and logs, confirm source access, then resubmit. Relay does not fall back to downloading |
| Files appear in the cache | Check for website tasks, other storage, or Telegram storage without `reuse_media` enabled |

See [configuration](../deployment/configuration/_index.md) for more options. Report reproducible problems in [Issues](https://github.com/muheheheh/SaveAny-Bot/issues), removing tokens and other secrets from logs first.
