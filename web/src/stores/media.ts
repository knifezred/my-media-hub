import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Media, MediaPageParams } from '../types'
import * as mediaApi from '../api/media'

export const useMediaStore = defineStore('media', () => {
  const items = ref<Media[]>([])
  const total = ref(0)
  const loading = ref(false)
  const currentPage = ref(1)
  const pageSize = ref(20)

  async function fetchPage(params: MediaPageParams = {}) {
    loading.value = true
    try {
      const data = await mediaApi.fetchMediaPage({
        page: currentPage.value,
        page_size: pageSize.value,
        ...params,
      })
      items.value = data.items
      total.value = data.total
    } finally {
      loading.value = false
    }
  }

  return { items, total, loading, currentPage, pageSize, fetchPage }
})
