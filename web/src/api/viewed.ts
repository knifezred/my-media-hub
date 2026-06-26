import type { Media, PaginationData } from '../types'
import { post, get } from './client'

export function markViewed(mediaId: number): Promise<void> {
  return post('/history', { media_id: mediaId })
}

export function fetchViewedPage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/history/page', params)
}

export function fetchHistory(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/history?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}
