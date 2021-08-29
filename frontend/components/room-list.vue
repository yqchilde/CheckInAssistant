<template>
    <view>
        <template v-if="tabIndex == 0">
            <view class="u-f-ajc" @click="getRoomDetails(item.room_id)" @longpress="delRoom(item.room_id)">
                <view class="classList" :style="{'border-left-color':item.leftColor}">
                    <view class="classList-title u-f-jsb">
                        <text class="classList-name" decode>{{item.roomName}}</text>
                    </view>
                    <view class="classList-info">
                        <text>创建时间：</text>
                        <text style="color: #57A3FB;">{{item.joinTime}}</text>
                    </view>
                </view>
            </view>
        </template>
        <template v-else-if="1">
            <view class="u-f-ajc" @click="getRoomDetails(item.room_id)">
                <view class="classList" :style="{'border-left-color':item.leftColor}">
                    <view class="classList-title u-f-jsb">
                        <text class="classList-name" decode>{{item.roomName}}</text>
                    </view>
                    <view class="classList-info">
                        <text>创建时间：</text>
                        <text style="color: #57A3FB;">{{item.joinTime}}</text>
                    </view>
                </view>
            </view>
        </template>
    </view>
	
</template>

<script>
	export default {
		props:{
            item: Object,
            tabIndex: Number
        },
		methods: {
			delRoom: function(room_id) {
			    this.$emit('delRoom', room_id)
			},
			getRoomDetails: function(room_id) {
			    this.$emit('getRoomDetails', room_id)
			}
		}
	}
</script>

<style>
    .classList{
        width: 94%;
        margin: 0 0 25rpx 0;
        border-radius: 15rpx;
        background-color: #fff;
        box-shadow:0rpx 2rpx 2rpx #aaaaaa;
        border-left-style: solid;
        border-left-width: 15rpx;
    }
    .classList .classList-title, .classList-info{
        padding: 12rpx;
        font-size: 38rpx;
    }
    .classList-title .classList-img{
        width: 50rpx;
        height: 50rpx;
    }
    .classList-info text{
        font-size: 30rpx;
    }
</style>
