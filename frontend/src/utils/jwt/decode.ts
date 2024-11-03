import { jwtDecode, JwtPayload } from 'jwt-decode'

type Result<T> = { success: true; data: T } | { success: false; error: Error }

const decode = <C extends object>(
  rawJwt: string,
): Result<(C & JwtPayload) | undefined> => {
  try {
    const decoded = jwtDecode<C & JwtPayload>(rawJwt)
    return { success: true, data: decoded }
  } catch (error) {
    console.log('JWTトークンが無効です。')
    return { success: false, error: error as Error }
  }
}

export { decode }
