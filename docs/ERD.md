# My Media Hub - 数据库设计文档（ERD）

版本：2.1

状态：Stable

数据库：

SQLite 3 + FTS5

---

# 一、设计目标

My Media Hub 是一个面向 NAS 的个人媒体发现中心。

数据库设计围绕以下目标展开：

* 支持百万级资源管理
* 支持推荐引擎
* 支持发现流
* 支持用户行为分析
* 支持未来媒体类型扩展
* 支持多来源元数据解析
* 支持搜索能力持续升级

---

# 二、设计原则

## 统一媒体模型

所有资源统一抽象为：

```text
Media
```

支持：

```text
image
video
novel
music
```

未来扩展：

```text
comic
pdf
course
archive
```

无需新增数据表。

---

## 状态与历史分离

当前状态：

存储在 media 表。

历史行为：

存储在 media_behavior 表。

避免频繁聚合计算。

---

## 元数据可扩展

不同媒体类型元数据差异极大：

视频：

```text
导演
演员
年份
```

音乐：

```text
专辑
歌手
时长
```

小说：

```text
作者
出版社
简介
```

图片：

```text
相机
镜头
ISO
```

采用：

```text
metadata_json
```

统一存储。

---

## 搜索独立

搜索索引不影响业务数据。

通过：

```text
media_fts
```

维护全文搜索能力。

---

# 三、实体关系图

```text
Media
 │
 ├── MediaContent
 │
 ├── MediaTag
 │      │
 │      └── Tag
 │
 ├── MediaCategory
 │      │
 │      └── Category
 │
 ├── MediaBehavior
 │
 ├── SearchHistory
 │
 ├── ScannerIndex
 │
 └── MediaFTS
```

---

# 四、数据表总览

| 表名             | 说明     |
| -------------- | ------ |
| media          | 媒体主表   |
| media_content  | 媒体内容扩展 |
| tag            | 标签字典   |
| media_tag      | 媒体标签关系 |
| category       | 分类字典   |
| media_category | 媒体分类关系 |
| media_behavior | 用户行为流水 |
| search_history | 搜索历史   |
| scanner_index  | 扫描缓存   |
| media_fts      | 全文搜索索引 |

---

# 五、媒体主表

## media

媒体资源核心表。

承载当前状态字段，避免聚合查询。

---

### 字段

| 字段               | 类型       | 说明                |
| ---------------- | -------- | ----------------- |
| id               | INTEGER  | 主键                |
| media_type       | TEXT     | 媒体类型              |
| title            | TEXT     | 标题                |
| description      | TEXT     | 描述                |
| path             | TEXT     | 文件路径              |
| hash             | TEXT     | 文件哈希              |
| size             | INTEGER  | 文件大小              |
| cover_path       | TEXT     | 封面路径              |
| status           | TEXT     | 当前状态              |
| last_error       | TEXT     | 最近一次错误信息           |
| metadata_json    | TEXT     | 元数据 JSON          |
| metadata_version | INTEGER  | 元数据版本             |
| favorite         | INTEGER  | 是否收藏（0/1）         |
| favorite_at      | TEXT     | 收藏时间              |
| rating           | REAL     | 当前评分              |
| rating_at        | TEXT     | 评分时间              |
| hidden           | INTEGER  | 是否隐藏（0/1）         |
| hidden_at        | TEXT     | 隐藏时间              |
| view_count       | INTEGER  | 浏览次数              |
| last_viewed_at   | TEXT     | 最近浏览时间            |
| created_at       | TEXT     | 创建时间              |
| updated_at       | TEXT     | 更新时间              |

### 状态机

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

合法状态值：

```text
new        # 新建尚未解析
parsing    # 正在解析
active     # 正常可用
missing    # 文件丢失
error      # 解析失败
deleted    # 已删除
```

---

# 六、媒体内容表

## media_content

用于存储：

* 小说内容预览
* OCR 文本
* AI 摘要
* 字幕内容
* NFO 内容

---

### 字段

| 字段           | 类型       | 说明   |
| ------------ | -------- | ---- |
| media_id     | INTEGER  | 媒体ID |
| content_type | TEXT     | 内容类型 |
| content      | TEXT     | 内容   |
| updated_at   | TEXT     | 更新时间 |

### 主键

```sql
PRIMARY KEY (media_id, content_type)
```

---

### content_type

```text
preview
ocr
summary
subtitle
nfo
```

---

# 七、标签

## tag

标签字典表。

---

### 字段

| 字段         | 类型       | 说明    |
| ---------- | -------- | ----- |
| id         | INTEGER  | 主键    |
| name       | TEXT     | 标签原始名 |
| name_norm  | TEXT     | 标准化名  |
| source     | TEXT     | 来源    |
| created_at | TEXT     | 创建时间  |

---

### source

```text
manual
parser
ai
```

---

### 唯一性

```sql
UNIQUE (name_norm)
```

`name_norm` 用于去重（如"波多野结衣" / "波多野結衣" 视为同一标签）。

---

# 八、媒体标签关系

## media_tag

---

### 字段

| 字段       | 类型      |
| -------- | ------- |
| media_id | INTEGER |
| tag_id   | INTEGER |

### 主键

```sql
PRIMARY KEY (media_id, tag_id)
```

---

# 九、分类

## category

分类字典表。

支持树形结构。

---

### 字段

| 字段         | 类型       |
| ---------- | -------- |
| id         | INTEGER  |
| name       | TEXT     |
| parent_id  | INTEGER  |
| level      | INTEGER  |
| path       | TEXT     |
| sort       | INTEGER  |
| created_at | TEXT     |

---

# 十、媒体分类关系

## media_category

---

### 字段

| 字段          | 类型      |
| ----------- | ------- |
| media_id    | INTEGER |
| category_id | INTEGER |
| is_primary  | INTEGER |

### 主键

```sql
PRIMARY KEY (media_id, category_id)
```

---

# 十一、用户行为流水

## media_behavior

记录所有用户行为。

用于：

* 推荐引擎
* 兴趣画像
* 用户分析

事件流模型。

---

### 字段

| 字段              | 类型       | 说明                |
| --------------- | -------- | ----------------- |
| id              | INTEGER  | 主键                |
| media_id        | INTEGER  | 媒体ID              |
| behavior_type   | TEXT     | 行为类型              |
| behavior_value  | TEXT     | 行为值（JSON 字符串）     |
| behavior_source | TEXT     | 行为来源              |
| created_at      | TEXT     | 创建时间              |

---

### behavior_type

```text
view
favorite
unfavorite
rate
hide
unhide
```

---

### behavior_value（JSON）

```json
// 收藏
{}

// 评分
{"rating": 4.5}

// 搜索点击
{"position": 3, "keyword": "三体"}
```

---

### behavior_source

```text
manual
search
recommendation
home_feed
```

---

# 十二、搜索历史

## search_history

记录用户搜索行为。

---

### 字段

| 字段           | 类型       |
| ------------ | -------- |
| id           | INTEGER  |
| keyword      | TEXT     |
| keyword_norm | TEXT     |
| use_count    | INTEGER  |
| last_used_at | TEXT     |
| created_at   | TEXT     |

---

# 十三、扫描索引

## scanner_index

增量扫描缓存。

---

### 字段

| 字段            | 类型        | 说明                          |
| ------------- | --------- | --------------------------- |
| id            | INTEGER   | 主键                          |
| media_id      | INTEGER   | 媒体 ID（未入库时为 NULL）           |
| file_path     | TEXT      | 文件路径                        |
| file_size     | INTEGER   | 文件大小                        |
| modified_time | TEXT      | 修改时间                        |
| file_hash     | TEXT      | 文件哈希                        |
| last_scan_at  | TEXT      | 最近扫描时间                      |

---

### 字段语义

```text
media_id = NULL     # 已扫到但尚未入库
media_id = <int>    # 已入库，关联到 media 表
```

---

# 十四、全文搜索

## media_fts

FTS5 虚拟表。

---

### 字段

| 字段          | 类型      |
| ----------- | ------- |
| rowid       | INTEGER |
| title       | TEXT    |
| description | TEXT    |
| content     | TEXT    |

---

### 同步策略

触发器同步：

```text
INSERT media
 ↓
INSERT media_fts

UPDATE media
 ↓
UPDATE media_fts

DELETE media
 ↓
DELETE media_fts
```

---

# 十五、推荐引擎数据来源

推荐引擎主要依赖：

## Media

```text
收藏状态
评分
浏览次数
最近浏览时间
```

---

## MediaBehavior

```text
收藏行为
评分行为
隐藏行为
浏览行为
```

---

## SearchHistory

```text
搜索兴趣
```

---

## Tag

```text
兴趣标签
```

---

## Category

```text
兴趣分类
```

---

# 十六、扩展能力

当前结构已支持：

| 扩展需求    | 是否支持 |
| ------- | ---- |
| 新媒体类型   | ✅    |
| 新元数据字段  | ✅    |
| 新标签体系   | ✅    |
| 新分类层级   | ✅    |
| 推荐引擎    | ✅    |
| 发现流     | ✅    |
| 搜索升级    | ✅    |
| AI摘要    | ✅    |
| OCR搜索   | ✅    |
| 兴趣画像    | ✅    |
| 多用户（未来） | ✅    |

---

# 十七、最终表数量

```text
业务表：

media
media_content

tag
media_tag

category
media_category

media_behavior

search_history

scanner_index

共 9 张
```

```text
搜索索引表：

media_fts

共 1 张
```

```text
总计：

10 张表
```

该结构已覆盖 M1（媒体库）、M2（用户行为）、M3（推荐引擎）以及未来发现流与 AI 能力扩展，无需再次进行数据库重构。
