import { defineStore } from 'pinia'

export const useStore = defineStore('lf', {
  state: () => {
    return {
      theme: true,
      userInfo: {
        id: null,
        nick_name: '默认用户名',
        user_name: '',
        password: '',
        avatar_id: 'https://q.qlogo.cn/g?b=qq&nk=2124427385&s=100',
        tel: '',
        addr: '',
        token: '',
        role: 2,
      }
    }
  },
  actions: {
    // 设置主题
    setTheme() {
      this.theme = !this.theme
      if (this.theme) {
        //白色主题
        document.documentElement.classList.remove('dark')
        localStorage.setItem('theme', "light")
      } else {
        //黑色主题
        document.documentElement.classList.add('dark')
        localStorage.setItem('theme', "dark")
      }
    },
    // 加载主题
    loadTheme() {
      const theme = localStorage.getItem('theme')
      if (theme === "dark") {
        //黑色主题
        this.theme = false
        document.documentElement.classList.add('dark')
        return
      }
      this.theme = true

    },
    // 加载userInfo
    loadUserInfo() {
      const user = localStorage.getItem('user');
      if (user) {
        this.userInfo = JSON.parse(user)
      }
    },
    setUserInfo(userinfo) {
      
      this.userInfo.nick_name = userinfo.nick_name
      this.userInfo.user_name = userinfo.user_name
      this.userInfo.password = userinfo.password
      this.userInfo.avatar_id = userinfo.avatar_id
      this.userInfo.tel = userinfo.tel

      localStorage.setItem('user', JSON.stringify(this.userInfo))
    }

  }
})
