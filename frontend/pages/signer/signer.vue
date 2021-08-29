<template>
	<view>
		<cu-custom bgColor="bg-blue" :isBack="true">
			<block slot="content" class="text-xl">签到</block>
		</cu-custom>
		<view class="top u-f-ac u-f-column">
			<view class="item u-f-ajc u-f-jsb">
				<text class="top-left">课堂名称</text>
				<text class="top-right">{{roomName}}</text>
			</view>
			<view class="item u-f-ajc u-f-jsb">
				<text class="top-left">发起人</text>
				<text class="top-right">{{uName}}</text>
			</view>
			<view class="item u-f-ajc u-f-jsb">
				<text class="top-left">时间</text>
				<text class="top-right">{{times}}</text>
			</view>
			<view class="item u-f-ajc u-f-jsb">
				<text class="top-left">签到方式</text>
				<text class="top-right">按{{checkInWayText}}方式签到</text>
			</view>
		</view>


		<!-- 创建者打开且为签到码签到方式 -->
		<template v-if="openMode=='create' && checkInWay == 'code'">
			<view class="signCode u-f-ac u-f-column">
				<view class="label">签到码</view>
				<view class="flex">
					<view v-for="(item, index) in checkInCode" :key="index" class="grid flex align-center justify-center">{{item}}</view>
				</view>
				<view class="signCode-btn u-f-ajc" @tap="viewSign(checkInID)">
					查
				</view>
			</view>
		</template>
		<!-- 创建者打开且为GPS签到方式 -->
		<template v-if="openMode=='create' && checkInWay == 'gps'">
			<view class="margin-top-xl u-f-ac u-f-column">
				<map style="width: 100%; height: 200px;" :latitude="latitude" :longitude="longitude" :markers="covers"></map>
				<view class="signCode-btn u-f-ajc" @tap="viewSign(checkInID)">
					查
				</view>
			</view>
		</template>



		<!-- 参与者打开且为签到码签到方式 -->
		<template v-else-if="openMode=='join' && checkInWay == 'code'">
			<view class="signCode u-f-ac u-f-column" style="margin-top: 50rpx;">
				<view class="label">请输入签到码，完成签到</view>
				<view class="text-lg" style="color:#9B9B9B;">（签到码由发起者获取并告知）</view>
				<view class="flex">
					<number-box :maxlength="4" @finish="getCode"></number-box>
				</view>
			</view>
			<view v-if="signStatus == 0" class="u-f-ajc">
				<view class="btnStart u-f-ajc" @tap="toSignIn(checkInID)">
					立即签到
				</view>
			</view>
			<view v-else-if="signStatus == 1 || signStatus == 3" class="u-f-ajc">
				<view class="btnEnd u-f-ajc u-f-column" @tap="viewSign(checkInID)">
					已签到
					<view class="text-sm">（点击查看签到情况）</view>
				</view>
			</view>
			<view v-else-if="signStatus == 2" class="u-f-ajc">
				<view class="btnEnd u-f-ajc">
					签到结束
				</view>
			</view>
		</template>

		<!-- 参与者打开且为GPS签到方式 -->
		<template v-else-if="openMode=='join' && checkInWay == 'gps'">
			<uni-notice-bar :scrollable="true" :single="true" :speed="40" :text=dizhi />
			<view class="signCode u-f-ac u-f-column" style="margin-top: 50rpx;">
				<view style="color: #F4645F;">温馨提示：为了更精确的定位，请开启WLAN</view>
				<view class="text-lg">[{{rangeTxt}}]</view>
				<image src="/static/images/inc/refresh.png" mode="widthFix" style="height: 60rpx;width: 60rpx;float: right;" @tap="refreshRange"></image>
			</view>
			<view v-if="signStatus == 0" class="u-f-ajc">
				<view class="btnStart u-f-ajc" @tap="toSignIn(checkInID)">
					立即签到
				</view>
			</view>
			<view v-else-if="signStatus == 1 || signStatus == 3" class="u-f-ajc">
				<view class="btnEnd u-f-ajc u-f-column" @tap="viewSign(checkInID)">
					已签到
					<view class="text-sm">（点击查看签到情况）</view>
				</view>
			</view>
			<view v-else-if="signStatus == 2" class="u-f-ajc">
				<view class="btnEnd u-f-ajc">
					签到结束
				</view>
			</view>
			<view v-else-if="signStatus == 4" class="u-f-ajc">
				<view class="btnEnd u-f-ajc">
					立即签到
				</view>
			</view>
		</template>


		<!-- 倒计时 -->
		<view v-if="signStatus == 2 || signStatus == 3" class="timerEnd flex justify-center align-center">
			已 结 束
		</view>
		<view v-else class="timerStart flex justify-center align-center">
			剩 余 时 间：{{ currentTime }}s
		</view>
	</view>
</template>

<script>
	var that;
	import numberBox from "../../components/number-box.vue";
	import gaodeDt from "../../common/gaode_dt.js";
	import uniNoticeBar from '@/components/uni-notice-bar/uni-notice-bar.vue';
	import dayjs from 'dayjs';
	export default {
		components: {
			numberBox,
			uniNoticeBar
		},
		data() {
			return {
				uName: '', //发起者名字
				roomName: '', //课堂名称
				times: '', //签到时间
				checkInWay: '', //签到方式
				checkInWayText: '', //签到方式文字
				checkInCode: '', //签到码
				totalTime: 0, //总时间，单位秒
				recordingTime: 0, //记录时间变量
				currentTime: 0, //显示时间变量
				timeOut: '',
				openMode: '', //打开方式（create是创建者打开，join是参与者打开）
				checkInID: '', //签到Id
				codeSign: '', //连续的签到码
				inputCode: '', //输入的数字
				signStatus: 0, //签到状态（0表示进行中、1表示已签到、2表示签到结束、3表示时间结束且已签到、4表示不在范围内）

				dizhi: '', //地址
				latitude: 0, //纬度
				longitude: 0, //经度
				rangeTxt: '', //范围地址信息
				covers: [{ //标记
					latitude: 0,
					longitude: 0,
					iconPath: './static/images/inc/location.png'
				}]
			}
		},
		onLoad(option) {
			that = this;
			uni.showLoading({
				title: '加载中...',
			});
			this.checkInID = option.checkInID;
			this.openMode = option.openMode;
			setTimeout((e) => {
				this.getSignInfo(this.checkInID);
			}, 1000)
		},
		methods: {
			getCode(val) {
				this.inputCode = val;
			},
			//设置计时器
			checking() {
				//把显示时间设为总时间
				// this.signStatus = 0;
				this.currentTime = this.totalTime;
				//执行倒计时
				this.checkingTime();
			},
			//计时器开始计时
			checkingTime() {
				clearTimeout(this.timeOut);
				//判断显示时间是否已到0，判断记录时间是否已到总时间
				if (this.currentTime > 0 && this.recordingTime <= this.totalTime) {
					//记录时间增加 1
					this.recordingTime = this.recordingTime + 1;
					//显示时间，用总时间 - 记录时间
					this.currentTime = this.totalTime - this.recordingTime;
					//1秒钟后，再次执行本方法
					this.timeOut = setTimeout(() => {
						//定时器内，this指向外部，找不到vue的方法，所以，需要用that变量。
						that.checkingTime();
					}, 1000);
				} else {
					//时间已完成，还原相关变量
					this.recordingTime = 0; //记录时间为0
					// 判断时间结束前时候完成了签到,完成就显示已签到,未完成就显示签到结束
					if (this.signStatus == 1) {
						this.signStatus = 3
					} else if (this.signStatus == 3) {
						this.signStatus == 3
					} else {
						this.signStatus = 2
					}
				}
			},
			//开始获取签到信息
			getSignInfo(checkInID) {
				let userId = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'checkIn/detail/' + checkInID + '/user/' + userId,
					method: 'GET',
					header: {
						'content-type': 'application/json',
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							res = res.data;
							this.roomName = res.room_name;
							this.uName = res.creator_name;
							this.signStatus = res.status;
							this.times = dayjs(res.begin_time).format('YYYY-MM-DD HH:mm') + '-' + dayjs(res.end_time).format('HH:mm');

							//倒计时时间
							this.totalTime = parseInt(dayjs(res.end_time).unix() - Math.floor(new Date / 1000));
							this.checking();
							uni.hideLoading();
							//签到方式
							switch (res.check_in_way) {
								case 0:
									this.checkInWay = 'code';
									this.checkInWayText = '签到码'
									// 正则处理签到码
									let code = res.check_in_code;
									this.codeSign = code;
									let regexp = /\d{1}/g,
										match;
									let signCode = new Array();
									while ((match = regexp.exec(code)) != null) {
										signCode.push(match[0]);
									}
									this.checkInCode = signCode;
									break;
								case 1:
									this.checkInWay = 'gps';
									this.checkInWayText = 'GPS'
									this.latitude = res.latitude;
									this.longitude = res.longitude;
									//获取原来签到的目的地
									var point = new plus.maps.Point(res.longitude, res.latitude);
									plus.maps.Map.reverseGeocode(
										point, {},
										function(event) {
											var address = event.address; // 转换后的地理位置
											var point = event.coord; // 转换后的坐标信息
											var coordType = event.coordType; // 转换后的坐标系类型
											that.dizhi = '目标地点：' + address;
										},

									);
									this.covers[0].latitude = res.latitude;
									this.covers[0].longitude = res.longitude;
									// 刷新签到范围
									this.refreshRange();
									break;
								default:
									break;
							}
						} else if (res.status == 'failed') {
							res = res.data;
						}
					},
				});
			},
			//点击签到
			toSignIn(signId) {
				var signInWay = this.checkInWay;
				switch (signInWay) {
					case 'code':
						if (this.inputCode === this.codeSign) {
							this.signOrder(signId);
						} else {
							uni.showToast({
								title: '签到码错误',
								icon: 'none'
							});
							return;
						}
						break;
					case 'gps':
						this.signOrder(signId);
						break;
					case 'face':
						break;
					default:
						break;
				}
			},
			//签到提交
			signOrder(signId) {
				let userId = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'checkIn/startCheckIn',
					method: 'POST',
					header: {
						'content-type': 'application/json'
					},
					data: {
						check_in_id: signId,
						user_id: userId
					},
					success: res => {
						res = res.data;
						if (res.status == 'failed') {
							uni.showToast({
								title: "签到失败",
								icon: 'none'
							});
						} else if (res.status == 'success') {
							uni.showToast({
								title: "签到成功",
								icon: 'none'
							});
							this.signStatus = 1;
						}
					}
				});
			},
			//查看签到情况
			viewSign(checkInID) {
				uni.navigateTo({
					url: '../viewSign/viewSign?checkInID=' + checkInID
				})
			},
			// 计算两个经纬度之间的距离
			computedDt(lon, lat) {
				var start = new gaodeDt.LngLat(this.longitude, this.latitude);
				var end = new gaodeDt.LngLat(lon, lat);
				let result = gaodeDt.calculateLineDistance(start, end)
				return result;
			},
			// 刷新签到范围
			refreshRange() {
				//获取当前坐标
				uni.showLoading({
					title: '正在计算距离中...'
				})
				setTimeout(e => {
					uni.getLocation({
						type: 'gcj02',
						success: (res) => {
							//计算当前距离
							let result = this.computedDt(res.longitude, res.latitude);
							if (result <= 50) {
								this.rangeTxt = '当前已进入签到范围内(50米)';
							} else {
								this.rangeTxt = '当前未进入签到范围';
								// 判断当前是否已经签到
								console.log(this.signStatus);
								if (this.signStatus != 3) {
									this.signStatus = 4;
								}

								this.checking();
							}
							uni.hideLoading()
						}
					});
				}, 500);
			}
		}
	}
</script>

<style>
	page {
		background: #FFFFFF;
	}

	.top {
		margin: 30rpx 0;
	}

	.item {
		height: auto;
		width: 95%;
		padding: 5rpx 0;
	}

	.top-left {
		font-size: 40rpx;
	}

	.top-right {
		font-size: 35rpx;
		color: #9B9B9B
	}

	.circle {
		width: 300rpx;
		height: 300rpx;
		margin-top: 50rpx;
		color: #FFFFFF;
		font-size: 45rpx;
		background: #2692ff;
		border-radius: 100%;
	}

	.line {
		width: 100%;
		height: 10rpx;
		margin-top: 50rpx;
		background: #CCCCCC;
	}

	/* 签到码 */
	.signCode {
		margin-top: 150rpx;
	}

	.signCode .label {
		font-size: 50rpx;
	}

	.signCode .grid {
		margin-top: 10rpx;
		width: 100rpx;
		height: 100rpx;
		background: #EBF5FF;
		border: solid 1rpx #C7E7FE;
		font-size: 50rpx;
		color: #6CB3F4;
	}

	.signCode-btn {
		width: 100rpx;
		height: 100rpx;
		border-radius: 50%;
		margin-top: 30rpx;
		background: #0081ff;
		font-size: 40rpx;
		color: #FFFFFF;
		box-shadow: 0 0 5px 0 rgba(0, 0, 0, 0.8);
	}

	.timerStart {
		width: 100%;
		height: 100rpx;
		background: #E35261;
		color: #FFFFFF;
		font-size: 48rpx;
		position: fixed;
		bottom: 0;
	}

	.timerEnd {
		width: 100%;
		height: 100rpx;
		background: #B2B2B2;
		color: #FFFFFF;
		font-size: 48rpx;
		position: fixed;
		bottom: 0;
	}

	/* 签到按钮 */
	.btnStart {
		margin-top: 50rpx;
		border-radius: 50%;
		width: 300rpx;
		height: 300rpx;
		background: #0081ff;
		font-size: 50rpx;
		color: #FFFFFF;
		box-shadow: 0 0 5px 0 rgba(0, 0, 0, 0.8);
	}

	.btnEnd {
		margin-top: 50rpx;
		border-radius: 50%;
		width: 300rpx;
		height: 300rpx;
		background: #B2B2B2;
		font-size: 50rpx;
		color: #FFFFFF;
		box-shadow: 0 0 5px 0 rgba(0, 0, 0, 0.8);
	}
</style>
