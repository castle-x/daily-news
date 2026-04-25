# Runtime Install and Run

本指南用于教 AI 如何引导用户在本地直接运行 Daily News（无需 npm/Node）。

## 1) 从 GitHub Release 下载

让用户打开仓库 Releases 页面，选择最新版本，并下载对应平台文件：

- Release 固定入口：`https://github.com/castle-x/daily-news/releases`
- 最新版本直达：`https://github.com/castle-x/daily-news/releases/latest`

- `daily-news-linux-amd64.tar.gz`
- `daily-news-linux-arm64.tar.gz`
- `daily-news-darwin-amd64.tar.gz`
- `daily-news-darwin-arm64.tar.gz`
- `daily-news-windows-amd64.zip`

## 2) 解压并运行

### Linux/macOS

```bash
tar -xzf daily-news-<platform>.tar.gz
chmod +x daily-news-<platform>
./daily-news-<platform>
```

默认监听：`http://localhost:17631`

### Windows

解压 `daily-news-windows-amd64.zip` 后，双击 `daily-news-windows-amd64.exe` 或在 PowerShell 运行：

```powershell
.\daily-news-windows-amd64.exe
```

## 严格限制（默认生效）

- 未经用户明确授权，禁止引导或执行源码开发流程。
- 禁止默认执行：
  - `go build`
  - `npm install`
  - `npm run dev`
  - `make dev`
- 默认只允许“下载 Release 二进制 -> 运行”这一路径。
- 若用户明确要求“源码开发/调试”，再切换到开发模式并提前说明原因。

## 3) 初始化数据目录

```bash
mkdir -p ~/.daily-news/data/ai \
         ~/.daily-news/data/social-trends \
         ~/.daily-news/data/miscellaneous
```

Windows 可在用户目录下创建同等路径结构（例如 `%USERPROFILE%\.daily-news\data\...`）。

## 4) 网络访问边界说明

- Daily News 默认是本地服务（`localhost`），不是公网服务。
- 若 AI 运行在云端环境，无法直接访问用户本机 `localhost`，需要用户自行处理网络连通。
- 常见处理方式（由用户自行实施）：
  - 隧道
  - 代理
  - 端口转发
