<template>
    <view class="nothing u-f-ajc" :style="{'height':swiperHeight + 'px'}">
        <image src="../static/images/inc/nothing.png" mode="widthFix"></image>
    </view>
</template>

<script>
    export default {
        props:{
            swiperHeight:Number
        }
    }
</script>

<style scoped>
    .nothing{
        background: #FFFFFF;
    }
    .nothing image{
        width: 50%;
    }
</style>
