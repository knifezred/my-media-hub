import type { FavoritePageParams, ListData, Media } from '../types'
import { post, del } from './client'

export function addFavorite(mediaId: number): Promise<void> {
  return post('/favorites', { media_id: mediaId })
}

export function removeFavorite(mediaId: number): Promise<void> {
  return del(`/favorites/${mediaId}`)
}

export function fetchFavoritePage(params: FavoritePageParams): Promise<ListData<Media>> {
  return post('/favorites/page', params)
}
