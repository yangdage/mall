// custom-tab-bar/index.js
Component({
  data: {
    active: 0,
    list: [
      {
        "url": "/pages/main/home/home",
        "icon": "home-o",
        "text": "首页"
      },
      {
        "url": "/pages/main/classify/classify",
        "icon": "apps-o",
        "text": "分类"
      },
      {
        "url": "/pages/main/cart/cart",
        "icon": "cart-o",
        "text": "购物车"
      },
      {
        "url": "/pages/main/mine/mine",
        "icon": "user-o",
        "text": "我的"
      }
    ]
  },
  methods: {
    onChange(e) {
        console.log(e,'e')
        this.setData({ active: e.detail });
        wx.switchTab({
          url: this.data.list[e.detail].url
        });
    },
    init() {
        const page = getCurrentPages().pop();
        this.setData({
        　  active: this.data.list.findIndex(item => item.url === `/${page.route}`)
        });
    }
   }
  })