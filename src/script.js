;(function main() {
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

    if (n < 0 || n > 4) {
      console.error(`out of bounds: ${n}`)
      return
    }

    current = n
    bioContainerEl.innerHTML = bios[n]
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
})()
