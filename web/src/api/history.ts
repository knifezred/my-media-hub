import type { PaginationData, SearchHistory } from '../types'
import { post, del } from './client'

export function fetchSearchHistory(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<SearchHistory>> {
  return post('/search/history/page', params)
}

export function deleteSearchHistory(id: number): Promise<void> {
  return del(`/search/history/${id}`)
}

export function clearSearchHistory(): Promise<void> {
  return del('/search/history')
}
