---
title: "Usage"
weight: 10
---

# Usage

SaveAny-Bot can automatically relay Telegram media to the current chat or save Telegram and website files to other storage. The [repository homepage](https://github.com/muheheheh/SaveAny-Bot/blob/main/README.md) provides the complete feature overview.

## First use

1. Use `config.docker.example.toml` and fill in the bot token and the authorized user's numeric Telegram ID.
2. Start the bot following [Docker deployment](../deployment/installation.md).
3. Open a private chat, press **Start** or send `/start`, then send or forward media.

The template has no storage configured. The bot only relays to the current chat, without requiring a default storage or silent mode.

## Telegram media relay

| Input | How to use it |
| --- | --- |
| A photo, video, or document | Send or forward it to the bot |
| An album or media group | Send or forward the group together, preserving its grouping where possible |
| A Telegram message link | Send a link such as `https://t.me/channel_username/123` to read its media |

Relay is always enabled and targets the chat that initiated the request. Album order and each source caption are preserved; captions are plain text.

Relay alone does not download, cache, or upload media again. Single files and albums create no progress messages. Message links show a temporary status, deleted when all media is successfully relayed. Errors or partial failures remain visible. Original messages and relayed media are retained.

Message links require source access. UserBot needs a separate user login and that account must also have access. Relay does not bypass Telegram permissions or content protection.

## Saving and task management

| Command | Purpose |
| --- | --- |
| `/storage` | Select the default storage |
| `/silent` | Toggle automatic saving |
| `/task` | List running tasks |
| `/task queued` | List queued tasks |
| `/cancel <task-ID>` | Cancel a task; progress messages also provide a cancel button |
| `/help` | Show bot help |

When the user has available storage, the bot relays Telegram media first, then asks where to save. Select a default with `/storage` and enable `/silent` to save automatically using the default and configured rules. `/silent` is a toggle; sending it again disables automatic saving. These preferences are stored per user.

Download and save tasks keep their progress and results. With no available storage, processing ends after relay; explicit commands such as `/save` show a no-storage error. Website links use their corresponding parser or downloader.

`/task` and `/cancel` manage queued or running download and save tasks. Direct relay does not enter the task queue.

## Other supported features

| Scenario | Command or requirement | Details |
| --- | --- | --- |
| Existing media or historical messages | Reply with `/save`, or use `/save <chat-ID-or-username> <first-message-ID>-<last-message-ID> [text-regex]` | Saves to configured storage without triggering another relay |
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

- Relay supports directly sent or forwarded Telegram media and Telegram media message links.
- Saving alongside relay, explicit save commands, website downloads, and chat monitoring retain their normal transfer behavior and may use temporary server files.
- Relay keeps the original media type and filename. Forced document sending, renaming, transcoding, and splitting do not apply; caption formatting is not reconstructed.
- Unavailable or rejected media references produce relay errors without a download fallback. Configured saving proceeds independently.
- Plain-text chat archives, polls, and other non-media content are outside the media relay scope.
- Configuration, sessions, user settings, and logs remain server-side runtime data.

## Troubleshooting

| Symptom | What to check |
| --- | --- |
| Permission denied | Ensure `[[users]].id` is the sender's numeric user ID; restart after editing configuration |
| Save command reports no available storage | Ensure storage is enabled and allowed by the user's filters; relay alone needs no storage |
| Destination prompt after every relay | For automatic saving, set `/storage` and `/silent`. For relay only, use `storages = []` and `blacklist = false` for the user |
| Cannot read a message link | Check the link and source access; a message visible to your user account may not be visible to the bot |
| Reference sending fails | Inspect the error and logs, confirm source access, then resubmit. Relay does not fall back to downloading |
| Files appear in the cache | Check for configured storage, explicit save commands, website downloads, or chat monitoring; relay alone creates no media cache |

See [configuration](../deployment/configuration/_index.md) for more options. Report reproducible problems in [Issues](https://github.com/muheheheh/SaveAny-Bot/issues), removing tokens and other secrets from logs first.
