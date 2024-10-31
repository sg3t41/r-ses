import { set, get } from './cookie/cookie'
import { hash } from './crypt/sha256'
import { decode } from './jwt/decode'

const sha256 = { hash }
const cookie = { set, get }
const jwt = { decode }

export { cookie, sha256, jwt }
