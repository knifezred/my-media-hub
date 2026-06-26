import type { Media, PaginationData, BehaviorStatistics } from '../types'
import { get, post, del } from './client'

export function recordBehavior(mediaId: number, behaviorType: string, behaviorValue: string = '{}', behaviorSource: string = 'manual'): Promise<void> {
  return post('/behavior', { media_id: mediaId, behavior_type: behaviorType, behavior_value: behaviorValue, behavior_source: behaviorSource })
}

export function fetchBehaviorStatistics(): Promise<BehaviorStatistics> {
  return get('/behavior/statistics')
}

export function addFavorite(mediaId: number): Promise<void> {
  return post('/favorites', { media_id: mediaId })
}

export function removeFavorite(mediaId: number): Promise<void> {
  return del(`/favorites/${mediaId}`)
}

export function fetchFavoritePage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/favorites/page', params)
}

export function fetchRatingForMedia(mediaId: number): Promise<{ rating: number }> {
  return get(`/ratings/${mediaId}`)
}

export function rateMedia(mediaId: number, rating: number): Promise<void> {
  return fetch(`/api/v1/ratings/${mediaId}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ rating }),
  }).then(r => r.json()).then(b => { if (b.code !== 0) throw new Error(b.message) })
}

export function markViewed(mediaId: number): Promise<void> {
  return post('/history', { media_id: mediaId })
}

export function fetchViewedPage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/history/page', params)
}

export function hideMedia(mediaId: number): Promise<void> {
  return post('/hidden', { media_id: mediaId })
}

export function unhideMedia(mediaId: number): Promise<void> {
  return del(`/hidden/${mediaId}`)
}

export function fetchHiddenPage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/hidden/page', params)
}
