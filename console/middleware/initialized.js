export default function ({ redirect, store, route }) {
  if (!store.state.isInit && route.path.indexOf('/init') === -1) {
    return redirect('/init')
  }
}
