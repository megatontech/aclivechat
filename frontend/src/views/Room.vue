<template>
  <chat-renderer ref="renderer" :maxNumber="config.maxNumber" :showGiftName="config.showGiftName"  :showGiftPrice="config.showGiftPrice"></chat-renderer>
</template>

<script>
import {mergeConfig, toBool, toInt} from '@/utils'
import * as config from '@/api/config'
import ChatRenderer from '@/components/ChatRenderer'
import * as constants from '@/components/ChatRenderer/constants'

const COMMAND_HEARTBEAT = 0
const COMMAND_JOIN_ROOM = 1
const COMMAND_ADD_TEXT = 2
const COMMAND_ADD_GIFT = 3
const COMMAND_ADD_MEMBER = 4
const COMMAND_ADD_SUPER_CHAT = 5
const COMMAND_DEL_SUPER_CHAT = 6
const COMMAND_UPDATE_TRANSLATION = 7
const COMMAND_ADD_LOVE = 8
const COMMAND_QUIT_ROOM = 9
const COMMAND_ADD_FOLLOW = 10

export default {
  name: 'Room',
  components: {
    ChatRenderer
  },
  data() {
    return {
      config: {...config.DEFAULT_CONFIG},

      websocket: null,
      retryCount: 0,
      isDestroying: false,
      heartbeatTimerId: null
    }
  },
  computed: {
    blockKeywords() {
      return this.config.blockKeywords.split('\n').filter(val => val)
    },
    blockUsers() {
      return this.config.blockUsers.split('\n').filter(val => val)
    }
  },
  created() {
    this.updateConfig()
    this.wsConnect()
    // 提示用户已加载
    this.$message({
      message: 'Loaded',
      duration: '500'
    })
  },
  beforeDestroy() {
    this.isDestroying = true
    this.websocket.close()
  },
  methods: {
    updateConfig() {
      let cfg = {}
      // 留空的使用默认值
      for (let i in this.$route.query) {
        if (this.$route.query[i] !== '') {
          cfg[i] = this.$route.query[i]
        }
      }
      cfg = mergeConfig(cfg, config.DEFAULT_CONFIG)
      cfg.minGiftPrice = toInt(cfg.minGiftPrice, config.DEFAULT_CONFIG.minGiftPrice)
      cfg.showDanmaku = toBool(cfg.showDanmaku)
      cfg.showLove = toBool(cfg.showLove)
      cfg.showFollow = toBool(cfg.showFollow)
      cfg.showJoin = toBool(cfg.showJoin)
      cfg.showQuit = toBool(cfg.showQuit)
      cfg.showGift = toBool(cfg.showGift)
      cfg.showGiftName = toBool(cfg.showGiftName)
      cfg.showGiftPrice = toBool(cfg.showGiftPrice)
      cfg.mergeSimilarDanmaku = toBool(cfg.mergeSimilarDanmaku)
      cfg.mergeSimilarOther = toBool(cfg.mergeSimilarOther)
      cfg.mergeGift = toBool(cfg.mergeGift)
      cfg.maxNumber = toInt(cfg.maxNumber, config.DEFAULT_CONFIG.maxNumber)
      cfg.blockGiftDanmaku = toBool(cfg.blockGiftDanmaku)
      cfg.blockLevel = toInt(cfg.blockLevel, config.DEFAULT_CONFIG.blockLevel)
      cfg.blockNewbie = toBool(cfg.blockNewbie)
      cfg.blockNotMobileVerified = toBool(cfg.blockNotMobileVerified)
      cfg.blockMedalLevel = toInt(cfg.blockMedalLevel, config.DEFAULT_CONFIG.blockMedalLevel)
      cfg.autoTranslate = toBool(cfg.autoTranslate)

      this.config = cfg
    },
    wsConnect() {
      const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
      // 开发时使用localhost:12450
      const host = process.env.NODE_ENV === 'development' ? 'localhost:12451' : window.location.host
      const url = `${protocol}://${host}/chat`
      this.websocket = new WebSocket(url)
      this.websocket.onopen = this.onWsOpen
      this.websocket.onclose = this.onWsClose
      this.websocket.onmessage = this.onWsMessage
      this.heartbeatTimerId = window.setInterval(this.sendHeartbeat, 10 * 1000)
    },
    sendHeartbeat() {
      this.websocket.send(JSON.stringify({
        cmd: COMMAND_HEARTBEAT
      }))
    },
    onWsOpen() {
      this.retryCount = 0
      this.websocket.send(JSON.stringify({
        cmd: COMMAND_JOIN_ROOM,
        data: {
          roomId: parseInt(this.$route.params.roomId),
          config: {
            autoTranslate: this.config.autoTranslate
          }
        }
      }))
    },
    onWsClose() {
      if (this.heartbeatTimerId) {
        window.clearInterval(this.heartbeatTimerId)
        this.heartbeatTimerId = null
      }
      if (this.isDestroying) {
        return
      }
      window.console.log(`掉线重连中${++this.retryCount}`)
      this.wsConnect()
    },
    onWsMessage(event) {
      let {cmd, data} = JSON.parse(event.data)
      let message = null
      switch (cmd) {
      case COMMAND_JOIN_ROOM:
        if (!this.config.showJoin || this.mergeSimilarOther(data.authorName, data.content)) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_JOIN,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          authorType: data.authorType,
          content: data.content,
          privilegeType: data.privilegeType,
          repeated: 1,
          translation: data.translation
        }
        break
      case COMMAND_QUIT_ROOM:
        if (!this.config.showQuit || this.mergeSimilarOther(data.authorName, data.content)) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_QUIT,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          authorType: data.authorType,
          content: data.content,
          privilegeType: data.privilegeType,
          repeated: 1,
          translation: data.translation
        }
        break
      case COMMAND_ADD_TEXT:
        if (!this.config.showDanmaku || !this.filterTextMessage(data) || this.mergeSimilarText(data.content)) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_TEXT,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          authorType: data.authorType,
          content: data.content,
          privilegeType: data.privilegeType,
          repeated: 1,
          translation: data.translation
        }
        break
      case COMMAND_ADD_GIFT: {
        if (!this.config.showGift) {
          break
        }
        let price = data.totalCoin / 1000
        if (this.mergeSimilarGift(data.authorName, price, data.giftName, data.num)) {
          break
        }
        if (price < this.config.minGiftPrice) { // 丢人
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_GIFT,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          price: price,
          giftName: data.giftName,
          num: data.num
        }
        break
      }
      case COMMAND_ADD_LOVE: {
        if (!this.config.showLove || this.mergeSimilarOther(data.authorName, data.content)) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_LOVE,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          authorType: data.authorType,
          content: data.content,
          price: 0,
          repeated: 1,
          giftName: data.giftName,
          num: data.num
        }
        break
      }
      case COMMAND_ADD_FOLLOW: {
        if (!this.config.showFollow) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_FOLLOW,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          content: data.content,
          price: 0,
          giftName: data.giftName,
          num: data.num
        }
        break
      }
      case COMMAND_ADD_MEMBER:
        if (!this.config.showGift || !this.filterNewMemberMessage(data)) {
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_MEMBER,
          avatarUrl: data.avatarUrl,
          time: new Date(data.timestamp * 1000),
          authorName: data.authorName,
          title: 'NEW MEMBER!',
          content: `Welcome ${data.authorName}!`
        }
        break
      case COMMAND_ADD_SUPER_CHAT:
        if (!this.config.showGift || !this.filterSuperChatMessage(data)) {
          break
        }
        if (data.price < this.config.minGiftPrice) { // 丢人
          break
        }
        message = {
          id: data.id,
          type: constants.MESSAGE_TYPE_SUPER_CHAT,
          avatarUrl: data.avatarUrl,
          authorName: data.authorName,
          price: data.price,
          time: new Date(data.timestamp * 1000),
          content: data.content.trim()
        }
        break
      case COMMAND_DEL_SUPER_CHAT:
        for (let id of data.ids) {
          this.$refs.renderer.delMessage(id)
        }
        break
      case COMMAND_UPDATE_TRANSLATION:
        if (!this.config.autoTranslate) {
          break
        }
        data = {
          id: data[0],
          translation: data[1]
        }
        this.$refs.renderer.updateMessage(data.id, {translation: data.translation})
        break
      }
      if (message) {
        this.$refs.renderer.addMessage(message)
      }
    
    },
    filterTextMessage(data) {
      if (this.config.blockGiftDanmaku && data.isGiftDanmaku) {
        return false
      } else if (this.config.blockLevel > 0 && data.authorLevel < this.config.blockLevel) {
        return false
      } else if (this.config.blockNewbie && data.isNewbie) {
        return false
      } else if (this.config.blockNotMobileVerified && !data.isMobileVerified) {
        return false
      } else if (this.config.blockMedalLevel > 0 && data.medalLevel < this.config.blockMedalLevel) {
        return false
      }
      return this.filterSuperChatMessage(data)
    },
    filterSuperChatMessage(data) {
      for (let keyword of this.blockKeywords) {
        if (data.content.indexOf(keyword) !== -1) {
          return false
        }
      }
      return this.filterNewMemberMessage(data)
    },
    filterNewMemberMessage(data) {
      for (let user of this.blockUsers) {
        if (data.authorName === user) {
          return false
        }
      }
      return true
    },
    mergeSimilarText(content) {
      if (!this.config.mergeSimilarDanmaku) {
        return false
      }
      return this.$refs.renderer.mergeSimilarText(content)
    },
    mergeSimilarOther(authorName, content) {
      if (!this.config.mergeSimilarOther) {
        return false
      }
      return this.$refs.renderer.mergeSimilarOther(authorName, content)
    },
    mergeSimilarGift(authorName, price, giftName, num) {
      if (!this.config.mergeGift) {
        return false
      }
      return this.$refs.renderer.mergeSimilarGift(authorName, price, giftName, num)
    }
  }
}
</script>
