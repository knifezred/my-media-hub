import type { Media, MediaDetail, MediaPageParams, PaginationData } from '../types'
import { post, get } from './client'

export function fetchMediaPage(params: MediaPageParams): Promise<PaginationData<Media>> {
  return post('/media/page', params)
}

export function fetchMediaDetail(id: number): Promise<MediaDetail> {
  return get(`/media/${id}`)
}
