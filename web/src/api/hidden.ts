import { post, del } from './client'

export function hideMedia(mediaId: number): Promise<void> {
  return post('/hidden', { media_id: mediaId })
}

export function unhideMedia(mediaId: number): Promise<void> {
  return del(`/hidden/${mediaId}`)
}
