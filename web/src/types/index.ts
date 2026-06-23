export interface Media {
  id: number
  media_type: string
  title: string
  description: string
  path: string
  cover_path: string
  size: number
  favorite: boolean
  rating: number
  viewed: boolean
  hidden: boolean
  tags: Tag[]
  categories: Category[]
  metadata: Record<string, string>
  created_at: string
  updated_at: string
}

export interface Tag {
  id: number
  name: string
}

export interface Category {
  id: number
  name: string
  parent_id: number
}

export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

export interface PaginationData<T> {
  items: T[]
  total: number
  page: number
  page_size: number
}

export interface ListData<T> {
  items: T[]
}

export interface StatsOverview {
  total_media: number
  total_images: number
  total_videos: number
  total_novels: number
  favorite_count: number
  viewed_count: number
}

export interface ScannerStatus {
  running: boolean
  processed: number
  total: number
  progress: number
}

export interface MediaPageParams {
  page?: number
  page_size?: number
  media_type?: string
  category_id?: number
  tag_id?: number
  sort?: string
}

export interface SearchParams {
  keyword: string
  media_type?: string
  page?: number
  page_size?: number
}

export interface MediaDetail extends Media {
  metadata: Record<string, string>
}

export interface FavoritePageParams {
  page?: number
  page_size?: number
  media_type?: string
}
