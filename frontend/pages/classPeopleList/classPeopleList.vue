<template>
	<view>
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">班级人员</block>
		</cu-custom>
		<template v-if="tabIndex == 0">
			<scroll-view :scroll-y="modalName==null" class="page" :class="modalName!=null?'show':''">
				<view class="cu-list menu-avatar margin-top">
					<view class="cu-item" :class="modalName=='move-box-'+ index?'move-cur':''" v-for="(item,index) in peoList" :key="index"
					 @touchstart="ListTouchStart" @touchmove="ListTouchMove" @touchend="ListTouchEnd" :data-target="'move-box-' + index">
						<view class="cu-avatar round lg" :style="[{backgroundImage:'url(/static/images/me/header-default.png)'}]"></view>
						<view class="content">
							<view class="text-black">{{item.real_name}}</view>
							<view class="text-gray text-sm">
								<text>{{item.created_at}}加入</text>
							</view>
						</view>
						<view class="action">
							<view class="cuIcon-roundcheckfill text-green"></view>
						</view>
						<view class="move">
							<view class="bg-red" @click="delUser(item.user_id)">删除</view>
						</view>
					</view>
				</view>
			</scroll-view>
		</template>
		<template v-else-if="tabIndex == 1">
			<scroll-view :scroll-y="modalName==null" class="page" :class="modalName!=null?'show':''">
				<view class="cu-list menu-avatar margin-top">
					<view class="cu-item" :class="modalName=='move-box-'+ index?'move-cur':''" v-for="(item,index) in peoList" :key="index">
						<view class="cu-avatar round lg" :style="[{backgroundImage:'url(/static/images/me/header-default.png)'}]"></view>
						<view class="content">
							<view class="text-black">{{item.real_name}}</view>
							<view class="text-gray text-sm">
								<text>{{item.created_at}}加入</text>
							</view>
						</view>
						<view class="action">
							<view class="cuIcon-roundcheckfill text-green"></view>
						</view>
					</view>
				</view>
			</scroll-view>
		</template>
	</view>
</template>

<script>
	var that, classId;
	import dayjs from 'dayjs';
	export default {
		data() {
			return {
				modalName: null,
				listTouchStart: 0,
				listTouchDirection: null,
				peoList: [],
				tabIndex: 0
			}
		},
		onLoad(option) {
			that = this;
			let loginRes = this.checkLogin('../classList/classList', 1);
			if (!loginRes) {
				return;
			}
			this.tabIndex = option.tabIndex;
			classId = option.classId;
			this.getPeoList();
		},
		methods: {
			getPeoList: function() {
				uni.request({
					url: that.apiServer + 'class/users/' + classId,
					method: 'GET',
					header: {
						'content-type': 'application/json'
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							res = res.data;
							if (res == null || res.length == 0) {
								uni.showToast({
									title: '当前班级还未有成员加入',
									icon: 'none'
								});
							} else {
								for (let i = 0; i < res.length; i++) {
									res[i].created_at = dayjs(res[i].created_at).format('YYYY-MM-DD HH:mm:ss');
									this.peoList.push(res[i])
								}
							}
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "获取数据失败",
								icon: 'none'
							});
						}
					},
				});
			},
			showModal(e) {
				this.modalName = e.currentTarget.dataset.target
			},
			hideModal(e) {
				this.modalName = null
			},
			// ListTouch触摸开始
			ListTouchStart(e) {
				this.listTouchStart = e.touches[0].pageX
			},
			// ListTouch计算方向
			ListTouchMove(e) {
				this.listTouchDirection = e.touches[0].pageX - this.listTouchStart > 0 ? 'right' : 'left'
			},
			// ListTouch计算滚动
			ListTouchEnd(e) {
				if (this.listTouchDirection == 'left') {
					this.modalName = e.currentTarget.dataset.target
				} else {
					this.modalName = null
				}
				this.listTouchDirection = null
			},
			delUser: function(userId) {
				let user_id = userId;
				uni.request({
					url: that.apiServer + 'class/out/' + classId + '/user/' + userId,
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
							setTimeout(() => {
								this.peoList = [];
								this.getPeoList(classId)
							}, 500);
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "删除失败",
								icon: 'none'
							});
						}
					},
				});
			}
		}
	}
</script>

<style>
	.page {
		height: 100Vh;
		width: 100vw;
	}

	.page.show {
		overflow: hidden;
	}

	.cu-list>.cu-item.move-cur {
		transform: translateX(-130upx)
	}

	.cu-list>.cu-item .move {
		width: 130upx;
	}
</style>
