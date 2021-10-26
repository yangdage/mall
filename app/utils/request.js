// 通用请求路径basePath
var basePath = 'http://localhost:8000/app'

// 发送GET请求
function get(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: {
        'Content-Type': 'application/json', // 默认值
      },
      method: 'GET',
      success: (res) => {
        console.log(res);
        resolve(res);
      },
      fail: (err) => {
        console.log(err);
        reject(err);
      }
    })
  })
}

// 发送POST请求
function post(url, data = {}){
  return new Promise((resolve, reject) => {
    wx.request({
      url: basePath + url,
      data: data,
      header: {
        'Content-Type': 'application/x-www-form-urlencoded',
        'token': wx.getStorageSync('token')
      },
      method: 'POST',
      success: (res) => {
        console.log(res);
        if(res.statusCode === 401){
          wx.navigateTo({
            url: '/pages/login/login',
          })
        }
        resolve(res);
      },
      fail: (err) => {
        console.log(err);
        reject(err);
      }
    })
  })
}

module.exports = {
  get, post
}
