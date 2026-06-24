# 系统架构设计（SAD）

# My Media Hub

版本：2.1

状态：Stable

最后更新：2026-06

---

# 一、架构目标

My Media Hub 是一个部署在 NAS 上的个人媒体发现中心。

系统核心定位：

```text
Discovery Engine
+
Recommendation Engine
+
Behavior Driven System
```

系统目标不是帮助用户管理资源。

而是帮助用户持续发现值得消费的内容。

---

## 架构设计目标

* 简单可靠
* 易于维护
* 易于 AI 辅助开发
* 支持百万级资源
* 支持未来扩展
* 单机高性能运行
* 不依赖外部基础设施

---

## 目标硬件

```text
Intel N100
8GB RAM
```

---

# 二、架构原则

## Discovery First

发现流是系统核心入口。

首页优先展示推荐内容。

---

## Recommendation First

推荐系统是核心能力。

所有行为数据最终服务于推荐引擎。

---

## Search Third

搜索仅作为辅助发现能力。

搜索引擎实现可替换。

但不是系统核心。

---

## Behavior Driven

用户行为是系统最重要的数据资产。

包括：

* 浏览
* 点击
* 收藏
* 评分
* 已看
* 不感兴趣
* 搜索

推荐系统优先基于行为构建兴趣画像。

---

## Simplicity First

优先简单架构。

避免：

* 微服务
* 分布式系统
* 复杂基础设施

---

# 三、总体架构

采用：

**模块化单体（Modular Monolith）**

---

## 选择原因

* 部署简单
* NAS 场景适配
* 资源占用低
* AI 开发效率高
* 易于长期维护
* 后续可拆分

---

# 四、系统架构图

```text
┌───────────────────────────┐
│       Vue3 Frontend       │
└─────────────┬─────────────┘
              │ HTTP API
              ▼
┌───────────────────────────┐
│          Gin API          │
└─────────────┬─────────────┘
              ▼
┌────────────────────────────────────┐
│              Service               │
│                                    │
│ Recommendation Service (核心)      │
│ Discovery Feed Service (核心)      │
│ User Behavior Service (核心)       │
│ Media Service                      │
│ Search Service                     │
│ Scanner Service                    │
└─────────────┬──────────────────────┘
              ▼
┌────────────────────────────────────┐
│            Repository              │
└─────────────┬──────────────────────┘
              ▼
┌────────────────────────────────────┐
│              SQLite                │
└────────────────────────────────────┘
```

---

# 五、核心业务闭环

系统围绕内容发现闭环构建。

```text
文件扫描
↓
媒体索引
↓
用户浏览
↓
用户行为
↓
兴趣画像
↓
推荐引擎
↓
发现流
↓
用户继续消费
```

这是系统最重要的数据流。

---

# 六、分层架构

采用严格四层架构。

```text
API
↓
Service
↓
Repository
↓
Database
```

---

## API层

职责：

* 参数校验
* 请求解析
* DTO转换
* 响应输出
* 错误转换

禁止：

* SQL
* Repository调用
* 业务逻辑

---

## Service层

职责：

* 推荐逻辑
* 发现流生成
* 用户行为处理
* 搜索逻辑
* 业务规则实现

禁止：

* SQL
* 数据库操作

---

## Repository层

职责：

* 数据访问
* 查询封装
* 数据持久化

禁止：

* 业务逻辑
* 推荐逻辑

---

## Database层

职责：

* 数据存储

当前实现：

```text
SQLite
```

---

# 七、核心模块设计

## Recommendation 模块（核心）

系统最重要模块。

负责：

* 推荐计算
* 推荐排序
* 推荐策略执行
* 推荐结果生成

---

### V1推荐策略

* Random
* Recent
* Favorite Based
* Tag Based
* Category Based

---

### V2推荐策略

* Similar
* Rediscover
* Interest Profile

---

### V3推荐策略

* Cross Media Recommendation
* AI Discovery

---

# Discovery Feed 模块（核心）

负责：

首页发现流生成。

---

支持：

* 今日推荐
* 猜你喜欢
* 最近新增
* 继续浏览
* 重新发现
* 随机探索

---

# User Behavior 模块（核心）

负责：

采集用户行为。

---

支持：

* 收藏
* 评分
* 已看
* 不感兴趣
* 点击
* 浏览
* 搜索

---

输出：

```text
User Behavior
↓
Interest Profile
↓
Recommendation
```

---

## 媒体资源状态机

媒体主表（media）通过 `status` 字段表达资源的生命周期状态。

这是数据流的关键节点——所有外部系统（Scanner、用户操作、文件系统事件）都通过状态机驱动资源状态。

```text
new
 ↓
parsing
 ↓
active

parsing
 ↓
error

active
 ↓
missing

missing
 ↓
active

任意状态
 ↓
deleted
```

### 状态语义

| 状态       | 含义                  | 触发                    |
| -------- | ------------------- | --------------------- |
| new      | 新建尚未解析              | Scanner 发现新文件         |
| parsing  | 正在解析元数据             | 后台任务开始解析              |
| active   | 正常可用                | 解析成功 / missing 文件恢复   |
| missing  | 文件丢失（磁盘文件被删/移走）     | 定时扫描发现文件不存在          |
| error    | 解析失败（last_error 记录原因） | 解析任务异常               |
| deleted  | 用户主动删除（软删）          | 用户在 UI 移除资源          |

### 设计要点

* `status` 是状态机的入口——所有状态变更都走 Service 层
* `last_error` 字段记录 `error` 状态的原因（解析器输出、文件 I/O 错误等）
* `active` ↔ `missing` 双向转移：文件被恢复时无需重新解析
* `deleted` 是终态——软删后不进入推荐池

### 状态与行为流水的关系

`media.status` 表示资源生命周期（资源现在怎么样了）；
`media_behavior` 表示用户行为历史（用户对它做了什么）。

两者正交：

* 资源可处于 `parsing` 但用户已经 `view` 过（前端先添加后解析）
* 资源可处于 `missing` 但用户的 `favorite` 行为仍保留（行为流水不丢）

---

# Media 模块

负责：

媒体基础能力。

---

支持：

* 图片
* 视频
* 小说

---

统一媒体模型：

```text
Media
```

---

未来扩展：

* 漫画
* 音乐
* PDF
* 课程
* 网页存档

---

# Search 模块

搜索属于辅助模块。

---

职责：

* 搜索
* 搜索建议
* 标签搜索
* 分类搜索
* 条件筛选

---

采用：

```text
Search Strategy
```

抽象。

---

当前实现：

* LIKE
* FTS

---

未来扩展：

* Bleve
* Elasticsearch
* Semantic Search

---

# History 模块

负责：

用户历史记录。

---

支持：

* 浏览历史
* 搜索历史
* 推荐点击历史

---

# Scanner 模块

负责：

文件扫描。

---

支持：

* 首次全量扫描
* 增量扫描
* 元数据解析

---

# 八、搜索架构

采用：

Strategy Pattern

---

接口定义：

```go
type Strategy interface {
    Search(...)
    Suggestions(...)
}
```

---

当前实现：

```text
LikeStrategy
FTSStrategy
```

---

未来扩展：

```text
BleveStrategy
ElasticStrategy
SemanticStrategy
```

---

调用链：

```text
Search Service
↓
Search Index
↓
Strategy
```

业务层无需感知具体搜索实现。

---

# 九、推荐架构

采用：

多策略混合推荐。

---

推荐流程：

```text
Candidate Generator
↓
Strategy Executor
↓
Score Calculator
↓
Recommendation Result
```

---

输出：

```text
Recommendation Service
↓
Discovery Feed
```

---

# 十、扫描架构

采用：

```text
全量扫描
+
增量扫描
```

---

增量判断依据：

* 文件路径
* 文件大小
* 修改时间

---

# 十一、任务系统

系统内置轻量任务队列。

用于：

* 文件扫描
* 元数据提取
* 推荐计算
* 统计分析

---

不引入：

* Redis
* Kafka
* RabbitMQ

---

# 十二、缓存策略

原则：

只缓存计算结果。

---

缓存：

* 推荐结果
* 首页发现流
* 热门统计

---

不缓存：

* 搜索结果
* 媒体详情
* 用户行为

---

# 十三、数据库架构

数据库：

```text
SQLite
```

---

职责：

* 媒体数据
* 用户行为
* 推荐数据
* 系统配置

---

全文搜索：

```text
FTS Virtual Table
```

作为可选搜索实现。

---

# 十四、项目结构

```text
backend/

cmd/

internal/

api/
service/
repository/

modules/
├── media
├── recommendation
├── discovery
├── user
├── history
├── search
├── scanner

pkg/

frontend/
```

---

# 十五、扩展设计

## 搜索可插拔

新增搜索引擎：

仅需实现：

```go
Strategy
```

接口。

---

## 推荐可扩展

新增推荐算法：

仅需新增：

```text
Recommendation Strategy
```

实现。

---

## 媒体类型可扩展

新增媒体类型：

仅需：

* 新增 MediaType
* 新增 Metadata Extractor

无需修改核心流程。

---

# 十六、非功能需求

## 性能

推荐：

```text
<100ms
```

搜索：

```text
<100ms
```

---

## 启动时间

```text
<5秒
```

---

## 内存占用

常驻：

```text
<1GB
```

---

## 资源规模

图片：

```text
300万+
```

小说：

```text
300万+
```

视频：

```text
20万+
```

---

## 稳定性

单模块异常不得导致系统崩溃。

---

# 十七、依赖原则

所有依赖优先：

```text
Pure Go
```

避免：

```text
CGO
```

减少 NAS 平台兼容性问题。

---

# 十八、最终架构目标

让系统从：

```text
媒体管理系统
```

演进为：

```text
个人内容发现平台
```

最终实现：

```text
打开首页
↓
无需搜索
↓
持续发现值得消费的内容
```
