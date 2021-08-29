<template>
	<view>
		<view class="page-one">
			<text class="cuIcon-close" @tap="backIndex"></text>
		</view>
		<view class="header padding">
			<view class="cu-avatar xl round margin-left" style="background-image:url(/static/images/me/header-default.png);"></view>
			<text class="text-xxl padding" @tap="nameModal" data-target="editName" v-if="realName == null || realName == ''">
				点击此处设置名字
			</text>
			<text class="text-xxl padding" v-else>
				{{realName}}
			</text>
		</view>
		<view class="cu-list menu sm-border">
			<view class="cu-item arrow" @tap="userModal" data-target="userInfo">
				<view class="content" hover-class="none">
					<text class="cuIcon-peoplelist text-9152fa"></text>
					<text class="text-grey">个人信息</text>
				</view>
			</view>
			<view class="cu-item arrow" @tap="classModal" data-target="setClass">
				<view class="content" hover-class="none">
					<text class="cuIcon-group text-9152fa"></text>
					<text class="text-grey">绑定班级</text>
				</view>
			</view>
			<view class="cu-item arrow" @tap="emailModal" data-target="setEmail">
				<view class="content" hover-class="none">
					<text class="cuIcon-punch text-9152fa"></text>
					<text class="text-grey">我的邮箱</text>
				</view>
			</view>
			<view class="cu-item arrow" @tap="pwdModal" data-target="editPwd">
				<view class="content" hover-class="none">
					<text class="cuIcon-unlock text-9152fa"></text>
					<text class="text-grey">修改密码</text>
				</view>
			</view>
			<view class="cu-item arrow" @tap="shareFid">
				<view class="content" hover-class="none">
					<text class="cuIcon-friendadd text-9152fa"></text>
					<text class="text-grey">邀请好友</text>
				</view>
			</view>
		</view>
		<!-- 修改密码 -->
		<form @submit="formEditPwd">
			<view class="cu-modal" :class="pwdModalShow=='editPwd'?'show':''">
				<view class="cu-dialog">
					<view class="cu-bar bg-white justify-end">
						<view class="content" style="color: #000; font-size: 35upx;">修改密码</view>
						<view class="action" @tap="pwdHide">
							<text class="cuIcon-close text-red"></text>
						</view>
					</view>
					<view>
						<view class="padding-sm">
							<input password name="oldpwd" style="font-size: 30upx;" placeholder="请输入您的原密码" />
						</view>
						<view class="bg-white" style="height: 25rpx;"></view>
						<view class="padding-sm">
							<input password name="newpwd" style="font-size: 30upx;" placeholder="请输入您的新密码" />
						</view>
						<view class="bg-white" style="height: 25rpx;"></view>
						<view class="padding-sm">
							<input password name="renewpwd" style="font-size: 30upx;" placeholder="请确认您的新密码" />
						</view>

					</view>
					<view class="cu-bar bg-white justify-end">
						<view class="action">
							<button class="cu-btn line-green text-green" @tap="pwdHide">取消</button>
							<button form-type="submit" class="cu-btn bg-green margin-left" @tap="pwdHide">确定</button>
						</view>
					</view>
				</view>
			</view>
		</form>

		<!-- 编辑真实姓名 -->
		<form @submit="formEditName">
			<view class="cu-modal" :class="nameModalShow=='editName'?'show':''">
				<view class="cu-dialog">
					<view class="cu-bar bg-white justify-end">
						<view class="content" style="color: #000; font-size: 35upx;">修改名字</view>
						<view class="action" @tap="nameHide">
							<text class="cuIcon-close text-red"></text>
						</view>
					</view>
					<view class="padding">
						<input type="text" name="realName" style="font-size: 30upx;" placeholder="请输入真实姓名,只有一次机会" />
					</view>
					<view class="cu-bar bg-white justify-end">
						<view class="action">
							<button class="cu-btn line-green text-green" @tap="nameHide">取消</button>
							<button form-type="submit" class="cu-btn bg-green margin-left" @tap="nameHide">确定</button>
						</view>
					</view>
				</view>
			</view>
		</form>

		<!-- 绑定班级 -->
		<form @submit="formSetClass">
			<view class="cu-modal" :class="classModalShow=='setClass'?'show':''">
				<view class="cu-dialog">
					<view class="cu-bar bg-white justify-end">
						<view class="content" style="color: #000; font-size: 35upx;">绑定班级</view>
						<view class="action" @tap="classHide">
							<text class="cuIcon-close text-red"></text>
						</view>
					</view>
					<view class="padding">
						<view v-if="className == null || className == ''">
							<input type="text" name="class" style="font-size: 30upx;" placeholder="绑定你的班级" />
						</view>
						<view v-else>
							<input type="text" name="class" style="font-size: 30upx;" placeholder="绑定你的班级" :value="className" />
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

		<!-- 绑定邮箱 -->
		<form @submit="formSetEmail">
			<view class="cu-modal" :class="emailModalShow=='setEmail'?'show':''">
				<view class="cu-dialog">
					<view class="cu-bar bg-white justify-end">
						<view class="content" style="color: #000; font-size: 35upx;">绑定邮箱</view>
						<view class="action" @tap="emailHide">
							<text class="cuIcon-close text-red"></text>
						</view>
					</view>
					<view class="padding">
						<view v-if="email == null || email == ''">
							<input type="text" name="email" style="font-size: 30upx;" placeholder="绑定你的邮箱" />
						</view>
						<view v-else>
							<input type="text" name="email" style="font-size: 30upx;" placeholder="绑定你的邮箱" :value="email" />
						</view>
					</view>
					<view class="cu-bar bg-white justify-end">
						<view class="action">
							<button class="cu-btn line-green text-green" @tap="emailHide">取消</button>
							<button form-type="submit" class="cu-btn bg-green margin-left" @tap="emailHide">确定</button>
						</view>
					</view>
				</view>
			</view>
		</form>

		<!-- 个人信息 -->
		<view class="cu-modal" :class="userModalShow=='userInfo'?'show':''">
			<view class="cu-dialog">
				<view class="cu-bar bg-white justify-end">
					<view class="content">个人信息</view>
					<view class="action" @tap="userHide">
						<text class="cuIcon-close text-red"></text>
					</view>
				</view>
				<view class="padding-xl text-left">
					<view style="font-size: 30upx;margin-bottom: 10upx;">
						手机号: {{phoneNumber}}
					</view>
					<view style="font-size: 30upx;margin-bottom: 10upx;">
						姓名: {{realName}}
					</view>
					<view style="font-size: 30upx;margin-bottom: 10upx;">
						邮箱: {{email}}
					</view>
					<view style="font-size: 30upx;margin-bottom: 10upx;">
						班级: {{className}}
					</view>
					<view style="font-size: 30upx;margin-bottom: 10upx;">
						注册时间: {{regTime}}
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
				phoneNumber: '', //手机号
				realName: '', //姓名
				email: '', //邮箱
				className: '', //班级
				regTime: '', //注册时间
				nameModalShow: null,
				userModalShow: null,
				classModalShow: null,
				emailModalShow: null,
				pwdModalShow: null
			}
		},
		onLoad() {
			that = this;
			let loginRes = this.checkLogin('../me/me', 1);
			if (!loginRes) {
				return;
			}
			let userid = uni.getStorageSync('USERID');
			let phoneNumber = uni.getStorageSync('PHONENUMBER');
			uni.request({
				url: that.apiServer + 'user/info',
				method: 'GET',
				header: {
					'content-type': 'application/json'
				},
				data: {
					user_id: userid,
				},
				success: res => {
					res = res.data;
					if (res.status == 'success') {
						this.phoneNumber = res.data.phone;
						this.realName = res.data.real_name;
						this.email = res.data.email;
						this.className = res.data.class_id;
						this.regTime = dayjs(res.data.created_at).format('YYYY-MM-DD HH:mm:ss');
					} else if (res.status == 'failed') {
						uni.showToast({
							title: "数据获取失败"
						})
					}
				},
				fail: (res) => {
					console.log(res);
				}
			});
		},
		methods: {
			nameModal: function(e) { //姓名modal
				this.nameModalShow = e.currentTarget.dataset.target
			},
			nameHide: function(e) {
				this.nameModalShow = null
			},
			userModal: function(e) { //信息modal
				this.userModalShow = e.currentTarget.dataset.target
			},
			userHide: function(e) {
				this.userModalShow = null
			},
			classModal: function(e) { //班级modal
				this.classModalShow = e.currentTarget.dataset.target
			},
			classHide: function(e) {
				this.classModalShow = null
			},
			emailModal: function(e) { //邮箱modal
				this.emailModalShow = e.currentTarget.dataset.target
			},
			emailHide: function(e) {
				this.emailModalShow = null
			},
			pwdModal: function(e) { //密码modal
				this.pwdModalShow = e.currentTarget.dataset.target
			},
			pwdHide: function(e) {
				this.pwdModalShow = null
			},
			formEditName: function(e) {
				let userId = uni.getStorageSync('USERID');
				let realName = e.detail.value.realName;
				let realNameReg = /^[\u4e00-\u9fa5]{2,4}$/;
				if (realName == '') {
					uni.showToast({
						title: '姓名不能为空',
						icon: 'none'
					});
					return false;
				} else if (!realNameReg.test(realName)) {
					uni.showToast({
						title: '请输入汉字,且长度为2-4位',
						icon: 'none'
					});
					return false;
				}
				uni.request({
					url: that.apiServer + 'user/info?type=realName',
					method: 'PUT',
					header: {
						'content-type': 'application/json'
					},
					data: {
						user_id: userId,
						real_name: realName
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "修改成功",
								icon: 'none'
							});
							this.realName = realName;
							if (uni.getStorageSync('REALNAME')) {
								uni.removeStorageSync('REALNAME');
							}
							uni.setStorageSync('REALNAME', realName);
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "修改失败",
								icon: 'none'
							})
						}
					},
				});
			},
			formSetClass: function(e) {
				let userId = uni.getStorageSync('USERID');
				let className = e.detail.value.class;
				if (className == '') {
					uni.showToast({
						title: '班级不能为空',
						icon: 'none'
					});
					return false;
				}
				uni.request({
					url: that.apiServer + 'user/info?type=class',
					method: 'PUT',
					header: {
						'content-type': 'application/json'
					},
					data: {
						user_id: userId,
						class_id: className
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "修改成功",
								icon: 'none'
							});
							this.className = className;
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "修改失败",
								icon: 'none'
							})
						}
					},
					fail: (res) => {
						console.log(res);
					}
				});
			},
			formSetEmail: function(e) {
				let userId = uni.getStorageSync('USERID');
				let email = e.detail.value.email;
				let emailReg = /^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$/;
				if (email == '') {
					uni.showToast({
						title: '邮箱不能为空',
						icon: 'none'
					});
					return false;
				} else if (!emailReg.test(email)) {
					uni.showToast({
						title: '邮箱格式不正确',
						icon: 'none'
					});
					return false;
				}
				uni.request({
					url: that.apiServer + 'user/info?type=email',
					method: 'PUT',
					header: {
						'content-type': 'application/json'
					},
					data: {
						user_id: userId,
						email: email
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "修改成功",
								icon: 'none'
							});
							this.email = email;
						} else if (res.status == 'failed') {
							console.log(res);
							uni.showToast({
								title: "修改失败",
								icon: 'none'
							})
						}
					},
					fail: (res) => {
						console.log(res);
					}
				});
			},
			formEditPwd: function(e) {
				let userId = uni.getStorageSync('USERID');
				let oldpwd = e.detail.value.oldpwd;
				let newpwd = e.detail.value.newpwd;
				let renewpwd = e.detail.value.renewpwd;
				let passwordReg = /^(\w){8,16}$/;

				if (oldpwd == '') {
					uni.showToast({
						title: '旧密码不能为空',
						icon: 'none',
						position: 'bottom'
					});
					return false;
				} else if (newpwd == '') {
					uni.showToast({
						title: '新密码不能为空',
						icon: 'none',
						position: 'bottom'
					});
					return false;
				} else if (renewpwd == '') {
					uni.showToast({
						title: '确认密码不能为空',
						icon: 'none',
						position: 'bottom'
					});
					return false;
				} else if (!passwordReg.test(newpwd)) {
					uni.showToast({
						title: '密码为8到16位',
						icon: 'none',
						position: 'bottom'
					});
					return false;
				} else if (renewpwd != newpwd) {
					uni.showToast({
						title: '两次密码不一致',
						icon: 'none',
						position: 'bottom'
					});
					return false;
				}
				uni.request({
					url: that.apiServer + 'user/info?type=password',
					method: 'PUT',
					header: {
						'content-type': 'application/json'
					},
					data: {
						user_id: userId,
						password: newpwd
					},
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.showToast({
								title: "修改成功",
								icon: 'none'
							});
							setTimeout((e) => {
								uni.reLaunch({
									url: '../login/login'
								})
							}, 1000);
						} else if (res.status == 'failed') {
							uni.showToast({
								title: "修改失败",
								icon: 'none'
							})
						}
					},
				});
			},
			// 右上角叉号返回
			backIndex() {
				uni.navigateBack({
					delta: 1
				})
			},
			// 分享好友下载地址
			shareFid() {
				uni.request({
					url: that.apiServer + 'app/share',
					method: 'GET',
					success: res => {
						res = res.data;
						if (res.status == 'success') {
							uni.setClipboardData({
								data: res.data.share_down,
								success: function() {
									uni.showToast({
										title: '已成功复制下载地址',
										icon: 'none',
										position: 'bottom'
									});
								}
							});
						}
					},
				});
			}
		},
	}
</script>

<style>
	page {
		background: white;
	}

	.page-one {
		height: 130rpx;
	}

	.page-one .cuIcon-close {
		position: fixed;
		right: 50rpx;
		top: 50rpx;
		font-size: 60rpx;
	}
</style>
