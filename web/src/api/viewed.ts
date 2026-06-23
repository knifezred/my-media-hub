import type { ListData, Media } from '../types'
import { post } from './client'

export function markViewed(mediaId: number): Promise<void> {
  return post('/viewed', { media_id: mediaId })
}

export function fetchViewedPage(): Promise<ListData<Media>> {
  return post('/viewed/page')
}
