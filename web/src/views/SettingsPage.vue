<script setup lang="ts">
import { ref } from 'vue'
import { startScan, fetchScanStatus } from '../api/scanner'
import type { ScannerStatus } from '../types'

const scanPaths = ref('')
const scanStatus = ref<ScannerStatus | null>(null)
const scanning = ref(false)

async function handleStartScan() {
  const dirs = scanPaths.value
    .split('\n')
    .map(s => s.trim())
    .filter(Boolean)
  if (!dirs.length) return

  scanning.value = true
  try {
    await startScan(dirs)
    scanStatus.value = { running: true, processed: 0, total: 0, progress: 0 }
    pollStatus()
  } catch {
    scanning.value = false
  }
}

async function pollStatus() {
  const interval = setInterval(async () => {
    try {
      const status = await fetchScanStatus()
      scanStatus.value = status
      if (!status.running) {
        clearInterval(interval)
        scanning.value = false
      }
    } catch {
      clearInterval(interval)
      scanning.value = false
    }
  }, 2000)
}

async function loadStatus() {
  try {
    scanStatus.value = await fetchScanStatus()
  } catch {}
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
      <div class="btn-group">
        <button class="btn-primary" :disabled="scanning || !scanPaths.trim()" @click="handleStartScan">
          {{ scanning ? '扫描中...' : '启动扫描' }}
        </button>
        <button class="btn-secondary" @click="loadStatus">刷新状态</button>
      </div>

      <div v-if="scanStatus" class="scan-status">
        <div class="status-row">
          <span class="status-label">状态</span>
          <span class="status-value" :class="{ active: scanStatus.running }">
            {{ scanStatus.running ? '扫描中' : '空闲' }}
          </span>
        </div>
        <div v-if="scanStatus.total > 0" class="progress-bar-bg">
          <div class="progress-bar-fill" :style="{ width: scanStatus.progress + '%' }" />
        </div>
        <div v-if="scanStatus.running" class="status-row">
          <span class="status-label">进度</span>
          <span class="status-value">{{ scanStatus.processed }}/{{ scanStatus.total }} ({{ Math.round(scanStatus.progress) }}%)</span>
        </div>
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

.btn-group {
  display: flex;
  gap: 8px;
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

.btn-primary:hover:not(:disabled) {
  opacity: 0.9;
}

.btn-primary:disabled {
  opacity: 0.5;
  cursor: not-allowed;
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
  padding: 16px;
  border-radius: 8px;
  background: var(--bg-hover);
}

.status-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.status-label {
  font-size: 14px;
  color: var(--text-secondary);
}

.status-value {
  font-size: 14px;
  font-weight: 500;
}

.status-value.active {
  color: var(--accent);
}

.progress-bar-bg {
  width: 100%;
  height: 6px;
  background: var(--border);
  border-radius: 3px;
  overflow: hidden;
  margin-bottom: 8px;
}

.progress-bar-fill {
  height: 100%;
  background: var(--accent);
  border-radius: 3px;
  transition: width 1s ease;
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
