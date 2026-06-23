<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue'
import Header from '../components/Header.vue'
import { useUiStore } from '../stores/ui'

const ui = useUiStore()
</script>

<template>
  <div class="layout" :class="{ collapsed: ui.sidebarCollapsed }">
    <Sidebar />
    <div class="main-area">
      <Header />
      <main class="content">
        <router-view />
      </main>
    </div>
  </div>
</template>

<style scoped>
.layout {
  display: flex;
  height: 100vh;
  overflow: hidden;
}

.main-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.content {
  flex: 1;
  overflow-y: auto;
  padding: 24px;
}
</style>
