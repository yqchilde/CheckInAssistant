<template>
	<view>
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">课堂列表</block>
		</cu-custom>

		<swiper-tab :tabBars="tabBars" :tabIndex="tabIndex" @tabtap="tabtap"></swiper-tab>

		<view class="swiper-top">
			<swiper class="swiper-box" :style="{'height':swiperHeight + 'px'}" :current="tabIndex" @change="tabChange">
				<swiper-item v-for="(items, index) in roomList" :key="index">
					<scroll-view scroll-y class="list" :style="{'height':swiperHeight + 'px'}">
						<template v-if="items.list.length > 0">
							<block v-for="item in items.list" :key="item.room_id">
								<room-list :item="item" :tabIndex="tabIndex" @getRoomDetails="getRoomDetails" @delRoom="delRoom"></room-list>
							</block>
						</template>
						<template v-else>
							<no-thing :swiperHeight="swiperHeight"></no-thing>
						</template>
					</scroll-view>
				</swiper-item>
			</swiper>
		</view>

		<uni-fab :style="{display: floatButton}" ref="fab" :pattern="pattern" :horizontal="horizontal" :vertical="vertical"
		 :popMenu=false @fabClick="toCreate" />
	</view>
</template>

<script>
	var that;
	import uniFab from '@/components/uni-fab/uni-fab.vue';
	import roomList from "../../components/room-list.vue";
	import swiperTab from "../../components/swiper-tab.vue";
	import noThing from "../../components/no-thing.vue";
	import dayjs from 'dayjs';
	export default {
		components: {
			uniFab,
			roomList,
			swiperTab,
			noThing
		},
		data() {
			return {
				tabIndex: 0,
				tabBars: [{
						name: '已创建',
						id: 'create'
					},
					{
						name: '已加入',
						id: 'join'
					},
				],
				swiperHeight: 0,
				leftColor: ['#9152fa', '#2692ff', '#fc7c1c', '#0cd452', '#fc424b'], //左框颜色
				roomList: [{
						list: []
					},
					{
						list: []
					}
				], //课堂列表数据
				horizontal: 'right', //水平向右
				vertical: 'bottom', //垂直向下
				direction: 'horizontal', //弹出方向为水平方向
				pattern: {
					color: '#7A7E83', //阴影颜色
					backgroundColor: '#fff', //加号颜色
					selectedColor: '#9152fa', //选中菜单颜色
					buttonColor: '#9152fa' //悬浮按钮背景颜色
				},
				floatButton: 'inline',
				flag: 0,
				lastX: 0,
				lastY: 0,
			}
		},
		onLoad() {
			that = this;
			let loginRes = this.checkLogin('../roomList/roomList', 1);
			if (!loginRes) {
				return;
			}
			this.tabIndex = 0;
			this.getInfo(this.tabIndex);
			uni.getSystemInfo({
				success: (res) => {
					let height = res.windowHeight - this.CustomBar - uni.upx2px(105);
					this.swiperHeight = height;
				}
			})
		},
		methods: {
			tabtap(index) {
				this.tabIndex = index
			},
			tabChange(e) {
				this.tabIndex = e.detail.current;
				uni.showLoading({
					title: '获取数据...'
				})
				setTimeout(() => {
					this.roomList[this.tabIndex].list = [];
					this.getInfo(this.tabIndex);
					uni.hideLoading();
				}, 500);
			},
			getInfo(tabIndex = 0) { //获取数据信息
				let userId = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'classRoom/user/' + userId,
					method: 'GET',
					header: {
						'content-type': 'application/json'
					},
					data: {
						tabIndex
					},
					success: res => {
						res = res.data;
						console.log(res);
						if (typeof(res.data) == 'undefined' || res.data == null) {
							this.roomList[this.tabIndex].list = [];
							return;
						} else {
							for (let i = 0; i < res.data.length; i++) { //处理一下拼装随机颜色
								var List = res.data[i]; // 接口获取到的所有课堂数据
								let j = i % this.leftColor.length; //排列颜色数组
								let month = dayjs(List.created_at).format('MM');
								let day = dayjs(List.created_at).format('DD');


								List.leftColor = this.leftColor[j];
								List.roomName = month + '.' + day + ' ' + List.room_name;
								List.joinTime = dayjs(List.created_at).format('YYYY-MM-DD HH:mm:ss');
								this.roomList[this.tabIndex].list.push(List);
							}
						}
					},
				});
			},
			toCreate: function(e) {
				uni.navigateTo({
					url: '../createRoom/createRoom',
					success: res => {},
					fail: () => {},
					complete: () => {}
				});
			},
			delRoom: function(roomId) { // 删除课堂
				console.log(roomId);
				uni.showModal({
					title: '温馨提示',
					content: '确定删除该课堂吗？',
					success: (res) => {
						if (res.confirm) {
							uni.request({
								url: that.apiServer + 'classRoom/' + roomId,
								method: 'DELETE',
								header: {
									'content-type': 'application/json'
								},
								success: res => {
									res = res.data;
									if (res.status == 'success') {
										uni.showToast({
											title: "删除成功",
											icon: 'none'
										});
										//重载页面数据
										this.roomList[this.tabIndex].list = [];
										this.getInfo(this.tabIndex);
									} else if (res.status == 'failed') {
										uni.showToast({
											title: "删除失败",
											icon: 'none'
										});
									}
								},
								fail: () => {},
								complete: () => {}
							});
						}
					}
				})
			},
			getRoomDetails: function(roomId) {
				uni.navigateTo({
					url: '../roomDetails/roomDetails?roomId=' + roomId
				})
			}
		}
	}
</script>
<style>
	/* 列表 */
	.swiper-top {
		position: relative;
		top: 90rpx;
	}

	.swiper-box {
		margin-top: 15rpx;
	}
</style>
