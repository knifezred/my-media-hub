import { defineStore } from 'pinia'
import { ref } from 'vue'
import type { Media, SearchParams } from '../types'
import * as searchApi from '../api/search'

export const useSearchStore = defineStore('search', () => {
  const results = ref<Media[]>([])
  const total = ref(0)
  const loading = ref(false)
  const suggestions = ref<string[]>([])

  async function search(params: SearchParams) {
    loading.value = true
    try {
      const data = await searchApi.searchMedia(params)
      results.value = data.items
      total.value = data.total
    } finally {
      loading.value = false
    }
  }

  async function fetchSuggestions(keyword: string) {
    if (!keyword.trim()) {
      suggestions.value = []
      return
    }
    try {
      const data = await searchApi.fetchSearchSuggestions(keyword)
      suggestions.value = data.items
    } catch {
      suggestions.value = []
    }
  }

  return { results, total, loading, suggestions, search, fetchSuggestions }
})
