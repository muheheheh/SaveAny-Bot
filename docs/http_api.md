# HTTP API 使用说明

配置、鉴权、请求参数、响应示例、错误码、进度和 Webhook 回调详见 [接口文档](./content/zh/usage/api.md)。

## 功能概览

| 方法与路径 | 功能 |
| --- | --- |
| `GET /health` | 检查 API 服务状态 |
| `GET /api/v1/storages` | 列出可用存储 |
| `GET /api/v1/task-types` | 列出支持的任务类型 |
| `GET /api/v1/media-metadata` | 查询媒体元数据 |
| `POST /api/v1/tasks` | 创建下载或转存任务 |
| `GET /api/v1/tasks` | 列出任务 |
| `GET /api/v1/tasks/{task_id}` | 查询指定任务及进度 |
| `DELETE /api/v1/tasks/{task_id}` | 取消任务 |

## 配置与调用

API 默认关闭。在 `config.toml` 的 `[api]` 中配置 `enable`、`host`、`port` 和 `token`。启用后，请求使用 `Authorization: Bearer <token>` 进行鉴权；Token 留空时不会执行鉴权。

个人快速回传的 VPS 模板默认关闭 API，并且不发布端口。需要外部调用时，按 [完整说明](./content/zh/usage/api.md) 配置监听地址和容器网络。

## 任务记录与回调

任务记录仅保存在内存中，重启后丢失。任务结束后，记录会在最后更新满 24 小时后清理，清理检查每 10 分钟运行一次。

创建任务时可指定 `webhook`。任务成功或失败后发送回调；取消任务不触发回调。请求与重试细节见 [Webhook 回调](./content/zh/usage/api.md#webhook-回调)。
