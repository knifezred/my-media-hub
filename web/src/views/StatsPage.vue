<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchStatsOverview } from '../api/stats'
import type { StatsOverview } from '../types'

const stats = ref<StatsOverview | null>(null)

onMounted(async () => {
  try {
    stats.value = await fetchStatsOverview()
  } catch {}
})

const statCards = [
  { key: 'total_media' as const, label: '媒体总数', icon: '📦' },
  { key: 'total_images' as const, label: '图片', icon: '🖼' },
  { key: 'total_videos' as const, label: '视频', icon: '🎬' },
  { key: 'total_novels' as const, label: '小说', icon: '📖' },
  { key: 'favorite_count' as const, label: '收藏', icon: '♥' },
  { key: 'viewed_count' as const, label: '已看', icon: '✓' },
]
</script>

<template>
  <div class="stats-page">
    <h3 class="page-title">统计概览</h3>

    <div v-if="stats" class="stats-grid">
      <div v-for="card in statCards" :key="card.key" class="stat-card">
        <div class="stat-icon">{{ card.icon }}</div>
        <div class="stat-value">{{ stats[card.key].toLocaleString() }}</div>
        <div class="stat-label">{{ card.label }}</div>
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

.page-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 24px;
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
  background: var(--bg-hover);
  transition: border-color 0.2s;
}

.stat-card:hover {
  border-color: var(--accent);
}

.stat-icon {
  font-size: 24px;
  margin-bottom: 8px;
}

.stat-value {
  font-size: 28px;
  font-weight: 700;
  color: var(--accent);
  margin-bottom: 4px;
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
