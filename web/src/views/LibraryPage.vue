<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchMediaPage } from '../api/media'
import type { Media } from '../types'
import MediaGrid from '../components/MediaGrid.vue'
import Pagination from '../components/Pagination.vue'

const items = ref<Media[]>([])
const total = ref(0)
const page = ref(1)
const loading = ref(false)
const mediaType = ref('')
const sort = ref('created_at')

const pageSize = 20

async function loadData() {
  loading.value = true
  try {
    const data = await fetchMediaPage({
      page: page.value,
      page_size: pageSize,
      media_type: mediaType.value || undefined,
      sort: sort.value,
    })
    items.value = data.items
    total.value = data.total
  } finally {
    loading.value = false
  }
}

function onFilterChange() {
  page.value = 1
  loadData()
}

onMounted(loadData)
</script>

<template>
  <div class="library-page">
    <div class="toolbar">
      <div class="tabs">
        <button
          :class="['tab', { active: mediaType === '' }]"
          @click="mediaType = ''; onFilterChange()"
        >全部</button>
        <button
          :class="['tab', { active: mediaType === 'image' }]"
          @click="mediaType = 'image'; onFilterChange()"
        >图片</button>
        <button
          :class="['tab', { active: mediaType === 'video' }]"
          @click="mediaType = 'video'; onFilterChange()"
        >视频</button>
        <button
          :class="['tab', { active: mediaType === 'novel' }]"
          @click="mediaType = 'novel'; onFilterChange()"
        >小说</button>
      </div>
      <select v-model="sort" class="sort-select" @change="onFilterChange()">
        <option value="created_at">最新</option>
        <option value="title">标题</option>
        <option value="size">大小</option>
        <option value="random">随机</option>
      </select>
    </div>

    <MediaGrid :items="items" :loading="loading" show-actions />

    <Pagination
      v-if="total > pageSize"
      :page="page"
      :page-size="pageSize"
      :total="total"
      @change="page = $event; loadData()"
    />
  </div>
</template>

<style scoped>
.library-page {
  max-width: 1200px;
  margin: 0 auto;
}

.toolbar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 24px;
  flex-wrap: wrap;
  gap: 12px;
}

.tabs {
  display: flex;
  gap: 4px;
  background: var(--bg-hover);
  padding: 4px;
  border-radius: 8px;
  width: fit-content;
}

.tab {
  padding: 8px 20px;
  border: none;
  border-radius: 6px;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  font-size: 14px;
}

.tab.active {
  background: var(--bg-active);
  color: var(--accent);
}

.sort-select {
  padding: 8px 12px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: var(--bg-input);
  color: var(--text);
  font-size: 14px;
}

.sort-select:focus {
  outline: none;
  border-color: var(--accent);
}
</style>
