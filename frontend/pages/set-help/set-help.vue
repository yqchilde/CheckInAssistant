<template>
	<view>
        <cu-custom bgColor="bg-white" :isBack="true">
        	<block slot="content" class="text-xl">帮助手册</block>
        </cu-custom>
		<uni-collapse>
			<block v-for="(item,index) in list" :key="index">
				<uni-collapse-item :title="item.title" :show-animation="true" :contentdata="item.content"></uni-collapse-item>
			</block>
		</uni-collapse>
	</view>
</template>

<script>
	import uniCollapse from "../../components/uni-collapse/uni-collapse.vue"
	import uniCollapseItem from "../../components/uni-collapse-item/uni-collapse-item.vue"
	export default {
		components:{
			uniCollapse,
			uniCollapseItem
		},
		data() {
			return {
				list:[
					{
						title:"怎么加入班级",
						content:"在首页左上角扫码加入班级，扫码是别人分享的班级二维码。"
					},
					{
						title:"如何解散创建的课堂",
						content:"长按已创建的课堂列表内容即可删除。"
					},
					{
						title:"如何退出已加入的课堂",
						content:"只要退出课堂对应的班级即可退出。"
					},
                    {
                    	title:"如何创建课堂及班级",
                    	content:"点击在班级或者课堂列表右下角的悬浮“+”号即可创建。"
                    },
                    {
                    	title:"如何删除历史签到记录",
                    	content:"长按即可删除，但只能删除已发起的记录，已参与的记录只能由发起者删除。"
                    }
				]
			}
		},
		methods: {
			
		}
	}
</script>

<style>
    .uni-collapse-cell__content view {
        color: #6e6a6a;
    }
</style>
