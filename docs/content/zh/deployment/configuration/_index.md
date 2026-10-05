---
title: "配置说明"
---

# 配置说明

SaveAny-Bot 使用 UTF-8 编码的 TOML 配置文件。Docker 部署时，在项目根目录创建 `config.toml`，容器会读取挂载后的 `/app/config.toml`。

## 配置文件结构

| 位置 | 内容 |
| --- | --- |
| 文件开头的全局字段 | `lang`、`workers`、`threads`、`retry`、`stream`、`proxy`、`no_clean_cache` |
| `[telegram]` | Bot Token、Telegram API 凭据和连接参数 |
| `[telegram.proxy]` | Telegram 连接使用的代理 |
| `[telegram.userbot]` | 用户账号登录及会话配置 |
| `[log]` | 日志级别 |
| `[[storages]]` | 一个存储或回传目标，可重复定义多个 |
| `[[users]]` | 一个授权用户及其可用存储，可重复定义多个 |
| `[api]` | HTTP API 开关、监听地址和认证 Token |
| `[aria2]` | Aria2 RPC 连接及下载文件保留设置 |
| `[ytdlp]` | 视频下载的清晰度、格式和文件名设置 |
| `[parser]` | JavaScript 解析器插件及解析器代理设置 |
| `[parser.twitter]`、`[parser.kemono]` | 对应内置解析器的设置 |
| `[temp]` | 下载临时目录 |
| `[db]` | 用户设置数据库和机器人会话文件路径 |
| `[cache]` | 内存缓存参数 |
| `[hook.exec]` | 任务开始、成功、失败或取消时执行的命令 |

- `[名称]` 定义一组配置；`[telegram.proxy]` 表示 `telegram` 下的 `proxy` 子配置。
- `[[名称]]` 定义列表中的一项。每增加一个存储或用户，都要再写一组 `[[storages]]` 或 `[[users]]`。
- 字段属于它上方最近的配置组，空行和注释不会结束配置组。全局字段应写在第一个配置组之前。
- 字符串加引号，数字和 `true` / `false` 不加引号；`#` 后是注释。更多语法见 [TOML 文档](https://toml.io/)。

## Telegram 回传配置

`reuse_media` 是每个 Telegram 存储的配置项，写在 `config.toml` 中 `type = "telegram"` 的 `[[storages]]` 组内，与 `chat_id` 同级。设为 `true` 开启直接回传；多个回传存储需要分别设置。

下面是可直接保存为 `config.toml` 的完整回传配置。填写 Token，并把两处 `123456789` 替换为自己的 Telegram 数字用户 ID。仓库中的 [config.docker.example.toml](https://github.com/muheheheh/SaveAny-Bot/blob/main/config.docker.example.toml) 提供同一结构的模板。

```toml
lang = "zh-Hans"
workers = 1
threads = 2
stream = false

[log]
level = "info"

[telegram]
token = "填入 BotFather 提供的 Token"

[api]
enable = false

[[storages]]
name = "回传到聊天"
type = "telegram"
enable = true
chat_id = 123456789
force_file = false
reuse_media = true

[[users]]
id = 123456789
storages = ["回传到聊天"]
blacklist = false
```

| 配置项 | 类型 | 作用 |
| --- | --- | --- |
| `lang` | 字符串 | 机器人消息语言，`"zh-Hans"` 为简体中文 |
| `workers` | 整数 | 同时执行的任务数量 |
| `threads` | 整数 | 普通下载任务的线程数；直接回传不下载媒体 |
| `stream` | 布尔值 | 普通下载任务是否使用流式传输；回传由 `reuse_media` 控制 |
| `[log]` 中的 `level` | 字符串 | 日志级别，示例使用 `"info"` |
| `[telegram]` 中的 `token` | 字符串 | BotFather 提供的机器人 Token |
| `[api]` 中的 `enable` | 布尔值 | 是否启用 HTTP API；通过 Telegram 使用机器人时可保持 `false` |
| `[[storages]]` 中的 `name` | 字符串 | 存储名称，在所有存储中唯一 |
| `[[storages]]` 中的 `type` | 字符串 | 回传到 Telegram 时填写 `"telegram"` |
| `[[storages]]` 中的 `enable` | 布尔值 | 设为 `true` 启用这个存储 |
| `[[storages]]` 中的 `chat_id` | 整数 | 固定的回传目标；回传到自己的私聊时填写自己的用户 ID |
| `[[storages]]` 中的 `reuse_media` | 布尔值 | 设为 `true` 复用 Telegram 媒体引用，直接回传 |
| `[[storages]]` 中的 `force_file` | 布尔值 | 普通上传时是否强制作为文件发送；直接回传保留原媒体类型 |
| `[[users]]` 中的 `id` | 整数 | 允许操作机器人的 Telegram 用户 ID |
| `[[users]]` 中的 `storages` | 字符串列表 | 引用上面定义的存储名称，名称必须一致 |
| `[[users]]` 中的 `blacklist` | 布尔值 | `false` 表示只允许使用列表中的存储；`true` 表示排除列表中的存储 |

`chat_id` 决定媒体发给谁，`users.id` 决定谁能使用机器人。它们在回传给自己时填写相同值。一个存储的 `chat_id` 是固定目标，多个用户分别回传到自己时，应各自定义一个存储，并在各自的 `users.storages` 中引用。

例如，新增用户 `987654321` 时，可在文件末尾追加：

```toml
[[storages]]
name = "第二位用户"
type = "telegram"
enable = true
chat_id = 987654321
reuse_media = true

[[users]]
id = 987654321
storages = ["第二位用户"]
blacklist = false
```

启动后，每个用户通过 `/storage` 选择自己的默认存储，用 `/silent` 开启自动处理。这两个设置保存在数据库中。配置中的 `chat_id` 也可以是机器人有发送权限的群组或频道 ID；相关参数见 [Telegram 存储配置](./storages.md#telegram)。

未使用的扩展配置组可以省略。修改 `config.toml` 后，执行 `docker compose restart` 使配置生效。

## 详细配置

### 全局配置

- `lang`: Bot 使用的语言, 默认为 `zh-Hans` (简体中文), 设为 `en` 则使用英语.
- `stream`: 是否启用 Stream 模式, 默认为 `false`. 启用后 Bot 将直接将文件流式传输到存储端(若存储端支持), 不需要下载到本地
{{< hint warning >}}
Stream 模式对于磁盘空间有限的部署环境十分有用, 但也有一些弊端:
<br />
<ul>
<li>无法使用多线程从 Telegram 下载文件, 速度较慢.</li>
<li>网络不稳定时, 任务失败率高.</li>
<li>无法在中间层对文件进行处理, 例如自动文件类型识别.</li>
<li>并非支持所有存储端, 不支持的存储端可能会降级为普通模式或无法上传.</li>
</ul>
{{< /hint >}}
- `workers`: 同时处理任务数量, 默认为 3
- `threads`: 下载文件时使用的线程数, 默认为 4. 仅在未启用 Stream 模式时生效.
- `retry`: 任务失败时的重试次数, 默认为 3.
- `proxy`: 全局代理配置, 配置后程序内一切网络连接将会尝试使用该代理, 可选.

```toml
lang = "zh-Hans"
stream = false
workers = 3
threads = 4
retry = 3
proxy = "socks5://127.0.0.1:7890"
```

### Telegram 配置

- `token`: 你的 Telegram Bot Token, 可以通过 [BotFather](https://t.me/botfather) 创建 Bot 并获取 Token.
- `app_id`, `app_hash`: Telegram API ID & Hash, 在 [Telegram API](https://my.telegram.org/apps) 创建应用获取, 若不提供则使用默认值.
- `rpc_retry`: RPC 请求重试次数, 默认为 5.
- `proxy`: 代理配置, 可选.
  - `enable`: 是否启用代理.
  - `url`: 代理地址
- `userbot`: userbot 配置, 可选.
  - `enable`: 启用 userbot 集成, 需要登录用户账号, 此时请务必使用自己的 api id & hash.
  - `session`: userbot 会话文件路径, 默认为 `data/usersession.db`.

{{< hint warning >}}
启用 userbot 集成后, bot 可以下载私密频道和群组的文件, 但具有无法避免的账号被封禁的风险.
<br />
开启 userbot 集成后第一次启动 bot 时需要通过终端交互输入手机号, 2FA 和验证码.
<br />
如果你使用 docker 部署, 请使用 -it 参数为容器提供交互式环境, 然后执行登录操作.
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

### Aria2 配置

Aria2 是一个强大的下载管理器，支持 HTTP/HTTPS、FTP、BitTorrent 等多种协议。启用后，Bot 可以使用 `/aria2dl` 命令通过 Aria2 下载文件。

- `enable`: 是否启用 Aria2 支持，默认为 `false`
- `url`: Aria2 RPC 地址，通常为 `http://localhost:6800/jsonrpc`
- `secret`: Aria2 RPC 密钥，如果你在 Aria2 中配置了 `rpc-secret`，需要在此填写
- `keep_file`: 转存完成后是否保留 Aria2 下载的本地文件，默认为 `false`

{{< hint info >}}
Aria2 需要单独安装和运行。你可以参考 [Aria2 官方文档](https://aria2.github.io/) 了解如何安装和配置 Aria2。
{{< /hint >}}

```toml
[aria2]
enable = true
url = "http://localhost:6800/jsonrpc"
secret = "your-rpc-secret"
keep_file = false
```

### yt-dlp 配置

用于配置 `/ytdlp` 命令以及 HTTP API 中 `ytdlp` 任务类型在未传自定义参数时的默认行为.

- `max_height`: 默认下载的最高视频清晰度 (按高度限制), 如 `1080`, `720`, `480`; `0` 表示不限制 (下载最佳画质). 当设置了 `format` 时此项被忽略.
- `format`: 直接指定 yt-dlp format 选择表达式, 设置后优先级高于 `max_height`, 例如 `bv*[height<=720]+ba/b`.
- `recode`: 下载后转封装的视频容器格式 (如 `mp4`), 留空则不转封装.
- `filename_template`: 下载文件名模板 (yt-dlp output template), 默认 `%(title)s.%(ext)s`. 模板中的目录会创建在 bot 的下载目录下, 并保留到存储路径中.
- `restrict_filenames`: 将文件名限制为 ASCII 字符 (yt-dlp `--restrict-filenames`), 同时会去掉空格和 `&`. 默认关闭.

{{< hint info >}}
`max_height`、`format` 和 `recode` 仅在使用 `/ytdlp` 命令 (或 API 的 `ytdlp` 任务类型) 且未传任何自定义参数时生效. `filename_template` 和 `restrict_filenames` 在传了其他参数时依然生效, 除非自己传了对应的参数 (`-o/--output`、`--restrict-filenames` 或 `--no-restrict-filenames`).
{{< /hint >}}

```toml
[ytdlp]
max_height = 1080
format = ""        # 例如 "bv*[height<=720]+ba/b"
recode = "mp4"     # 留空则不转封装
filename_template = "%(title)s.%(ext)s"  # 例如 "%(uploader)s - %(title)s.%(ext)s"
restrict_filenames = false
```

### HTTP API 配置

启用后, SaveAny-Bot 会暴露一套 HTTP API, 用于以编程方式创建/查询/取消任务. 完整的接口说明见 [HTTP API](../../usage/api).

- `enable`: 是否启用 HTTP API 服务, 默认为 `false`.
- `host`: 监听地址, 默认 `0.0.0.0`.
- `port`: 监听端口, 默认 `8080`.
- `token`: 鉴权 Token，启用 API 时必须填写。

```toml
[api]
enable = false
host = "0.0.0.0"
port = 8080
token = "your-token"
```

### 日志配置

- `level`: 日志级别, 可选 `debug`, `info`, `warn`, `error`, `fatal`. 默认为 `debug`，回传模板设为 `info`.

```toml
[log]
level = "info"
```

### 存储端列表

存储端列表用于定义 Bot 支持的存储位置, 每个存储端需要指定名称、类型和相关配置, 使用双中括号语法 `[[storages]]` 定义.

每一个存储端至少需要以下字段:

- `name`: 存储端名称, 用于在 Bot 中识别, 需要唯一
- `enable`: 是否启用该存储端，设为 `true` 才会加载
- `type`: 存储端类型, 目前支持以下类型:
  - `local`: 本地磁盘
  - `alist`: Alist
  - `webdav`: WebDAV
  - `s3`: aws S3 及其他兼容 S3 的服务
  - `minio`: MinIO 对象存储
  - `rclone`: 调用 rclone 实现上传
  - `telegram`: 上传到 Telegram

示例, 这是一个包含本地存储和 webdav 存储的配置:

```toml
[[storages]]
name = "本地存储"
type = "local"
enable = true
# 以下是 local 类型存储的自定义配置
base_path = "./downloads"

[[storages]]
name = "WebDAV"
type = "webdav"
enable = true
# 以下是 webdav 类型存储的自定义配置
url = "https://example.com/webdav"
base_path = "/path/to/webdav"
username = "your_username"
password = "your_password"
```

所有存储端的自定义配置项可查看 [存储端配置](./storages) 

### 用户列表

用户列表用于定义对存储端的访问控制, 每个用户需要指定 Telegram 上的用户 ID, 使用双中括号语法 `[[users]]` 定义.

- `id`: 用户的 Telegram User ID
- `storages`: 过滤的存储端列表, 使用存储端名称定义, 默认为白名单模式 (即只允许访问列表中的存储端)
- `blacklist`: 是否启用黑名单模式, 默认为 `false`. 若启用黑名单模式, 则仅允许访问**没有**在列表中的存储端.

示例, 这是一个包含三个用户的配置, 用户 `123123` 只能访问本地存储, 用户 `456456` 只能访问除 WebDAV 以外的存储, 用户 `789789` 启用黑名单模式但没有指定存储端, 因此可以访问所有存储:

```toml
[[users]]
id = 123123
storages = ["本地存储"]

[[users]]
id = 456456
storages = ["WebDAV"]
blacklist = true

[[users]]
id = 789789
storages = []
blacklist = true
```

### 事件触发

事件触发提供了在 Bot 处理任务时根据任务状态执行自定义操作的能力, 目前仅支持任意命令执行. 使用 `[hook.exec]` 配置.

目前具有以下几种事件类型:

- `task_before_start`: 任务即将开始前
- `task_success`: 任务成功完成后
- `task_fail`: 任务失败后
- `task_cancel`: 任务被取消后

提供的配置值需要为完整的命令行命令, Bot 会在事件发生时执行该命令. 示例:

```toml
[hook.exec]
task_before_start = "echo '任务即将开始'"
task_success = "bash /path/to/success_script.sh"
task_fail = "curl -X POST https://example.com/api/notify -d '任务失败'"
task_cancel = "bash /path/to/cancel_script.sh"
```

### 解析器

解析器为 Bot 提供了处理非 Telegram 文件的能力, 例如从其他网站下载文件. 使用 `[parser]` 配置.

```toml
[parser]
plugin_enable = true # 是否启用解析器插件
plugin_dirs = ["./plugins"] # 插件目录, 可以是多个目录
```

上述两个配置项只用于控制以 JavaScript 编写的解析器插件, Bot 还有内置的使用 Go 实现的解析器, 目前默认开启.

### 杂项

```toml
no_clean_cache = false # 是否在退出时不清空缓存文件夹
# 临时下载文件夹配置
[temp]
base_path = "./cache"
```
