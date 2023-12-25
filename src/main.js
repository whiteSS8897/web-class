import { createApp } from 'vue'
import './style.css'
import page from './homepage.vue'

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives,
})

createApp(page).use(vuetify).mount('#app')
