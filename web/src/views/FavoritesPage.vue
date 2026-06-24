<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchFavoritePage, removeFavorite } from '../api/favorite'
import type { Media } from '../types'
import MediaGrid from '../components/MediaGrid.vue'
import Pagination from '../components/Pagination.vue'

const items = ref<Media[]>([])
const total = ref(0)
const page = ref(1)
const loading = ref(false)
const pageSize = 20

async function loadData() {
  loading.value = true
  try {
    const data = await fetchFavoritePage({ page: page.value, page_size: pageSize })
    items.value = data.items
    total.value = data.total
  } finally {
    loading.value = false
  }
}

async function handleRemoveFavorite(id: number) {
  try {
    await removeFavorite(id)
    items.value = items.value.filter(i => i.id !== id)
    total.value--
  } catch {}
}

onMounted(loadData)
</script>

<template>
  <div class="favorites-page">
    <h3 class="page-title">收藏</h3>

    <div v-if="items.length" class="favorites-content">
      <MediaGrid :items="items" :loading="loading" @favorite="handleRemoveFavorite" show-actions />
      <Pagination
        v-if="total > pageSize"
        :page="page"
        :page-size="pageSize"
        :total="total"
        @change="page = $event; loadData()"
      />
      <p class="count-info">共 {{ total }} 个收藏</p>
    </div>

    <div v-else-if="!loading" class="empty-state">
      <p>暂无收藏</p>
      <p class="hint">浏览资源时点击收藏按钮即可添加</p>
    </div>
  </div>
</template>

<style scoped>
.favorites-page {
  max-width: 1200px;
  margin: 0 auto;
}

.page-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 16px;
}

.count-info {
  text-align: center;
  padding: 8px;
  color: var(--text-secondary);
  font-size: 14px;
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}

.hint {
  font-size: 14px;
  margin-top: 8px;
  opacity: 0.7;
}
</style>
