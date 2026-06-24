import type { Media, PaginationData } from '../types'
import { post, del, get } from './client'

export function hideMedia(mediaId: number): Promise<void> {
  return post('/hidden', { media_id: mediaId })
}

export function unhideMedia(mediaId: number): Promise<void> {
  return del(`/hidden/${mediaId}`)
}

export function fetchHidden(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/hidden?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}
