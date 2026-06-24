# My Media Hub - 数据库设计文档（ERD）

版本：2.0

状态：Draft

数据库：

SQLite 3

---

# 一、设计目标

数据库围绕以下核心目标设计：

* Discovery First
* Recommendation First
* Behavior Driven

支持：

* 图片
* 视频
* 小说

支持未来：

* 漫画
* 音乐
* PDF
* 课程

支持：

* 百万级资源
* 推荐计算
* 用户兴趣画像
* 发现流生成

---

# 二、核心数据流

```text
Media
↓
User Behavior
↓
Interest Profile
↓
Recommendation
↓
Discovery Feed
```

---

# 三、实体关系图

```text
Media
│
├── MediaTag
├── MediaCategory
├── MediaMetadata
│
├── UserBehavior
│
├── SearchHistory
├── SearchClickHistory
│
├── InterestProfile
│
├── RecommendationCache
│
└── DiscoveryFeedLog
```

---

# 四、媒体资源

## media

媒体主表

| 字段          | 类型      |
| ----------- | ------- |
| id          | INTEGER |
| media_type  | TEXT    |
| title       | TEXT    |
| description | TEXT    |
| path        | TEXT    |
| hash        | TEXT    |
| size        | INTEGER |
| cover_path  | TEXT    |

统计字段：

| 字段             | 类型       |
| -------------- | -------- |
| favorite_count | INTEGER  |
| view_count     | INTEGER  |
| rating_count   | INTEGER  |
| avg_rating     | REAL     |
| last_viewed_at | DATETIME |

审计字段：

| 字段         | 类型       |
| ---------- | -------- |
| created_at | DATETIME |
| updated_at | DATETIME |

---

# 五、媒体元数据

## media_metadata

KV扩展字段

| 字段         | 类型      |
| ---------- | ------- |
| id         | INTEGER |
| media_id   | INTEGER |
| meta_key   | TEXT    |
| meta_value | TEXT    |

示例：

* author
* actor
* director
* publisher
* resolution

---

# 六、标签

## tag

| 字段   | 类型      |
| ---- | ------- |
| id   | INTEGER |
| name | TEXT    |

---

## media_tag

| 字段       | 类型      |
| -------- | ------- |
| media_id | INTEGER |
| tag_id   | INTEGER |

---

# 七、分类

## category

| 字段        | 类型      |
| --------- | ------- |
| id        | INTEGER |
| name      | TEXT    |
| parent_id | INTEGER |

---

## media_category

| 字段          | 类型      |
| ----------- | ------- |
| media_id    | INTEGER |
| category_id | INTEGER |

---

# 八、统一用户行为

## user_behavior

推荐系统核心表

| 字段            | 类型       |
| ------------- | -------- |
| id            | INTEGER  |
| media_id      | INTEGER  |
| behavior_type | TEXT     |
| score         | REAL     |
| created_at    | DATETIME |

behavior_type：

* favorite
* rating
* view
* click
* hidden
* search_click

示例：

favorite → score=5

rating(5星) → score=5

view → score=1

click → score=1

hidden → score=-5

---

# 九、搜索历史

## search_history

| 字段            | 类型       |
| ------------- | -------- |
| id            | INTEGER  |
| keyword       | TEXT     |
| search_source | TEXT     |
| result_count  | INTEGER  |
| created_at    | DATETIME |

search_source：

* search
* recommendation
* discovery

---

# 十、搜索点击历史

## search_click_history

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| keyword    | TEXT     |
| media_id   | INTEGER  |
| position   | INTEGER  |
| created_at | DATETIME |

---

# 十一、兴趣画像

## interest_profile

用户兴趣聚合结果

| 字段          | 类型       |
| ----------- | -------- |
| id          | INTEGER  |
| profile_key | TEXT     |
| score       | REAL     |
| updated_at  | DATETIME |

示例：

* 科幻
* 摄影
* 推理
* 动漫
* AI

---

# 十二、推荐缓存

## recommendation_cache

| 字段            | 类型       |
| ------------- | -------- |
| id            | INTEGER  |
| strategy_name | TEXT     |
| media_id      | INTEGER  |
| score         | REAL     |
| generated_at  | DATETIME |

strategy_name：

* random
* recent
* favorite_based
* tag_based
* category_based
* similar

---

# 十三、发现流记录

## discovery_feed_log

记录曾推荐给用户的内容

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| feed_type  | TEXT     |
| media_id   | INTEGER  |
| score      | REAL     |
| created_at | DATETIME |

feed_type：

* daily
* discover
* guess_like
* rediscover

作用：

避免首页长期推荐相同内容。

---

# 十四、索引原则

所有列表查询必须命中索引。

禁止全表扫描。

重点索引：

media：

* idx_media_type
* idx_media_created_at
* idx_media_hash

user_behavior：

* idx_behavior_media
* idx_behavior_type
* idx_behavior_created_at

interest_profile：

* idx_profile_score

recommendation_cache：

* idx_strategy_score

---

# 十五、未来扩展

新增媒体类型：

无需新增表。

仅扩展：

media.media_type

即可接入系统。

数据库结构长期保持稳定。
