import * as utils from '@/utils'

const Header = async () => {
  const currentUser = await utils.jwt.isLogin()

  if (currentUser.isLogin) {
    const { username } = currentUser.data
    return <header>{`${username} でログイン中`}</header>
  } else {
    return <header>ログインしていません</header>
  }
}

export default Header
