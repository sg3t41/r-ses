import { cookies } from 'next/headers'
import { redirect } from 'next/navigation'
import { NextRequest } from 'next/server'

export async function GET(request: NextRequest) {
  const searchParams = request.nextUrl.searchParams
  const jwtAccessToken = searchParams.get('token')
  const cookieStore = await cookies()
  if (jwtAccessToken) {
    cookieStore.set('token', jwtAccessToken)
  }

  redirect('http://localhost:3000/test2')
}
