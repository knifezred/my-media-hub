# Current Milestone

AI-Owned File

Governed by:

.ai/project-governance.md

AI agents MAY update this file.

AI agents SHOULD update this file after completing implementation tasks.

---

# Active Milestone

Milestone 1 + Milestone 2: 媒体库基础能力 + 用户行为系统

**变更任务**：ERD v2.0 → v2.1 升级（Phase 10）+ 行为表 score 字段重构（Phase 11）

---

# 来源

docs/Roadmap.md
docs/ERD.md v2.1

说明：

搜索能力（M9）作为辅助能力一并保留。

---

# 状态

M1+M2 主任务已完成
v2.1 升级变更任务待执行（Phase 10 / 11）

---

# M1 交付标准

能够扫描资源并展示媒体列表。

# M2 交付标准

能够完整记录用户行为（收藏 / 评分 / 已看 / 不感兴趣 / 搜索历史）。

# v2.1 升级交付标准

完成 Phase 10 + Phase 11 所有任务，schema 与 v2.1 文档完全对齐。

---

# M1 完成清单

- [x] 项目框架（Go Module + Gin + 分层架构）
- [x] SQLite 数据库初始化
- [x] Media 模块（模型 + 增删改查 + 列表分页）
- [x] Scanner 模块（文件扫描 + 类型识别 + 去重）
- [x] Tag 模块（模型 + 增删查 + 分页）
- [x] Category 模块（模型 + 增删查 + 分页）
- [x] API 规范（统一响应 + 错误码）
- [x] Vue 前端框架（Vite + Vue3 + TS + Pinia + Router）
- [x] Media 详情页（基础信息 + 标签 + 分类 + 元数据）
- [x] Stats 模块（基础统计）
- [x] GetTagByID API

# M2 完成清单

- [x] 统一 user_behavior 表（替换 user_favorite / user_rating / user_viewed / user_hidden）
- [x] user_behavior 增删查 Repository
- [x] BehaviorService（Record + 各行为快捷方法）
- [x] BehaviorAPI（统一 POST /api/v1/behavior + 旧路径兼容）
- [x] search_history Repository + Service + API
- [x] 前端收藏页
- [x] 前端历史页（浏览历史 + 搜索历史）
- [x] 媒体详情页用户操作（收藏 / 评分 / 已看 / 不感兴趣）

# v2.1 升级清单

详见 `.ai/project-task-breakdown.md` Phase 10 / 11

- [x] media.last_error 字段落地
- [x] media.rating INTEGER → REAL
- [x] media_content 复合主键 (media_id, content_type)
- [x] media.status 状态机落地
- [x] tag UNIQUE(name_norm) 约束
- [x] scanner_index.media_id 允许 NULL
- [x] media_behavior.behavior_value 文档化 JSON
- [x] 行为表 score 字段废弃（迁移到 behavior_value JSON）

# M9（辅助保留）

- [x] Search Strategy 接口定义
- [x] LIKE 搜索实现
- [x] FTS5 搜索实现
- [x] Search Suggestions
- [x] 搜索前端

---

# Backend Progress

| Module | Status |
|--------|--------|
| 项目骨架 | 已完成 |
| 数据库 Schema (M1+M2) | 已完成 |
| **数据库 Schema (v2.1 升级)** | **已完成** |
| Media Repository | 已完成 |
| Media Service | 已完成 |
| Media API | 已完成 |
| Tag Repository | 已完成 |
| Tag API | 已完成（含 GetTagByID） |
| Category Repository | 已完成 |
| Scanner Repository | 已完成（scanner_index 表） |
| Scanner Service | 已完成（增量扫描缓存） |
| Scanner API | 已完成 |
| User Behavior (M2) | 已完成（统一 behavior 表 + 旧路径兼容） |
| **User Behavior (v2.1 JSON 化)** | **已完成** |
| Search History | 已完成 |
| Search Strategy | 已完成 |
| LIKE / FTS5 | 已完成 |
| Stats | 已完成 |
| 前端 - 媒体库 | 已完成 |
| 前端 - 搜索 | 已完成 |
| 前端 - 媒体详情 | 已完成 |
| 前端 - 收藏页 | 已完成 |
| 前端 - 历史页 | 已完成 |

---

# Next Actions

1. **M3 推荐引擎开发启动**（v2.1 schema 已就绪）
2. 前端适配（新字段 favorite/favorited_at/rating/rating_at/hidden/hidden_at/view_count/last_viewed_at）
3. API 兼容期维护（旧字段 mapping）
