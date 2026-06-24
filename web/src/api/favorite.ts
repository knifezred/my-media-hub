import type { Media, PaginationData } from '../types'
import { post, del, get } from './client'

export function addFavorite(mediaId: number): Promise<void> {
  return post('/favorites', { media_id: mediaId })
}

export function removeFavorite(mediaId: number): Promise<void> {
  return del(`/favorites/${mediaId}`)
}

export function fetchFavoritePage(params: { page?: number; page_size?: number; media_type?: string } = {}): Promise<PaginationData<Media>> {
  return post('/favorites/page', params)
}

export function fetchFavorites(params: { page?: number; page_size?: number } = {}): Promise<PaginationData<Media>> {
  return get(`/favorites?page=${params.page || 1}&page_size=${params.page_size || 20}`)
}
