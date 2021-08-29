<template>
    <view class="uni-tab-bar" :style="{'top':tabBarHeight + 'px'}">
        <scroll-view scroll-x="true" class="uni-swiper-tab">
            <block v-for="(tab,index) in tabBars" :key="index">
                <view @tap="tabtap(index)" class="swiper-tab-list" :class="{'active':tabIndex==index}">
                    {{tab.name}}
                    <view class="swiper-tab-line"></view>
                </view>
            </block>
        </scroll-view>
    </view>
</template>

<script>
    export default {
        props:{
            tabBars:Array,
            tabIndex:Number,
        },
        data() {
            return {
                tabBarHeight:this.CustomBar
            }
        },
        methods:{
            tabtap(index) {
                this.$emit('tabtap',index)
            },
        }
    }
</script>

<style scoped>
    .uni-tab-bar{
        height: 90rpx;
        width: 100%;
        position: fixed; 
        z-index: 1;
    }
    .uni-swiper-tab{
        border-bottom: 1rpx solid #EEEEEE;
        background: #FFFFFF;
    }
    .swiper-tab-list{
        color: #969696;
        font-weight: bold;
        width: 50%;
    }
    .uni-tab-bar .active{
        color: #343434;
    }
    .active .swiper-tab-line {
        border-bottom: 6rpx solid #FEDE33;
        width: 70rpx;
        margin: auto;
        border-top: 6rpx solid #FEDE33;
        border-radius: 20rpx;
    }
</style>
