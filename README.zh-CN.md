# Daily News

[English](./README.md) | [中文](./README.zh-CN.md)

`Daily News` 是一个面向 AI 工作流的每日新闻摘要项目。

它的核心目标是：

- 让 AI 基于真实来源生成结构化新闻摘要
- 用统一 JSON 规范沉淀内容，供站点直接读取
- 支持双语（中文/英文）与分栏组织
- 支持把“详细内容”沉淀到外部渠道（如飞书文档），平台侧保留摘要与入口链接

## One-Line AI Quick Start

把下面这句话发给你的 AI 助手，即可快速开始安装并使用本项目技能：

```text
请加载并使用 https://github.com/castle-x/daily-news/tree/main/skills/daily-news-castle-x 这个技能，先检查本地可用工具（follow-builders、larksuite/cli），然后按技能规范生成并写入今日新闻到 ~/.daily-news/data/<category>/YYYY-MM-DD.json。
```

## 项目定位

Daily News 不只是一个前端页面，而是一套“AI 可执行的新闻生产协议”：

- 有固定 schema
- 有固定落盘路径
- 有明确的渠道协作规则
- 有可追溯来源要求

## License

本项目采用 [MIT License](./LICENSE)。

## Acknowledgements

感谢以下开源项目与作者提供的能力支持与启发：

- [zarazhangrui/follow-builders](https://github.com/zarazhangrui/follow-builders)（信息获取与追踪）
- [larksuite/cli](https://github.com/larksuite/cli)（飞书渠道自动化）
