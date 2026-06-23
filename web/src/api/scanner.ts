import type { ScannerStatus } from '../types'
import { post, get } from './client'

export function startScan(): Promise<void> {
  return post('/scanner/start')
}

export function fetchScanStatus(): Promise<ScannerStatus> {
  return get('/scanner/status')
}
