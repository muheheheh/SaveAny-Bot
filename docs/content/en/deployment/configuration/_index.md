---
title: "Configuration Guide"
---

# Configuration Guide

SaveAny-Bot uses a UTF-8 TOML configuration file. For Docker deployment, create `config.toml` in the repository root; the container reads it at `/app/config.toml` through the bind mount.

## Configuration structure

| Location | Contents |
| --- | --- |
| Global fields at the beginning of the file | `lang`, `workers`, `threads`, `retry`, `stream`, `proxy`, `no_clean_cache` |
| `[telegram]` | Bot token, Telegram API credentials, and connection settings |
| `[telegram.proxy]` | Proxy used for Telegram connections |
| `[telegram.userbot]` | User account login and session settings |
| `[log]` | Log level |
| `[[storages]]` | One storage or relay destination; repeat for additional destinations |
| `[[users]]` | One authorized user and their available storages; repeat for additional users |
| `[api]` | HTTP API switch, listen address, and authentication token |
| `[aria2]` | Aria2 RPC connection and downloaded file retention |
| `[ytdlp]` | Video quality, format, and filename settings |
| `[parser]` | JavaScript parser plugins and parser proxy settings |
| `[parser.twitter]`, `[parser.kemono]` | Settings for the corresponding built-in parsers |
| `[temp]` | Temporary download directory |
| `[db]` | User settings database and bot session file paths |
| `[cache]` | In-memory cache settings |
| `[hook.exec]` | Commands run when tasks start, succeed, fail, or are cancelled |

- `[name]` defines a configuration table; `[telegram.proxy]` defines the `proxy` subtable of `telegram`.
- `[[name]]` adds an entry to an array of tables. Repeat `[[storages]]` or `[[users]]` for each additional storage or user.
- Fields belong to the most recent table header. Blank lines and comments do not end a table. Put global fields before the first table header.
- Quote strings; leave numbers and `true` / `false` unquoted. Text after `#` is a comment. See the [TOML documentation](https://toml.io/) for more syntax.

## Telegram relay configuration

Set `reuse_media` inside each `[[storages]]` table whose `type` is `"telegram"` in `config.toml`, at the same level as `chat_id`. Set it to `true` to enable direct relay. Configure it separately for each relay storage.

Save the following complete relay configuration as `config.toml`. Fill in the token and replace both instances of `123456789` with your numeric Telegram user ID. The repository's [config.docker.example.toml](https://github.com/muheheheh/SaveAny-Bot/blob/main/config.docker.example.toml) provides a template with the same structure.

```toml
lang = "en"
workers = 1
threads = 2
stream = false

[log]
level = "info"

[telegram]
token = "YOUR_BOTFATHER_TOKEN"

[api]
enable = false

[[storages]]
name = "Relay to chat"
type = "telegram"
enable = true
chat_id = 123456789
force_file = false
reuse_media = true

[[users]]
id = 123456789
storages = ["Relay to chat"]
blacklist = false
```

| Setting | Type | Purpose |
| --- | --- | --- |
| `lang` | String | Bot message language; `"en"` selects English |
| `workers` | Integer | Number of tasks executed concurrently |
| `threads` | Integer | Threads used for regular downloads; direct relay does not download media |
| `stream` | Boolean | Streaming for regular downloads; `reuse_media` controls direct relay |
| `level` in `[log]` | String | Log level; this example uses `"info"` |
| `token` in `[telegram]` | String | Bot token provided by BotFather |
| `enable` in `[api]` | Boolean | Enables the HTTP API; leave `false` when using the bot through Telegram |
| `name` in `[[storages]]` | String | Storage name, unique across all storages |
| `type` in `[[storages]]` | String | Use `"telegram"` to relay to Telegram |
| `enable` in `[[storages]]` | Boolean | Set to `true` to enable this storage |
| `chat_id` in `[[storages]]` | Integer | Fixed relay destination; use your user ID to relay to your private chat |
| `reuse_media` in `[[storages]]` | Boolean | Set to `true` to send existing Telegram media references directly |
| `force_file` in `[[storages]]` | Boolean | Forces regular uploads to be sent as files; direct relay preserves the original media type |
| `id` in `[[users]]` | Integer | Telegram user ID authorized to operate the bot |
| `storages` in `[[users]]` | List of strings | References the storage names defined above; names must match |
| `blacklist` in `[[users]]` | Boolean | `false` allows only listed storages; `true` excludes listed storages |

`chat_id` determines who receives media; `users.id` determines who can use the bot. Use the same ID for both when relaying to yourself. Each storage has a fixed destination. To relay separately for multiple users, define a storage for each user and reference it in that user's `storages` list.

For example, append the following to add user `987654321`:

```toml
[[storages]]
name = "Second user"
type = "telegram"
enable = true
chat_id = 987654321
reuse_media = true

[[users]]
id = 987654321
storages = ["Second user"]
blacklist = false
```

After startup, each user selects their default storage with `/storage` and enables automatic processing with `/silent`. These preferences are stored in the database. A destination `chat_id` can also identify a group or channel where the bot has permission to send messages. See [Telegram storage configuration](./storages.md#telegram) for related settings.

Unused optional configuration tables can be omitted. Run `docker compose restart` after editing `config.toml` to apply the changes.

## Detailed Configuration

### Global Configuration

- `lang`: The language used by the Bot, default is `zh-Hans` (Simplified Chinese). `en` is used for English.
- `stream`: Whether to enable Stream mode, default is `false`. When enabled, the Bot will stream files directly to storage endpoints (if supported), without downloading them locally.
{{< hint warning >}}
Stream mode is very useful for deployment environments with limited disk space, but it also has some drawbacks:
<br />
<ul>
<li>Cannot use multi-threading to download files from Telegram, resulting in slower speeds.</li>
<li>Higher task failure rate when the network is unstable.</li>
<li>Cannot process files in the middle layer, such as automatic file type identification.</li>
<li>Not supported by all storage endpoints; unsupported endpoints may downgrade to normal mode or fail to upload.</li>
</ul>
{{< /hint >}}
- `workers`: Number of tasks to process simultaneously, default is 3.
- `threads`: Number of threads used when downloading files, default is 4. Only effective when Stream mode is not enabled.
- `retry`: Number of retries when a task fails, default is 3.
- `proxy`: Global proxy configuration. After setting this, all network connections inside the program will try to use this proxy. Optional.

```toml
lang = "en"
stream = false
workers = 3
threads = 4
retry = 3
proxy = "socks5://127.0.0.1:7890"
```

### Telegram Configuration

- `token`: Your Telegram Bot Token, which can be obtained by creating a Bot through [BotFather](https://t.me/botfather).
- `app_id`, `app_hash`: Telegram API ID & Hash, obtained by creating an application at [Telegram API](https://my.telegram.org/apps). Default values will be used if not provided.
- `rpc_retry`: Number of retries for RPC requests, default is 5.
- `proxy`: Proxy configuration, optional.
  - `enable`: Whether to enable the proxy.
  - `url`: Proxy address
- `userbot`: Userbot configuration, optional.
  - `enable`: Enable userbot integration. Requires logging in with a user account; you should use your own API ID & Hash when enabling this.
  - `session`: Path to the userbot session file, default is `data/usersession.db`.

{{< hint warning >}}
After enabling userbot integration, the bot can download files from private channels and groups, but there is an unavoidable risk of the account being banned.
<br />
On the first start after enabling userbot, you need to input phone number, 2FA and verification code in the terminal.
<br />
If you deploy with Docker, please run the container with `-it` for an interactive environment, then perform the login.
{{< /hint >}}

```toml
[telegram]
token = "1234567890:ABCDEFGHIJKLMNOPQRSTUVWXYZ"
app_id = 1025907
app_hash = "452b0359b988148995f22ff0f4229750"
rpc_retry = 5
[telegram.proxy]
enable = false
url = "socks5://127.0.0.1:7890"
[telegram.userbot]
enable = false
session = "data/usersession.db"
```

### Aria2 Configuration

Aria2 is a powerful download manager that supports HTTP/HTTPS, FTP, BitTorrent, and other protocols. When enabled, the bot can use the `/aria2dl` command to download files via Aria2.

- `enable`: Whether to enable Aria2 support, default is `false`
- `url`: Aria2 RPC address, typically `http://localhost:6800/jsonrpc`
- `secret`: Aria2 RPC secret, if you configured `rpc-secret` in Aria2, you need to fill it in here
- `keep_file`: Whether to keep local files downloaded by Aria2 after transfer, default is `false`

{{< hint info >}}
Aria2 needs to be installed and running separately. You can refer to the [Aria2 official documentation](https://aria2.github.io/) to learn how to install and configure Aria2.
{{< /hint >}}

```toml
[aria2]
enable = true
url = "http://localhost:6800/jsonrpc"
secret = "your-rpc-secret"
keep_file = false
```

### yt-dlp Configuration

Configures the behavior of the `/ytdlp` command and the `ytdlp` HTTP-API task type when no custom flags are passed.

- `max_height`: Default maximum video resolution by height in pixels (e.g. `1080`, `720`). `0` means no limit (best available). Ignored when `format` is set.
- `format`: A raw yt-dlp format selector (`-f`). When set, it takes precedence over `max_height` and gives you full control, e.g. `bv*[height<=720]+ba/b`.
- `recode`: The target video container yt-dlp recodes into after download (e.g. `mp4`). Leave empty to disable recoding.
- `filename_template`: yt-dlp output template for downloaded file names, default `%(title)s.%(ext)s`. Directories in the template are created below the bot's download directory and kept in the storage path.
- `restrict_filenames`: Restrict file names to ASCII characters (yt-dlp `--restrict-filenames`), which also drops spaces and `&`. Disabled by default.

{{< hint info >}}
`max_height`, `format` and `recode` only apply when using the `/ytdlp` command (or the `ytdlp` API task type) without passing any custom flags. `filename_template` and `restrict_filenames` also apply when other flags are passed, unless you pass the corresponding flag yourself (`-o/--output`, `--restrict-filenames` or `--no-restrict-filenames`).
{{< /hint >}}

```toml
[ytdlp]
max_height = 1080
format = ""        # e.g. "bv*[height<=720]+ba/b"
recode = "mp4"     # empty disables recoding
filename_template = "%(title)s.%(ext)s"  # e.g. "%(uploader)s - %(title)s.%(ext)s"
restrict_filenames = false
```

### HTTP API Configuration

When enabled, SaveAny-Bot exposes an HTTP API for creating/querying/canceling tasks programmatically. See [HTTP API](../../usage/api) for the full endpoint reference.

- `enable`: Whether to enable the HTTP API server, default is `false`.
- `host`: Bind address, default `0.0.0.0`.
- `port`: Listen port, default `8080`.
- `token`: Authentication token, required when the API is enabled.

```toml
[api]
enable = false
host = "0.0.0.0"
port = 8080
token = "your-token"
```

### Log Configuration

- `level`: Log level. One of `debug`, `info`, `warn`, `error`, `fatal`. Default is `debug`; the relay template uses `info`.

```toml
[log]
level = "info"
```

### Storage Endpoints List

The storage endpoints list is used to define the storage locations supported by the Bot. Each storage endpoint needs to specify a name, type, and related configuration, using the double bracket syntax `[[storages]]`.

Each storage endpoint requires at least the following fields:

- `name`: Storage endpoint name, used for identification in the Bot, must be unique.
- `enable`: Set to `true` to load this storage endpoint.
- `type`: Storage endpoint type, currently supports the following types:
  - `local`: Local disk
  - `alist`: Alist
  - `webdav`: WebDAV
  - `s3`: aws S3 and other S3 compatible services
  - `minio`: MinIO object storage
  - `rclone`: Uses rclone to implement uploads
  - `telegram`: Upload to Telegram

Example, this is a configuration that includes local storage and webdav storage:

```toml
[[storages]]
name = "Local Storage"
type = "local"
enable = true
# Custom configuration for local type storage
base_path = "./downloads"

[[storages]]
name = "WebDAV"
type = "webdav"
enable = true
# Custom configuration for webdav type storage
url = "https://example.com/webdav"
base_path = "/path/to/webdav"
username = "your_username"
password = "your_password"
```

For custom configuration items for all storage endpoints, see [Storage Configuration](./storages)

### User List

The user list is used to define access control for storage endpoints. Each user needs to specify a Telegram User ID, defined using the double bracket syntax `[[users]]`.

- `id`: The user's Telegram User ID
- `storages`: Filtered list of storage endpoints, defined by storage endpoint names, default is whitelist mode (i.e., only allows access to storage endpoints in the list)
- `blacklist`: Whether to enable blacklist mode, default is `false`. If blacklist mode is enabled, the user is allowed to access only storage endpoints that are **not** in the list.

Example, this is a configuration containing three users: user `123123` can only access local storage, user `456456` can only access storage other than WebDAV, and user `789789` has blacklist mode enabled but no storage endpoints specified, so they can access all storage:

```toml
[[users]]
id = 123123
storages = ["Local Storage"]

[[users]]
id = 456456
storages = ["WebDAV"]
blacklist = true

[[users]]
id = 789789
storages = []
blacklist = true
```

### Events

Event hooks allow you to run custom commands based on task status while the bot is processing tasks. Currently only arbitrary command execution is supported, configured via `[hook.exec]`.

Supported event types:

- `task_before_start`: Before a task starts
- `task_success`: After a task completes successfully
- `task_fail`: After a task fails
- `task_cancel`: After a task is cancelled

The configured value must be a full shell command line. The bot will execute this command when the event occurs. Example:

```toml
[hook.exec]
task_before_start = "echo 'task is about to start'"
task_success = "bash /path/to/success_script.sh"
task_fail = "curl -X POST https://example.com/api/notify -d 'task failed'"
task_cancel = "bash /path/to/cancel_script.sh"
```

### Parsers

Parsers give the bot the ability to handle non-Telegram files, such as downloading files from other websites. Configure them via `[parser]`.

```toml
[parser]
plugin_enable = true # Whether to enable parser plugins
plugin_dirs = ["./plugins"] # Plugin directories, can be multiple
```

The above settings only control JavaScript-based parser plugins. The bot also has built-in parsers implemented in Go, which are enabled by default.

### Miscellaneous

```toml
no_clean_cache = false # Whether not to clear the cache folder when exiting
# Temporary download folder configuration
[temp]
base_path = "./cache"
```
