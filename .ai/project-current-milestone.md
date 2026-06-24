# Current Milestone

AI-Owned File

Governed by:

.ai/project-governance.md

AI agents MAY update this file.

AI agents SHOULD update this file after completing implementation tasks.

---

# Active Milestone

Milestone 1 + Milestone 2: 媒体库基础能力 + 用户行为系统

---

# 来源

docs/Roadmap.md

说明：

搜索能力（M9）作为辅助能力一并保留。

---

# 状态

已完成

---

# M1 交付标准

能够扫描资源并展示媒体列表。

# M2 交付标准

能够完整记录用户行为（收藏 / 评分 / 已看 / 不感兴趣 / 搜索历史）。

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
| 数据库 Schema | 已完成（M1+M2 表 + FTS5） |
| Media Repository | 已完成 |
| Media Service | 已完成 |
| Media API | 已完成 |
| Tag Repository | 已完成 |
| Tag API | 已完成（含 GetTagByID） |
| Category Repository | 已完成 |
| Scanner Repository | 已完成 |
| Scanner Service | 已完成 |
| Scanner API | 已完成 |
| User Behavior | 已完成（统一 behavior 表 + 旧路径兼容） |
| Search History | 已完成 |
| Search Strategy | 已完成 |
| LIKE / FTS5 | 已完成 |
| Stats | 已完成 |
| 前端 - 媒体库 | 已完成 |
| 前端 - 搜索 | 已完成 |
| 前端 - 媒体详情 | 已完成 |
| 前端 - 收藏页 | 待适配 |
| 前端 - 历史页 | 待适配 |

---

# Next Actions

1. 前端适配（收藏页 / 历史页 / 媒体详情 behavior 统一调用）
