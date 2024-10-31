'use server'
import * as utils from '@/utils'
import { jwtDecode, JwtPayload } from 'jwt-decode'
export default async function TestPage2() {
  const jwtToken = await utils.cookie.get('token')

  type CustomJwtPayload = {
    session_id: string
    user_id: string
    username: string
    avatar_url: string
    exp: string
    iss: string
  } & JwtPayload

  let decodedJwtToken
  if (jwtToken) {
    decodedJwtToken = jwtDecode<CustomJwtPayload>(jwtToken)
    console.log(decodedJwtToken)
  }
  return (
    <>
      <div>{jwtToken}</div>
      <div>{decodedJwtToken?.session_id}</div>

      <div>{decodedJwtToken?.user_id}</div>
      <div>{decodedJwtToken?.username}</div>
    </>
  )
}
