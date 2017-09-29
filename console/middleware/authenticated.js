export default function ({ redirect, store }) {
  // TODO:
  if (document.cookie.indexOf('solo.go') > -1) {
    // logout
    store.commit('setUserInfo', null)
  } else {
    // login
    store.commit('getUserInfo')
  }
}
