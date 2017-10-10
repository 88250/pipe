export function required (v) {
  return !!v || this.$t('required', this.$store.state.locale)
}

export function maxSize (v, size) {
  return v.length <= size || this.$t('validateRule', this.$store.state.locale).replace('{{size}}', size)
}

export function numberOnly (v) {
  return /^\d+$/.test(v) || this.$t('validateRule3', this.$store.state.locale)
}

export function email (v) {
  return /^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$/.test(v) ||
    this.$t('emailRule', this.$store.state.locale)
}
