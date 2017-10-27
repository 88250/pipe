export default function ({ redirect, store, route }) {
  const isLogin = store.state.name !== '' && store.state.isInit

  if (route.path.indexOf('/admin') > -1) {
    if (!isLogin) {
      redirect('/login')
    }
  } else if (route.path === '/login') {
    if (isLogin) {
      redirect('/')
      return
    }
    window.location.href = 'https://hacpai.com/login'
  } else if (route.path === '/init' && store.state.isInit) {
    if (isLogin) {
      redirect('/')
      return
    }
    window.location.href = 'https://hacpai.com/register?r=Vanessa'
  }
}
