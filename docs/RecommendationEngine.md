# My Media Hub - 推荐引擎设计文档

版本：2.0

状态：Draft

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

## 收藏

behavior_type

```text
favorite
```

权重：

```text
+5
```

---

## 评分

```text
rating
```

权重：

```text
1~5
```

---

## 浏览

```text
view
```

权重：

```text
+1
```

---

## 点击

```text
click
```

权重：

```text
+1
```

---

## 不感兴趣

```text
hidden
```

权重：

```text
-5
```

---

## 搜索点击

```text
search_click
```

权重：

```text
+2
```

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
