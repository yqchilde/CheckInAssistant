<template>
	<form @submit="formSubmit">
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">创建课堂</block>
		</cu-custom>
		<view class="u-f-ajc u-f-column">
			<!-- 课堂名称 -->
			<view class="className">
				<view class="u-f u-f-jsb">
					<text class="className-left">课堂名称</text>
				</view>
				<view class="className-input">
					<input type="text" name="roomName" placeholder="请输入名称" />
				</view>
			</view>
			<!-- 上课班级 -->
			<view class="class">
				<view class="u-f u-f-jsb u-f-ajc">
					<text class="classTitle">上课班级</text>
					<view @tap.stop="showModal" data-target="modalSelectClass">
						<text class="classTitle cuIcon-roundadd text-9152fa"></text>
					</view>
				</view>
				<view class="submit-class u-f-column">
					<view class="classSubmit" v-for="item in submitClass" :key='item.class_id'>
						{{item.class_name}}
					</view>
				</view>
			</view>
			<!-- 选择班级模态框 -->
			<view class="cu-modal padding " :class="modalSelectClass=='modalSelectClass'?'show':''" @tap="hideModal">
				<view class="cu-dialog" @tap.stop>
					<view class="cu-bar solid-bottom ">
						<view class="action">
							<text class="cuIcon-title text-green"></text>选择班级
						</view>
					</view>
					<checkbox-group class="block" @change="CheckboxChange">
						<view class="cu-list menu text-left">
							<view class="cu-item cu-item-error" v-for="item in classList" :key='item.class_id'>
								<view class="title text-black margin-left">{{item.class_name}}</view>
								<checkbox :value="item.class_id" :checked="item.checked"></checkbox>
							</view>
						</view>
					</checkbox-group>
					<view class="padding flex flex-direction ">
						<button class="cu-btn bg-green margin-tb-sm lg" @click="SubmitClass">确 定</button>
					</view>
				</view>
			</view>
			<!-- 发起签到按钮 -->
			<button class="btnOk" formType="submit">创建</button>
		</view>
	</form>
</template>

<script>
	var that, user_id;
	export default {
		data() {
			return {
				classList: [],
				classShow: false,
				modalSelectClass: null,
				selectClass: [],
				submitClass: []
			}
		},
		onLoad() {
			that = this;
			let loginRes = this.checkLogin('../roomList/roomList', 1);
			if (!loginRes) {
				return;
			}
			that.getClass();
		},
		methods: {
			getClass: function(e) { // 获取班级
				user_id = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'class/user/' + user_id,
					method: 'GET',
					header: {
						'content-type': 'application/json',
					},
					data: {
						tabIndex: 0
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							this.classList = res.data;
						}
					},
				});
			},
			showModal: function(e) { //显示模态框
				that.modalSelectClass = e.currentTarget.dataset.target
			},
			hideModal: function(e) { //隐藏模态框
				that.modalSelectClass = null
			},
			SubmitClass: function(e) { //确定班级
				that.modalSelectClass = null;
			},
			CheckboxChange: function(e) { //复选选中
				var classList = this.classList,
					values = e.detail.value;
				var classMap = new Map();
				for (var i = 0, lenI = classList.length; i < lenI; ++i) {
					const item = classList[i];
					if (values.includes(item.class_id)) {
						this.$set(item, 'checked', true);
						classMap.set(item.class_id, item.class_name);
					} else {
						this.$set(item, 'checked', false);
						classMap.delete(item.class_id)
					}
				}

				// 遍历去给map放入array
				this.submitClass = [];
				classMap.forEach(function(value, key) {
					var data = {};
					data["class_id"] = key;
					data["class_name"] = value;
					that.submitClass.push(data);
				});
			},
			// 表单提交代码
			formSubmit: function(e) { // 创建班级
				// 用户ID
				let userID = uni.getStorageSync('USERID');
				// 创建课堂名称
				let roomName = e.detail.value.roomName;
				// 课堂上课班级
				let classList = [];
				this.submitClass.forEach(e => {
					classList.push(e.class_id);
				});
				// 判断是否为空
				if (roomName == '') {
					uni.showToast({
						title: '课堂名称不能为空',
						duration: 1500,
						icon: 'none'
					});
					return false;
				} else if (classList.length == 0) {
					uni.showToast({
						title: '请选择班级',
						duration: 1500,
						icon: 'none'
					});
					return false;
				}
				uni.request({
					url: that.apiServer + 'classRoom',
					method: 'POST',
					header: {
						'content-type': 'application/json',
					},
					data: {
						creator: userID,
						room_name: roomName,
						class_id_list: classList
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "创建成功",
								image: '/static/images/inc/icon_success.png',
								icon: 'none'
							});
							setTimeout(function(e) {
								uni.navigateBack({
									delta: 1
								});
							}, 1500);
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "创建失败",
								image: '/static/images/inc/icon_warning.png',
								icon: 'none'
							});
						}
					},
					fail: (res) => {
						console.log(res);
					}
				});
			}
		}
	}
</script>

<style scoped>
	page {
		background-color: #F2F2F2;
	}

	.className {
		width: 95%;
		height: auto;
		margin: 25rpx 0 0 0;
		border-radius: 15rpx;
		background-color: #fff;
		box-shadow: 0rpx 2rpx 2rpx #aaaaaa;
	}

	.className text {
		padding: 20rpx;
		font-size: 38rpx;
	}

	.className-input {
		padding: 20rpx;
		font-size: 35rpx;
	}

	.class {
		width: 95%;
		height: auto;
		margin: 25rpx 0 0 0;
		border-radius: 15rpx;
		background-color: #fff;
		box-shadow: 0rpx 2rpx 2rpx #aaaaaa;
	}

	.submit-class {
		padding: 0rpx 20rpx 5rpx 20rpx;
		font-size: 35rpx;
		color: #8C8C8C;
	}

	.classSubmit {
		margin: 10rpx 0;
	}

	.class .classTitle,
	.class-choice {
		padding: 20rpx;
		font-size: 38rpx;
	}

	.iconxuanze {
		color: #9152FA;
		display: flex;
	}

	.class .uni-list-cell-db {
		font-size: 35rpx;
	}

	.btnOk {
		width: 300rpx;
		height: auto;
		background: linear-gradient(to left, #9152fa, #67A2FA);
		border-radius: 70rpx;
		box-shadow: 0rpx 5rpx 2rpx #aaaaaa;
		margin-top: 500rpx;
		font-size: 38rpx;
		color: #fff;
		margin: 200rpx 0 50rpx 0;
	}
</style>
