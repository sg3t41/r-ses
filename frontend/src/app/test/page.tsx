export default function TestPage() {
  return (
    <>
      <a href={'http://localhost:8080/api/auth/github/login'}>
        GitHubでログイン
      </a>
      <br />
      <a href={'http://localhost:8080/api/auth/linkedin/login'}>
        LinkedInでログイン
      </a>
    </>
  )
}
