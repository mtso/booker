import * as T from './types'
import request from 'superagent'
import { getBook } from './books'

export const postTrade = (book_id) => (dispatch) => request
  .post('/api/trade')
  .send({ book_id })
  .then(({ body }) => body)
  .then(({ ok, message }) => {
    if (!ok) {
      throw new Error(message)
    }
  })

export const postTradeFromPreview = (book_id) => (dispatch) => request
  .post('/api/trade')
  .send({ book_id })
  .then(({ body }) => body)
  .then(({ ok, message }) => {
    if (!ok) {
      throw new Error(message)
    }
  })

export const acceptTrade = (id) => (dispatch) => request
  .put('/api/trade/'+id)
  .then(({ body }) => body)
  .then(({ ok, message }) => {
    if (ok) {
      dispatch(getIncoming())
    } else {
      throw new Error(message)
    }
  })

export const cancelTradeFromPreview = (id) => (dispatch) => request
  .del('/api/trade/'+id)
  .then(({ body }) => body)
  .then(({ ok, message }) => {
    if (!ok) {
      throw new Error(message)
    }
  })

export const cancelTrade = (id) => (dispatch) => request
  .del('/api/trade/'+id)
  .set('Accept', 'application/json')
  .then(({ body }) => body)
  .then(({ ok, message }) => {
    if (ok) {
      dispatch(getOutgoing())
    } else {
      throw new Error(message)
    }
  })

export const getIncoming = () => (dispatch) => request
  .get('/api/trades/incoming')
  .then(({ body }) => body)
  .then(({ ok, trades, message }) => {
    if (ok) {
      dispatch(receiveIncoming(trades))
    } else {
      throw new Error(message)
    }
  })

export const getOutgoing = () => (dispatch) => request
  .get('/api/trades/outgoing')
  .then(({ body }) => body)
  .then(({ ok, trades, message }) => {
    if (ok) {
      dispatch(receiveOutgoing(trades))
    } else {
      throw new Error(message)
    }
  })

export const receiveIncoming = (trades) => ({
  type: T.RECEIVE_INCOMING,
  trades,
})

export const receiveOutgoing = (trades) => ({
  type: T.RECEIVE_OUTGOING,
  trades,
})

