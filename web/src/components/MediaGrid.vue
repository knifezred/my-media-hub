<script setup lang="ts">
import type { Media } from '../types'
import MediaCard from './MediaCard.vue'

defineProps<{
  items: Media[]
  loading?: boolean
  showActions?: boolean
}>()

const emit = defineEmits<{
  favorite: [id: number]
  viewed: [id: number]
  hidden: [id: number]
}>()
</script>

<template>
  <div v-if="loading" class="grid-loading">
    <div v-for="i in 8" :key="i" class="skeleton-card">
      <div class="skeleton-cover" />
      <div class="skeleton-body">
        <div class="skeleton-line w-70" />
        <div class="skeleton-line w-50" />
      </div>
    </div>
  </div>
  <div v-else-if="items.length" class="media-grid">
    <MediaCard
      v-for="item in items"
      :key="item.id"
      :media="item"
      :show-actions="showActions"
      @favorite="emit('favorite', $event)"
      @viewed="emit('viewed', $event)"
      @hidden="emit('hidden', $event)"
    />
  </div>
  <div v-else class="grid-empty">
    <p>暂无媒体资源</p>
  </div>
</template>

<style scoped>
.media-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 16px;
}

.grid-loading {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(220px, 1fr));
  gap: 16px;
}

.skeleton-card {
  border: 1px solid var(--border);
  border-radius: 10px;
  overflow: hidden;
}

.skeleton-cover {
  aspect-ratio: 16 / 9;
  background: var(--bg-hover);
}

.skeleton-body {
  padding: 12px;
}

.skeleton-line {
  height: 12px;
  border-radius: 4px;
  background: var(--bg-hover);
  margin-bottom: 8px;
}

.skeleton-line:last-child {
  margin-bottom: 0;
}

.w-70 { width: 70%; }
.w-50 { width: 50%; }

.grid-empty {
  text-align: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}
</style>
