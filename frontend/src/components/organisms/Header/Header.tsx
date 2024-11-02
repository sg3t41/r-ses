import * as utils from '@/utils'
import { JwtPayload } from 'jwt-decode'

type JwtFields = {
  session_id: string
  user_id: string
  username: string
  avatar_url: string
} & JwtPayload

const Header = async () => {
  const jwt = (await utils.cookie.get('token')) ?? ''
  const { username } = utils.jwt.decode<JwtFields>(jwt)

  return <header>{username}でログイン中</header>
}

export default Header
