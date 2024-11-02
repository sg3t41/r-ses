'use client'
import * as hooks from '@/hooks'
import * as utils from '@/utils'
import RepositoryCard from './RepositoryCard'
import React from 'react'

type FormState = {
  error?: string
}

export const PublicRepositories = () => {
  const { repositories } = hooks.usePublicRepositories()
  const [state, action] = React.useActionState(register, {})
  const [selectedRepoIds, setSelectedRepoIds] = React.useState<string[]>([])

  async function register(previousState: FormState, formData: FormData) {
    const entriesArray = Array.from(formData.entries())
    const selectedRepoIds: Array<string> = entriesArray
      .map(([, value]) => value)
      .filter(value => typeof value === 'string')

    const jwtToken = (await utils.cookie.get('token')) ?? ''
    const response = await fetch(
      `http://localhost:8080/api/v1/github/repositories/portfolio`,
      {
        method: 'POST',
        headers: {
          Authorization: `Bearer ${jwtToken}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ portfolioRepoIds: selectedRepoIds }),
      },
    )
    if (!response.ok) {
      const data = await response.json()
      throw new Error(data.error || 'Failed to post selected repository IDs')
    }

    const responseData = await response.json()
    console.log('Response:', responseData.status)

    return { error: '' } as FormState
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
