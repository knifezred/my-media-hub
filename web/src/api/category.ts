import type { Category } from '../types'
import { post } from './client'

export function fetchCategoryPage(): Promise<{ items: Category[] }> {
  return post('/categories/page')
}
