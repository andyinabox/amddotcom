;(function () {
  // elements
  const btnLessEl = document.getElementById('btn-less')
  const btnMoreEl = document.getElementById('btn-more')
  const sliderEl = document.getElementById('slider')
  const bioContainerEl = document.getElementById('bio-container')
  const publishDateEl = document.getElementById('publish-date')

  // gather templates
  const bioNames = ['bio-xs', 'bio-sm', 'bio-md', 'bio-lg', 'bio-xl']
  const bios = bioNames.map(
    (name) => document.getElementById(`tmpl-${name}`).innerHTML
  )

  // default value for current
  let current = 2

  // set to whatever current slider value is
  setBio(Math.floor(sliderEl.value - 1), true)

  function setBio(n, skipUpdateSlider) {
    if (current == n) return // ignore if already set

    // log error if n is out of bounds
    if (n < 0 || n > 4) {
      console.error(`out of bounds: ${n}`)
      return
    }

    // set current index
    current = n

    // remove current bio name classes
    bioNames.forEach((n) => bioContainerEl.classList.remove(n))

    // set current bio and class
    bioContainerEl.innerHTML = bios[n]
    bioContainerEl.classList.add(bioNames[current])

    // set value of slider
    if (!skipUpdateSlider) sliderEl.value = n + 1
  }

  // set listeners

  btnLessEl.addEventListener('click', () => {
    if (current > 0) setBio(current - 1)
  })

  btnMoreEl.addEventListener('click', () => {
    if (current < 4) setBio(current + 1)
  })

  sliderEl.addEventListener('input', (evt) => {
    const v = sliderEl.value
    if (v <= 1.5) {
      setBio(0)
    } else if (v <= 2.5) {
      setBio(1)
    } else if (v <= 3.5) {
      setBio(2)
    } else if (v <= 4.5) {
      setBio(3)
    } else {
      setBio(4)
    }
  })

  // update date formatting
  const pubDate = new Date(publishDateEl.getAttribute('datetime'))
  publishDateEl.innerText = getTimeSince(pubDate, publishDateEl.innerText)

  // not very scientific night mode logic
  const hour = new Date().getHours()
  if (hour < 7 || hour >= 7) enableNightMode()
})()

function enableNightMode() {
  const linkEl = document.createElement('link')
  linkEl.rel = 'stylesheet'
  linkEl.href = '/night.css'
  document.body.appendChild(linkEl)
}

function getTimeSince(pubDate, defaultString) {
  const nowDate = new Date()
  const diff = nowDate - pubDate

  // more than a year old
  if (diff > 1000 * 60 * 60 * 24 * 365) {
    const d = Math.floor(diff / (1000 * 60 * 60 * 24 * 365))
    if (d == 1) {
      return `${d} year ago`
    } else if (d > 1) {
      return `${d} years ago`
    }
  }

  // more than a month old
  if (diff > 1000 * 60 * 60 * 24 * 28) {
    const nowMonth = nowDate.getMonth()
    const pubMonth = pubDate.getMonth()
    let d = nowMonth - pubMonth
    // if it was last year
    if (pubMonth > nowMonth) {
      d = nowMonth + (12 - pubMonth)
    }

    if (d == 1) {
      return `${d} month ago`
    } else if (d > 1) {
      return `${d} months ago`
    }
  }

  // more than a day old
  if (diff > 1000 * 60 * 60 * 24) {
    const d = Math.floor(diff / (1000 * 60 * 60 * 24))
    if (d == 1) {
      return `${d} day ago`
    } else if (d > 1) {
      return `${d} days ago`
    }
  }

  // more than an hour old
  if (diff > 1000 * 60 * 60) {
    const d = Math.floor(diff / (1000 * 60 * 60))
    if (d == 1) {
      return `${d} hour ago`
    } else if (d > 1) {
      return `${d} hours ago`
    }
  }

  // more than a minute old
  if (diff > 1000 * 60) {
    const d = Math.floor(diff / (1000 * 60))
    if (d == 1) {
      return `${d} minute ago`
    } else if (d > 1) {
      return `${d} minutes ago`
    }
  }

  // less than a minute old
  if (diff <= 1000 * 60) {
    return 'just now'
  }

  // when in doubt, use the default string
  return defaultString
}
