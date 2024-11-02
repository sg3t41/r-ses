import React from 'react'
import { Repository } from './types'

type Props = {
  repository: Repository
  isChecked: boolean
  onChange: (id: string) => void
}

const RepositoryCard: React.FC<Props> = ({
  repository,
  isChecked,
  onChange,
}) => {
  return (
    <div className='repository-card' style={styles.card}>
      <input
        name='repository'
        value={repository.id}
        type='checkbox'
        checked={isChecked}
        onChange={() => onChange(repository.id)}
        style={styles.checkbox}
      />
      <h3 style={styles.title}>{repository.id}</h3>
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
  checkbox: {
    marginRight: '8px',
  },
}

export default RepositoryCard
