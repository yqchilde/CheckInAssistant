<template>
	<view>
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">历史签到</block>
		</cu-custom>

		<swiper-tab :tabBars="tabBars" :tabIndex="tabIndex" @tabtap="tabtap"></swiper-tab>

		<view class="swiper-top">
			<swiper class="swiper-box" :style="{'height':swiperHeight + 'px'}" :current="tabIndex" @change="tabChange">
				<swiper-item v-for="(items,index) in signList" :key="index">
					<scroll-view scroll-y class="list" :style="{'height':swiperHeight + 'px'}">
						<template v-if="items.list.length>0">
							<block v-for="(item, index1) in items.list" :key="index1">
								<template v-if="tabIndex == 0">
									<view class="u-f-ajc" @click="goSigner(item.check_in_id)">
										<view class="signList flex justify-between">
											<text class="signList-name" decode>{{item.room_name}}</text>
											<text v-if="item.status == '已结束'">
												<text style="color: #fc424b;">{{item.status}}</text>
											</text>
											<text v-else>
												<text style="color: #57A3FB;">{{item.status}}</text>
											</text>
										</view>
									</view>
								</template>
								<template v-else-if="tabIndex == 1">
									<view class="u-f-ajc" @click="joinSigner(item.check_in_id)">
										<view class="signList flex justify-between">
											<text class="signList-name" decode>{{item.room_name}}</text>
											<text v-if="item.status == '已结束'">
												<text style="color: #fc424b;">{{item.status}}</text>
											</text>
											<text v-else-if="item.status == '已完成'">
												<text style="color: #57A3FB;">{{item.status}}</text>
											</text>
											<text v-else>
												<text style="color: #67C23A;">{{item.status}}</text>
											</text>
										</view>
									</view>
								</template>
							</block>
						</template>
						<template v-else>
							<no-thing :swiperHeight="swiperHeight"></no-thing>
						</template>
					</scroll-view>
				</swiper-item>
			</swiper>
		</view>
	</view>
</template>

<script>
	var that;
	import swiperTab from "../../components/swiper-tab.vue";
	import noThing from "../../components/no-thing.vue";
	import dayjs from 'dayjs';
	export default {
		components: {
			swiperTab,
			noThing
		},
		data() {
			return {
				tabIndex: 0,
				tabBars: [{
						name: '已发起',
						id: 'create'
					},
					{
						name: '已参与',
						id: 'join'
					},
				],
				swiperHeight: 0,
				signList: [{
						list: []
					},
					{
						list: []
					}
				], //课堂列表数据
			}
		},
		onLoad() {
			that = this;
			this.getSignList(this.tabIndex);
			uni.getSystemInfo({
				success: (res) => {
					let height = res.windowHeight - this.CustomBar - uni.upx2px(105);
					this.swiperHeight = height;
				}
			})
		},
		methods: {
			// tabbar点击切换
			tabtap(index) {
				this.tabIndex = index
			},
			// tabbar滑动切换
			tabChange(e) {
				this.tabIndex = e.detail.current;
				uni.showLoading({
					title: '获取数据...'
				})
				setTimeout(() => {
					this.signList[this.tabIndex].list = [];
					this.getSignList(this.tabIndex);
					uni.hideLoading();
				}, 500);
			},
			getSignList(tabIndex = 0) { //获取数据信息
				let userId = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'checkIn/user/' + userId,
					method: 'GET',
					header: {
						'content-type': 'application/json'
					},
					data: {
						tabIndex
					},
					success: res => {
						res = res.data;
						if (res.data == null) {
							uni.showToast({
								title: '还未有签到记录',
								icon: 'none'
							});
						} else {
							for (let i = 0; i < res.data.length; i++) {
								var List = res.data[i];
								console.log(List);
								List.room_name = dayjs(List.begin_time).format('MM.DD') + ' ' + List.room_name;
								
								if (List.status == 0) {
									List.status = '进行中';
								} else if (List.status == 2) {
									List.status = '已结束';
								} else if (List.status == 1 || List.status == 3) {
									List.status = '已完成';
								}
								
								this.signList[this.tabIndex].list.push(List);
							}
						}
					},
					fail: () => {},
					complete: () => {}
				});
			},
			goSigner(checkInID) {
				uni.navigateTo({
					url: '../signer/signer?checkInID=' + checkInID + '&openMode=create'
				})
			},
			joinSigner(checkInID) {
				uni.navigateTo({
					url: '../signer/signer?checkInID=' + checkInID + '&openMode=join'
				})
			}
		}
	}
</script>

<style>
	.signList {
		width: 95%;
		margin: 0 0 25rpx 0;
		border-radius: 15rpx;
		background-color: #fff;
		box-shadow: 0rpx 2rpx 2rpx #aaaaaa;
		padding: 20rpx;
		font-size: 34rpx;
	}

	/* 列表 */
	.swiper-top {
		position: relative;
		top: 90rpx;
	}

	.swiper-box {
		margin-top: 15rpx;
	}
</style>
