export default function ({ redirect, store, route }) {
  const userInfo = localStorage.getItem('userInfo')
  if (userInfo) {
    const userInfoJSON = JSON.parse(userInfo)
    if (userInfoJSON.name === '') {
      // logout
      store.commit('setUserInfo', null)
      redirect('/login')
    } else {
      store.commit('getUserInfo')
      if (route.path === '/login' || route.path === '/init') {
        redirect('/')
      }
    }
  } else if (route.path.indexOf('/admin') > -1) {
    // logout
    redirect('/login')
    store.commit('getUserInfo')
  }
}
