import { jwtDecode, JwtPayload } from 'jwt-decode'

const decode = <C extends object>(rawJwt: string): C & JwtPayload => {
  type CustomJwtPayload = C & JwtPayload
  const decoded = jwtDecode<CustomJwtPayload>(rawJwt)
  return decoded
}

export { decode }
