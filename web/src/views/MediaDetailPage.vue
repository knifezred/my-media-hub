<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchMediaDetail } from '../api/media'
import { addFavorite, removeFavorite } from '../api/favorite'
import { rateMedia } from '../api/rating'
import { markViewed } from '../api/viewed'
import { hideMedia, unhideMedia } from '../api/hidden'
import type { MediaDetail } from '../types'
import StarRating from '../components/StarRating.vue'
import TagChip from '../components/TagChip.vue'

const route = useRoute()
const router = useRouter()
const media = ref<MediaDetail | null>(null)
const loading = ref(true)
const localRating = ref(0)

onMounted(async () => {
  try {
    const data = await fetchMediaDetail(Number(route.params.id))
    media.value = data
    localRating.value = data.rating
  } finally {
    loading.value = false
  }
})

async function toggleFavorite() {
  if (!media.value) return
  try {
    if (media.value.favorite) {
      await removeFavorite(media.value.id)
    } else {
      await addFavorite(media.value.id)
    }
    media.value.favorite = !media.value.favorite
  } catch {}
}

async function onRate(rating: number) {
  if (!media.value) return
  try {
    await rateMedia(media.value.id, rating)
    media.value.rating = rating
    localRating.value = rating
  } catch {}
}

async function toggleViewed() {
  if (!media.value) return
  try {
    await markViewed(media.value.id)
    media.value.viewed = true
  } catch {}
}

async function toggleHidden() {
  if (!media.value) return
  try {
    if (media.value.hidden) {
      await unhideMedia(media.value.id)
    } else {
      await hideMedia(media.value.id)
    }
    media.value.hidden = !media.value.hidden
  } catch {}
}

function formatSize(bytes: number): string {
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(0)} KB`
  if (bytes < 1024 * 1024 * 1024) return `${(bytes / 1024 / 1024).toFixed(1)} MB`
  return `${(bytes / 1024 / 1024 / 1024).toFixed(2)} GB`
}

function formatDate(d: string): string {
  return new Date(d).toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' })
}
</script>

<template>
  <div v-if="loading" class="loading">加载中...</div>

  <div v-else-if="media" class="detail-page">
    <button class="back-btn" @click="router.back()">← 返回</button>

    <div class="detail-layout">
      <div class="cover-section">
        <div class="cover-placeholder">
          <span class="cover-icon">{{ media.media_type === 'video' ? '🎬' : media.media_type === 'image' ? '🖼' : '📖' }}</span>
        </div>
        <span class="type-label">{{ media.media_type }}</span>
      </div>

      <div class="info-section">
        <h1 class="title">{{ media.title }}</h1>

        <p v-if="media.description" class="description">{{ media.description }}</p>

        <div class="meta-strip">
          <div class="meta-cell">
            <span class="meta-cell-label">文件大小</span>
            <span class="meta-cell-value">{{ formatSize(media.size) }}</span>
          </div>
          <div class="meta-cell">
            <span class="meta-cell-label">添加时间</span>
            <span class="meta-cell-value">{{ formatDate(media.created_at) }}</span>
          </div>
        </div>

        <div v-if="media.tags?.length" class="info-block">
          <h4 class="info-block-title">标签</h4>
          <TagChip :tags="media.tags" />
        </div>

        <div v-if="media.categories?.length" class="info-block">
          <h4 class="info-block-title">分类</h4>
          <TagChip :tags="media.categories.map(c => ({ id: c.id, name: c.name }))" />
        </div>

        <div v-if="Object.keys(media.metadata).length" class="info-block">
          <h4 class="info-block-title">元数据</h4>
          <div class="meta-grid">
            <div v-for="(val, key) in media.metadata" :key="key" class="meta-item">
              <span class="meta-label">{{ key }}</span>
              <span class="meta-value">{{ val }}</span>
            </div>
          </div>
        </div>

        <div class="action-bar">
          <button class="action-btn big" :class="{ active: media.favorite }" @click="toggleFavorite">
            {{ media.favorite ? '♥ 已收藏' : '♡ 收藏' }}
          </button>
          <button class="action-btn big" :class="{ active: media.viewed }" @click="toggleViewed">
            {{ media.viewed ? '✓ 已看' : '○ 标记已看' }}
          </button>
          <button class="action-btn big" :class="{ active: media.hidden }" @click="toggleHidden">
            {{ media.hidden ? '✕ 已隐藏' : '✕ 不感兴趣' }}
          </button>
        </div>

        <div class="rating-section">
          <span class="rating-label">评分：</span>
          <StarRating v-model="localRating" @update:model-value="onRate" />
        </div>
      </div>
    </div>
  </div>

  <div v-else class="not-found">资源不存在</div>
</template>

<style scoped>
.detail-page {
  max-width: 1000px;
  margin: 0 auto;
}

.back-btn {
  padding: 8px 16px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  font-size: 14px;
  margin-bottom: 24px;
}

.back-btn:hover {
  color: var(--text);
  border-color: var(--text-secondary);
}

.detail-layout {
  display: flex;
  gap: 40px;
}

.cover-section {
  flex-shrink: 0;
  width: 320px;
}

.cover-placeholder {
  width: 100%;
  aspect-ratio: 2 / 3;
  background: var(--bg-hover);
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid var(--border);
}

.cover-icon {
  font-size: 72px;
  opacity: 0.6;
}

.type-label {
  display: inline-block;
  margin-top: 8px;
  padding: 4px 12px;
  border-radius: 6px;
  background: var(--bg-hover);
  font-size: 13px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.info-section {
  flex: 1;
  min-width: 0;
}

.title {
  font-size: 28px;
  font-weight: 700;
  margin: 0 0 12px;
  line-height: 1.3;
}

.description {
  color: var(--text-secondary);
  line-height: 1.6;
  margin-bottom: 24px;
  font-size: 15px;
}

.meta-strip {
  display: flex;
  gap: 16px;
  margin-bottom: 24px;
}

.meta-cell {
  padding: 12px 16px;
  border: 1px solid var(--border);
  border-radius: 8px;
  min-width: 140px;
}

.meta-cell-label {
  display: block;
  font-size: 12px;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

.meta-cell-value {
  font-size: 14px;
  font-weight: 500;
}

.info-block {
  margin-bottom: 20px;
}

.info-block-title {
  font-size: 14px;
  font-weight: 600;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.meta-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(180px, 1fr));
  gap: 8px;
}

.meta-item {
  padding: 8px 12px;
  border: 1px solid var(--border);
  border-radius: 6px;
}

.meta-label {
  display: block;
  font-size: 11px;
  color: var(--text-secondary);
  text-transform: uppercase;
  letter-spacing: 0.5px;
  margin-bottom: 2px;
}

.meta-value {
  font-size: 13px;
}

.action-bar {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
  flex-wrap: wrap;
}

.action-btn {
  padding: 10px 20px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: transparent;
  color: var(--text);
  cursor: pointer;
  font-size: 14px;
  transition: all 0.15s;
}

.action-btn:hover {
  background: var(--bg-hover);
}

.action-btn.active {
  border-color: var(--accent);
  color: var(--accent);
}

.rating-section {
  display: flex;
  align-items: center;
  gap: 8px;
}

.rating-label {
  font-size: 14px;
  color: var(--text-secondary);
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
