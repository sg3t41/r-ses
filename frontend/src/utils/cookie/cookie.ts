'use server'
import { cookies } from 'next/headers'

const set = async (key: string, value: string) => {
  const cookieStore = await cookies()
  cookieStore.set({
    name: key,
    value: value,
    httpOnly: true,
    path: '/',
  })
}

const get = async (key: string) => {
  const cookieStore = await cookies()
  const value = cookieStore.get(key)?.value
  return value
}

export { set, get }
