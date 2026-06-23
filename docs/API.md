# API - 接口设计文档

# My Media Hub

版本：1.0

状态：Draft

API 风格：

RESTful API

数据格式：

JSON

---

# 一、设计原则

## 稳定优先

API 一旦发布：

* 不删除字段
* 不重命名字段
* 不修改字段含义

优先新增字段。

---

## 统一响应格式

所有 API 响应必须使用统一的 `code` / `message` / `data` 包装。

`code` 为 `0` 表示成功，非 `0` 表示错误码。

成功响应：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

失败响应：

```json
{
  "code": 10001,
  "message": "media not found",
  "data": {}
}
```

---

## 分页统一格式

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

## 数组规范

禁止：

```json
{
  "items": null
}
```

必须：

```json
{
  "items": []
}
```

---

## JSON规范

禁止：

omitempty

必须返回稳定字段结构。

---

# 二、错误码规划

## 编码规则

错误码为数字，按模块划分范围。

| 范围 | 分类 |
|------|------|
| 0 | 成功 |
| 10001-19999 | 媒体错误 |
| 20001-29999 | 用户行为错误 |
| 30001-39999 | 搜索错误 |
| 40001-49999 | 请求错误 |
| 50001-59999 | 扫描错误 |
| 90001-99999 | 系统错误 |

---

## 已注册错误码

| Code | Name | Description |
|------|------|-------------|
| 0 | SUCCESS | 成功 |
| 10001 | MEDIA_NOT_FOUND | 媒体资源不存在 |
| 10002 | MEDIA_TYPE_INVALID | 无效的媒体类型 |
| 10003 | MEDIA_ALREADY_EXISTS | 媒体哈希已存在 |
| 20001 | FAVORITE_ALREADY_EXISTS | 已收藏 |
| 20002 | FAVORITE_NOT_FOUND | 收藏不存在 |
| 20003 | RATING_INVALID | 评分超出范围 |
| 20004 | RATING_NOT_FOUND | 评分不存在 |
| 20005 | MEDIA_ALREADY_HIDDEN | 已隐藏 |
| 20006 | HIDDEN_NOT_FOUND | 隐藏记录不存在 |
| 30001 | SEARCH_KEYWORD_EMPTY | 搜索关键词不能为空 |
| 40001 | VALIDATION_ERROR | 请求校验失败 |
| 40002 | PARAMETER_INVALID | 请求参数错误 |
| 50001 | SCAN_IN_PROGRESS | 扫描正在进行中 |
| 50002 | SCAN_NOT_RUNNING | 当前没有扫描任务 |
| 90001 | INTERNAL_ERROR | 服务器内部错误 |

新增错误码时请在此表登记，遵循模块范围分配。

---

# 三、媒体模型

## Media

```json
{
  "id": 1,
  "media_type": "video",
  "title": "示例资源",
  "description": "",
  "path": "",
  "cover_path": "",
  "size": 0,
  "favorite": false,
  "rating": 0,
  "viewed": false,
  "hidden": false,
  "tags": [],
  "categories": [],
  "created_at": "2026-01-01T00:00:00Z",
  "updated_at": "2026-01-01T00:00:00Z"
}
```

---

# 四、媒体接口

## 获取媒体列表

POST

```http
/api/v1/media/page
```

---

### Request Body

| 参数          | 类型     |
| ----------- | ------ |
| page        | int    |
| page_size   | int    |
| media_type  | string |
| category_id | int    |
| tag_id      | int    |
| sort        | string |

---

### Response

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

## 获取媒体详情

GET

```http
/api/v1/media/{id}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "media_type": "video",
    "title": "",
    "description": "",
    "path": "",
    "cover_path": "",
    "size": 0,
    "favorite": false,
    "rating": 0,
    "viewed": false,
    "hidden": false,
    "tags": [],
    "categories": [],
    "metadata": {},
    "created_at": "",
    "updated_at": ""
  }
}
```

---

# 五、搜索接口

## 全文搜索

POST

```http
/api/v1/search/page
```

---

### Request Body

| 参数         | 类型     |
| ---------- | ------ |
| keyword    | string |
| media_type | string |
| page       | int    |
| page_size  | int    |

---

### Response

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

## 搜索建议

GET

```http
/api/v1/search/suggestions
```

---

### Query

```text
keyword
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": [
      "三体",
      "流浪地球"
    ]
  }
}
```

---

# 六、标签接口

## 标签列表

POST

```http
/api/v1/tags/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 标签详情

GET

```http
/api/v1/tags/{id}
```

---

# 七、分类接口

## 分类列表

POST

```http
/api/v1/categories/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

# 八、收藏接口

## 收藏资源

POST

```http
/api/v1/favorites
```

---

### Request

```json
{
  "media_id": 1
}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 取消收藏

DELETE

```http
/api/v1/favorites/{media_id}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 收藏列表

POST

```http
/api/v1/favorites/page
```

---

### Response

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

# 九、评分接口

## 评分

POST

```http
/api/v1/ratings
```

---

### Request

```json
{
  "media_id": 1,
  "rating": 5
}
```

---

### rating

范围：

```text
1~5
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

# 十、已看接口

## 标记已看

POST

```http
/api/v1/viewed
```

---

### Request

```json
{
  "media_id": 1
}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 已看列表

POST

```http
/api/v1/viewed/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

# 十一、不感兴趣接口

## 标记隐藏

POST

```http
/api/v1/hidden
```

---

### Request

```json
{
  "media_id": 1
}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 取消隐藏

DELETE

```http
/api/v1/hidden/{media_id}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

# 十二、搜索历史接口

## 搜索历史

POST

```http
/api/v1/search/history/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 删除搜索历史

DELETE

```http
/api/v1/search/history/{id}
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 清空搜索历史

DELETE

```http
/api/v1/search/history
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

# 十三、推荐接口

## 首页推荐

POST

```http
/api/v1/recommendations/home/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 随机推荐

POST

```http
/api/v1/recommendations/random/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 最近新增

POST

```http
/api/v1/recommendations/recent/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 相似推荐

POST

```http
/api/v1/recommendations/similar/{media_id}/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 标签推荐

POST

```http
/api/v1/recommendations/tag/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

## 分类推荐

POST

```http
/api/v1/recommendations/category/page
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "items": []
  }
}
```

---

# 十四、统计接口

## 总览统计

GET

```http
/api/v1/stats/overview
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total_media": 0,
    "total_images": 0,
    "total_videos": 0,
    "total_novels": 0,
    "favorite_count": 0,
    "viewed_count": 0
  }
}
```

---

# 十五、扫描接口

## 启动扫描

POST

```http
/api/v1/scanner/start
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

---

## 获取扫描状态

GET

```http
/api/v1/scanner/status
```

---

### Response

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "running": false,
    "processed": 0,
    "total": 0,
    "progress": 0
  }
}
```

---

# 十六、版本规划

## V1

支持：

* 媒体管理
* 搜索
* 收藏
* 评分
* 已看
* 历史记录

---

## V2

支持：

* 推荐引擎
* 首页发现流
* 相似推荐

---

## V3

支持：

* 兴趣画像
* AI发现引擎
* 语义搜索

---

# 十七、接口分层原则

API层负责：

* 参数校验
* 请求解析
* 响应转换

禁止：

* SQL
* Repository调用链绕过Service
* 业务逻辑

必须遵循：

API

↓

Service

↓

Repository

↓

Database

架构规范。
