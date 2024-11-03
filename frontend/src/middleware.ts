export async function middleware() {
  console.log('middleware is running')
  //  const token = request.cookies.get('jwttoken')?.value
  //  // TODO: tokenが存在しない処理
  //  if (!token) {
  //    return NextResponse.next()
  //  }
  //
  //  // TODO
  //  const decoded = jwtDecode(token)
  //  console.log('***jwt token***')
  //  console.log(decoded)
  //
  //  return NextResponse.next()
}

export const config = {
  // tmp
  matcher: ['/(.*)'],
}
