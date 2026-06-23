<script setup lang="ts">
</script>

<template>
  <div class="favorites-page">
    <div class="empty-state">
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
