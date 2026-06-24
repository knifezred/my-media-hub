# Task Breakdown

AI-Owned File

Governed by:

.ai/project-governance.md

AI agents MAY generate and update tasks within the active milestone.

---

# Active Milestone

Milestone 1: 媒体库基础能力
Milestone 2: 用户行为系统

M9（搜索）作为辅助能力一并维护。

---

# Phase 1: 数据库 Schema 重建

- [x] database.go migration 重写（M1+M2 表 + FTS5，不含 M3+）
- [x] 索引创建（按 ERD v2.0）
- [x] 旧数据库删除重建

# Phase 2: Model 更新

- [x] media.go: 新增 FavoriteCount, ViewCount, RatingCount, AvgRating, LastViewedAt
- [x] media.go: MediaMetadata 字段名 Key→MetaKey, Value→MetaValue
- [x] user.go: 删除旧4表结构体，新增 UserBehavior
- [x] dto.go: 新增 BehaviorRequest, BehaviorStatistics
- [x] dto.go: 保留 FavoritePageRequest, RateRequest, MediaIDRequest（前端兼容）
- [x] search.go: SearchHistory 新增 SearchSource 字段

# Phase 3: Repository 重构

- [x] behavior_repository.go 创建（Insert, GetByMediaID, ListByType, 统计方法）
- [x] favorite_repository.go 删除
- [x] rating_repository.go 删除
- [x] viewed_repository.go 删除
- [x] hidden_repository.go 删除
- [x] media_repository.go: 统计字段扫描 + UpdateStats
- [x] media_metadata_repository.go: key→meta_key
- [x] stats_repository.go: 改用 user_behavior 查询
- [x] search_history_repository.go: 新增 search_source 字段

# Phase 4: Service 重构

- [x] behavior_service.go 创建（Record, GetStatistics, IsFavorited, ListFavorites 等）
- [x] user_service.go 删除
- [x] media_service.go: UserService→BehaviorService
- [x] search_service.go: Search 方法传入 search_source
- [x] stats_service.go: 使用 behaviorRepo

# Phase 5: API 重构

- [x] behavior_api.go 创建（统一 behavior + 旧路径兼容）
- [x] user_api.go 删除
- [x] media_api.go: GetTagByID 实现
- [x] search_api.go: 保留 history 端点

# Phase 6: Error Code

- [x] 保留 M2 相关错误码
- [x] 新增 NotFound 错误码（40003）

# Phase 7: Router + Main

- [x] router.go: userAPI→behaviorAPI
- [x] main.go: 更新依赖

# Phase 8: 前端适配（已完成）

- [x] api/behavior.ts → 新建（统一行为 API，含所有端点方法）
- [x] api/favorite.ts → 保留，增加 GET /favorites 端点
- [x] api/rating.ts → 保留，增加 GET /ratings/:id + PUT /ratings/:id 端点
- [x] api/viewed.ts → 保留，增加 GET /history + POST /history 端点
- [x] api/hidden.ts → 保留，增加 GET /hidden 端点
- [x] api/history.ts → 修复 SearchHistory 导入（从 types 导出）
- [x] types/index.ts → Media 新增统计字段 + BehaviorStatistics + SearchHistory 导出

# Phase 9: 编译验证

- [x] go mod tidy
- [x] go build ./...
- [x] go vet ./...
- [x] 启动测试（数据库自动重建 + 无报错）
- [x] Media 列表 API ✅
- [x] Media 详情 API ✅（空数据库返回 NotFound）
- [x] Behavior Record API ✅
- [x] Behavior Statistics API ✅
- [x] 旧路径兼容（/favorites）✅
- [x] Stats Overview API ✅

---

# Phase 10: ERD v2.0 → v2.1 升级（变更任务）

> 触发：docs/ERD.md v2.0 已升级到 v2.1
> 涉及 7 处 schema 调整 + 文档同步
> 阻塞 M2 验收 / M3 启动前必须完成

## 10.1 Schema 调整（database.go）

- [x] `media` 表新增 `last_error TEXT` 字段
- [x] `media.rating` 由 INTEGER 改为 REAL（应用层做范围校验）
- [x] `media_content` 表删除 `id` 自增主键
- [x] `media_content` 表改用 `PRIMARY KEY (media_id, content_type)` 复合主键
- [x] `tag` 表明确 `UNIQUE (name_norm)` 唯一约束
- [x] `tag` 表新增 `INDEX (name)` 用于原始名查询
- [x] `scanner_index.media_id` 允许 NULL（已扫但未入库）
- [x] `media_behavior.behavior_value` 文档化为 JSON 字符串

## 10.2 Model 同步（model/）

- [x] `media.go`: `Media.Rating` 类型 `int` → `float64`
- [x] `media.go`: `Media` 新增 `LastError string` 字段
- [x] `media.go`: `Media` 新增 `Status` 字段（active/missing/parsing/error/deleted）
- [x] `media.go`: `Media` 新增 `MetadataJSON string` 字段
- [x] `media.go`: `Media` 新增 `MetadataVersion int` 字段
- [x] `content.go`: 新建 `MediaContent` 结构体，去 `ID`，主键 `(MediaID, ContentType)`
- [x] `behavior.go`: 新建 `MediaBehavior` 结构体，`BehaviorValue` 为 JSON string
- [x] `user.go` 删除（拆为 behavior.go + media.go 字段）
- [x] `dto.go`: `RateRequest.Rating` int → float64

## 10.3 Service / Repository 同步

- [x] `media_repository.go`: 新增 `UpdateStatus(id, status, lastError)` 方法
- [x] `media_repository.go`: 新增 `SetFavorite` / `SetRating` / `SetHidden` / `IncViewCount` 方法
- [x] `media_repository.go`: 删除 `UpdateStats` / `UpdateViewStats`（双写炸弹已拆除）
- [x] `media_service.go`: 删除 `media_metadata` 依赖（已无 metadata 查询）
- [x] `media_metadata_repository.go` 删除（表已移除）
- [x] `media_tag_repository.go` / `media_category_repository.go` 合并到 `relation_repository.go`
- [x] `tag_repository.go`: Create/GetByName 改用 `name_norm` 唯一性
- [x] `tag_repository.go`: 新增 `NormalizeName()` helper（放入 `helpers.go`）
- [x] `behavior_service.go`: 评分范围校验支持 0.5 步进（0.5 / 1.0 / ... / 5.0）
- [x] `behavior_service.go`: 双写逻辑改为直接更新 media 字段（favorite / rating / hidden / view_count）
- [x] `scanner_repository.go`: 重写为 `scanner_index` 表操作（支持 `media_id IS NULL` 查询）
- [x] `scanner_service.go`: 增量扫描使用 `scanner_index` 缓存（而非直接查 media）
- [x] `search_history_repository.go`: 改为 `InsertOrUpdate`（keyword 去重聚合）
- [x] `stats_repository.go`: 适配新 Media 字段
- [x] `search_repository.go`: 适配新 Media 字段

## 10.4 API / Router 同步

- [x] `behavior_api.go`: 统一端点 + 便捷端点（favorite/unfavorite/rate/view/hide/unhide）
- [x] `behavior_api.go`: 废弃旧 `score` 参数，改用 `behavior_value` JSON
- [x] `media_api.go`: 适配新 MediaDetail（无 Metadata map）
- [x] `router.go`: 适配新依赖关系
- [ ] `behavior_api.go`: Request/Response 改用 `behavior_value` JSON
- [ ] `router.go`: 兼容性兜底（旧字段名 → 新字段名映射）

## 10.5 文档同步（已完成）

- [x] `docs/ERD.md` 升级到 v2.1
- [x] `docs/API.md` 媒体模型 + 行为 API 改用新字段
- [x] `docs/SAD.md` 补充状态机说明
- [x] `docs/RecommendationEngine.md` 行为事件 JSON 化 + 权重公式

## 10.6 验证

- [ ] `go build ./...` 通过
- [ ] `go vet ./...` 通过
- [ ] 数据库自动重建无报错
- [ ] 行为 API 测试：rate=4.5 成功 / rate=4.7 失败
- [ ] 媒体 API 返回新字段（status / metadata_json / last_error）
- [ ] Scanner API 测试：未入库文件 `media_id = NULL`
- [ ] 旧 API 兼容测试：老前端收藏页能正常加载
- [ ] media_content 复合主键测试：同 media_id + 不同 content_type 可共存

---

# Phase 11: 行为表 score 字段重构（变更任务）

> 触发：ERD v2.0 `behavior_value` 取代 v1 的单 `score` 字段
> 阻塞 M3 推荐引擎启动

## 11.1 数据迁移

- [ ] 评估现有数据规模（user_behavior 表行数）
- [ ] 写 migration 脚本：`behavior_value` 字段填充
  - `behavior_type=favorite` → `behavior_value = "{}"`
  - `behavior_type=rating` → `behavior_value = "{\"rating\": <score>}"`
  - `behavior_type=hidden` → `behavior_value = "{}"`
  - `behavior_type=view` → `behavior_value = "{}"`
- [ ] 旧 `score` 字段保留读兼容，逐步废弃
- [ ] 验证迁移后无数据丢失

## 11.2 Service 层重写

- [ ] `behavior_service.go`: 删除 `Record(mediaID, behaviorType, score)` 方法
- [ ] `behavior_service.go`: 新增 `Record(mediaID, behaviorType, valueJSON, source)` 方法
- [ ] `behavior_service.go`: 新增 `Favorite(mediaID)` / `Unfavorite(mediaID)` / `Rate(mediaID, rating)` / `Hide(mediaID)` / `View(mediaID)` 命名方法
- [ ] `behavior_service.go`: 旧方法保留为 deprecated，转调新方法

## 11.3 推荐引擎准备（M3 启动前）

- [ ] `interest_profile` 重建时直接消费 `behavior_value` JSON
- [ ] 权重公式按 RecommendationEngine.md v2.1 实现
- [ ] 时间衰减 `e^(-Δdays/30)` 公式落地

---

# 阻塞 / 风险

| 项 | 状态 | 备注 |
|----|------|------|
| Phase 10 schema 变更 | 未启动 | ERD v2.1 已稳定，需要落地 |
| Phase 11 行为表迁移 | 未启动 | M3 启动前必做 |
| 旧 API 兼容期 | 待定 | 建议保留 1 个版本周期 |
| M3 推荐引擎 | 阻塞中 | 等 Phase 10 + 11 完成 |

---

# 验证清单（M2 验收标准）

- [x] 媒体扫描入库 ✅
- [x] 媒体列表分页 ✅
- [x] 媒体详情（含标签/分类/元数据）✅
- [x] 收藏 / 取消收藏 ✅
- [x] 评分（整数 1-5）✅
- [ ] 评分（0.5 步进）⏳ Phase 10
- [ ] 状态机（active/missing/parsing/error/deleted）⏳ Phase 10
- [x] 搜索历史 ✅
- [x] 行为统计 ✅
- [x] 前端媒体详情页 behavior 集成 ✅
- [x] 前端收藏页 ✅
- [x] 前端历史页 ✅
