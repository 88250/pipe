export default function ({ redirect, store, route }) {
  const isLogin = store.state.name !== ''

  if (route.path.indexOf('/admin') > -1) {
    if (!isLogin) {
      redirect('/login')
    }
  } else if (route.path === '/login') {
    if (isLogin) {
      redirect('/')
    }
    window.location.href = 'https://hacpai.com/login'
  } else if (route.path === '/init') {
    if (isLogin) {
      redirect('/')
    }
    window.location.href = 'https://hacpai.com/register?r=Vanessa'
  }
}
