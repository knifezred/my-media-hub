<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  page: number
  pageSize: number
  total: number
}>()

const emit = defineEmits<{
  change: [page: number]
}>()

const totalPages = computed(() => Math.max(1, Math.ceil(props.total / props.pageSize)))

function go(p: number) {
  if (p >= 1 && p <= totalPages.value) {
    emit('change', p)
  }
}
</script>

<template>
  <div v-if="totalPages > 1" class="pagination">
    <button class="page-btn" :disabled="page <= 1" @click="go(page - 1)">‹</button>
    <template v-for="p in totalPages" :key="p">
      <button
        v-if="p === 1 || p === totalPages || Math.abs(p - page) <= 1"
        class="page-btn"
        :class="{ active: p === page }"
        @click="go(p)"
      >{{ p }}</button>
      <span v-else-if="p === page - 2 || p === page + 2" class="page-ellipsis">…</span>
    </template>
    <button class="page-btn" :disabled="page >= totalPages" @click="go(page + 1)">›</button>
  </div>
</template>

<style scoped>
.pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 4px;
  padding: 24px 0;
}

.page-btn {
  min-width: 36px;
  height: 36px;
  padding: 0 10px;
  border: 1px solid var(--border);
  border-radius: 6px;
  background: transparent;
  color: var(--text);
  cursor: pointer;
  font-size: 14px;
  transition: all 0.15s;
}

.page-btn:hover:not(:disabled):not(.active) {
  border-color: var(--accent);
  color: var(--accent);
}

.page-btn.active {
  background: var(--accent);
  border-color: var(--accent);
  color: #fff;
}

.page-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.page-ellipsis {
  padding: 0 4px;
  color: var(--text-secondary);
}
</style>
