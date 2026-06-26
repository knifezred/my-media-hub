import { post, del } from './client'
import type { Media, PaginationData } from '../types'

export function hideMedia(mediaId: number): Promise<void> {
  return post('/hidden', { media_id: mediaId })
}

export function unhideMedia(mediaId: number): Promise<void> {
  return del(`/hidden/${mediaId}`)
}

export function fetchHiddenPage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/hidden/page', params)
}
