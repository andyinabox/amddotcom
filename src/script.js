;(function () {
  // elements
  const btnLessEl = document.getElementById('btn-less')
  const btnMoreEl = document.getElementById('btn-more')
  const sliderEl = document.getElementById('slider')
  const bioContainerEl = document.getElementById('bio-container')

  // gather templates
  const bioNames = ['bio-xs', 'bio-sm', 'bio-md', 'bio-lg', 'bio-xl']
  const bios = bioNames.map(
    (name) => document.getElementById(`tmpl-${name}`).innerHTML
  )

  let current = 2 // we start with the middle ('md') bio

  function setBio(n) {
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
    sliderEl.value = n + 1
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

  // not very scientific night mode logic
  const hour = new Date().getHours()
  if (hour < 8 || hour > 6) {
    const linkEl = document.createElement('link')
    linkEl.rel = 'stylesheet'
    linkEl.href = '/night.css'
    document.body.appendChild(linkEl)
  }
})()
