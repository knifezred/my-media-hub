<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchStatsOverview } from '../api/stats'
import type { StatsOverview } from '../types'

const stats = ref<StatsOverview | null>(null)

onMounted(async () => {
  try {
    stats.value = await fetchStatsOverview()
  } catch {
    // API not ready yet
  }
})
</script>

<template>
  <div class="stats-page">
    <div v-if="stats" class="stats-grid">
      <div class="stat-card">
        <div class="stat-value">{{ stats.total_media.toLocaleString() }}</div>
        <div class="stat-label">媒体总数</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.total_images.toLocaleString() }}</div>
        <div class="stat-label">图片</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.total_videos.toLocaleString() }}</div>
        <div class="stat-label">视频</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.total_novels.toLocaleString() }}</div>
        <div class="stat-label">小说</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.favorite_count.toLocaleString() }}</div>
        <div class="stat-label">收藏</div>
      </div>
      <div class="stat-card">
        <div class="stat-value">{{ stats.viewed_count.toLocaleString() }}</div>
        <div class="stat-label">已看</div>
      </div>
    </div>

    <div v-else class="empty-state">
      <p>暂无统计数据</p>
    </div>
  </div>
</template>

<style scoped>
.stats-page {
  max-width: 900px;
  margin: 0 auto;
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.stat-card {
  padding: 24px;
  border: 1px solid var(--border);
  border-radius: 12px;
  text-align: center;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: var(--accent);
  margin-bottom: 8px;
}

.stat-label {
  font-size: 14px;
  color: var(--text-secondary);
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}
</style>
