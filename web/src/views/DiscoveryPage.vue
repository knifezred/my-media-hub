<script setup lang="ts">
</script>

<template>
  <div class="discovery-page">
    <section class="section">
      <h3 class="section-title">最近新增</h3>
      <div class="empty-state">
        <p>暂无媒体资源</p>
        <p class="hint">前往设置页配置扫描目录并启动扫描</p>
      </div>
    </section>
  </div>
</template>

<style scoped>
.discovery-page {
  max-width: 1200px;
  margin: 0 auto;
}

.section {
  margin-bottom: 32px;
}

.section-title {
  font-size: 20px;
  font-weight: 600;
  margin-bottom: 16px;
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
