export default function ({ redirect, store }) {
  // TODO:
  if (document.cookie.indexOf('solo.go') > -1) {
    return redirect('/login')
  }
}
