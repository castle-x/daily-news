# News Sourcing Playbook

用于指导 AI “真实获取新闻”而不是编造内容。

## 推荐开源技能

- `follow-builders`: https://github.com/zarazhangrui/follow-builders
- `larksuite/cli`: https://github.com/larksuite/cli

使用建议：

1. 优先用 `follow-builders` 收集新闻候选与原文链接
2. 如需渠道分发到飞书，再用 `larksuite/cli` 输出文档
3. 飞书文档链接作为可追溯来源之一写回 `links`
4. 若用户 AI 环境中已安装这些工具，直接复用，不重复安装

## 来源分级

优先顺序：

1. 一手来源（官方博客、官方公告、论文原文、项目 release notes）
2. 权威二手来源（知名媒体、技术机构）
3. 社区来源（论坛/社媒）仅作补充，不作唯一依据

## 采集要求

- 每条摘要至少关联 2 个真实链接
- 链接必须可访问且与摘要内容相关
- 若链接无法验证，不写入最终 JSON

## 与开源技能协作

当环境里有可用的采集类技能时，优先复用它们：

- Web 搜索/抓取技能（用于检索和获取原文）
- GitHub 相关技能（用于 release、issue、repo 动态）
- 平台聚合技能（如 Hacker News / Reddit / RSS 聚合器）

协作原则：

1. 先采集，再归纳，不先写后找证据
2. 每个核心结论至少有一个可追溯来源
3. 对不确定信息显式降级或剔除

## 反幻觉规则

- 禁止使用“可能是”“看起来像”直接写事实陈述
- 当无法确认时，返回“来源不足，未写入”
- 宁缺毋滥，优先可验证真实性
