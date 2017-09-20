export default {
  name: 'icon',
  functional: true,
  render (createElement, content) {
    const props = content.props
    const data = { style: {} }

    const classNames = {}
    classNames['icon'] = true

    if (props.size) {
      if (isNaN(parseInt(props.size, 10))) {
        classNames['icon--medium'] = props.size === 'medium'
        classNames['icon--large'] = props.size === 'large'
        classNames['icon--x-large'] = props.size === 'x-large'
      } else {
        const sizeArray = props.size.split(' ')
        if (sizeArray.length === 2) {
          data.style.width = `${sizeArray[0]}px`
          data.style.height = `${sizeArray[1]}px`
        } else {
          data.style.width = `${props.size}px`
          data.style.height = `${props.size}px`
        }
      }
    }

    if (props.className) {
      let classNameTemp = []

      if (typeof props.className === 'string') {
        classNameTemp.push(props.className)
      } else {
        classNameTemp = props.className
      }

      classNameTemp.forEach((d) => {
        classNames[d] = true
      })
    }
    data.class = classNames

    if (props.color) {
      data.style.color = props.color
    }

    return createElement('svg', data, [
      createElement('use', {
        attrs: {
          'xlink:href': `#${props.icon}`
        }
      })
    ])
  }
}
