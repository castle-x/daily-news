# Daily News

[English](./README.md) | [中文](./README.zh-CN.md)

`Daily News` is an AI-first daily news summarization project.

Its core goals are:

- Let AI generate structured daily news from real, verifiable sources
- Persist content with a stable JSON contract for direct web consumption
- Support bilingual output (Chinese/English) with category-based organization
- Allow detailed content to live in external channels (for example, Feishu docs) while the platform keeps concise summaries and entry links

## One-Line AI Quick Start

Send the following prompt to your AI assistant to get started quickly:

```text
Please load and use the skill at https://github.com/castle-x/daily-news/tree/main/skills/daily-news-castle-x, first check whether follow-builders and larksuite/cli are already available, then generate today's news and write it to ~/.daily-news/data/<category>/YYYY-MM-DD.json according to the skill schema.
```

## Project Positioning

Daily News is more than a frontend page. It is an executable AI content protocol with:

- A fixed schema
- A fixed disk output path
- Clear channel collaboration rules
- Source traceability requirements

## License

This project is licensed under the [MIT License](./LICENSE).

## Acknowledgements

Thanks to these open-source projects and authors for inspiration and practical tooling:

- [zarazhangrui/follow-builders](https://github.com/zarazhangrui/follow-builders) (information sourcing and tracking)
- [larksuite/cli](https://github.com/larksuite/cli) (Feishu channel automation)
