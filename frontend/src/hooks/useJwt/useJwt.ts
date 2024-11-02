// hooks/useJwtToken.js
import { useLayoutEffect, useState } from 'react'
import * as utils from '@/utils'

export const useJwt = () => {
  const [token, setToken] = useState('')

  useLayoutEffect(() => {
    const fetchToken = async () => {
      const token = await utils.cookie.get('token')
      setToken(token || '')
    }
    fetchToken()
  }, [])

  return { token }
}
