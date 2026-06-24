# API - 接口设计文档

# My Media Hub

版本：2.0

状态：Draft

协议：

HTTP RESTful API

数据格式：

JSON

统一前缀：

```http
/api/v1
```

---

# 一、设计原则

## Discovery First

首页入口必须围绕 Discovery Feed 设计。

禁止以媒体列表作为首页核心接口。

---

## Recommendation First

推荐接口属于核心接口。

优先级高于搜索接口。

---

## Behavior Driven

所有推荐能力来源于用户行为。

行为数据必须完整记录。

---

## API First

前后端通过 API Contract 解耦。

禁止前端依赖数据库结构。

---

## Backward Compatibility

已发布 API：

* 不删除字段
* 不修改字段含义
* 不修改字段类型

允许：

* 新增字段

---

# 二、统一响应格式

成功：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

失败：

```json
{
  "code": 10001,
  "message": "media not found",
  "data": {}
}
```

---

# 三、分页响应格式

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [],
    "total": 0,
    "page": 1,
    "page_size": 20
  }
}
```

---

# 四、JSON规范

禁止：

```go
json:",omitempty"
```

必须返回稳定字段。

---

数组：

```json
[]
```

禁止：

```json
null
```

---

对象：

```json
{}
```

禁止缺失字段。

---

# 五、错误码规范

| 范围          | 模块             |
| ----------- | -------------- |
| 0           | Success        |
| 10001~19999 | Media          |
| 20001~29999 | Behavior       |
| 30001~39999 | Recommendation |
| 40001~49999 | Search         |
| 50001~59999 | Scanner        |
| 90001~99999 | System         |

---

# 六、媒体模型

```json
{
  "id": 1,
  "media_type": "video",
  "title": "",
  "description": "",
  "path": "",
  "cover_path": "",
  "size": 0,

  "favorite_count": 0,
  "view_count": 0,
  "rating_count": 0,
  "avg_rating": 0,

  "tags": [],
  "categories": [],

  "created_at": "",
  "updated_at": ""
}
```

---

# 七、Discovery API（核心）

首页发现流。

---

## 获取发现流

```http
GET /api/v1/discovery/feed
```

Response：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "daily": [],
    "guess_like": [],
    "recent": [],
    "rediscover": [],
    "explore": []
  }
}
```

---

## 今日推荐

```http
GET /api/v1/discovery/daily
```

---

## 猜你喜欢

```http
GET /api/v1/discovery/guess-like
```

---

## 最近新增

```http
GET /api/v1/discovery/recent
```

---

## 重新发现

```http
GET /api/v1/discovery/rediscover
```

---

## 随机探索

```http
GET /api/v1/discovery/explore
```

---

# 八、Recommendation API（核心）

---

## 获取媒体推荐

```http
GET /api/v1/recommendation/media/{id}
```

---

## 随机推荐

```http
GET /api/v1/recommendation/random
```

---

## 标签推荐

```http
GET /api/v1/recommendation/tag
```

---

## 分类推荐

```http
GET /api/v1/recommendation/category
```

---

## 相似推荐

```http
GET /api/v1/recommendation/similar/{media_id}
```

---

## 推荐状态

```http
GET /api/v1/recommendation/status
```

---

## 重建推荐缓存

```http
POST /api/v1/recommendation/rebuild
```

---

# 九、Behavior API（核心）

统一行为接口。

对应：

user_behavior

---

## 记录行为

```http
POST /api/v1/behavior
```

Request：

```json
{
  "media_id": 1,
  "behavior_type": "favorite",
  "score": 5
}
```

---

behavior_type：

```text
favorite
rating
view
click
hidden
search_click
```

---

## 获取行为统计

```http
GET /api/v1/behavior/statistics
```

---

## 获取最近行为

```http
GET /api/v1/behavior/recent
```

---

# 十、Media API

媒体管理能力。

---

## 获取媒体列表

```http
POST /api/v1/media/page
```

---

## 获取媒体详情

```http
GET /api/v1/media/{id}
```

---

## 更新标签

```http
PUT /api/v1/media/{id}/tags
```

---

## 更新分类

```http
PUT /api/v1/media/{id}/categories
```

---

## 获取媒体元数据

```http
GET /api/v1/media/{id}/metadata
```

---

# 十一、Search API

搜索属于辅助能力。

---

## 搜索

```http
POST /api/v1/search/page
```

Request：

```json
{
  "keyword": "",
  "media_type": "",
  "page": 1,
  "page_size": 20
}
```

---

## 搜索建议

```http
GET /api/v1/search/suggestions
```

---

## 搜索历史

```http
POST /api/v1/search/history/page
```

---

## 删除搜索历史

```http
DELETE /api/v1/search/history/{id}
```

---

## 清空搜索历史

```http
DELETE /api/v1/search/history
```

---

# 十二、Tag API

---

## 标签列表

```http
POST /api/v1/tags/page
```

---

## 标签详情

```http
GET /api/v1/tags/{id}
```

---

# 十三、Category API

---

## 分类列表

```http
POST /api/v1/categories/page
```

---

## 分类详情

```http
GET /api/v1/categories/{id}
```

---

# 十四、Scanner API

---

## 启动扫描

```http
POST /api/v1/scanner/start
```

---

## 停止扫描

```http
POST /api/v1/scanner/stop
```

---

## 扫描状态

```http
GET /api/v1/scanner/status
```

---

# 十五、System API

---

## 系统信息

```http
GET /api/v1/system/info
```

---

## 系统统计

```http
GET /api/v1/system/statistics
```

Response：

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "media_count": 0,
    "image_count": 0,
    "video_count": 0,
    "novel_count": 0,
    "favorite_count": 0,
    "view_count": 0
  }
}
```

---

# 十六、版本规划

## V1

* Media
* Scanner
* Tag
* Category
* Search
* Behavior

---

## V2

* Recommendation
* Discovery Feed
* Interest Profile

---

## V3

* Similar Recommendation
* Rediscover
* Cross Media Recommendation

---

## V4

* AI Discovery
* Semantic Recommendation
* AI Tagging

---

# 十七、接口分层原则

严格遵守：

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

API层职责：

* 参数校验
* DTO转换
* 响应封装

禁止：

* SQL
* Repository直调
* 业务逻辑

---

# 十八、API优先级

系统开发优先级：

```text
1. Discovery API

2. Recommendation API

3. Behavior API

4. Media API

5. Search API

6. Scanner API

7. System API
```

该优先级必须与：

* PRD
* SAD
* ERD
* RecommendationEngine

保持一致。
