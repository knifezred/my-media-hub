# ERD - 数据库设计文档

# My Media Hub

版本：1.0

状态：Draft

数据库：

SQLite 3 + FTS5

---

# 一、设计目标

数据库设计目标：

* 支持百万级资源
* 支持高速搜索
* 支持推荐计算
* 支持用户行为分析
* 支持未来媒体扩展

---

# 二、核心设计原则

## 媒体统一模型

图片

视频

小说

统一抽象为：

Media

---

## 行为驱动推荐

推荐系统主要依赖：

用户行为

而不是目录结构。

---

## 少 JOIN

优先：

冗余字段

避免复杂关联查询。

---

# 三、实体关系图

```text id="8yz1n9"
Media
 │
 ├── MediaTag
 │
 ├── MediaCategory
 │
 ├── MediaMetadata
 │
 ├── UserFavorite
 │
 ├── UserRating
 │
 ├── UserViewed
 │
 ├── UserHidden
 │
 ├── SearchHistory
 │
 └── RecommendationCache
```

---

# 四、媒体资源

## media

媒体主表

### 字段

| 字段          | 类型       |
| ----------- | -------- |
| id          | INTEGER  |
| media_type  | TEXT     |
| title       | TEXT     |
| description | TEXT     |
| path        | TEXT     |
| hash        | TEXT     |
| size        | INTEGER  |
| cover_path  | TEXT     |
| created_at  | DATETIME |
| updated_at  | DATETIME |

---

### media_type

支持：

```text id="1zhnd6"
image
video
novel
```

未来：

```text id="u0a4xg"
comic
music
pdf
course
```

---

### 索引

```sql id="3d4hkr"
idx_media_type

idx_media_created_at

idx_media_hash

idx_media_path
```

---

# 五、媒体元数据

## media_metadata

媒体扩展信息

### 字段

| 字段       | 类型      |
| -------- | ------- |
| id       | INTEGER |
| media_id | INTEGER |
| key      | TEXT    |
| value    | TEXT    |

---

### 示例

```text id="e9hzli"
author

publisher

actor

director

camera

resolution
```

---

### 索引

```sql id="ftl8uv"
idx_media_metadata_media

idx_media_metadata_key
```

---

# 六、标签

## tag

标签表

### 字段

| 字段   | 类型      |
| ---- | ------- |
| id   | INTEGER |
| name | TEXT    |

---

### 索引

```sql id="1jwvyr"
uniq_tag_name
```

---

# 七、媒体标签关系

## media_tag

### 字段

| 字段       | 类型      |
| -------- | ------- |
| media_id | INTEGER |
| tag_id   | INTEGER |

---

### 索引

```sql id="1dtpsf"
idx_media_tag_media

idx_media_tag_tag
```

---

# 八、分类

## category

### 字段

| 字段        | 类型      |
| --------- | ------- |
| id        | INTEGER |
| name      | TEXT    |
| parent_id | INTEGER |

---

# 九、媒体分类关系

## media_category

### 字段

| 字段          | 类型      |
| ----------- | ------- |
| media_id    | INTEGER |
| category_id | INTEGER |

---

# 十、收藏

## user_favorite

### 字段

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| media_id   | INTEGER  |
| created_at | DATETIME |

---

### 说明

单用户系统。

无需 user_id。

---

# 十一、评分

## user_rating

### 字段

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| media_id   | INTEGER  |
| rating     | INTEGER  |
| created_at | DATETIME |

---

### rating

范围：

```text id="ckmy8t"
1 ~ 5
```

---

# 十二、已看

## user_viewed

### 字段

| 字段        | 类型       |
| --------- | -------- |
| id        | INTEGER  |
| media_id  | INTEGER  |
| viewed_at | DATETIME |

---

# 十三、不感兴趣

## user_hidden

### 字段

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| media_id   | INTEGER  |
| created_at | DATETIME |

---

# 十四、搜索历史

## search_history

### 字段

| 字段           | 类型       |
| ------------ | -------- |
| id           | INTEGER  |
| keyword      | TEXT     |
| result_count | INTEGER  |
| created_at   | DATETIME |

---

### 用途

行为分析

推荐优化

---

# 十五、搜索点击历史

## search_click_history

### 字段

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| keyword    | TEXT     |
| media_id   | INTEGER  |
| position   | INTEGER  |
| created_at | DATETIME |

---

### 用途

分析：

用户真正选择了什么。

---

# 十六、推荐缓存

## recommendation_cache

### 字段

| 字段                  | 类型       |
| ------------------- | -------- |
| id                  | INTEGER  |
| recommendation_type | TEXT     |
| media_id            | INTEGER  |
| score               | REAL     |
| generated_at        | DATETIME |

---

### recommendation_type

```text id="h5i5z5"
random

recent

favorite

rating

tag

category

similar
```

---

# 十七、全文搜索

## media_fts

FTS5 虚拟表

---

索引字段：

```text id="l9f15y"
title

description

tags

author

metadata
```

---

查询：

```sql id="ukg7vt"
MATCH
```

---

# 十八、推荐引擎数据来源

推荐计算来源：

## 收藏

user_favorite

---

## 评分

user_rating

---

## 已看

user_viewed

---

## 搜索

search_history

---

## 点击

search_click_history

---

## 标签

media_tag

---

## 分类

media_category

---

# 十九、性能设计

目标：

图片：

300万+

小说：

300万+

视频：

20万+

---

原则：

所有列表查询必须命中索引。

禁止全表扫描。

---

FTS5负责：

全文搜索。

---

普通索引负责：

过滤

排序

推荐计算。

---

# 二十、未来扩展

新增媒体类型：

无需新增表。

仅新增：

```text id="0m0utl"
media_type
```

枚举值。

即可接入系统。

数据库结构保持稳定。
