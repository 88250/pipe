/**
 * @description 图片延迟加载
 * @returns {boolean}
 */
export const LazyLoadImage = () => {
  const loadImg = (it) => {
    const testImage = document.createElement('img')
    testImage.src = it.getAttribute('data-src')
    testImage.addEventListener('load', () => {
      it.src = testImage.src
      it.style.backgroundImage = 'none'
      it.style.backgroundColor = 'transparent'
    })
    it.removeAttribute('data-src')
  }

  if (!('IntersectionObserver' in window)) {
    document.querySelectorAll('.vditor-reset img').forEach((data) => {
      if (data.getAttribute('data-src')) {
        loadImg(data)
      }
    })
    return false
  }

  if (window.imageIntersectionObserver) {
    window.imageIntersectionObserver.disconnect()
    document.querySelectorAll('.vditor-reset img').forEach(function (data) {
      window.imageIntersectionObserver.observe(data)
    })
  } else {
    window.imageIntersectionObserver = new IntersectionObserver((entries) => {
      entries.forEach((entrie) => {
        if ((typeof entrie.isIntersecting === 'undefined' ? entrie.intersectionRatio !== 0 : entrie.isIntersecting) && entrie.target.getAttribute('data-src')) {
          loadImg(entrie.target)
        }
      })
    })
    document.querySelectorAll('.vditor-reset img').forEach(function (data) {
      window.imageIntersectionObserver.observe(data)
    })
  }
}

// 1 - supre admin, 2 - blog admin, 3 - blog user, 4 - prohibit user, 0 - un login user
export const genMenuData = (app, locale) => [
  {
    title: app.$t('home', locale),
    icon: 'home',
    link: '/admin/',
    role: 3
  },
  {
    title: app.$t('postArticle', locale),
    icon: 'add',
    link: '/admin/articles/post',
    role: 3
  },
  {
    title: app.$t('manage', locale),
    icon: 'manage',
    active: true,
    role: 3,
    items: [
      {
        title: app.$t('articleList', locale),
        link: '/admin/articles',
        role: 3
      },
      {
        title: app.$t('commentList', locale),
        link: '/admin/comments',
        role: 3
      },
      {
        title: app.$t('categoryList', locale),
        link: '/admin/categories',
        role: 2
      },
      {
        title: app.$t('navigationList', locale),
        link: '/admin/navigations',
        role: 2
      },
      {
        title: app.$t('tagList', locale),
        link: '/admin/tags',
        role: 2
      }
      /*,
      {
        title: app.$t('userList', locale),
        link: '/admin/users',
        role: 3
      },
      {
        title: app.$t('blogManage', locale),
        link: '/admin/blogs',
        role: 3
      } */
    ]
  },
  {
    title: app.$t('setting', locale),
    icon: 'setting',
    active: app.$route.path.indexOf('settings') > -1,
    role: 2,
    items: [
      {
        title: app.$t('baseInfo', locale),
        link: '/admin/settings/basic',
        role: 2
      },
      {
        title: app.$t('themeList', locale),
        icon: 'theme',
        link: '/admin/themes',
        role: 2
      },
      {
        title: app.$t('preference', locale),
        link: '/admin/settings/preference',
        role: 2
      },
      {
        title: app.$t('signs', locale),
        link: '/admin/settings/sign',
        role: 2
      },
      {
        title: app.$t('account', locale),
        link: '/admin/settings/account',
        role: 2
      },
      {
        title: app.$t('internationalization', locale),
        link: '/admin/settings/i18n',
        role: 2
      },
      {
        title: app.$t('feed', locale),
        link: '/admin/settings/feed',
        role: 2
      },
      {
        title: app.$t('3rdStatistic', locale),
        link: '/admin/settings/3rd-statistic',
        role: 2
      },
      {
        title: app.$t('ad', locale),
        link: '/admin/settings/ad',
        role: 2
      }
    ]
  },
  {
    title: app.$t('others', locale),
    icon: 'inbox',
    link: '/admin/others',
    role: 3
  },
  {
    title: app.$t('about', locale),
    icon: 'info',
    link: '/admin/about',
    role: 3
  }
]

export const initParticlesJS = (id) => {
  window.particlesJS(id, {
    particles: {
      number: {
        value: 80,
        density: {
          enable: true,
          value_area: 800
        }
      },
      color: {
        value: '#1976d2'
      },
      shape: {
        type: 'circle',
        stroke: {
          width: 0,
          color: '#000000'
        },
        polygon: {
          nb_sides: 5
        }
      },
      opacity: {
        value: 0.5,
        random: false,
        anim: {
          enable: false,
          speed: 1,
          opacity_min: 0.1,
          sync: false
        }
      },
      size: {
        value: 3,
        random: true,
        anim: {
          enable: false,
          speed: 40,
          size_min: 0.1,
          sync: false
        }
      },
      line_linked: {
        enable: true,
        distance: 150,
        color: '#1976d2',
        opacity: 0.4,
        width: 1
      },
      move: {
        enable: true,
        speed: 3,
        direction: 'none',
        random: false,
        straight: false,
        out_mode: 'out',
        bounce: false,
        attract: {
          enable: false,
          rotateX: 600,
          rotateY: 1200
        }
      }
    },
    interactivity: {
      detect_on: 'canvas',
      events: {
        onhover: {
          enable: true,
          mode: 'repulse'
        },
        onclick: {
          enable: true,
          mode: 'push'
        },
        resize: true
      },
      modes: {
        grab: {
          distance: 400,
          line_linked: {
            opacity: 1
          }
        },
        bubble: {
          distance: 400,
          size: 40,
          duration: 2,
          opacity: 8,
          speed: 3
        },
        repulse: {
          distance: 200,
          duration: 0.4
        },
        push: {
          particles_nb: 4
        },
        remove: {
          particles_nb: 2
        }
      }
    },
    retina_detect: true
  })
}
