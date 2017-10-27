export default function ({redirect, store, route}) {
  if (!store.state.isInit) {
    return
  }
  const isLogin = store.state.role !== 0
  if (route.path.indexOf('/admin') > -1) {
    if (!isLogin || store.state.role === 4) {
      redirect('/')
    }
  } else if (route.path === '/login' || route.path === '/init') {
    redirect('/')
  }
}
