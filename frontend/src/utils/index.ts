import { set, get } from './cookie/cookie'
import { hash } from './crypt/sha256'

const sha256 = { hash }
const cookie = { set, get }

export { cookie, sha256 }
