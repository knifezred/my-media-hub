import type { StatsOverview } from '../types'
import { get } from './client'

export function fetchStatsOverview(): Promise<StatsOverview> {
  return get('/stats/overview')
}
