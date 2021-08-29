<template>
    <form @submit="formSubmit">
        <cu-custom bgColor="bg-white" :isBack="true">
        	<block slot="content" class="text-xl">创建班级</block>
        </cu-custom>
        <view class="u-f-ajc u-f-column">
            <!-- 班级名称模块 -->
            <view class="className">
                <view class="u-f u-f-jsb">
                    <text class="className-left">班级名称</text>
                </view>
                <view class="className-input">
                    <input type="text" name="setName" placeholder="请输入名称" />
                </view>
            </view>
            <!-- 容纳人数模块 -->
            <view class="peopleNum">
                <view class="u-f u-f-jsb">
                    <text class="peopleNum-left">容纳人数</text>
                </view>
                <view class="peopleNum-input">
                    <input type="number" name="setPeoNum" placeholder="请输入人数" />
                </view>
            </view>
            <!-- 确定按钮 -->
            <button class="btnOk" formType="submit">确定</button>
        </view>
    </form>
</template>

<script>
    var that,setName,setPeoNum;
	export default {
		data() {
			return {
				
			}
		},
        onLoad() {
            that = this;
            let loginRes = this.checkLogin('../classList/classList', 1);
            if(!loginRes){return ;}
        },
		methods: {
            // 表单提交代码
            formSubmit: function (e) { // 创建班级
                let creator = uni.getStorageSync('USERID');
                let class_name = e.detail.value.setName;
                let class_capacity = Number(e.detail.value.setPeoNum);
                let setPeoNumReg = /^\d{1,3}$/;
                // 判断是否为空
                if (class_name == '') {
                    uni.showToast({title:'班级名称不能为空',duration: 1500,icon: 'none'});
                    return false;
                } else if (class_capacity == '') {
                    uni.showToast({title:'容纳人数不能为空',duration: 1500,icon: 'none'});
                    return false;
                } else if (!setPeoNumReg.test(class_capacity)) {
                    uni.showToast({title:'容纳人数不能超过1000人',duration: 1500,icon: 'none'});
                    return false;
                } else {
                    // request是uni封装的异步请求
                    uni.request({
                        url: that.apiServer + 'class',
                        method: 'POST',
                        header : {
							'content-type': 'application/json',
						},
                        data: {
                            creator,
                            class_name,
                            class_capacity
                        },
                        // 请求成功执行的代码
                        success: res => {
                            res = res.data;
                            if (res.status == 'success') {
                                uni.showToast({title:"创建成功",image: '/static/images/inc/icon_success.png',icon: 'none'});
                                setTimeout(function(e) {
                                    uni.navigateBack({
                                        delta: 1
                                    });
                                },1500);
                            } else if (res.status == 'failed') {
                                uni.showToast({title:res.data,image: '/static/images/inc/icon_warning.png',icon: 'none'});
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
        background-color:#F2F2F2;
    }
    .className, .peopleNum{
        width: 95%;
        height: auto;
        margin: 25rpx 0 0 0;
        border-radius: 15rpx;
        background-color: #fff;
        box-shadow:0rpx 2rpx 2rpx #aaaaaa;
    }
    .className text, .peopleNum text{
        padding: 20rpx;
        font-size: 38rpx;
    }
    .className-input, .peopleNum-input{
        padding: 20rpx;
        font-size: 35rpx;
    }
    .btnOk{
        width: 300rpx;
        height: auto;
        background:linear-gradient(to left,#9152fa,#67A2FA);
        border-radius: 70rpx;
        box-shadow:0rpx 5rpx 2rpx #aaaaaa;
        margin-top: 500rpx;
        font-size: 38rpx;
        color: #fff;
    }
</style>
