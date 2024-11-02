'use client'
import * as hooks from '@/hooks'
import RepositoryCard from './RepositoryCard'
import React from 'react'

type FormState = {
  selectedRepoIds?: Array<string>
  error?: string
}

export const PublicRepositories = () => {
  const { repositories } = hooks.usePublicRepositories()
  const [state, action] = React.useActionState(register, {})
  const [selectedRepoIds, setSelectedRepoIds] = React.useState<string[]>([])

  async function register(previousState: FormState, formData: FormData) {
    console.log(formData.get('repo'))

    return { error: 'error!' } as FormState
  }

  const handleSelect = (id: string) => {
    setSelectedRepoIds(
      prevSelected =>
        prevSelected.includes(id)
          ? prevSelected.filter(repoId => repoId !== id) // 選択解除
          : [...prevSelected, id], // 新規選択
    )
  }

  return (
    <form action={action}>
      <ul>
        {repositories.map(repository => (
          <li key={repository.id}>
            <RepositoryCard
              repository={repository}
              isChecked={selectedRepoIds.includes(repository.id)}
              onChange={handleSelect}
            />
          </li>
        ))}
        <button type={'submit'}>登録</button>
        {state.error}
      </ul>
    </form>
  )
}
