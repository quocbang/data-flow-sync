import { Component, Vue, Watch } from 'vue-property-decorator'
import { Dictionary } from 'vue-router/types/router'
import { Route } from 'vue-router'
import { Form as ElForm, Input } from 'element-ui'
import { UserModule } from '@/store/modules/user'

@Component({
  name: 'Login'
})

export default class extends Vue {
  private loginForm = {
    username: '',
    password: ''
  }
  private passwordType = 'password'
  private loading = false
  private redirect?: string
  private otherQuery: Dictionary<string> = {}

  @Watch('$route', { immediate: true })
  private onRouteChange(route: Route) {
    // TODO: remove the "as Dictionary<string>" hack after v4 release for vue-router
    // See https://github.com/vuejs/vue-router/pull/2050 for details
    const query = route.query as Dictionary<string>
    if (query) {
      this.redirect = query.redirect
      this.otherQuery = this.getOtherQuery(query)
    }
  }

  private getOtherQuery(query: Dictionary<string>) {
    return Object.keys(query).reduce((acc, cur) => {
      if (cur !== 'redirect') {
        acc[cur] = query[cur]
      }
      return acc
    }, {} as Dictionary<string>)
  }

  private handleLogin() {
    (this.$refs.loginForm as ElForm).validate(async (valid: boolean) => {
      if (valid) {
        this.loading = true
        try {
          const checkTag = false
            await UserModule.Login(this.loginForm)
            // eslint-disable-next-line
            this.$router.push({ path: '/' }, () => {})
            setTimeout(() => {
              this.loading = false
            }, 0.5 * 1000)
        } catch (e) {
          console.log(e)
          this.loading = false
        }
      } else {
        return false
      }
    })
    this.loading = false
  }
}