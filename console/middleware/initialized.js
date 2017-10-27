export default function ({ redirect, store, route }) {
  if (!store.state.isInit) {
    if (route.path !== '/init') {
      return redirect('/init')
    }
  } else {
    if (route.path === '/init') {
      return redirect('/')
    }
  }
}
