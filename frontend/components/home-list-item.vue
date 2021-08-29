<template>
	<view class="home-list-item u-f-ac u-f-jsb animated fadeIn fast" 
	hover-class="home-list-hover"
	@tap="clickevent">
		<view class="u-f-ac">
			<view v-if="item.icon"
			class="icon iconfont"
			:class="['icon-'+item.icon]"></view>
			{{item.name}}
		</view>
		
        <template v-if="item.info">
            <view>{{item.info}}</view>
        </template>
        <template v-else>
            <view class="cuIcon-right"></view>
        </template>
	</view>
</template>

<script>
	export default {
		props:{
			item:Object,
			index:Number
		},
		methods:{
			clickevent(){
				switch (this.item.clicktype){
					case "navigateTo":
                        if(this.item.url){ uni.navigateTo({ url:this.item.url}); }
						break;
					case "switchTab":
                        if(this.item.url){ uni.switchTab({url:this.item.url}) }
						break;
                    case "update":
                        this.lib.Update(true);
                        break;
				}
			}
		}
	}
</script>

<style scoped>
.home-list-item{
	padding: 20upx;
	border-top: 1upx solid #F4F4F4;
	border-bottom: 1upx solid #F4F4F4;
}
.home-list-item>view:first-child{
	color: #333333;
}
.home-list-item>view:first-child>view{
	margin-right: 10upx;
}
.home-list-item>view:last-child{
	color: #9B9B9B
}
.home-list-hover{
	background: #f4f4f4;
}
</style>
