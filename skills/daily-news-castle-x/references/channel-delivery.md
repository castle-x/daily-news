# Channel Delivery (Feishu First)

本规范定义“渠道输出”流程，当前重点支持飞书（Feishu）。

## 目标

将每日新闻摘要同步到飞书文档，并把文档链接回填到 Daily News JSON 的 `links` 模块，形成闭环。

## 推荐工具

- `larksuite/cli`: https://github.com/larksuite/cli

## 标准流程

1. 生成当日摘要（按 `news-item.schema.json`）
2. 使用 `larksuite/cli` 创建或更新飞书文档
3. 获取文档链接（可访问）
4. 将文档链接写入 JSON：
   - `links[].title.en`: `Feishu Daily Brief`
   - `links[].title.zh`: `飞书日报文档`
   - `links[].url`: 飞书文档链接
   - `links[].domain`: 文档域名（如 `feishu.cn`）
5. 再次校验 JSON 后写盘

## 回填规则

- 飞书文档链接可作为“相关链接”中的一条
- 若已启用“飞书承载详细内容”模式，允许仅保留飞书文档链接
- 未启用该模式时，建议保留 1-2 条原始来源链接
- 链接必须真实可访问，禁止占位符

## 摘要模式（Feishu Detail Mode）

适用场景：详细内容已经写入飞书文档，希望 Daily News 平台只显示简版摘要。

执行要求：

1. `summary` 仅保留精简摘要（中英双语）
2. `links` 放飞书文档链接作为详情入口
3. `observations/quote` 保持简短提炼，不重复大段正文
4. 不得伪造飞书链接；获取失败则回退常规模式

## 常规模式（无飞书文档）

适用场景：用户不使用飞书文档，或飞书插件不可用。

执行要求：

1. 按 `news-item.schema.json` 完整写入每日 JSON
2. `links` 使用真实新闻来源链接（非飞书链接）
3. 保持平台内可独立阅读，不依赖外部文档承载详情

## 失败处理

- 飞书 API 或 CLI 失败时，不要伪造文档链接
- 明确返回错误并保留原始新闻 JSON 结果（不写飞书链接）
