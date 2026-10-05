---
title: "安装与更新"
---

# 安装与更新

## 源码构建

安装 Go 1.25 或更新版本，然后从本仓库构建：

```sh
git clone https://github.com/muheheheh/SaveAny-Bot.git
cd SaveAny-Bot
CGO_ENABLED=0 go build -trimpath -o saveany-bot .
```

复制仓库内的 `deploy/vps/config.example.toml` 为 `config.toml`，填写 Bot Token、目标聊天 ID 和允许使用的用户 ID。模板默认启用快速回传。更多选项见 [配置说明](./configuration/_index.md)。

运行:

```bash
chmod +x saveany-bot
./saveany-bot
```

### 进程守护

{{< tabs "daemon" >}}
{{< tab "systemd (常规 Linux)" >}}

创建文件 <code>/etc/systemd/system/saveany-bot.service</code> 并写入以下内容:

{{< codeblock >}}
[Unit]
Description=SaveAnyBot
After=systemd-user-sessions.service

[Service]
Type=simple
WorkingDirectory=/yourpath/
ExecStart=/yourpath/saveany-bot
Restart=always

[Install]
WantedBy=multi-user.target
{{< /codeblock >}}

设为开机启动并启动服务:

{{< codeblock >}}
systemctl enable --now saveany-bot
{{< /codeblock >}}

{{< /tab >}}

{{< tab "procd (OpenWrt)" >}}

<h4>添加开机自启动服务</h4>

创建文件 <code>/etc/init.d/saveanybot</code> ，参考 <a href="https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/confs/wrt_init" target="_blank">wrt_init</a> 并自行修改:

{{< codeblock >}}
#!/bin/sh /etc/rc.common

#This is the OpenWRT init.d script for SaveAnyBot

START=99 
STOP=10
description="SaveAnyBot"

WORKING_DIR="/mnt/mmc1-1/SaveAnyBot"
EXEC_PATH="$WORKING_DIR/saveany-bot"
start() {
    echo "Starting SaveAnyBot..."
    cd $WORKING_DIR
    $EXEC_PATH &
}
stop() {
    echo "Stopping SaveAnyBot..."
    killall saveany-bot
}
reload() {
    stop
    start
}

{{< /codeblock >}}

赋予权限:

{{< codeblock >}}
chmod +x /etc/init.d/saveanybot
{{< /codeblock >}}

然后将文件复制到 <code>/etc/rc.d</code> 并重命名为 <code>S99saveanybot</code>, 同样赋予权限:

{{< codeblock >}}
chmod +x /etc/rc.d/S99saveanybot
{{< /codeblock >}}

<h4>添加快捷指令</h4>

创建文件 <code>/usr/bin/sabot</code> ，参考 <a href="https://github.com/muheheheh/SaveAny-Bot/blob/main/docs/confs/wrt_bin" target="_blank">wrt_bin</a>  并自行修改，注意此处文件编码仅支持 ANSI 936 .

随后赋予权限:

{{< codeblock >}}
chmod +x /usr/bin/sabot
{{< /codeblock >}}

使用: <code>sudo sabot start|stop|restart|status|enable|disable</code>

{{< /tab >}}
{{< /tabs >}}


## 使用 Docker 部署

在本仓库的完整源码目录中，按上文准备好 `config.toml` 后运行：

```sh
docker compose up -d --build
```

根目录的 Compose 文件会从源码构建镜像。
小内存 VPS 可参考 [本地编译后部署](https://github.com/muheheheh/SaveAny-Bot/blob/main/deploy/vps/README.md)，在开发机器编译后上传。

如需精简版本，可使用 `Dockerfile.micro` 或 `Dockerfile.pico` 构建。各版本的依赖和构建标签见对应 Dockerfile。

## 更新

拉取更新后，重新构建并重建容器：

```sh
git pull --ff-only
docker compose up -d --build
```

如果使用本地编译后部署的 VPS 模板，请重新编译、上传二进制，再在服务器部署目录运行 `docker compose up -d --build`。

原生二进制部署则重新运行上面的 Go 编译命令，替换可执行文件并重启服务。
`/update` 和 `./saveany-bot up` 需要可用的 GitHub Release。Docker 部署使用上述重建步骤更新。
