<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchViewedPage } from '../api/viewed'
import { fetchSearchHistory, deleteSearchHistory, clearSearchHistory } from '../api/history'
import type { Media, SearchHistory } from '../types'
import MediaGrid from '../components/MediaGrid.vue'
import Pagination from '../components/Pagination.vue'

const viewedItems = ref<Media[]>([])
const viewedTotal = ref(0)
const viewedPage = ref(1)
const viewedLoading = ref(false)

const searchHistory = ref<SearchHistory[]>([])
const historyLoading = ref(false)

const activeTab = ref<'viewed' | 'search'>('viewed')
const pageSize = 20

async function loadViewed() {
  viewedLoading.value = true
  try {
    const data = await fetchViewedPage({ page: viewedPage.value, page_size: pageSize })
    viewedItems.value = data.items
    viewedTotal.value = data.total
  } finally {
    viewedLoading.value = false
  }
}

async function loadSearchHistory() {
  historyLoading.value = true
  try {
    const data = await fetchSearchHistory({ page: 1, page_size: 50 })
    searchHistory.value = data.items
  } finally {
    historyLoading.value = false
  }
}

async function handleDeleteHistory(id: number) {
  try {
    await deleteSearchHistory(id)
    searchHistory.value = searchHistory.value.filter(h => h.id !== id)
  } catch {}
}

async function handleClearHistory() {
  try {
    await clearSearchHistory()
    searchHistory.value = []
  } catch {}
}

onMounted(() => {
  loadViewed()
  loadSearchHistory()
})
</script>

<template>
  <div class="history-page">
    <div class="tabs">
      <button :class="['tab', { active: activeTab === 'viewed' }]" @click="activeTab = 'viewed'">浏览历史</button>
      <button :class="['tab', { active: activeTab === 'search' }]" @click="activeTab = 'search'">搜索历史</button>
    </div>

    <div v-if="activeTab === 'viewed'">
      <div v-if="viewedItems.length">
        <MediaGrid :items="viewedItems" :loading="viewedLoading" show-actions />
        <Pagination
          v-if="viewedTotal > pageSize"
          :page="viewedPage"
          :page-size="pageSize"
          :total="viewedTotal"
          @change="viewedPage = $event; loadViewed()"
        />
      </div>
      <div v-else-if="!viewedLoading" class="empty-state">
        <p>暂无浏览记录</p>
      </div>
    </div>

    <div v-else>
      <div class="history-header">
        <span class="history-count">{{ searchHistory.length }} 条搜索记录</span>
        <button v-if="searchHistory.length" class="clear-btn" @click="handleClearHistory">清空全部</button>
      </div>

      <div v-if="searchHistory.length" class="search-history-list">
        <div v-for="h in searchHistory" :key="h.id" class="search-history-item">
          <router-link :to="`/search?q=${encodeURIComponent(h.keyword)}`" class="history-keyword">
            {{ h.keyword }}
          </router-link>
          <span class="history-results">{{ h.result_count }} 条结果</span>
          <span class="history-time">{{ new Date(h.created_at).toLocaleDateString('zh-CN') }}</span>
          <button class="history-delete" @click="handleDeleteHistory(h.id)">✕</button>
        </div>
      </div>

      <div v-else-if="!historyLoading" class="empty-state">
        <p>暂无搜索记录</p>
      </div>
    </div>
  </div>
</template>

<style scoped>
.history-page {
  max-width: 900px;
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

.history-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 16px;
}

.history-count {
  font-size: 14px;
  color: var(--text-secondary);
}

.clear-btn {
  padding: 6px 16px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  font-size: 13px;
}

.clear-btn:hover {
  border-color: var(--accent);
  color: var(--accent);
}

.search-history-list {
  display: flex;
  flex-direction: column;
  gap: 0;
  border: 1px solid var(--border);
  border-radius: 8px;
  overflow: hidden;
}

.search-history-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-bottom: 1px solid var(--border);
  font-size: 14px;
}

.search-history-item:last-child {
  border-bottom: none;
}

.history-keyword {
  flex: 1;
  color: var(--accent);
  text-decoration: none;
}

.history-keyword:hover {
  text-decoration: underline;
}

.history-results {
  color: var(--text-secondary);
  font-size: 13px;
  white-space: nowrap;
}

.history-time {
  color: var(--text-secondary);
  font-size: 13px;
  white-space: nowrap;
  min-width: 80px;
  text-align: right;
}

.history-delete {
  width: 28px;
  height: 28px;
  border: none;
  border-radius: 4px;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  font-size: 14px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.history-delete:hover {
  background: var(--bg-hover);
  color: var(--accent);
}

.empty-state {
  text-align: center;
  padding: 80px 20px;
  color: var(--text-secondary);
}
</style>
