import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false

// 挂载全局方法
// 全局标题栏
import cuCustom from './colorui/components/cu-custom.vue'
Vue.component('cu-custom',cuCustom)

// 全局判断登录
Vue.prototype.checkLogin  = function(backpage, backtype){
	var USERID  = uni.getStorageSync('USERID');
	var PHONE = uni.getStorageSync('PHONE');
	if(USERID == '' || PHONE == ''){
        uni.showToast({title:'未登录即将跳转',image:'/static/images/inc/icon_warning.png'})
        setTimeout(function() {
            uni.redirectTo({
                url:"../login/login?backpage="+backpage+"&backtype="+backtype
            });
        },1000);
		return false;
	}
	return [USERID, PHONE];
}

// 配置全局接口
// Vue.prototype.apiServer = 'https://ktdm.yqqy.top/?token=api2019&c=';
// Vue.prototype.apiServer = 'http://localhost:8000/';
Vue.prototype.apiServer = 'http://192.168.0.100:8000/';

// 配置全局类
import lib from "./common/lib.js"
Vue.prototype.lib = lib


App.mpType = 'app'

const app = new Vue({
    ...App,
    lib
})
app.$mount()
