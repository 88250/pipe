export default function ({ redirect, store, route }) {
  if (!store.state.isInit) {
    if (route.path.indexOf('/init') === -1) {
      return redirect('/init')
    }
  }
}
