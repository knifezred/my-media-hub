<script setup lang="ts">
defineProps<{
  tags: { id: number; name: string }[]
}>()
</script>

<template>
  <div class="tag-chips">
    <span v-for="tag in tags" :key="tag.id" class="tag-chip">{{ tag.name }}</span>
  </div>
</template>

<style scoped>
.tag-chips {
  display: flex;
  gap: 6px;
  flex-wrap: wrap;
}

.tag-chip {
  padding: 3px 10px;
  border-radius: 12px;
  background: var(--bg-hover);
  border: 1px solid var(--border);
  font-size: 12px;
  color: var(--text-secondary);
  white-space: nowrap;
}
</style>
