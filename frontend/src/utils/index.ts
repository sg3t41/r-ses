import { set, get } from './cookie/cookie'
import { hash } from './crypt/sha256'
import { decode } from './jwt/decode'
import { isLogin } from './jwt/isLogin'

const sha256 = { hash }
const cookie = { set, get }
const jwt = { decode, isLogin }

export { cookie, sha256, jwt }
