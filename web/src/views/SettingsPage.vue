<script setup lang="ts">
import { ref } from 'vue'
import { startScan, fetchScanStatus } from '../api/scanner'
import type { ScannerStatus } from '../types'

const scanPaths = ref('')
const scanStatus = ref<ScannerStatus | null>(null)

async function handleStartScan() {
  try {
    await startScan()
  } catch {
    // API not ready yet
  }
}

async function loadStatus() {
  try {
    scanStatus.value = await fetchScanStatus()
  } catch {
    // API not ready yet
  }
}
</script>

<template>
  <div class="settings-page">
    <section class="settings-section">
      <h3>扫描配置</h3>
      <div class="form-group">
        <label>媒体目录</label>
        <textarea
          v-model="scanPaths"
          class="form-textarea"
          placeholder="每行一个目录路径&#10;例如:&#10;/mnt/nas/media/movies&#10;/mnt/nas/media/books"
          rows="5"
        />
      </div>
      <button class="btn-primary" @click="handleStartScan">启动扫描</button>
      <button class="btn-secondary" @click="loadStatus" style="margin-left: 8px;">刷新状态</button>

      <div v-if="scanStatus" class="scan-status">
        <p>状态: {{ scanStatus.running ? '扫描中' : '空闲' }}</p>
        <p v-if="scanStatus.running">
          进度: {{ scanStatus.processed }}/{{ scanStatus.total }}
          ({{ scanStatus.progress }}%)
        </p>
      </div>
    </section>

    <section class="settings-section">
      <h3>系统信息</h3>
      <div class="info-list">
        <div class="info-item">
          <span class="info-label">版本</span>
          <span class="info-value">0.1.0</span>
        </div>
      </div>
    </section>
  </div>
</template>

<style scoped>
.settings-page {
  max-width: 600px;
}

.settings-section {
  margin-bottom: 32px;
}

.settings-section h3 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 16px;
  padding-bottom: 8px;
  border-bottom: 1px solid var(--border);
}

.form-group {
  margin-bottom: 16px;
}

.form-group label {
  display: block;
  font-size: 14px;
  color: var(--text-secondary);
  margin-bottom: 8px;
}

.form-textarea {
  width: 100%;
  padding: 12px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: var(--bg-input);
  color: var(--text);
  font-size: 14px;
  resize: vertical;
  font-family: inherit;
}

.form-textarea:focus {
  outline: none;
  border-color: var(--accent);
}

.btn-primary {
  padding: 10px 24px;
  border: none;
  border-radius: 8px;
  background: var(--accent);
  color: #fff;
  cursor: pointer;
  font-size: 14px;
}

.btn-primary:hover {
  opacity: 0.9;
}

.btn-secondary {
  padding: 10px 24px;
  border: 1px solid var(--border);
  border-radius: 8px;
  background: transparent;
  color: var(--text);
  cursor: pointer;
  font-size: 14px;
}

.btn-secondary:hover {
  background: var(--bg-hover);
}

.scan-status {
  margin-top: 16px;
  padding: 12px;
  border-radius: 8px;
  background: var(--bg-hover);
  font-size: 14px;
  color: var(--text-secondary);
}

.info-list {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
}

.info-label {
  color: var(--text-secondary);
  font-size: 14px;
}

.info-value {
  font-size: 14px;
}
</style>
