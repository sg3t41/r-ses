import * as utils from '@/utils'

export const PublicRepositories = async () => {
  const jwtToken = (await utils.cookie.get('token')) ?? ''

  const response = await fetch(`http://api:8080/api/v1/repositories/public`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${jwtToken}`,
    },
  })

  const { repositories } = await response.json()

  if (!response.ok) {
    throw new Error('Failed to get repositories')
  }

  return (
    <ul>
      {repositories &&
        repositories.map((v: string, i: number) => <li key={i}>{v}</li>)}
    </ul>
  )
}
