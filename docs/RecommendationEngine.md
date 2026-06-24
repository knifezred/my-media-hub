# My Media Hub - 推荐引擎设计文档

版本：2.1

状态：Stable

最后更新：2026-06

---

# 一、设计目标

My Media Hub 的核心目标不是管理媒体资源。

而是帮助用户持续发现值得消费的内容。

因此推荐引擎是整个系统的核心能力。

---

推荐引擎负责：

* 内容发现
* 兴趣挖掘
* 资源再利用
* 个性化推荐
* 首页发现流生成

---

# 二、设计原则

## Discovery First

推荐系统服务于发现。

目标：

让用户无需搜索即可发现内容。

---

## Behavior Driven

用户行为是唯一可信数据来源。

包括：

* 收藏
* 评分
* 浏览
* 点击
* 已看
* 不感兴趣
* 搜索

---

## Recommendation First

推荐优先于搜索。

搜索属于辅助能力。

---

## Simple First

优先使用：

* 规则推荐
* 权重推荐
* 标签推荐

避免过早引入：

* AI模型
* 向量数据库
* LLM推荐

---

# 三、推荐系统总体架构

```text
Media
↓
User Behavior
↓
Interest Profile
↓
Candidate Generator
↓
Strategy Executor
↓
Score Calculator
↓
Recommendation Result
↓
Discovery Feed
```

---

# 四、推荐流程

## Step1

候选集生成

Candidate Generation

从多个策略产生候选内容。

---

## Step2

策略评分

Strategy Scoring

不同推荐策略独立计算分数。

---

## Step3

统一排序

Ranking

统一转换为推荐分数。

---

## Step4

结果输出

Recommendation Result

返回推荐内容列表。

---

# 五、用户行为体系

推荐引擎的数据来源：

## 行为类型（v2.1）

| behavior_type | 语义        | 权重贡献    |
| ------------- | --------- | ------- |
| view          | 浏览        | 弱正向     |
| favorite      | 收藏        | 强正向     |
| unfavorite    | 取消收藏      | 弱负向     |
| rate          | 评分        | 按 score |
| hide          | 不感兴趣      | 强负向     |
| unhide        | 取消不感兴趣    | 弱正向     |

权重计算在 v2.1 中已不再使用单 `score` 字段（v2.0 的方式已废弃）——

v2.0（旧）：

```text
favorite   score=5
rating     score=5
view       score=1
hidden     score=-5
```

v2.1（新）：行为类型 + JSON value 共同决定权重，由推荐引擎在运行时计算。

---

## 行为事件模型（v2.1）

每条 `media_behavior` 记录是一次完整行为快照，符合事件流（Event）标准做法。

字段：

* `behavior_type`：行为类型（view / favorite / rate / hide / ...）
* `behavior_value`：JSON 字符串，携带行为具体数据
* `behavior_source`：行为来源（manual / search / recommendation / home_feed）
* `created_at`：发生时间

---

## 行为值示例

不同行为类型的 `behavior_value` 结构：

### 收藏 / 取消收藏 / 不感兴趣 / 取消不感兴趣

```json
{}
```

或携带额外上下文（推荐原因 / 来源）：

```json
{
  "reason": "similar-tag"
}
```

### 评分

```json
{
  "rating": 4.5
}
```

支持 0.5 步进（如 3.5、4.5）。

### 浏览

```json
{
  "duration_sec": 3600
}
```

或仅记录访问：

```json
{}
```

### 搜索点击

```json
{
  "position": 3,
  "keyword": "三体"
}
```

---

## 为什么用 JSON 而非拆分多列

v2.0 方案是单 `score` 字段表达所有行为——6 种行为共用一列，归因不可逆（收藏 5 分 = 评分 5 分）。

v2.1 改用 JSON 原因：

* 字段组合不可枚举（不同行为携带不同上下文）
* 未来扩展无需改 schema（如新增"分享""下载"等行为）
* 行为流符合 Event Sourcing 范式
* Go 端用 `encoding/json` 反序列化，业务层做结构体映射

---

## 状态-历史 双轨

推荐引擎读取两类数据：

### 当前状态（来自 media 表）

```text
favorite         # 当前是否收藏
favorite_at      # 收藏时间
rating           # 当前评分
rating_at        # 评分时间
hidden           # 当前是否隐藏
view_count       # 浏览次数
last_viewed_at   # 最近浏览时间
```

用途：快速判断"用户对该资源的态度"，避免每次都聚合行为流水。

### 行为历史（来自 media_behavior 表）

```text
view         浏览次数
favorite     收藏时间线
rate         评分变化（重评）
hide         不感兴趣时间
```

用途：时序分析（最近 7 天行为 / 评分变化趋势 / 浏览频次）。

---

## 权重计算（v2.1 推荐）

推荐引擎按行为类型分别计算权重，公式示意：

```text
Score = Σ (behavior_weight × time_decay × type_weight)

behavior_weight:
  favorite  = +50
  rate ≥ 4  = +30
  rate 3    = +10
  rate ≤ 2  = -10
  view      = +1
  hide      = -100
  search_click = +5

time_decay:
  e^(-Δdays / 30)   # 30 天半衰期
```

权重公式由 Service 层实现，DB 只负责存储原始事件。

---

# 六、兴趣画像

兴趣画像来源：

```text
Behavior
+
Tag
+
Category
```

---

示例：

```text
科幻      120

推理      80

动漫      60

摄影      30
```

---

数据来源：

interest_profile

---

更新方式：

异步聚合。

---

# 七、推荐策略体系

推荐采用：

多策略混合推荐。

---

# V1 推荐策略

## Random

随机推荐

目标：

探索未知内容。

---

## Recent

最近新增

目标：

优先发现新资源。

---

## Favorite Based

收藏驱动推荐

目标：

推荐与收藏资源相似内容。

---

## Tag Based

标签推荐

目标：

推荐相同标签资源。

---

## Category Based

分类推荐

目标：

推荐相同分类资源。

---

# V2 推荐策略

## Similar Media

相似资源推荐

依据：

* 标签
* 分类
* 元数据

---

## Rediscover

重新发现

目标：

找回被遗忘资源。

---

规则：

长期未访问

高评分

高收藏

---

## Interest Profile

兴趣画像推荐

根据用户兴趣分布推荐。

---

# V3 推荐策略

## Cross Media

跨媒体推荐

示例：

```text
小说
↓
影视改编

影视
↓
原著小说
```

---

## AI Discovery

AI推荐

基于：

* LLM
* Embedding
* Semantic Analysis

---

# 八、候选集生成

Candidate Generator

负责：

生成候选内容。

---

候选来源：

```text
Random

Recent

Favorite

Tag

Category

Similar
```

---

输出：

```text
Candidate List
```

---

# 九、评分体系

统一评分公式：

```text
Final Score

=

Behavior Score

+

Tag Score

+

Category Score

+

Freshness Score

+

Random Score
```

---

示例：

```text
Favorite Score     50

Tag Score          20

Category Score     15

Freshness Score    10

Random Score        5

Final Score        100
```

---

# 十、推荐结果缓存

表：

recommendation_cache

---

缓存内容：

```text
strategy_name

media_id

score
```

---

缓存时间：

默认：

```text
30分钟
```

---

更新方式：

异步任务。

---

# 十一、发现流生成

推荐系统最终输出：

Discovery Feed

---

发现流组成：

## 今日推荐

Daily

---

## 猜你喜欢

Guess Like

---

## 最近新增

Recent

---

## 随机探索

Explore

---

## 重新发现

Rediscover

---

统一输出：

```text
Discovery Feed
```

---

# 十二、防重复推荐机制

记录：

discovery_feed_log

---

避免：

```text
连续多天推荐同一资源
```

---

规则：

最近7天推荐过：

降低权重。

最近30天频繁推荐：

直接过滤。

---

# 十三、异步任务

推荐相关任务：

## Interest Profile Rebuild

兴趣画像重建

---

## Recommendation Rebuild

推荐结果重建

---

## Rediscover Calculation

沉睡资源计算

---

## Recommendation Cache Refresh

缓存刷新

---

# 十四、性能目标

目标硬件：

```text
Intel N100
8GB RAM
```

---

首页推荐：

```text
<100ms
```

---

发现流生成：

```text
<100ms
```

---

兴趣画像重建：

```text
<5分钟
```

---

推荐缓存刷新：

```text
<1分钟
```

---

# 十五、未来演进路线

V1

规则推荐

---

V2

兴趣画像

---

V3

跨媒体推荐

---

V4

语义推荐

---

V5

AI Discovery

---

# 十六、最终目标

用户打开 My Media Hub 后：

无需搜索。

无需筛选。

无需记忆资源位置。

系统主动帮助用户发现：

```text
已经拥有

却被遗忘

但值得消费的内容
```

让 NAS 从资源仓库演进为个人内容发现平台。
