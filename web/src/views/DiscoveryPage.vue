<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { fetchMediaPage } from '../api/media'
import type { Media } from '../types'
import MediaGrid from '../components/MediaGrid.vue'
import Pagination from '../components/Pagination.vue'

const recentItems = ref<Media[]>([])
const recentTotal = ref(0)
const recentPage = ref(1)
const loading = ref(false)

const pageSize = 20

async function loadRecent() {
  loading.value = true
  try {
    const data = await fetchMediaPage({ page: recentPage.value, page_size: pageSize, sort: 'created_at' })
    recentItems.value = data.items
    recentTotal.value = data.total
  } finally {
    loading.value = false
  }
}

onMounted(loadRecent)
</script>

<template>
  <div class="discovery-page">
    <section class="section">
      <h3 class="section-title">最近新增</h3>
      <MediaGrid :items="recentItems" :loading="loading" />
      <Pagination
        v-if="recentTotal > pageSize"
        :page="recentPage"
        :page-size="pageSize"
        :total="recentTotal"
        @change="recentPage = $event; loadRecent()"
      />
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
</style>
