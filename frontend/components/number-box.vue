<template>
	<view class="code-area">
		<view class="flex-box">
			<input
				type="number"
				:maxlength="maxlength"
				class="hide-input"
				@input="getVal"
			/>
			<view class="u-f-ajc" :class="['item', { active: codeIndex == 1 }]">
				<view class="line"></view>
				<block>{{codeArr[0]}}</block>
			</view>
			<view class="u-f-ajc" :class="['item', { active: codeIndex == 2 }]">
				<view class="line"></view>
				<block>{{codeArr[1]}}</block>
			</view>
			<view class="u-f-ajc" :class="['item', { active: codeIndex == 3 }]">
				<view class="line"></view>
				<block>{{codeArr[2]}}</block>
			</view>
			<view class="u-f-ajc" :class="['item', { active: codeIndex == 4 }]">
				<view class="line"></view>
				<block>{{codeArr[3]}}</block>
			</view>
		</view>
	</view>
</template>

<script>
export default {
	props: {
		//最大长度 值为4或者6
		maxlength: {
			type: Number,
			default: 4
		}
	},
	data() {
		return {
			codeIndex: 1, //下标
			codeArr: []
		};
	},
	methods: {
		//取值
		getVal(e) {
			let { value } = e.detail;
			let arr = value.split('');
			this.codeIndex = arr.length + 1;
			this.codeArr = arr;
			if (this.codeIndex > Number(this.maxlength)) {
				//输入完成
				this.$emit('finish',this.codeArr.join(''));
			}
		}
	}
};
</script>

<style lang="scss">
.code-area {
	text-align: center;
	.flex-box {
		display: flex;
		flex-wrap: wrap;
		position: relative;
		justify-content: center;
	}

	.item {
		position: relative;
		width: 100upx;
		height: 100upx;
		margin-right: 18upx;
		font-size: 30upx;
		font-weight: bold;
		color: #333333;
		line-height: 100upx;
		box-sizing: border-box;
		border: 2upx solid #cccccc;
	}

	.item:last-child {
		margin-right: 0;
	}

	.active {
		border-color: #ff4b4b;
	}
	.active .line {
		display: block;
	}

	.line {
		display: none;
		transform: translate(-50%, -50%);
		width: 2upx;
		height: 40upx;
		background: #ff4b4b;
		animation: twinkling 1s infinite ease;
	}

	.hide-input {
		position: absolute;
		top: 0;
		left: -100%;
		width: 200%;
		height: 100%;
		text-align: left;
		z-index: 9;
		opacity: 1;
	}

	@keyframes twinkling {
		0% {
			opacity: 0.2;
		}
		50% {
			opacity: 0.5;
		}
		100% {
			opacity: 0.2;
		}
	}
}
</style>
