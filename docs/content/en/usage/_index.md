---
title: "Usage"
weight: 10
---

# Usage

## First use

1. Use `config.docker.example.toml` and fill in the bot token and the authorized user's numeric Telegram ID.
2. Start the bot following [Docker deployment](../deployment/installation.md).
3. Open a private chat, press **Start** or send `/start`, then send or forward media.

## Telegram media relay

Send any of the following in a private chat. The bot will send the photos, videos, or files back to you.

| Input | How to use it |
| --- | --- |
| A photo, video, or document | Send or forward it to the bot |
| An album | Send or forward the group of photos or videos together |
| A Telegram message link | Send a link such as `https://t.me/channel_username/123` to read its media |

## Saving and task management

Add a destination using the [storage configuration](../deployment/configuration/_index.md#file-storage). After sending a file, select where to save it.

Set a default with `/storage`, then send `/silent` to enable automatic saving. Send `/silent` again to disable it.

| Command | Purpose |
| --- | --- |
| `/storage` | Select the default storage |
| `/silent` | Toggle automatic saving |
| `/task` | List running tasks |
| `/task queued` | List queued tasks |
| `/cancel <task-ID>` | Cancel a task; progress messages also provide a cancel button |
| `/help` | Show bot help |

Reply to a media message with `/save` to save its file. To save a range of messages, use:

```text
/save <chat-ID-or-username> <first-message-ID>-<last-message-ID> [text-regex]
```

## Downloads and file management

| Scenario | Command or requirement | Details |
| --- | --- | --- |
| HTTP/HTTPS file links | `/dl <URL1> [URL2]` | [Direct links](./directlinks.md) |
| Website images and media | Send Telegraph articles or URLs supported by enabled parsers, including Twitter/X and Kemono | [Parsers](./parsers.md) |
| Website video and audio | `/ytdlp <URL> [flags]`; requires yt-dlp and FFmpeg for some processing | [yt-dlp](./ytdlp.md) |
| Aria2 downloads | `/aria2dl <URL>`; requires enabled and configured Aria2 RPC | [Aria2](./aria2.md) |
| Storage and directory rules | Use `/rule` for rules and `/dir` for directories | [Rules](./rules.md) |
| Naming and duplicate handling | `/config`, `/fnametmpl` | [Naming](./config.md) |
| Cross-storage transfer | `/transfer`; the source must support listing and reading files | [Transfer](./transfer.md) |
| Chat monitoring | `/watch`, `/unwatch`, `/lswatch`; requires UserBot | [Watch and notifications](./watch.md) |
| Local uploads and directory watching | Run `saveany-bot upload` or `saveany-bot watch` on the server | [CLI](./cli.md) |
| Programmatic tasks and callbacks | Configure the HTTP API to create, query, and cancel tasks and receive webhooks | [HTTP API](./api.md) |
| Additional websites | Enable and load JavaScript parser plugins | [Plugin development](https://github.com/muheheheh/SaveAny-Bot/blob/main/plugins/README.md) |

## Troubleshooting

| Symptom | What to check |
| --- | --- |
| Permission denied | Ensure `[[users]].id` is the sender's numeric user ID; restart after editing configuration |
| No available storage | Check the storage's `enable` setting and the user's `storages` list |
| Cannot read a message link | Check that the link works and that the bot or UserBot account can access the source message |

See [configuration](../deployment/configuration/_index.md) for more options. Report reproducible problems in [Issues](https://github.com/muheheheh/SaveAny-Bot/issues), removing tokens and other secrets from logs first.
