export default function ({ redirect }) {
  // TODO: 
  if (document.cookie.indexOf('solo.go') > -1) {
    return redirect('/login')
  }
}
