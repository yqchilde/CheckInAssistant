<template>
    <form @submit="formSubmit">
        <cu-custom></cu-custom> 
        <view class="zai-box">
            <image src="../../static/images/login/register.png" mode='aspectFit' class="zai-logo"></image>
            <view class="zai-title">签到助手</view>
            <view class="zai-form">
                <input type="number" class="zai-input" name="phoneNumber" placeholder-class placeholder="请输入手机号码" />
                <input password class="zai-input" name="password" placeholder-class placeholder="请输入密码" />
                <button class="zai-btn" form-type="submit">立即注册</button>
                <navigator url="../login/login" open-type='navigateBack' hover-class="none" class="zai-label">已有账号，点此去登录.</navigator>
            </view>
        </view>
    </form>
</template>

<script>
    var phoneNumber, password, verifyCode, that;
    export default {
        data() {
            return {
            }
        },
        onLoad() {
            that = this;
        },
        methods: {
            formSubmit: function(e) { // 检验表单数据
                phoneNumber = e.detail.value.phoneNumber;
                password = e.detail.value.password;
                // 正则表达式验证
                let phoneNumberReg = /^1[3456789]\d{9}$/;
                let passwordReg = /^(\w){8,16}$/;
                
                if (phoneNumber == '') {
                    uni.showToast({
                        title: '手机号不能为空',
                        duration: 1500,
                        icon: 'none'
                    });
                    return false; 
                } else if (!phoneNumberReg.test(phoneNumber)) {
                    uni.showToast({
                        title: '手机号格式不正确',
                        duration: 1500,
                        icon: 'none'
                    });
                    return false;
                } else if (password == '') {
                    uni.showToast({
                        title: '密码不能为空',
                        duration: 1500,
                        icon: 'none'
                    });
                    return false;
                } else if (!passwordReg.test(password)) {
                    uni.showToast({
                        title: '密码长度6~20且只能由字母、数字、下划线组成',
                        duration: 1500,
                        icon: 'none'
                    });
                    return false;
                } else {
                    uni.showLoading();
                    uni.request({
                        url: that.apiServer + 'user/register',
                        method: 'POST',
                        header: {
                            'content-type': 'application/json'
                        },
                        data: {
                            phone: phoneNumber,
                            password: password,
                        },
                        // 请求成功执行的代码
                        success: res => {
                            res = res.data;
							console.log(res);
                            if (res.status == "failed") {
                                // 注册失败
                                uni.hideLoading();
                                uni.showToast({
                                    title:"注册失败",
                                    image: '../../static/images/inc/icon_warning.png',
                                    duration: 2000
                                })
                                return false;
                            } else if (res.status == "success") {
                                // 注册失败
                                uni.hideLoading();
                                uni.showToast({
                                    title:"注册成功",
                                    image: '../../static/images/inc/icon_success.png',
                                    duration: 2000
                                });
                                setTimeout(function() {
                                    uni.reLaunch({
                                    url:'../login/login?mode=register&phone='+phoneNumber+'&password='+password,
                                    })
                                },2500);
                            }
                        },
                    });
                }
            },
        }
    }
</script>

<style>
    page{
        background: white;
    }
    .zai-box {
        padding: 0 100upx;
        position: relative;
    }

    .zai-logo {
        width: 100%;
        width: 100%;
        height: 300upx;
    }

    .zai-title {
        position: absolute;
        top: 0;
        line-height: 360upx;
        font-size: 68upx;
        color: #fff;
        text-align: center;
        width: 100%;
        margin-left: -100upx;
    }

    .zai-form {
        margin-top: 250upx;
    }

    .zai-input {
        background: #e2f5fc;
        margin-top: 30upx;
        border-radius: 100upx;
        padding: 20upx 40upx;
        font-size: 36upx;
        -webkit-box-sizing: content-box;
        box-sizing: content-box;
    }

    .input-placeholder,
    .zai-input {
        color: #94afce;
    }

    .zai-label {
        padding: 60upx 0 30upx 0;
        text-align: center;
        font-size: 30upx;
        color: #a7b6d0;
    }

    .zai-btn {
        background: #ff65a3;
        color: #fff;
        border: 0;
        border-radius: 100upx;
        font-size: 36upx;
        margin-top: 50upx;
    }

    .zai-btn:after {
        border: 0;
    }

    /*验证码输入框*/
    .zai-input-btn {
        position: relative;
    }

    .zai-input-btn .zai-input {
        padding-right: 260upx;
    }

    .zai-checking {
        position: absolute;
        right: 0;
        top: 0;
        background: #ff65a3;
        color: #fff;
        border: 0;
        border-radius: 110upx;
        font-size: 36upx;
        margin-left: auto;
        margin-right: auto;
        padding-left: 28upx;
        padding-right: 28upx;
        box-sizing: border-box;
        text-align: center;
        text-decoration: none;
        line-height: 2.55555556;
        -webkit-tap-highlight-color: transparent;
        overflow: hidden;
        padding-top: 2upx;
        padding-bottom: 2upx;
    }

    .zai-checking.zai-time {
        background: #a7b6d0;
    }

    /*按钮点击效果*/
    .zai-btn.button-hover {
        transform: translate(1upx, 1upx);
    }

    .verifyCode {
        display: flex;
        flex-flow: nowrap;
        justify-content: center;
    }

    #verifyCode {
        height: 100upx;
        width: 200upx;
        margin-bottom: 10upx;
    }
</style>
