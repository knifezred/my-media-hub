<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchMediaDetail } from '../api/media'
import type { MediaDetail } from '../types'

const route = useRoute()
const media = ref<MediaDetail | null>(null)
const loading = ref(true)

onMounted(async () => {
  try {
    media.value = await fetchMediaDetail(Number(route.params.id))
  } finally {
    loading.value = false
  }
})
</script>

<template>
  <div v-if="loading" class="loading">加载中...</div>

  <div v-else-if="media" class="detail-page">
    <div class="cover-section">
      <div class="cover-placeholder">
        {{ media.media_type === 'video' ? '🎬' : media.media_type === 'image' ? '🖼' : '📖' }}
      </div>
    </div>

    <div class="info-section">
      <h1 class="title">{{ media.title }}</h1>
      <p v-if="media.description" class="description">{{ media.description }}</p>

      <div class="tags">
        <span v-for="tag in media.tags" :key="tag.id" class="tag">{{ tag.name }}</span>
      </div>

      <div class="meta-grid">
        <div class="meta-item">
          <span class="meta-label">类型</span>
          <span class="meta-value">{{ media.media_type }}</span>
        </div>
        <div class="meta-item">
          <span class="meta-label">大小</span>
          <span class="meta-value">{{ (media.size / 1024 / 1024).toFixed(1) }} MB</span>
        </div>
      </div>

      <div class="actions">
        <button class="action-btn" :class="{ active: media.favorite }">
          {{ media.favorite ? '♥ 已收藏' : '♡ 收藏' }}
        </button>
        <button class="action-btn" :class="{ active: media.viewed }">
          {{ media.viewed ? '✓ 已看' : '○ 标记已看' }}
        </button>
      </div>
    </div>
  </div>

  <div v-else class="not-found">
    资源不存在
  </div>
</template>

<style scoped>
.detail-page {
  max-width: 900px;
  margin: 0 auto;
  display: flex;
  gap: 32px;
}

.cover-section {
  flex-shrink: 0;
  width: 300px;
}

.cover-placeholder {
  width: 100%;
  aspect-ratio: 2/3;
  background: var(--bg-hover);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 64px;
}

.info-section {
  flex: 1;
}

.title {
  font-size: 24px;
  font-weight: 600;
  margin: 0 0 12px;
}

.description {
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 20px;
}

.tags {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
  margin-bottom: 20px;
}

.tag {
  padding: 4px 12px;
  border-radius: 16px;
  background: var(--bg-hover);
  font-size: 13px;
  color: var(--text-secondary);
}

.meta-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 12px;
  margin-bottom: 24px;
}

.meta-item {
  padding: 12px;
  border-radius: 8px;
  background: var(--bg-hover);
}

.meta-label {
  display: block;
  font-size: 12px;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

.meta-value {
  font-size: 14px;
  font-weight: 500;
}

.actions {
  display: flex;
  gap: 12px;
}

.action-btn {
  padding: 10px 20px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: transparent;
  color: var(--text);
  cursor: pointer;
  font-size: 14px;
}

.action-btn.active {
  border-color: var(--accent);
  color: var(--accent);
}

.action-btn:hover {
  background: var(--bg-hover);
}

.loading {
  text-align: center;
  padding: 80px;
  color: var(--text-secondary);
}

.not-found {
  text-align: center;
  padding: 80px;
  color: var(--text-secondary);
}
</style>
