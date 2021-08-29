<template>
    <view>
        <template v-if="tabIndex == 0">
            <view class="u-f-ajc" @click="goPeoList(item.class_id)">
                <view class="classList" :style="{'border-left-color':item.leftColor}">
                    <view class="classList-title u-f u-f-jsb">
                        <text class="classList-name">{{item.class_name}}</text>
                        <view v-if="item.peonum == item.realPeoNum">
                            <image class="classList-img" src="../static/images/classList/manyuan.png"></image>
                        </view>
                    </view>
                    <view class="classList-info u-f">
                        <text class="info-one">容纳人数：{{item.class_capacity}}人 | 实际人数：</text>
                        <text style="color: #fc7c1c;" class="info-one">{{item.realPeoNum}}人</text>
                    </view>
                    <view class="classList-btn u-f u-f-jsb">
                        <view class="classBtn">
                            <text class="cuIcon-writefill"></text>
                            <text @tap.stop="classModal(index)" data-target="editClass">编辑</text>
                        </view>
                        <view class="classBtn">
                            <text class="cuIcon-forwardfill"></text>
                            <text @tap.stop="shareFc(index)">邀请</text>
                        </view>
                        <view class="classBtn">
                            <text class="cuIcon-deletefill"></text>
                            <text @tap.stop="delClass(item.class_id)">移除</text>
                        </view>
                    </view>
                </view>
            </view>
        </template>
        <template v-else-if="tabIndex == 1">
            <view class="u-f-ajc" @click="goPeoList(item.class_id)">
                <view class="classList" :style="{'border-left-color':item.leftColor}">
                    <view class="classList-title u-f u-f-jsb">
                        <text class="classList-name">{{item.class_name}}</text>
                        <view v-if="item.peonum == item.realPeoNum">
                            <image class="classList-img" src="/static/images/classList/manyuan.png"></image>
                        </view>
                    </view>
                    <view class="classList-info u-f">
                        <text class="info-one">容纳人数：{{item.class_capacity}}人 | 实际人数：</text>
                        <text style="color: #fc7c1c;" class="info-one">{{item.realPeoNum}}人</text>
                    </view>
                    <view class="classList-btn u-f u-f-jsb">
                        <view class="classBtn">
                            <text class="cuIcon-forwardfill"></text>
                            <text @tap.stop="shareFc(index)">邀请</text>
                        </view>
                        <view class="classBtn">
                            <text class="cuIcon-deletefill"></text>
                            <text @tap.stop="outClass(item.class_id)">退出</text>
                        </view>
                    </view>
                </view>
            </view>
        </template>
    </view>
</template>

<script>
	export default {
        props:{
            item:Object,
            index: Number,
            tabIndex: Number
        },
        methods:{
            goPeoList: function(class_id) {
                this.$emit('goPeoList', class_id)
            },
            classModal: function(index) { //班级modal
                this.$emit('classModal', index)
            },  
            delClass: function(class_id) { //删除班级
                this.$emit('delClass', class_id)
            },
            outClass: function(class_id) { //退出班级
                this.$emit('outClass', class_id)
            },
            async shareFc(index) {
                this.$emit('shareFc', index)
            }
        }
	}
</script>

<style scoped>
    .classList{
        width: 94%;
        margin: 0 0 20rpx 0;
        border-radius: 15rpx;
        background-color: #fff;
        box-shadow:0rpx 2rpx 2rpx #aaaaaa;
        border-left-style: solid;
        border-left-width: 15rpx;
    }
    .classList-title, .classList-info, .classList-btn{
        padding: 12rpx;
        font-size: 38rpx;
    }
    .classList-img{
        width: 50rpx;
        height: 50rpx;
    }
    .classList-info text{
        font-size: 30rpx;
    }
    .classBtn{
        background: #F2F2F2;
        border-radius: 50rpx;
        padding: 2rpx 15rpx;
        font-size: 28rpx;
        color: #969696;
    }
</style>
