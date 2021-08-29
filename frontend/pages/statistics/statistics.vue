<template>
	<view>
		<template v-if="isHaveData == false">
			<cu-custom bgColor="bg-white" :isBack="true">
				<block slot="content" class="text-xl">数据统计</block>
			</cu-custom>
			<no-thing :swiperHeight="swiperHeight"></no-thing>
		</template>
		<template v-else>
			<view class="page-one">
				<text class="cuIcon-close" @tap="backIndex"></text>
			</view>
			<view class="page-two">
				<view class="page-title">数据统计</view>
				<view class="u-f u-f-jsb">
					<view class="class-info">
						<text>{{roomName}}</text><text class="cuIcon-right"></text>
					</view>
					<view class="u-f-ajc">
						<view class="oval" @tap="switcher"><text>切换</text><text class="cuIcon-unfold"></text></view>
					</view>
				</view>
			</view>
			<view class="page-three qiun-charts">
				<canvas canvas-id="canvasPie" id="canvasPie" class="charts" @touchstart="touchPie"></canvas>
			</view>
			<view class="page-four u-f-ajc">
				<view class="card u-f u-f-jsb">
					<view class="card-item u-f-ajc u-f-column">
						<view class="" style="color: #9152FA;">{{headCount}}人</view>
						<view class="">总人数</view>
					</view>
					<view class="card-item u-f-ajc u-f-column">
						<view class="" style="color: #2692FF;">{{signCount}}人</view>
						<view class="">已签到</view>
					</view>
					<view class="card-item u-f-ajc u-f-column">
						<view class="" style="color: #fc7c1c;">{{headCount - signCount}}人</view>
						<view class="">未签到</view>
					</view>
					<view class="card-item u-f-ajc u-f-column">
						<view class="" style="color: #0cd452;">{{signCount / headCount * 100}}%</view>
						<view class="">签到率</view>
					</view>
				</view>
			</view>
		</template>
	</view>
</template>

<script>
	var that, canvaPie = null;;
	import uCharts from '@/components/u-charts/u-charts.js';
	import noThing from "../../components/no-thing.vue";
	import dayjs from 'dayjs';
	export default {
		components: {
			noThing
		},
		data() {
			return {
				cWidth: '',
				cHeight: '',
				pixelRatio: 1,
				serverData: '',
				listIndex: 0, //列表索引
				roomName: '', //课堂名字
				headCount: 0, //总人数
				signCount: 0, //签到人数
				infoList: [], //展示列表信息
				swiperHeight: 0,
				isHaveData: false
			}
		},
		onLoad() {
			that = this;
			
			uni.getSystemInfo({
				success: (res) => {
					console.log(res);
					let height = res.windowHeight - this.CustomBar;
					this.swiperHeight = height;
				}
			})
			
			this.cWidth = uni.upx2px(750);
			this.cHeight = uni.upx2px(500);
			this.getStatistics();
		},
		methods: {
			getStatistics() {
				let userId = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'checkIn/user/' + userId,
					method: 'GET',
					header: {
						'content-type': 'application/json'
					},
					data: {
						tabIndex: 0
					},
					success: (res) => {
						res = res.data;
						// 判断是否等于null，等于null说明没有数据
						if (res.data == null) {
							this.isHaveData = false;
							uni.showToast({
								title: '还没有签到数据，签到以后才可查看',
								icon: 'none'
							});
							return;
						}
						this.isHaveData = true;
						if (res.status == 'success') {
							console.log(res);
							this.infoList = res.data;;
							let InfoList = res.data[0];
							console.log(InfoList);
							this.roomName = InfoList.room_name;
							this.headCount = InfoList.people_total;
							this.signCount = InfoList.check_in_total;
							let Pie = {
								series: []
							};
							Pie.series = [{
									"name": "已签到",
									"data": InfoList.check_in_total
								},
								{
									"name": "未签到",
									"data": InfoList.people_total - InfoList.check_in_total
								},
							];
							that.showPie("canvasPie", Pie);
						}
					},
					fail: (res) => {
						console.log(res);
					}
				});
			},
			showPie(canvasId, chartData) {
				canvaPie = new uCharts({
					$this: that,
					canvasId: canvasId,
					type: 'pie',
					fontSize: 13,
					padding:[15,15,0,15],
					legend:{
						show:true,
						padding:5,
						lineHeight:11,
						margin:0,
					},
					background: '#FFFFFF',
					pixelRatio: that.pixelRatio,
					series: chartData.series,
					animation: true,
					width: that.cWidth * that.pixelRatio,
					height: that.cHeight * that.pixelRatio,
					dataLabel: true,
					extra: {
						pie: {
							border:true,
							borderColor:'#FFFFFF',
							borderWidth:3
							// lableWidth: 10
						}
					},
				});
			},
			touchPie(e) {
				canvaPie.showToolTip(e, {
					format: function(item) {
						return item.name + ':' + item.data
					}
				});
				canvaPie.touchLegend(e,{animation:true});
			},
			backIndex() {
				uni.navigateBack({
					delta: 1
				})
			},
			switcher() {
				var InfoList = this.infoList;
				var itemList = new Array();
				for (let i = 0; i < InfoList.length; i++) {
					itemList.push(dayjs(InfoList[i].begin_time).format('YYYY-MM-DD HH:mm:ss') + ' ' + InfoList[i].room_name);
				};
				uni.showActionSheet({
					itemList: itemList,
					success: function(res) {
						that.roomName = InfoList[res.tapIndex].room_name;
						that.headCount = InfoList[res.tapIndex].people_total;
						that.signCount = InfoList[res.tapIndex].check_in_total;
						let Pie = {
							series: []
						};
						Pie.series = [{
								"name": "已签到",
								"data": InfoList[res.tapIndex].check_in_total
							},
							{
								"name": "未签到",
								"data": InfoList[res.tapIndex].people_total - InfoList[res.tapIndex].check_in_total
							},
						];
						that.showPie("canvasPie", Pie);
					},
					fail: function(res) {
						console.log(res.errMsg);
					}
				});
			}
		}
	}
</script>

<style>
	page {
		background: #FFFFFF;
	}

	.page-one {
		height: 130rpx;
	}

	.cuIcon-close {
		position: fixed;
		right: 50rpx;
		top: 50rpx;
		font-size: 60rpx;
	}

	.page-two {
		height: auto;
		padding: 0 30rpx;
	}

	.page-title {
		font-size: 55rpx;
		font-weight: bold;
	}

	.class-info {
		font-size: 38rpx;
	}

	.oval {
		padding: 0 10rpx;
		border-radius: 40rpx;
		border: solid 1rpx #9152FA;
		color: #9152FA;
	}

	.page-three {
		height: 400rpx;
	}

	.card {
		margin-top: 20rpx;
		height: 480rpx;
		width: 82%;
		flex-wrap: wrap;
	}

	.card-item {
		width: 48%;
		height: 45%;
		box-shadow: 0 0 4px 0 rgba(0, 0, 0, 0.2);
	}

	.card-item>view:first-child {
		font-size: 50rpx;
	}

	.card-item>view:last-child {
		font-size: 30rpx;
	}

	.qiun-charts {
		width: 750upx;
		height: 500upx;
		background-color: #FFFFFF;
	}
	
	.charts {
		width: 750upx;
		height: 500upx;
		background-color: #FFFFFF;
	}
</style>
