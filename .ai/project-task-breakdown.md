# Task Breakdown

AI-Owned File

Governed by:

.ai/project-governance.md

AI agents MAY generate and update tasks within the active milestone.

Milestone definitions come from:

docs/Roadmap.md

Task status should remain synchronized with:

.ai/project-current-milestone.md

---

# Active Milestone

Milestone 1: 可搜索媒体库

---

# Phase 1: 项目骨架搭建

## P1.1 后端项目初始化

- [ ] 创建 backend/ 目录结构
- [ ] 初始化 Go module (go.mod)
- [ ] 引入依赖：Gin, go-sqlite3, etc.
- [ ] 创建 cmd/main.go 入口文件
- [ ] 创建 config 包（配置管理）
- [ ] 创建 database 包（SQLite 连接初始化 + 自动迁移）
- [ ] 创建 router 包（Gin 路由注册）
- [ ] 创建 response 包（统一响应格式：code/message/data）
- [ ] 创建 middleware 包（错误恢复、日志）
- [ ] 创建 errorcode 包（集中管理错误码常量）
- [ ] 验证：应用可启动

## P1.2 前端项目初始化

- [ ] 使用 Vite 创建 Vue3 + TypeScript 项目
- [ ] 配置文件：tsconfig, vite.config
- [ ] 引入 Pinia 状态管理
- [ ] 引入 Vue Router
- [ ] 创建目录结构：views/, components/, stores/, api/, types/
- [ ] 实现基础布局组件（Header + Sidebar + Content）
- [ ] 实现路由配置（所有页面路由占位）
- [ ] 实现 API 请求封装（基于 fetch/axios，统一错误处理）
- [ ] 实现主题切换（Dark/Light）
- [ ] 验证：前端可运行，页面可导航

---

# Phase 2: 数据库 Schema 实现

## P2.1 数据库模型定义

- [ ] 创建 model/ 包
- [ ] 定义 Media 模型（id, media_type, title, description, path, hash, size, cover_path, timestamps）
- [ ] 定义 Tag 模型（id, name）
- [ ] 定义 MediaTag 模型（media_id, tag_id）
- [ ] 定义 Category 模型（id, name, parent_id）
- [ ] 定义 MediaCategory 模型（media_id, category_id）
- [ ] 定义 MediaMetadata 模型（id, media_id, key, value）
- [ ] 定义 UserFavorite 模型
- [ ] 定义 UserRating 模型
- [ ] 定义 UserViewed 模型
- [ ] 定义 UserHidden 模型
- [ ] 定义 SearchHistory 模型（id, keyword, result_count）
- [ ] 定义 SearchClickHistory 模型（id, keyword, media_id, position）
- [ ] 定义 RecommendationCache 模型（id, recommendation_type, media_id, score）

## P2.2 数据库初始化

- [ ] 实现 AutoMigrate 逻辑
- [ ] 创建 FTS5 虚拟表 media_fts
- [ ] 创建 FTS5 触发器（同步 media 表到 media_fts）
- [ ] 创建所有必要索引
- [ ] 初始化种子数据（可选）

---

# Phase 3: Media 模块

## P3.1 Repository

- [ ] media_repository.go: CRUD + 列表查询（分页、按类型过滤）
- [ ] tag_repository.go: 标签创建、查询
- [ ] category_repository.go: 分类创建、查询
- [ ] media_tag_repository.go: 关联查询
- [ ] media_category_repository.go: 关联查询
- [ ] media_metadata_repository.go: 元数据读写

## P3.2 Service

- [ ] media_service.go: 媒体业务逻辑
- [ ] tag_service.go: 标签业务逻辑
- [ ] category_service.go: 分类业务逻辑

## P3.3 API

- [ ] POST /api/v1/media/page - 媒体列表（分页、类型过滤）
- [ ] GET /api/v1/media/{id} - 媒体详情（含 tags, categories, metadata, 用户行为状态）
- [ ] POST /api/v1/tags/page - 标签列表
- [ ] GET /api/v1/tags/{id} - 标签详情
- [ ] POST /api/v1/categories/page - 分类列表

---

# Phase 4: Search 模块

## P4.1 Repository

- [ ] 实现 FTS5 全文搜索查询（MATCH）
- [ ] 实现 FTS5 搜索建议查询（前缀匹配）
- [ ] 实现按媒体类型/标签/分类的组合搜索

## P4.2 Service

- [ ] search_service.go: 搜索业务逻辑
- [ ] 搜索历史记录自动插入

## P4.3 API

- [ ] POST /api/v1/search/page - 全文搜索
- [ ] GET /api/v1/search/suggestions - 搜索建议

---

# Phase 5: User 行为模块

## P5.1 Favorite

- [ ] Repository: 增删查
- [ ] Service: 收藏/取消收藏逻辑
- [ ] API: POST /api/v1/favorites, DELETE /api/v1/favorites/{media_id}, POST /api/v1/favorites/page

## P5.2 Rating

- [ ] Repository: 增改查
- [ ] Service: 评分业务逻辑
- [ ] API: POST /api/v1/ratings

## P5.3 Viewed

- [ ] Repository: 增查
- [ ] Service: 标记已看逻辑
- [ ] API: POST /api/v1/viewed, POST /api/v1/viewed/page

## P5.4 Hidden

- [ ] Repository: 增删查
- [ ] Service: 隐藏/取消隐藏逻辑
- [ ] API: POST /api/v1/hidden, DELETE /api/v1/hidden/{media_id}

## P5.5 History

- [ ] Repository: 搜索历史增删查
- [ ] Service: 搜索历史业务
- [ ] API: POST /api/v1/search/history/page, DELETE /api/v1/search/history/{id}, DELETE /api/v1/search/history

---

# Phase 6: Scanner 模块

## P6.1 Repository

- [ ] 实现按路径查询（增量扫描判断）
- [ ] 实现按 hash 查询（去重判断）
- [ ] 批量插入实现

## P6.2 Service

- [ ] 文件递归扫描逻辑
- [ ] 文件类型识别（image/video/novel）
- [ ] 文件 hash 计算
- [ ] 元数据提取（基础版：文件名 -> title）
- [ ] 增量扫描逻辑（对比路径/修改时间/大小）
- [ ] 扫描进度管理

## P6.3 API

- [ ] POST /api/v1/scanner/start - 启动扫描
- [ ] GET /api/v1/scanner/status - 扫描状态

---

# Phase 7: Stats 模块

## P7.1 Repository

- [ ] 总览统计数据查询
- [ ] 按 media_type 分组计数

## P7.2 Service

- [ ] stats_service.go: 统计业务

## P7.3 API

- [ ] GET /api/v1/stats/overview - 总览统计

---

# Phase 8: Frontend - API 层与 Store

## P8.1 API 封装

- [ ] api/media.ts - 媒体相关 API
- [ ] api/search.ts - 搜索相关 API
- [ ] api/tag.ts - 标签相关 API
- [ ] api/category.ts - 分类相关 API
- [ ] api/favorite.ts - 收藏相关 API
- [ ] api/rating.ts - 评分相关 API
- [ ] api/viewed.ts - 已看相关 API
- [ ] api/hidden.ts - 隐藏相关 API
- [ ] api/history.ts - 搜索历史相关 API
- [ ] api/scanner.ts - 扫描相关 API
- [ ] api/stats.ts - 统计相关 API

## P8.2 Pinia Store

- [ ] stores/media.ts
- [ ] stores/search.ts
- [ ] stores/ui.ts（侧边栏、主题等）

---

# Phase 9: Frontend - 页面实现

## P9.1 通用组件

- [ ] MediaCard.vue - 媒体卡片（封面 + 标题 + 标签 + hover 操作）
- [ ] MediaGrid.vue - 网格展示组件
- [ ] MediaList.vue - 列表展示组件（适用于小说、搜索结果）
- [ ] Pagination.vue - 分页组件
- [ ] TagChip.vue - 标签组件
- [ ] StarRating.vue - 评分组件
- [ ] Sidebar.vue - 侧边栏导航
- [ ] Header.vue - 顶部栏

## P9.2 首页 - 发现页

- [ ] 推荐内容展示区
- [ ] 最近新增展示区
- [ ] 随机发现展示区
- [ ] 热门标签展示区
- [ ] 支持空状态展示

## P9.3 搜索页

- [ ] 搜索输入框（支持即时搜索建议）
- [ ] 搜索建议下拉
- [ ] 筛选器：媒体类型
- [ ] 搜索结果展示（Grid/List 切换）
- [ ] 搜索结果分页

## P9.4 媒体库页

- [ ] 分类 Tab（全部/图片/视频/小说）
- [ ] 标签筛选
- [ ] 排序选择（最新/评分/收藏/随机）
- [ ] 媒体网格/列表展示

## P9.5 媒体详情页

- [ ] 封面展示
- [ ] 标题/描述
- [ ] 标签和分类展示
- [ ] 元数据展示
- [ ] 操作区：收藏、评分（1-5星）、已看、不感兴趣、打开资源
- [ ] 相似推荐区
- [ ] 用户行为状态同步

## P9.6 收藏页

- [ ] 收藏列表展示
- [ ] 取消收藏
- [ ] 筛选、排序

## P9.7 历史页

- [ ] 浏览历史
- [ ] 搜索历史
- [ ] 单条删除、清空全部

## P9.8 统计页

- [ ] 总览统计：媒体总数、各类别数量
- [ ] 收藏数、已看数

## P9.9 设置页

- [ ] 扫描目录配置
- [ ] 启动扫描按钮
- [ ] 扫描状态展示
- [ ] 系统信息

---

# Phase 10: 集成与测试

## P10.1 集成验证

- [ ] 前后端联调
- [ ] 端到端流程测试：扫描 -> 搜索 -> 收藏 -> 评分 -> 已看
- [ ] 性能验证（100ms 搜索延迟目标）

## P10.2 Bug 修复

- [ ] 根据联调结果修复问题
- [ ] 边界情况处理

---

# Task Status Legend

- [ ] 待开始
- [x] 已完成
- [-] 进行中
- [!] 受阻
