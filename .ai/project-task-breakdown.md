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
