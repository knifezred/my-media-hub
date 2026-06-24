# My Media Hub - 产品路线图

版本：2.0

状态：Draft

---

# 产品愿景

让用户：

无需搜索

即可发现值得消费的内容。

最终将 NAS 从：

资源仓库

变成：

内容发现平台

---

# Milestone 1

基础媒体库

目标：

建立统一媒体模型。

完成：

* 项目框架
* SQLite
* Media模块
* Scanner模块
* Tag模块
* Category模块
* API规范
* Vue前端框架

交付标准：

能够扫描资源并展示媒体列表。

---

# Milestone 2

用户行为系统

目标：

建立行为驱动体系。

完成：

* 收藏
* 评分
* 已看
* 不感兴趣
* 搜索历史
* 点击历史

新增：

* UserBehavior

交付标准：

能够完整记录用户行为。

---

# Milestone 3

推荐引擎 V1

目标：

构建推荐系统基础能力。

完成：

* Random
* Recent
* Favorite Based
* Tag Based
* Category Based

新增：

* Recommendation Service
* Recommendation Cache

交付标准：

首页可生成推荐内容。

---

# Milestone 4

发现流

目标：

首页成为发现入口。

完成：

* 今日推荐
* 最近新增
* 随机探索
* 发现流缓存

新增：

* Discovery Feed Service
* Discovery Feed Log

交付标准：

用户打开首页即可看到推荐内容。

---

# Milestone 5

兴趣画像

目标：

建立用户兴趣模型。

完成：

* Interest Profile
* Tag Weight
* Category Weight

新增：

* Interest Profile Service

交付标准：

推荐结果开始个性化。

---

# Milestone 6

相似推荐

目标：

实现内容关联发现。

完成：

* Similar Media
* Similar Tag
* Similar Category

交付标准：

每个媒体详情页均可推荐相关内容。

---

# Milestone 7

Rediscover

目标：

重新发现被遗忘资源。

完成：

* 长期未访问资源
* 收藏未观看资源
* 历史高评分资源

交付标准：

持续挖掘沉睡资源。

---

# Milestone 8

跨媒体推荐

目标：

统一内容发现。

完成：

* 图片推荐小说
* 小说推荐视频
* 视频推荐图片

交付标准：

打通不同媒体类型。

---

# Milestone 9

搜索增强

目标：

优化搜索体验。

完成：

* Search Strategy
* LIKE
* FTS
* Search Suggestion

可选：

* Bleve

说明：

搜索属于辅助能力。

优先级低于推荐系统。

---

# Milestone 10

AI Discovery

目标：

最终发现引擎。

完成：

* AI兴趣分析
* AI标签生成
* AI内容关联
* AI推荐解释

交付标准：

形成真正的个人内容发现平台。
