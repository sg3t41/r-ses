import { PublicRepositories } from './PublicRepositories'

export default async function TestPage2() {
  return (
    <>
      <div>ポートフォリオとして公開するプロジェクトを選択してください。</div>
      <PublicRepositories />
    </>
  )
}
