<template>
	<view>
		<cu-custom bgColor="bg-white" :isBack="true">
			<block slot="content" class="text-xl">班级列表</block>
		</cu-custom>

		<swiper-tab :tabBars="tabBars" :tabIndex="tabIndex" @tabtap="tabtap"></swiper-tab>

		<view class="swiper-top">
			<swiper class="swiper-box" :style="{'height':swiperHeight + 'px'}" :current="tabIndex" @change="tabChange">
				<swiper-item v-for="(items,index) in classList" :key="index">
					<scroll-view scroll-y class="list" :style="{'height':swiperHeight + 'px'}">
						<template v-if="items.list.length>0">
							<block v-for="(item, index1) in items.list" :key="index1">
								<class-list :item="item" :index="index1" :tabIndex="tabIndex" @outClass="outClass" @goPeoList="goPeoList"
								 @classModal="classModal" @delClass="delClass" @shareFc="shareFc"></class-list>
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
		<!-- 编辑班级 -->
		<form @submit="formEditClass">
			<view class="cu-modal" :class="classModalShow=='editClass'?'show':''">
				<view class="cu-dialog">
					<view class="cu-bar bg-white justify-end">
						<view class="content" style="color: #000; font-size: 35rpx;">修改班级信息</view>
						<view class="action" @tap="classHide">
							<text class="cuIcon-close text-red"></text>
						</view>
					</view>
					<!-- input框 -->
					<view class="bg-white padding-xs">
						<view style="display: flex;justify-content: center">
							<view style="margin-right: 15rpx;">班级名称：</view>
							<view><input type="text" name="newName" class="text-left inputInfo" :value="inputName" placeholder="此处填写班级名称" /></view>
						</view>
						<view style="display: flex;justify-content: center;margin-top: 20rpx;">
							<view style="margin-right: 15rpx;">容纳人数：</view>
							<view><input type="text" name="newNum" class="text-left inputInfo" :value="inputPeoNum" placeholder="此处填写容纳人数" /></view>
						</view>
					</view>
					<view class="cu-bar bg-white justify-end">
						<view class="action">
							<button class="cu-btn line-green text-green" @tap="classHide">取消</button>
							<button form-type="submit" class="cu-btn bg-green margin-left" @tap="classHide">确定</button>
						</view>
					</view>
				</view>
			</view>
		</form>
		<!-- 邀请二维码 -->
		<view class="flex_row_c_c modalView" :class="qrShow?'show':''" @tap="hideQr()">
			<view class="flex_column">
				<view class="backgroundColor-white padding1vh border_radius_10px">
					<image :src="poster.finalPath" mode="widthFix" class="posterImage"></image>
				</view>
				<view class="flex_row marginTop2vh">
					<button type="primary" size="mini" @tap.prevent.stop="saveImage()">保存图片</button>
					<button type="primary" size="mini" @tap.prevent.stop="share()">分享图片</button>
				</view>
			</view>
		</view>
		<view class="hideCanvasView">
			<canvas class="hideCanvas" canvas-id="ShareCanvasId" :style="{width: (poster.width||10) + 'px', height: (poster.height||10) + 'px'}"></canvas>
		</view>
	</view>
</template>

<script>
	var that;
	import _app from '@/common/QS-SharePoster/app.js';
	import {
		getSharePoster
	} from '@/common/QS-SharePoster/QS-SharePoster.js';
	import classList from "../../components/class-list.vue";
	import swiperTab from "../../components/swiper-tab.vue";
	import uniFab from '@/components/uni-fab/uni-fab.vue';
	import noThing from "../../components/no-thing.vue";

	export default {
		components: {
			uniFab,
			classList,
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
				classList: [{
						list: []
					},
					{
						list: []
					}
				], //班级列表数据
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
				classModalShow: null,
				inputName: '',
				inputPeoNum: '',
				postClassId: '',
				poster: {},
				qrShow: false,
				canvasId: 'ShareCanvasId'

			}
		},
		onLoad(option) {
			that = this;
			let loginRes = this.checkLogin('../classList/classList', 1);
			if (!loginRes) {
				return;
			}
			// tabbar定位
			if (typeof(option.tabIndex) != 'undefined') {
				this.tabIndex = option.tabIndex;
			} else {
				this.tabIndex = 0;
			}
			this.getInfo(this.tabIndex);
			uni.removeStorageSync('tmpIndex');
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
					this.classList[this.tabIndex].list = [];
					this.getInfo(this.tabIndex);
					uni.hideLoading();
				}, 500);
			},
			// 获取信息
			getInfo(tabIndex = 0) {
				let user_id = uni.getStorageSync('USERID');
				uni.request({
					url: that.apiServer + 'class/user/' + user_id,
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
						// 判断是否获取到数据
						if (typeof(res.data) == 'undefined' || res.data == null) {
							this.classList[this.tabIndex].list = [];
							return;
						} else {
							for (let i = 0; i < res.data.length; i++) { //处理一下拼装随机颜色
								var List = res.data[i]; // 接口获取到的所有班级数据
								List['realPeoNum'] = List.current_people_number;
								let j = i % this.leftColor.length; //排列颜色数组
								List['leftColor'] = this.leftColor[j];
								this.classList[this.tabIndex].list.push(List);
							}
						}
					},
				});
			},
			// 班级modal
			classModal(index) {
				this.classModalShow = 'editClass';
				this.inputName = this.classList[this.tabIndex].list[index].class_name;
				this.inputPeoNum = this.classList[this.tabIndex].list[index].class_capacity;
				this.postClassId = this.classList[this.tabIndex].list[index].class_id;
			},
			// 班级modal隐藏
			classHide(e) {
				this.classModalShow = null
			},
			// 跳转创建班级页面
			toCreate(e) {
				uni.navigateTo({
					url: '../createClass/createClass',
					success: res => {
						console.log("success");
					},
					fail: (res) => {
						console.log("failed");
					},
					complete: () => {}
				});
			},
			// 修改班级信息
			formEditClass(e) {
				let class_id = this.postClassId;
				let class_name = e.detail.value.newName;
				let class_capacity = Number(e.detail.value.newNum);
				uni.request({
					url: that.apiServer + 'class/' + class_id,
					method: 'PUT',
					header: {
						'content-type': 'application/json'
					},
					data: {
						class_name,
						class_capacity
					},
					success: res => {
						console.log(res);
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "修改成功",
								icon: 'none'
							});
							//重载页面数据
							this.classList[this.tabIndex].list = [];
							this.getInfo();
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "修改失败",
								icon: 'none'
							});
						}
					},
				});
			},
			// 获取班级成员
			goPeoList(classId) {
				uni.navigateTo({
					url: '../classPeopleList/classPeopleList?classId=' + classId + '&tabIndex=' + this.tabIndex
				})
			},
			// 删除班级
			delClass(classId) {
				console.log(classId);
				uni.showModal({
					title: '温馨提示',
					content: '确定删除该班级吗？',
					success: (res) => {
						if (res.confirm) {
							uni.request({
								url: that.apiServer + 'class/del/' + classId,
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
										this.classList[this.tabIndex].list = [];
										this.getInfo(this.tabIndex);
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
				})
			},
			// 退出班级
			outClass(classId) {
				let userId = uni.getStorageSync('USERID');
				uni.showModal({
					title: '温馨提示',
					content: '确定退出该班级吗？',
					success: (res) => {
						if (res.confirm) {
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
											title: "退出成功",
											icon: 'none'
										});
										//重载页面数据
										this.classList[this.tabIndex].list = [];
										this.getInfo(this.tabIndex);
									} else if (res.status == 'failed') {
										uni.showToast({
											title: "退出失败",
											icon: 'none'
										});
									}
								},
							});
						}
					}
				})
			},
			// 分享邀请二维码，参数index是列表点击的索引
			async shareFc(index) {
				let user_id = uni.getStorageSync('USERID');
				let user_name = uni.getStorageSync('REALNAME');
				if (uni.getStorageSync('tmpIndex') === index) { //判断临时索引是否是点击的索引
					this.qrShow = true;
				} else {
					try {
						uni.setStorageSync('tmpIndex', index); //将点击的索引绑定给本地sync，作为临时索引
						var qrMsg = JSON.stringify({
							cid: this.classList[this.tabIndex].list[index].class_id
						});
						var class_name = this.classList[this.tabIndex].list[index].class_name;
						_app.log('准备生成:' + new Date())
						const d = await getSharePoster({
							_this: this, //若在组件中使用 必传
							type: 'testShareType',
							formData: {
								//访问接口获取背景图携带自定义数据

							},
							posterCanvasId: this.canvasId, //canvasId
							delayTimeScale: 20, //延时系数
							drawArray: ({
								bgObj,
								type,
								bgScale
							}) => {
								const dx = bgObj.width * 0.3;
								const fontSize = bgObj.width * 0.075;
								const lineHeight = bgObj.height * 0.04;
								//可直接return数组，也可以return一个promise对象, 但最终resolve一个数组, 这样就可以方便实现后台可控绘制海报
								return new Promise((rs, rj) => {
									rs([{
											type: 'custom',
											setDraw(Context) {
												Context.setFillStyle('black');
												Context.setGlobalAlpha(0.3);
												Context.fillRect(0, bgObj.height - 400, bgObj.width, 400);
												Context.setGlobalAlpha(1);
												Context.setFillStyle('white');
												Context.setFontSize(100);
												//输出班级
												Context.setFillStyle('orange');
												Context.setFontSize(110);
												let textTitle = class_name;
												Context.fillText(textTitle, (bgObj.width - textTitle.length * 110) / 2, bgObj.height - 430);
											}
										},
										{
											type: 'qrcode',
											text: qrMsg, //传入userid和classid，
											size: 580,
											dx: (bgObj.width - 580) / 2,
											dy: bgObj.height - 1200
										},
										{
											type: 'image',
											url: '/static/images/3.jpg',
											alpha: .5,
											dx: 10,
											dy: bgObj.height - 300,
											infoCallBack(imageInfo) {
												let scale = bgObj.width * 0.2 / imageInfo.height;
												return {
													circleSet: {
														x: imageInfo.width * scale / 2,
														y: bgObj.width * 0.2 / 2,
														r: bgObj.width * 0.2 / 2
													}, // 圆形图片 , 若circleSet与roundRectSet一同设置 优先circleSet设置
													dWidth: imageInfo.width * scale, // 因为设置了圆形图片 所以要乘以2
													dHeight: bgObj.width * 0.2,
												}
											}
										},
										{
											type: 'text',
											text: user_name,
											fontWeight: 'bold',
											size: fontSize,
											color: 'white',
											alpha: 0.8,
											textAlign: 'left',
											textBaseline: 'middle',
											infoCallBack(textLength) {
												_app.log('index页面的text的infocallback ，textlength:' + textLength);
												return {
													dx: bgObj.width - textLength - fontSize,
													dy: bgObj.height - lineHeight * 3 + lineHeight
												}
											},
											serialNum: 0
										}
									]);
								})
							},
							setCanvasWH: ({
								bgObj,
								type,
								bgScale
							}) => { // 为动态设置画布宽高的方法，
								this.poster = bgObj;
							}
						});
						_app.log('海报生成成功, 时间:' + new Date() + '， 临时路径: ' + d.poster.tempFilePath)
						this.poster.finalPath = d.poster.tempFilePath;
						this.qrShow = true;
					} catch (e) {
						_app.hideLoading();
						_app.showToast(JSON.stringify(e));
						console.log(JSON.stringify(e));
					}
				}
			},
			// 保存图片
			saveImage() {
				// #ifndef H5
				uni.saveImageToPhotosAlbum({
					filePath: this.poster.finalPath,
					success(res) {
						_app.showToast('保存成功');
					}
				})
				// #endif
				// #ifdef H5
				_app.showToast('保存了');
				// #endif
			},
			// 分享
			share() {
				// #ifdef APP-PLUS
				_app.getShare(false, false, 2, '', '', '', this.poster.finalPath, false, false);
				// #endif

				// #ifndef APP-PLUS
				_app.showToast('分享了');
				// #endif
			},
			// 隐藏生成的图片
			hideQr() {
				this.qrShow = false;
			}
		}
	}
</script>
<style>
	.classList {
		width: 90%;
		margin: 0 0 25rpx 0;
		border-radius: 15rpx;
		background-color: #fff;
		box-shadow: 0rpx 2rpx 2rpx #aaaaaa;
		border-left-style: solid;
		border-left-width: 15rpx;
	}

	.classList-title,
	.classList-info,
	.classList-btn {
		padding: 20rpx;
		font-size: 38rpx;
	}

	.classList-img {
		width: 50rpx;
		height: 50rpx;
	}

	.classList-info text {
		font-size: 30rpx;
	}

	.classBtn {
		background: #F2F2F2;
		border-radius: 50rpx;
		padding: 10rpx;
		font-size: 28rpx;
		color: #969696;
	}

	.inputInfo {
		height: 40rpx;
		border: solid 1rpx #eee;
		border-radius: 10rpx;
		padding-left: 10rpx;
		font-size: 28rpx;
	}

	/* 邀请二维码 */
	.hideCanvasView {
		position: relative;
	}

	.hideCanvas {
		position: fixed;
		top: -99999upx;
		left: -99999upx;
		z-index: -99999;
	}

	.flex_row_c_c {
		display: flex;
		flex-direction: row;
		justify-content: center;
		align-items: center;
	}

	.modalView {
		width: 100%;
		height: 100%;
		position: fixed;
		top: 0;
		left: 0;
		right: 0;
		bottom: 0;
		opacity: 0;
		outline: 0;
		transform: scale(1.2);
		perspective: 2500upx;
		background: rgba(0, 0, 0, 0.6);
		transition: all .3s ease-in-out;
		pointer-events: none;
		backface-visibility: hidden;
		z-index: 999;
	}

	.modalView.show {
		opacity: 1;
		transform: scale(1);
		pointer-events: auto;
	}

	.flex_column {
		display: flex;
		flex-direction: column;
	}

	.backgroundColor-white {
		background-color: white;
	}

	.border_radius_10px {
		border-radius: 10px;
	}

	.padding1vh {
		padding: 1vh;
	}

	.posterImage {
		width: 60vw;
	}

	.flex_row {
		display: flex;
		flex-direction: row;
	}

	.marginTop2vh {
		margin-top: 2vh;
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
