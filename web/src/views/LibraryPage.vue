<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchMediaPage } from '../api/media'

const loading = ref(false)
const mediaType = ref('')

async function loadData() {
  loading.value = true
  try {
    await fetchMediaPage({ media_type: mediaType.value || undefined })
  } finally {
    loading.value = false
  }
}

onMounted(loadData)
</script>

<template>
  <div class="library-page">
    <div class="tabs">
      <button
        :class="['tab', { active: mediaType === '' }]"
        @click="mediaType = ''; loadData()"
      >全部</button>
      <button
        :class="['tab', { active: mediaType === 'image' }]"
        @click="mediaType = 'image'; loadData()"
      >图片</button>
      <button
        :class="['tab', { active: mediaType === 'video' }]"
        @click="mediaType = 'video'; loadData()"
      >视频</button>
      <button
        :class="['tab', { active: mediaType === 'novel' }]"
        @click="mediaType = 'novel'; loadData()"
      >小说</button>
    </div>

    <div v-if="loading" class="loading">加载中...</div>

    <div v-else class="empty-state">
      <p>暂无媒体资源</p>
    </div>
  </div>
</template>

<style scoped>
.library-page {
  max-width: 1200px;
  margin: 0 auto;
}

.tabs {
  display: flex;
  gap: 4px;
  margin-bottom: 24px;
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

.loading {
  text-align: center;
  padding: 40px;
  color: var(--text-secondary);
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}
</style>
