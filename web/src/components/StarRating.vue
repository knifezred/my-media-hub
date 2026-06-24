<script setup lang="ts">
defineProps<{
  modelValue: number
}>()

const emit = defineEmits<{
  'update:modelValue': [value: number]
}>()
</script>

<template>
  <div class="star-rating">
    <span
      v-for="i in 5"
      :key="i"
      class="star"
      :class="{ filled: i <= modelValue }"
      @click="emit('update:modelValue', i === modelValue ? 0 : i)"
    >★</span>
  </div>
</template>

<style scoped>
.star-rating {
  display: flex;
  gap: 4px;
}

.star {
  font-size: 22px;
  cursor: pointer;
  color: var(--border);
  transition: color 0.15s, transform 0.15s;
}

.star.filled {
  color: #f5c518;
}

.star:hover {
  transform: scale(1.2);
}
</style>
