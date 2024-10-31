'use server'
import * as utils from '@/utils'
import { PublicRepositories } from './PublicRepositories'
export default async function TestPage2() {
  const jwtRaw = (await utils.cookie.get('token')) ?? ''
  const { session_id, user_id, username, avatar_url, exp, iss } =
    utils.jwt.decode<{
      session_id: string
      user_id: string
      username: string
      avatar_url: string
      exp: string
      iss: string
    }>(jwtRaw)

  return (
    <>
      <div>{session_id}</div>
      <div>{user_id}</div>
      <div>{username}</div>
      <div>{avatar_url}</div>
      <div>{exp}</div>
      <div>{iss}</div>
      <div>ポートフォリオとして公開するプロジェクトを選択してください。</div>
      <PublicRepositories />
    </>
  )
}
