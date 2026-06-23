<script setup lang="ts">
import { useUiStore } from '../stores/ui'
import { useRoute } from 'vue-router'

const ui = useUiStore()
const route = useRoute()

const navItems = [
  { name: 'discovery', label: '发现', icon: '✦', path: '/' },
  { name: 'search', label: '搜索', icon: '⌕', path: '/search' },
  { name: 'library', label: '媒体库', icon: '▦', path: '/library' },
  { name: 'favorites', label: '收藏', icon: '♥', path: '/favorites' },
  { name: 'history', label: '历史', icon: '↻', path: '/history' },
  { name: 'stats', label: '统计', icon: '📊', path: '/stats' },
  { name: 'settings', label: '设置', icon: '⚙', path: '/settings' },
]
</script>

<template>
  <aside class="sidebar" :class="{ collapsed: ui.sidebarCollapsed }">
    <div class="logo">
      <span class="logo-icon">◆</span>
      <span v-if="!ui.sidebarCollapsed" class="logo-text">Media Hub</span>
    </div>
    <nav class="nav">
      <router-link
        v-for="item in navItems"
        :key="item.name"
        :to="item.path"
        class="nav-item"
        :class="{ active: route.name === item.name }"
      >
        <span class="nav-icon">{{ item.icon }}</span>
        <span v-if="!ui.sidebarCollapsed" class="nav-label">{{ item.label }}</span>
      </router-link>
    </nav>
    <button class="collapse-btn" @click="ui.toggleSidebar">
      {{ ui.sidebarCollapsed ? '▶' : '◀' }}
    </button>
  </aside>
</template>

<style scoped>
.sidebar {
  width: 220px;
  background: var(--bg-sidebar);
  border-right: 1px solid var(--border);
  display: flex;
  flex-direction: column;
  transition: width 0.2s;
  flex-shrink: 0;
}

.sidebar.collapsed {
  width: 60px;
}

.logo {
  padding: 20px 16px;
  display: flex;
  align-items: center;
  gap: 10px;
  border-bottom: 1px solid var(--border);
}

.logo-icon {
  font-size: 20px;
  color: var(--accent);
}

.logo-text {
  font-size: 16px;
  font-weight: 600;
  white-space: nowrap;
}

.nav {
  flex: 1;
  padding: 12px 8px;
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 12px;
  border-radius: 8px;
  color: var(--text-secondary);
  text-decoration: none;
  transition: background 0.15s, color 0.15s;
}

.nav-item:hover {
  background: var(--bg-hover);
  color: var(--text);
}

.nav-item.active {
  background: var(--bg-active);
  color: var(--accent);
}

.nav-icon {
  font-size: 18px;
  width: 24px;
  text-align: center;
  flex-shrink: 0;
}

.nav-label {
  font-size: 14px;
  white-space: nowrap;
}

.collapse-btn {
  padding: 12px;
  border: none;
  background: transparent;
  color: var(--text-secondary);
  cursor: pointer;
  border-top: 1px solid var(--border);
  font-size: 12px;
}

.collapse-btn:hover {
  color: var(--text);
}
</style>
