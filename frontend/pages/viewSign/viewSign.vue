<template>
	<view class="u-f-column u-f-ajc">
		<!-- 头部 -->
		<view class="header u-f-ajc">
			<text class="header-text">签到情况</text>
		</view>
		<view class="header-logo u-f-ajc"><text class="cuIcon-rankfill text-9152fa"></text></view>
		<view class="card-head u-f u-f-jsa">
			<view>
				<view class="circle u-f-ajc" style="background: #9152FA;">{{headCount}}</view>
				<view class="" style="color:#9152FA;">总人数</view>
			</view>
			<view>
				<view class="circle u-f-ajc" style="background: #fc7c1c;">{{signCount}}</view>
				<view class="" style="color:#fc7c1c;">已签到</view>
			</view>
			<view>
				<view class="circle u-f-ajc" style="background: #fc424b;">{{headCount - signCount}}</view>
				<view class="" style="color:#fc424b;">未签到</view>
			</view>
		</view>
		<view class="card-list">
			<view class="cu-list menu-avatar">
				<view class="cu-item card-list-item margin-bottom-sm" v-for="(item,index) in signList" :key="index">
					<view class="cu-avatar round lg" :style="[{backgroundImage:'url(/static/images/me/header-default.png)'}]"></view>
					<view class="content">
						<view class="text-black text-lg">{{item.real_name}}</view>
						<view class="text-gray text-sm">
							<text>{{item.signTime}}签到</text>
						</view>
					</view>
					<view class="action" v-if="item.signTime == '未'">
						<text class="cuIcon-roundclosefill text-red" style="font-size: 37rpx;"></text>
					</view>
					<view class="action" v-else>
						<text class="cuIcon-roundcheckfill text-green" style="font-size: 37rpx;"></text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	var that;
	import dayjs from 'dayjs';
	export default {
		data() {
			return {
				checkInID: '',
				headCount: 0,
				signCount: 0,
				signTime: '',
				signList: []
			}
		},
		onLoad(option) {
			that = this;
			this.checkInID = option.checkInID;
			this.viewSign(this.checkInID);
		},
		onPullDownRefresh() {
			setTimeout(() => {
				this.viewSign(this.checkInID);
				uni.stopPullDownRefresh();
			}, 2000);
		},
		methods: {
			viewSign(signId) {
				uni.request({
					url: that.apiServer + 'checkIn/detail/' + signId + '/checkUsers',
					method: 'GET',
					header: {
						'content-type': 'application/json'
					},
					success: res => {
						res = res.data;
						if (res.status == 'failed') {
							uni.showToast({
								title: "数据获取失败",
								icon: 'none'
							});
							return;
						} else if (res.status == 'success') {
							this.headCount = res.data.people_total;
							this.signCount = res.data.check_in_total;
							this.signList = res.data.user_list;
							for (let i = 0; i < this.signList.length; i++) {
								if (this.signList[i].created_at == '0001-01-01T00:00:00Z') {
									this.signList[i].signTime = '未'
								} else {
									this.signList[i].signTime = dayjs(this.signList[i].created_at).format('YYYY-MM-DD HH:mm:ss');
								}
							}
						}
					}
				});
			}
		}
	}
</script>

<style>
	.header {
		height: 300rpx;
		width: 100%;
		background: linear-gradient(to left, #9152fa, #67A2FA);
		border-radius: 0 0 60% 60%/0 0 25% 25%;
		border-top: none;
	}

	.header-text {
		position: fixed;
		top: 70rpx;
		font-size: 40rpx;
		color: #fff;
	}

	.header-logo {
		width: 90rpx;
		height: 90rpx;
		border-radius: 50%;
		background: #FFFFFF;
		box-shadow: 0 0 2px 0 rgba(0, 0, 0, 0.2);
		font-size: 45rpx;
		position: relative;
		top: -135rpx;
		z-index: 1;
	}

	.card-head {
		width: 94%;
		height: auto;
		padding: 80rpx 0 20rpx 0;
		background: #FFFFFF;
		border-radius: 8rpx;
		box-shadow: 0 0 2px 0 rgba(0, 0, 0, 0.2);
		position: relative;
		top: -175rpx;
	}

	.circle {
		width: 90rpx;
		height: 90rpx;
		border-radius: 50%;
		color: #FFFFFF;
		font-size: 35rpx;
	}

	.card-list {
		width: 94%;
		position: relative;
		top: -160rpx;
	}

	.card-list-item {
		margin-top: 12rpx;
		background: #FFFFFF;
		border-radius: 8rpx;
		box-shadow: 0 0 2px 0 rgba(0, 0, 0, 0.2);
	}

	/* scroll */
	.page.show {
		overflow: hidden;
	}

	.cu-list {
		border-radius: 8rpx;
	}
</style>
