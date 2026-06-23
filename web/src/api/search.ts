import type { Media, SearchParams, PaginationData } from '../types'
import { post, get } from './client'

export function searchMedia(params: SearchParams): Promise<PaginationData<Media>> {
  return post('/search/page', params)
}

export function fetchSearchSuggestions(keyword: string): Promise<{ items: string[] }> {
  return get(`/search/suggestions?keyword=${encodeURIComponent(keyword)}`)
}
