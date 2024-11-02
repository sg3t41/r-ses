import { useEffect, useState } from 'react'
import { Repository } from '@/app/test2/types'
import { useJwt } from '../useJwt/useJwt'

export const usePublicRepositories = () => {
  const [repositories, setRepositories] = useState<Repository[]>([])
  const [error, setError] = useState<string | null>(null)
  const { token } = useJwt()

  useEffect(() => {
    const fetchRepositories = async () => {
      if (!token) return

      try {
        const response = await fetch(
          `http://localhost:8080/api/v1/github/repositories/public`,
          {
            method: 'GET',
            headers: {
              Authorization: `Bearer ${token}`,
            },
          },
        )

        if (!response.ok) {
          const data = await response.json()
          throw new Error(data.error || 'Failed to get repositories')
        }

        const data = await response.json()
        setRepositories(data.repositories)
      } catch (err: any) {
        setError(err.message)
        console.error(err)
      }
    }

    fetchRepositories()
  }, [token])

  return { repositories, error }
}
