<script setup lang="ts">
import { useUiStore } from '../stores/ui'

const ui = useUiStore()
</script>

<template>
  <header class="header">
    <div class="header-left">
      <h2 class="page-title">
        <slot name="title">发现</slot>
      </h2>
    </div>
    <div class="header-right">
      <button class="theme-btn" @click="ui.toggleTheme" :title="ui.theme === 'dark' ? '切换亮色' : '切换暗色'">
        {{ ui.theme === 'dark' ? '☀' : '☾' }}
      </button>
    </div>
  </header>
</template>

<style scoped>
.header {
  height: 56px;
  border-bottom: 1px solid var(--border);
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 24px;
  flex-shrink: 0;
  background: var(--bg-header);
}

.page-title {
  font-size: 18px;
  font-weight: 600;
  margin: 0;
}

.header-right {
  display: flex;
  align-items: center;
  gap: 12px;
}

.theme-btn {
  width: 36px;
  height: 36px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: transparent;
  cursor: pointer;
  font-size: 18px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--text);
}

.theme-btn:hover {
  background: var(--bg-hover);
}
</style>
