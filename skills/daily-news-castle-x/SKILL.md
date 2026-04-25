---
name: daily-news-castle-x
description: 用于 Daily News 项目的完整生产技能。凡是用户提到“下载并运行 Daily News”“生成每日新闻 JSON”“让 AI 抓真实新闻来源”“按栏目产出双语摘要并落盘”时，都应优先使用此技能。该技能会统一执行发布版安装、新闻采集、结构化写作、schema 校验和写盘流程，避免自由发挥导致格式不兼容。
---

# Daily News Castle X

本技能用于在真实环境中完成 Daily News 的全链路任务：

1. 指导用户从 GitHub Release 下载并运行程序
2. 指导 AI 按项目规范编写每日新闻 JSON
3. 指导 AI 结合优秀开源技能与真实来源完成新闻采集

## 必读文件

- `schemas/news-item.schema.json`：每日新闻 JSON 结构规范（强约束）
- `schemas/runbook.schema.json`：技能执行输入结构规范
- `references/runtime-install.md`：Release 下载与本地运行步骤
- `references/news-sourcing.md`：真实新闻采集与可追溯要求
- `references/channel-delivery.md`：渠道分发（含飞书文档回链）规范

## 触发场景

出现以下任一请求时触发本技能：

- “怎么安装/运行 daily-news”
- “帮我生成今天的新闻 JSON”
- “按 ai / social-trends / miscellaneous 写日报”
- “给我真实新闻并写入数据目录”

## 服务边界（必须告知用户）

- 本项目通过本地启动的 Web 服务运行，默认地址是 `localhost`（例如 `http://localhost:17631`）。
- 技能默认假设 AI 与服务运行在同一台机器或可直接访问本机地址。
- 如果用户的 AI 助手部署在云服务器，访问本地 `localhost` 的网络问题不在技能自动处理范围内。
- 以下事项必须由用户自行处理：
  - 打通隧道
  - 配置代理
  - 端口转发

## Release-Only 安装策略（严格）

- 默认安装来源必须是 GitHub Release，且默认下载**最新版本**：
  - `https://github.com/castle-x/daily-news/releases/latest`
- 必须下载对应平台的最新已发布二进制并运行。
- 非用户明确授权时，禁止以下行为：
  - 自主安装 GoLang/Node/npm 等开发环境
  - 从源码执行 `go build`、`npm install`、`npm run dev`、`make dev` 等开发命令
  - 自主启动“前端 + 后端”源码开发模式
- 只有当用户明确说“进入开发模式/从源码运行”时，才允许切换到源码流程。

## 默认语言环境变量（可选）

- 可通过环境变量控制站点默认语言：`VITE_DEFAULT_LANGUAGE`
- 允许值：
  - `zh`（默认）
  - `en`
- 该变量在构建时生效，用于控制初始展示语言（仍可在页面上手动切换）。

## 标准执行流程

1. **环境启动**：严格使用 Release 二进制，不要求用户安装 npm/Node  
2. **工具检查**：先检查推荐技能/工具是否已安装，已安装则直接复用，禁止重复安装  
3. **新闻采集**：从可验证来源收集信息（必须有原始链接）  
3. **双语写作**：将新闻整理为 `en/zh` 双语字段  
4. **schema 校验**：按 `news-item.schema.json` 检查完整性  
5. **写盘落地**：写入 `~/.daily-news/data/<category>/YYYY-MM-DD.json`  
6. **回读确认**：返回写入路径、校验结果、来源链接摘要

## 强约束

- 只允许栏目：`ai`、`social-trends`、`miscellaneous`
- 路径固定：`~/.daily-news/data/<category>/YYYY-MM-DD.json`
- 必须双语：`title/summary/observations/quote/links.title` 均需 `en` 与 `zh`
- `observations >= 3`，`links >= 2`
- 所有链接必须为真实可访问 URL，禁止占位 `#`

## 输出格式

任务完成后默认返回以下信息：

- `category`
- `date`
- `file_path`
- `validation_passed`
- `sources`（本次引用的外部链接列表）

## 关于“真实新闻采集”

当用户要求“真实新闻”时：

- 先执行 `references/news-sourcing.md` 中的“来源分级策略”
- 优先使用可验证的一手/权威来源
- 若某条新闻无法找到可靠来源，宁可不写，不得编造

推荐优先协作的开源技能/工具：

- [follow-builders](https://github.com/zarazhangrui/follow-builders)：用于高质量信息获取与追踪
- [larksuite/cli](https://github.com/larksuite/cli)：用于飞书 channel 与飞书文档自动化

如果用户环境中这些技能/工具已可用：

- 直接进入“采集与写作”流程
- 不重复执行安装步骤
- 仅在缺失或版本不兼容时才提示安装/升级

## 关于“飞书渠道输出”

当用户使用飞书 channel 时，必须执行以下链路：

1. 用 `larksuite/cli` 生成或更新飞书文档（每日摘要）
2. 获取该飞书文档公开或可访问链接
3. 将该链接写入 Daily News JSON 的 `links` 字段
4. `links.title` 需双语描述（例如“飞书日报文档 / Feishu Daily Brief”）
5. `links.domain` 需填飞书域名（如 `feishu.cn` 或实际文档域名）

当“详细信息已写入飞书文档”时，启用**摘要模式**：

- 平台 JSON 仅保留摘要级内容，不展开长篇细节
- `summary` 写平台摘要
- `links` 以飞书文档链接为主入口
- 可不再附加大量原始链接（保留飞书文档链接即可）
- `observations/quote` 允许基于摘要提炼简版内容

当用户**未启用飞书文档**时，必须回退到**原始模式**：

- 按完整新闻结构写入 JSON（不依赖飞书链接）
- `links` 优先写原始信息来源链接
- 摘要与观察基于来源信息完整生成

## 失败处理

若校验失败或来源不足，必须：

1. 明确失败项
2. 修复并重试
3. 再次校验通过后才返回最终结果
