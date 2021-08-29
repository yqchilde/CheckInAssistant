<template>
	<view>
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">设置</block>
		</cu-custom>
        
        <view class="body">
            <block v-for="(item,index) in list" :key="index">
                <home-list-item :item="item"
                :index="index"></home-list-item>
            </block>
            <button class="user-set-btn" type="primary" @tap="logout">退出登陆</button>
        </view>
        
	</view>
</template>

<script>
    import homeListItem from "../../components/home-list-item.vue";
	export default {
        components:{
            homeListItem
        },
		data() {
			return {
				list:[
                    { icon:"",name:"资料编辑",clicktype:"navigateTo",url:"../../pages/me/me" },
                    { icon:"",name:"帮助手册",clicktype:"navigateTo",url:"../../pages/set-help/set-help" },
                    { icon:"",name:"版本检测",clicktype:"update",url:"",info:"" },
                    { icon:"",name:"开发作者",clicktype:"",url:"",info:"于波" },
                ]
			}
		},
        onLoad() {
        	// #ifdef APP-PLUS
        	plus.runtime.getProperty(plus.runtime.appid, (widgetInfo)=>{ 
        		this.list[2].info = widgetInfo.version
        	})
        	// #endif
        },
		methods: {
			logout(e) {
                uni.clearStorageSync();
                uni.reLaunch({
                    url:'../login/login'
                })
            }
		}
	}
</script>

<style>
    page {
        background: #FFFFFF;
    }
    .user-set-btn{
    	width: 100%;
    	margin: 20upx 0;
    	background: #FFE933!important;
    	border: 0!important;
    	color: #333333!important;
    }
    .user-set-btn-disable{
    	background: #F4F4F4!important;
    	border: 1upx solid #EEEEEE!important;
    	color: #909090!important;
    }
    
    .body{
    	padding: 0 20upx!important;
    }
    .common-input{
    	font-size: 30upx;
    	border-bottom: 1upx solid #F4F4F4;
    }
</style>
