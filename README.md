# My Media Hub

> Personal Media Discovery Center for NAS

一个专为 NAS 打造的个人媒体发现中心。

My Media Hub 不关注文件管理，而专注于帮助用户从海量媒体资源中持续发现值得消费的内容。

支持：

* 图片（Images）
* 视频（Videos）
* 小说（Novels）

未来支持：

* 漫画
* 音乐
* PDF
* 课程
* 网页存档

---

# 项目愿景

大多数 NAS 用户都会遇到同一个问题：

资源越来越多，消费越来越少。

随着资源积累：

* 收藏越来越多
* 下载越来越多
* 遗忘越来越多

最终变成：

> 我知道 NAS 里有很多好东西，但我不知道看什么。

My Media Hub 希望解决的正是这个问题。

---

# 项目定位

My Media Hub 是：

✅ 发现引擎（Discovery Engine）

✅ 推荐引擎（Recommendation Engine）

✅ 个人媒体中心（Personal Media Hub）

---

My Media Hub 不是：

❌ 播放器

❌ 下载器

❌ 文件管理器

❌ 网盘系统

---

# 核心理念

```text
发现优先

↓

推荐其次

↓

搜索第三

↓

管理最后
```

用户打开应用后：

不应该先搜索。

而应该先发现内容。

---

# 核心能力

## 媒体库

支持：

* 图片
* 视频
* 小说

统一媒体模型。

---

## 高性能搜索

支持：

* 全文搜索
* 标签搜索
* 分类搜索
* 条件筛选

目标：

百万级资源毫秒级响应。

---

## 用户行为分析

支持：

* 收藏
* 评分
* 已看
* 不感兴趣
* 搜索历史
* 点击历史

---

## 推荐引擎

支持：

* 随机推荐
* 最近新增
* 收藏推荐
* 标签推荐
* 分类推荐
* 相似推荐
* 重新发现推荐

---

## 发现流

首页即推荐。

持续帮助用户发现：

* 新内容
* 高价值内容
* 被遗忘内容

---

# 技术架构

## Backend

```text
Go
Gin
SQLite
FTS5
```

---

## Frontend

```text
Vue3
TypeScript
Pinia
```

---

## Deployment

```text
UGREEN NAS Native App
```

不依赖：

```text
Docker
Redis
Kafka
RabbitMQ
```

---

# 系统架构

```text
Vue3

↓

Gin API

↓

Service

↓

Repository

↓

SQLite
```

采用：

**Modular Monolith（模块化单体）**

架构。

---

# 性能目标

目标设备：

```text
Intel N100
8GB RAM
```

支持规模：

```text
图片     300万+

小说     300万+

视频      20万+
```

目标响应时间：

```text
搜索      <100ms

推荐      <100ms

启动时间   <5秒
```

---

# 项目文档

所有项目知识文档位于：

```text
docs/
```

## 产品文档

```text
docs/PRD.md
```

产品需求文档。

---

## 架构文档

```text
docs/SAD.md
```

系统架构设计。

---

## 数据库设计

```text
docs/ERD.md
```

数据库设计。

---

## API设计

```text
docs/API.md
```

接口规范。

---

## 推荐引擎

```text
docs/RecommendationEngine.md
```

推荐系统设计。

---

## UI设计

```text
docs/UI.md
```

界面设计规范。

---

## 路线图

```text
docs/Roadmap.md
```

项目演进路线图。

---

# AI开发

项目支持：

* Trae
* Cursor
* Claude Code
* Roo Code
* Cline
* GitHub Copilot

---

AI开发入口：

```text
AGENTS.md
```

---

AI开发说明：

```text
AI_README.md
```

---

AI治理规则：

```text
.ai/
```

---

# 开发原则

* Discovery First
* Recommendation First
* Simplicity First
* Performance First

---

禁止：

* 微服务
* 过度设计
* 无意义抽象

---

---

# 贡献

欢迎提交：

* Bug Report
* Feature Request
* Pull Request
* Documentation Improvement

请先阅读：

```text
CONTRIBUTING.md
```

---

# License

Apache License 2.0

---

# 最终目标

让 NAS 从：

```text
存储中心
```

升级为：

```text
内容发现中心
```

帮助用户持续发现自己已经拥有却被遗忘的内容。
