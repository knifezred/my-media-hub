import type { Media, PaginationData, BehaviorStatistics } from '../types'
import { get, post, del } from './client'

export function recordBehavior(mediaId: number, behaviorType: string, score: number = 1): Promise<void> {
  return post('/behavior', { media_id: mediaId, behavior_type: behaviorType, score })
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

export function fetchFavorites(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/favorites?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}

export function fetchFavoritePage(params: { page?: number; page_size?: number; media_type?: string } = {}): Promise<PaginationData<Media>> {
  return post('/favorites/page', params)
}

export function rateMedia(mediaId: number, rating: number): Promise<void> {
  return post('/ratings', { media_id: mediaId, rating })
}

export function rateMediaPut(mediaId: number, rating: number): Promise<void> {
  return fetch(`/api/v1/ratings/${mediaId}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ rating }),
  }).then(r => r.json()).then(b => { if (b.code !== 0) throw new Error(b.message) })
}

export function fetchRatingForMedia(mediaId: number): Promise<{ rating: number }> {
  return get(`/ratings/${mediaId}`)
}

export function markViewed(mediaId: number): Promise<void> {
  return post('/viewed', { media_id: mediaId })
}

export function markHistory(mediaId: number): Promise<void> {
  return post('/history', { media_id: mediaId })
}

export function fetchHistory(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/history?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}

export function fetchViewedPage(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return post('/viewed/page', params)
}

export function hideMedia(mediaId: number): Promise<void> {
  return post('/hidden', { media_id: mediaId })
}

export function unhideMedia(mediaId: number): Promise<void> {
  return del(`/hidden/${mediaId}`)
}

export function fetchHidden(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/hidden?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}
