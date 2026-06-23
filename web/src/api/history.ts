import type { ListData } from '../types'
import { post, del } from './client'

export interface SearchHistory {
  id: number
  keyword: string
  result_count: number
  created_at: string
}

export function fetchSearchHistory(): Promise<ListData<SearchHistory>> {
  return post('/search/history/page')
}

export function deleteSearchHistory(id: number): Promise<void> {
  return del(`/search/history/${id}`)
}

export function clearSearchHistory(): Promise<void> {
  return del('/search/history')
}
