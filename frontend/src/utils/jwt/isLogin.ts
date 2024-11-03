import * as utils from '@/utils'
import { JwtPayload } from 'jwt-decode'

type JwtFields = {
  session_id: string
  user_id: string
  username: string
  avatar_url: string
} & JwtPayload

type Result<T> = { isLogin: true; data: T } | { isLogin: false }

const isLogin = async (): Promise<Result<JwtFields>> => {
  const jwt = (await utils.cookie.get('token')) ?? ''
  const decoded = utils.jwt.decode<JwtFields>(jwt)

  if (decoded.success) {
    return { isLogin: true, data: { ...decoded.data! } }
  } else {
    return { isLogin: false }
  }
}

export { isLogin }
