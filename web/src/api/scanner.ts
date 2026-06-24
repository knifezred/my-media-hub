import type { ScannerStatus } from '../types'
import { post, get } from './client'

export function startScan(directories: string[]): Promise<void> {
  return post('/scanner/start', { directories })
}

export function fetchScanStatus(): Promise<ScannerStatus> {
  return get('/scanner/status')
}
