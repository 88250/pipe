export default function ({ redirect }) {
  if (document.cookie.indexOf('solo.go') > -1) {
    return redirect('/login')
  }
}
