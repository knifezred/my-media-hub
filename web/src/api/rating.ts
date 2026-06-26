import { get } from './client'

export function rateMedia(mediaId: number, rating: number): Promise<void> {
  return fetch(`/api/v1/ratings/${mediaId}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ rating }),
  }).then(r => r.json()).then(b => { if (b.code !== 0) throw new Error(b.message) })
}

export function fetchRatingForMedia(mediaId: number): Promise<{ rating: number }> {
  return get(`/ratings/${mediaId}`)
}
