// pages/login/login.js
import request from '../../utils/request'

Page({

  data: {

  },

  userLogin(e) {
    wx.login({
      async success (res) {
        if (res.code) {
          //发起登录请求
          let response = await request.post('/login', { 
            code: res.code
          })
          if(response.data.code === 200){
            wx.setStorage({ key: "uid", data: response.data.data});
          }
        }
      }
    })
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log(res)
        wx.setStorage({ key: "avatarUrl", data: res.userInfo.avatarUrl });
        wx.setStorage({ key: "nickName", data: res.userInfo.nickName });
        wx.switchTab({ url: '/pages/main/home/home' })
      }
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {

  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})