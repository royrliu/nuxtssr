{{{{raw}}}}
<template>
  <a-tabs defaultActiveKey="1" @change="callback" class="header-tabs">
    <a-tab-pane tab="記事" key="1">
      <div style="padding-bottom: 15px;border-bottom: 1px solid #ddd;">
        Featured Tags:
         <a-tag v-for="(user, index) in users" :key="index" class="user">
        <nuxt-link :to="{ name: 'id', params: { id: user.Id }}">
          {{ user.Username }}
        </nuxt-link>
      </a-tag>
      </div>
      <a-layout-content>
        <a-row class="title">記事</a-row>
        <a-list style="margin-bottom: 20px;" itemLayout="vertical" size="large" :pagination="pagination" :dataSource="listData">
          <a-list-item slot="renderItem" slot-scope="item" key="item.title">
            <template slot="actions" v-for="{type, text} in actions">
              <span :key="type">
                <a-icon :type="type" style="margin-right: 8px"/>
                {{text}}
              </span>
            </template>
            <img slot="extra" width="272" alt="logo" src="https://gw.alipayobjects.com/zos/rmsportal/mqaQswcyDLcXyDKnZfES.png">
            <a-list-item-meta :description="item.description">
              <a slot="title" :href="item.href">{{item.title}}</a>
              <a-avatar slot="avatar" :src="item.avatar"/>
            </a-list-item-meta>
            {{item.content}}
          </a-list-item>
        </a-list>
      </a-layout-content>
    </a-tab-pane>
    <a-tab-pane tab="イベント" key="2" forceRender>イベント</a-tab-pane>
    <a-tab-pane tab="プレミアムコンテンツ" key="3">プレミアムコンテンツ</a-tab-pane>
  </a-tabs>
</template>
{{{{/raw}}}}
<script>
const listData = []
for (let i = 0; i < 23; i++) {
  listData.push({
    href: 'https://www.antdv.com/',
    title: `ant design vue part ${i}`,
    avatar: 'https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png',
    description:
      'Ant Design, a design language for background applications, is refined by Ant UED Team.',
    content:
      'We supply a series of design principles, practical patterns and high quality design resources (Sketch and Axure), to help people create their product prototypes beautifully and efficiently.',
  })
}
import axios from '~/plugins/axios'
export default {
  async asyncData () {
    let { data } = await axios.get('/api/users')
    console.log('data1111===>', data)
    return { users: data }
  },
  layout: 'basicLayout',
  head() {
    return {
      title: 'Users',
    }
  },
  data() {
    return {
      listData,
      pagination: {
        onChange: page => {
          console.log(page)
        },
        pageSize: 3,
      },
      actions: [
        { type: 'star-o', text: '156' },
        { type: 'like-o', text: '156' },
        { type: 'message', text: '2' },
      ],
    }
  },
  methods: {
    callback(key) {
      console.log(key)
    },
  },
}
</script>
<style scoped>
.title {
  margin-top: 20px;
  font-weight: Bold;
  font-size: 24px;
  font-family: 'Noto Sans CJK JP';
  color: #333333;
}
.header-tabs {
  padding: 0 20%;
}
</style>
