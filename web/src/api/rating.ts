import { post } from './client'

export function rateMedia(mediaId: number, rating: number): Promise<void> {
  return post('/ratings', { media_id: mediaId, rating })
}
