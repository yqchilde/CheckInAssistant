import Vue from 'vue'
const NetWork = {
	// 网络状态
	isConnect:false,
	// 监听网络状态
	On(){
		// 获取当前网络状态
		uni.getNetworkType({
			success: (res) => {
				if(res.networkType!=='none'){ this.isConnect=true; return;}
				uni.showToast({
					icon:"none",
					title: '请先连接网络',
				});
			}
		})
		// 监听网络状态变化
		uni.onNetworkStatusChange((res)=>{
			this.isConnect = res.isConnected;
			if(!res.isConnected){
				uni.showToast({
					icon:"none",
					title: '您目前处于断网状态',
				});
			}
		})
	}
}

// app更新
const Update = function(showToast = false){
    var apiServer = Vue.prototype.apiServer;
	// #ifdef APP-PLUS  
	plus.runtime.getProperty(plus.runtime.appid, function(widgetInfo) {  
	    uni.request({  
	        url: apiServer + 'index&m=checkVersion',
            method:"POST",
            header: {'content-type': 'application/json'},
	        data: {  
	            version: widgetInfo.version,  
	        },  
	        success: (res) => {
                res = res.data
                if (res.status == 'error') {
                    if (showToast) {
                        return uni.showToast({ title: res.data,icon:"none" })
                    }
                } else if (res.status == 'ok') {
                    // 为了换行
                    let notes = res.data.notes.replace(/\\n/g,"\n");
                    uni.showModal({
                        title: '发现新的版本 ' + res.data.version,
                        content: notes,
                        cancelText: '放弃更新',
                        confirmText: '立即更新',
                        success: (result) => {
                            if(result.confirm){
                                uni.showLoading({
                                    title: '开始下载',
                                    mask: false
                                });
                                var new_version = res.data.version.split(".");
                                var old_version = widgetInfo.version.split(".");
                                // 判断是否是大版本更新
                                if (old_version[0] < new_version[0]) {
                                    // 大更新默认是整包更新，通常是多了模块等
                                    plus.runtime.openURL(res.data.url);  
                                } else {
                                    // 小更新默认都是热更新
                                    uni.downloadFile({
                                        url: res.data.url,  
                                        success: (downloadResult) => {  
                                            if (downloadResult.statusCode === 200) {  
                                                plus.runtime.install(downloadResult.tempFilePath, {  
                                                    force: false  
                                                }, function() {  
                                                    uni.hideLoading();
                                                    plus.runtime.restart();  
                                                }, function(e) {
                                                    uni.hideLoading();
                                                    uni.showToast({title:'下载失败',icon:'none'});
                                                });  
                                            }  
                                        }  
                                    }); 
                                }
                            }
                        }
                    });
                } 
	        }  
	    });  
	});  
	// #endif 
}

export default {
	NetWork,
	Update
}