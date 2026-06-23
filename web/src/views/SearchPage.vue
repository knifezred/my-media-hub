<script setup lang="ts">
import { ref } from 'vue'
import { useSearchStore } from '../stores/search'

const store = useSearchStore()
const keyword = ref('')
const mediaType = ref('')

function doSearch() {
  if (!keyword.value.trim()) return
  store.search({ keyword: keyword.value, media_type: mediaType.value || undefined })
}

function onInput() {
  store.fetchSuggestions(keyword.value)
}
</script>

<template>
  <div class="search-page">
    <div class="search-bar">
      <input
        v-model="keyword"
        type="text"
        placeholder="搜索标题、标签、作者..."
        class="search-input"
        @input="onInput"
        @keyup.enter="doSearch"
      />
      <button class="search-btn" @click="doSearch">搜索</button>
    </div>

    <div v-if="store.suggestions.length" class="suggestions">
      <div
        v-for="s in store.suggestions"
        :key="s"
        class="suggestion-item"
        @click="keyword = s; doSearch()"
      >
        {{ s }}
      </div>
    </div>

    <div class="filters">
      <select v-model="mediaType" class="filter-select">
        <option value="">全部类型</option>
        <option value="image">图片</option>
        <option value="video">视频</option>
        <option value="novel">小说</option>
      </select>
    </div>

    <div v-if="store.loading" class="loading">搜索中...</div>

    <div v-else-if="store.results.length" class="results">
      <div v-for="item in store.results" :key="item.id" class="result-item">
        <router-link :to="`/media/${item.id}`" class="result-title">{{ item.title }}</router-link>
        <div class="result-meta">
          <span class="type-badge">{{ item.media_type }}</span>
          <span>{{ item.tags.map(t => t.name).join(', ') }}</span>
        </div>
      </div>
      <div class="result-info">共 {{ store.total }} 条结果</div>
    </div>

    <div v-else-if="keyword" class="empty-state">
      未找到相关资源
    </div>

    <div v-else class="empty-state hint">
      输入关键词开始搜索
    </div>
  </div>
</template>

<style scoped>
.search-page {
  max-width: 800px;
  margin: 0 auto;
}

.search-bar {
  display: flex;
  gap: 8px;
  margin-bottom: 16px;
}

.search-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: var(--bg-input);
  color: var(--text);
  font-size: 16px;
}

.search-input:focus {
  outline: none;
  border-color: var(--accent);
}

.search-btn {
  padding: 12px 24px;
  border: none;
  border-radius: 8px;
  background: var(--accent);
  color: #fff;
  font-size: 14px;
  cursor: pointer;
}

.suggestions {
  border: 1px solid var(--border);
  border-radius: 8px;
  margin-bottom: 16px;
  overflow: hidden;
}

.suggestion-item {
  padding: 10px 16px;
  cursor: pointer;
  color: var(--text);
}

.suggestion-item:hover {
  background: var(--bg-hover);
}

.filters {
  margin-bottom: 16px;
}

.filter-select {
  padding: 8px 12px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: var(--bg-input);
  color: var(--text);
}

.results {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.result-item {
  padding: 16px;
  border: 1px solid var(--border);
  border-radius: 8px;
}

.result-title {
  font-size: 16px;
  font-weight: 500;
  color: var(--accent);
  text-decoration: none;
}

.result-title:hover {
  text-decoration: underline;
}

.result-meta {
  display: flex;
  gap: 12px;
  margin-top: 8px;
  font-size: 13px;
  color: var(--text-secondary);
}

.type-badge {
  padding: 2px 8px;
  border-radius: 4px;
  background: var(--bg-hover);
  font-size: 12px;
}

.result-info {
  padding: 12px;
  text-align: center;
  color: var(--text-secondary);
  font-size: 14px;
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

.hint {
  opacity: 0.6;
}
</style>
