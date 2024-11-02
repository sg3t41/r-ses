export type Repository = {
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
    key: string
    name: string
  } | null
  private: boolean
  stargazers_count: number
  forks_count: number
  language: string | null
  name: string
  fork: boolean
  full_name: string
}
