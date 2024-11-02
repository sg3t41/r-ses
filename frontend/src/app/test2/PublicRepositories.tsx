import * as utils from '@/utils'
import RepositoryCard from './RepositoryCard'

// リポジトリのフィールドに基づくインターフェースを定義
interface Repository {
  id: string
  allow_forking: boolean
  archive_url: string
  archived: boolean
  assignees_url: string
  blobs_url: string
  branches_url: string
  clone_url: string
  collaborators_url: string
  comments_url: string
  commits_url: string
  compare_url: string
  contents_url: string
  contributors_url: string
  created_at: string // ISO 8601形式の日時
  updated_at: string // 最後の更新日時
  default_branch: string // デフォルトブランチ
  deployments_url: string
  description: string | null // 説明がない場合もあるのでnullable
  disabled: boolean
  downloads_url: string
  license: {
    // ライセンス情報
    key: string
    name: string
  } | null
  private: boolean // プライベートかパブリックか
  stargazers_count: number // スター数
  forks_count: number // フォーク数
  language: string | null // 主なプログラミング言語
  name: string // リポジトリ名
  fork: boolean
  full_name: string
}

export const PublicRepositories = async () => {
  // JWTトークンを取得
  const jwtToken = (await utils.cookie.get('token')) || ''

  // トークンがない場合はエラーを投げる
  if (!jwtToken) {
    throw new Error('No JWT token found')
  }

  // リポジトリを取得するAPIにリクエスト
  const response = await fetch(`http://api:8080/api/v1/github/repositories`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${jwtToken}`,
    },
  })

  // レスポンスをJSON形式で解析
  const data = await response.json()

  // ステータスコードがOKでない場合のエラーハンドリング
  if (!response.ok) {
    console.error(data.error)
    throw new Error(data.error || 'Failed to get repositories')
  }

  const repositories: Repository[] = data.repositories

  return (
    <ul>
      {repositories.map((repository: Repository, index: number) => (
        <li key={index}>
          <RepositoryCard {...{ repository }} />
        </li>
      ))}
    </ul>
  )
}
