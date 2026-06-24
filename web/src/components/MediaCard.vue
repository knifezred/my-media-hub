<script setup lang="ts">
import type { Media } from '../types'
import { useRouter } from 'vue-router'

const props = defineProps<{
  media: Media
  showActions?: boolean
}>()

const emit = defineEmits<{
  favorite: [id: number]
  viewed: [id: number]
  hidden: [id: number]
}>()

const router = useRouter()

function goDetail() {
  router.push(`/media/${props.media.id}`)
}

const typeIcon: Record<string, string> = {
  image: '🖼',
  video: '🎬',
  novel: '📖',
}

function formatSize(bytes: number): string {
  if (bytes < 1024 * 1024) return `${(bytes / 1024).toFixed(0)} KB`
  if (bytes < 1024 * 1024 * 1024) return `${(bytes / 1024 / 1024).toFixed(1)} MB`
  return `${(bytes / 1024 / 1024 / 1024).toFixed(2)} GB`
}
</script>

<template>
  <div class="media-card" @click="goDetail">
    <div class="card-cover">
      <div class="cover-placeholder">
        {{ typeIcon[media.media_type] || '📁' }}
      </div>
      <div class="card-type-badge">{{ media.media_type }}</div>
      <div v-if="media.favorite" class="card-favorite-badge">♥</div>
    </div>
    <div class="card-body">
      <h4 class="card-title" :title="media.title">{{ media.title }}</h4>
      <div class="card-meta">
        <span>{{ formatSize(media.size) }}</span>
        <span v-if="media.tags?.length" class="card-tags">{{ media.tags.slice(0, 2).map(t => t.name).join(', ') }}</span>
      </div>
      <div v-if="showActions" class="card-actions" @click.stop>
        <button
          class="card-action-btn"
          :class="{ active: media.favorite }"
          :title="media.favorite ? '取消收藏' : '收藏'"
          @click="emit('favorite', media.id)"
        >♥</button>
        <button
          class="card-action-btn"
          :class="{ active: media.viewed }"
          :title="media.viewed ? '已看' : '标记已看'"
          @click="emit('viewed', media.id)"
        >✓</button>
        <button
          class="card-action-btn"
          :class="{ active: media.hidden }"
          title="不感兴趣"
          @click="emit('hidden', media.id)"
        >✕</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.media-card {
  border: 1px solid var(--border);
  border-radius: 10px;
  overflow: hidden;
  cursor: pointer;
  transition: border-color 0.2s, transform 0.2s;
  background: var(--bg-hover);
}

.media-card:hover {
  border-color: var(--accent);
  transform: translateY(-2px);
}

.card-cover {
  position: relative;
  aspect-ratio: 16 / 9;
  background: var(--bg-input);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.cover-placeholder {
  font-size: 48px;
  opacity: 0.5;
}

.card-type-badge {
  position: absolute;
  top: 8px;
  left: 8px;
  padding: 2px 8px;
  border-radius: 4px;
  background: rgba(0,0,0,0.6);
  color: #fff;
  font-size: 11px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.card-favorite-badge {
  position: absolute;
  top: 8px;
  right: 8px;
  color: var(--accent);
  font-size: 18px;
  text-shadow: 0 1px 4px rgba(0,0,0,0.5);
}

.card-body {
  padding: 12px;
}

.card-title {
  font-size: 14px;
  font-weight: 500;
  margin: 0 0 6px;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-meta {
  display: flex;
  gap: 8px;
  font-size: 12px;
  color: var(--text-secondary);
}

.card-tags {
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.card-actions {
  display: flex;
  gap: 4px;
  margin-top: 10px;
  padding-top: 10px;
  border-top: 1px solid var(--border);
}

.card-action-btn {
  flex: 1;
  padding: 6px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  font-size: 14px;
  transition: all 0.15s;
}

.card-action-btn:hover {
  background: var(--bg-active);
  border-color: var(--accent);
}

.card-action-btn.active {
  color: var(--accent);
  border-color: var(--accent);
}
</style>
