<template>
	<view>
		<cu-custom bgColor="bg-blue" :isBack="true">
			<block slot="content" class="text-xl">找回密码</block>
		</cu-custom>
        
        
        <view class="bg-white padding">
        	<view class="cu-steps">
        		<view class="cu-item" :class="index>num?'':'text-blue'" v-for="(item,index) in numList" :key="index">
        			<text class="num" :data-index="index + 1"></text> {{item.name}}
        		</view>
        	</view>
        </view>
        
        <form @submit="formSubmit">
            <view class="margin-tb-sm bg-white">
                <view class="u-f-ac input-box">
                    <text class="margin-left margin-right">手机号</text>
                    <view class="line-shu-gray margin-right"></view>
                    <view class="">
                        <input type="number" name="phoneNumber" @input="nextStep2" placeholder="请输入你的手机号"/>
                    </view>
                </view>
                <!-- 横线 -->
                <view class="u-f-ajc">
                    <view class="line-heng-gray"></view>
                </view>
                
                <view class="u-f-ac input-box">
                    <text class="margin-left margin-right" space="nbsp">邮    箱</text>
                    <view class="line-shu-gray margin-right"></view>
                    <view class="">
                        <input type="text" name="email" @input="nextStep3" placeholder="请输入你设置的邮箱"/>
                    </view>
                </view>
            </view>
            
            <!-- 按钮 -->
            <view class="">
                <button class="btnBack bg-blue" form-type="submit">找回</button>
            </view>
        </form>
        
        
        
        
	</view>
</template>

<script>
    var that;
	export default {
		data() {
			return {
				num: 0,
                numList: [{
                	name: '开始'
                }, {
                	name: '手机号'
                }, {
                	name: '邮箱'
                }, {
                	name: '完成'
                }, ],
			}
		},
        onLoad() {
            that = this;
        },
		methods: {
			NumSteps() {
				this.num= this.num == this.numList.length - 1 ? 0 : this.num + 1				
			},
            nextStep2(e) {
                if (e.target.value != "") {
                    this.num = 1;
                } else {
                    this.num = 0;
                }
            },
            nextStep3(e) {
                if (e.target.value != "") {
                    this.num = 2;
                } else {
                    this.num = 1;
                }
            },
            formSubmit(e) {
                this.num = 3;
                let phoneNumber = e.detail.value.phoneNumber;
                let email = e.detail.value.email;           
                let phoneNumberReg = /^1[3456789]\d{9}$/;
                let emailReg = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
                
                if (phoneNumber == '') {
                    uni.showToast({title:'手机号不能为空',icon:'none',position:'bottom'});
                    return false;
                } else if (!phoneNumberReg.test(phoneNumber)) {
                    uni.showToast({title: '手机格式不正确',icon: 'none',position:'bottom'});
                    return false;
                } else if (email == '') {
                    uni.showToast({title: '邮箱不能为空',icon: 'none',position:'bottom'});
                    return false;
                } else if (!emailReg.test(email)) {
                    uni.showToast({title: '邮箱格式不正确',icon: 'none',position:'bottom'});
                    return false;
                }
                uni.showLoading({title:"请求中..."})
                //向服务器提交找回密码
                uni.request({
                    url: that.apiServer + 'user&m=getbackPwd',
                    method: 'POST',
                    header: {'content-type': 'application/json'},
                    data: {phoneNumber,email},
                    success: res => {
                        res = res.data;
                        if (res.status == 'ok') {
                            setTimeout((e) => {
                                uni.hideLoading();
                                uni.showToast({title: res.data,icon: 'none',position:'bottom'});
                            },1000);
                        } else if (res.status == 'error') {
                            setTimeout((e) => {
                                uni.hideLoading();
                                uni.showToast({title: res.data,icon: 'none',position:'bottom'});
                            },1000);
                            return;
                        }
                    },
                });
            }
		}
	}
</script>

<style>
    .input-box {
        height: 100rpx;
    }
    .input-box>text:first-child,.input-box>view:last-child {
        font-size: 35rpx;
    }
    .line-shu-gray {
        height: 85rpx;
        width: 2rpx;
        background: #EEEEEE;
    }
    .line-heng-gray {
        height: 2rpx;
        width: 90%;
        background: #EEEEEE;
    }
    .btnBack {
        border-radius: 50rpx;
        width: 60%;
        height: auto;
        margin-top: 50rpx;
        font-size: 38rpx;
    }
</style>
