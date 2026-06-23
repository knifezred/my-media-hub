# Recommendation Engine - 推荐引擎设计文档

# My Media Hub

版本：1.0

状态：Draft

---

# 一、设计目标

My Media Hub 的目标不是帮助用户管理资源。

而是帮助用户持续发现资源。

推荐系统需要解决的问题：

* 收藏越来越多
* 消费越来越少
* 不知道看什么
* 遗忘已有资源
* 重复下载相同资源

推荐引擎应成为系统核心能力。

---

# 二、设计原则

## Discovery First

发现优先。

推荐系统不是帮助用户找到已知内容。

而是帮助用户发现未知内容。

---

## Personal First

优先使用用户自身资源。

不依赖外部内容源。

---

## Lightweight First

必须适配：

* Intel N100
* 8GB RAM

不依赖：

* GPU
* 向量数据库
* 大模型

---

## Explainable First

推荐结果应尽可能可解释。

例如：

```text
因为你收藏过：
《三体》

所以推荐：

《球状闪电》
```

---

# 三、推荐架构

采用：

多策略混合推荐架构

---

```text
             Recommendation Service

                       │

     ┌─────────────────┼─────────────────┐

     ▼                 ▼                 ▼

 Content          Behavior         Discovery

     ▼                 ▼                 ▼

 Tag            Favorite         Random

 Category       Rating           Rediscover

 Similar        Viewed           Recent

 Metadata       Search
```

---

最终统一输出：

Recommendation Feed

---

# 四、推荐类型

## Random

随机推荐

---

目标：

帮助用户发现被忽略内容。

---

特点：

* 完全随机
* 去重
* 排除已隐藏

---

适用：

首页发现

今日推荐

---

# 五、Recent

最近新增推荐

---

目标：

让用户快速消费新资源。

---

排序：

```text
created_at DESC
```

---

支持：

图片

视频

小说

---

# 六、Favorite

收藏驱动推荐

---

数据来源：

user_favorite

---

逻辑：

分析用户收藏资源。

提取：

* 标签
* 分类
* 元数据

寻找相似资源。

---

示例：

```text
收藏：

科幻小说

↓

推荐：

更多科幻小说
```

---

# 七、Rating

评分驱动推荐

---

数据来源：

user_rating

---

评分权重：

```text
5分 = 100

4分 = 70

3分 = 40

2分 = 10

1分 = -50
```

---

高评分：

增加兴趣权重。

低评分：

降低兴趣权重。

---

# 八、Viewed

观看驱动推荐

---

数据来源：

user_viewed

---

目的：

识别真实消费行为。

---

权重：

低于收藏。

高于搜索。

---

# 九、Search

搜索驱动推荐

---

数据来源：

search_history

search_click_history

---

目标：

识别用户当前兴趣。

---

示例：

```text
最近搜索：

AI
机器学习
大模型

↓

增加相关资源推荐
```

---

# 十、Tag Recommendation

标签推荐

---

推荐核心之一。

---

计算方式：

统计：

用户互动最多的标签。

---

行为包括：

收藏

评分

观看

搜索

---

形成：

Tag Interest Profile

---

示例：

```text
科幻

AI

历史

编程
```

---

# 十一、Category Recommendation

分类推荐

---

与标签类似。

---

粒度更粗。

---

示例：

```text
小说

影视

图片
```

---

# 十二、Metadata Recommendation

元数据推荐

---

基于：

作者

导演

演员

出版社

摄影设备

系列名称

等字段。

---

示例：

```text
东野圭吾

↓

更多东野圭吾作品
```

---

# 十三、Similar Recommendation

相似资源推荐

---

用于：

资源详情页。

---

输入：

Media A

---

输出：

与 Media A 相似的资源。

---

# 十四、相似度计算

## Tag Similarity

权重：

40%

---

计算：

共同标签数量。

---

## Category Similarity

权重：

20%

---

计算：

共同分类。

---

## Metadata Similarity

权重：

20%

---

计算：

作者

导演

系列

等字段。

---

## Behavior Similarity

权重：

20%

---

计算：

用户行为聚合。

---

最终：

```text
Similarity Score

0 ~ 100
```

---

# 十五、Rediscover

重新发现推荐

---

My Media Hub 核心特色。

---

目标：

找回被遗忘资源。

---

条件：

超过90天未访问。

---

优先推荐：

高评分

高收藏

高价值内容。

---

示例：

```text
你已经180天没看过：

《银河帝国》

重新看看？
```

---

# 十六、兴趣画像

## Interest Profile

系统自动维护。

---

结构：

```text
Tag Score

Category Score

Author Score

Media Type Score
```

---

示例：

```json
{
  "science_fiction": 120,
  "ai": 95,
  "history": 60
}
```

---

# 十七、推荐评分模型

推荐分数：

```text
Final Score

=
Interest Score

+
Behavior Score

+
Freshness Score

+
Diversity Score
```

---

## Interest Score

兴趣匹配度。

---

## Behavior Score

收藏

评分

观看

搜索

综合计算。

---

## Freshness Score

鼓励新内容。

避免总推荐旧内容。

---

## Diversity Score

鼓励多样性。

避免推荐结果完全相同。

---

# 十八、冷启动

新用户：

无行为数据。

---

策略：

Recent

*

Random

*

Popular Tags

---

新资源：

无行为数据。

---

策略：

Tag

*

Metadata

推荐。

---

# 十九、推荐结果去重

必须避免：

连续推荐相同内容。

---

规则：

最近推荐7天内：

降低权重。

---

已隐藏资源：

永不推荐。

---

# 二十、性能设计

目标：

推荐响应：

<100ms

---

首页推荐：

<50ms

---

支持：

300万图片

300万小说

20万视频

---

推荐结果应优先从：

recommendation_cache

读取。

---

复杂计算采用：

异步任务预生成。

---

# 二十一、未来规划

## V2

兴趣画像增强

推荐反馈闭环

推荐解释

推荐质量评估

---

## V3

Embedding

语义推荐

跨媒体推荐

---

示例：

```text
喜欢：

《三体》

↓

推荐：

宇宙主题纪录片

NASA图片集

相关科幻小说
```

---

# 二十二、最终目标

用户打开 My Media Hub 后：

不需要搜索。

系统主动帮助用户发现：

* 新资源
* 被遗忘资源
* 感兴趣资源
* 高价值资源

让 NAS 从存储中心升级为内容发现中心。
