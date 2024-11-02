// RepositoryCard.tsx
import React from 'react'

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
  fork: boolean // フォークされたかどうか
  full_name: string
}

interface RepositoryCardProps {
  repository: Repository
}

const RepositoryCard: React.FC<RepositoryCardProps> = ({ repository }) => {
  console.log(repository.id)
  return (
    <div className='repository-card' style={styles.card}>
      <h3 style={styles.title}>{repository.full_name}</h3>
      <p style={styles.description}>
        {repository.description || 'No description provided'}
      </p>
      <p>
        <strong>Created at:</strong>{' '}
        {new Date(repository.created_at).toLocaleDateString()}
      </p>
      <p>
        <strong>Updated at:</strong>{' '}
        {new Date(repository.updated_at).toLocaleDateString()}
      </p>
      <p>
        <strong>Clone URL:</strong>{' '}
        <a
          href={repository.clone_url}
          target='_blank'
          rel='noopener noreferrer'>
          {repository.clone_url}
        </a>
      </p>
      <p>
        <strong>Archived:</strong> {repository.archived ? 'Yes' : 'No'}
      </p>
      <p>
        <strong>Forking allowed:</strong>{' '}
        {repository.allow_forking ? 'Yes' : 'No'}
      </p>
      <p>
        <strong>Forked:</strong> {repository.fork ? 'Yes' : 'No'}
      </p>
      <p>
        <strong>License:</strong>{' '}
        {repository.license ? repository.license.name : 'No license'}
      </p>
      <p>
        <strong>Stars:</strong> {repository.stargazers_count}
      </p>
      <p>
        <strong>Forks:</strong> {repository.forks_count}
      </p>
      <p>
        <strong>Language:</strong> {repository.language || 'Not specified'}
      </p>
    </div>
  )
}

// スタイルの定義
const styles = {
  card: {
    border: '1px solid #ddd',
    borderRadius: '8px',
    padding: '16px',
    margin: '16px 0',
    boxShadow: '0 2px 4px rgba(0, 0, 0, 0.1)',
  },
  title: {
    fontSize: '1.5em',
    margin: '0 0 10px',
  },
  description: {
    fontStyle: 'italic',
    color: '#555',
  },
}

export default RepositoryCard
