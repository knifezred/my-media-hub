import type { Tag, PaginationData } from '../types'
import { post, get } from './client'

export function fetchTagPage(): Promise<PaginationData<Tag>> {
  return post('/tags/page')
}

export function fetchTagDetail(id: number): Promise<Tag> {
  return get(`/tags/${id}`)
}
