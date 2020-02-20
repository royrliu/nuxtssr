<template>
  <div>
     <link href="https://fonts.googleapis.com/css?family=Noto+Sans+JP:400,700%7CNoto+Sans+KR:400,700%7CNoto+Sans+SC:400,700%7CNoto+Sans+TC:400,700%7CNoto+Sans:400,700&display=swap&subset=chinese-simplified,chinese-traditional,japanese,korean" rel="stylesheet">
    <header class="header">
      <a-row style="margin: 0 auto;">
        <a-col :span="6"> Salesforce Forum  </a-col>
        <a-col :span="1" :offset="10" class="headerIcon">
          <a-icon type="search" />
        </a-col>
        <a-col :span="1" class="headerIcon">
        <a-icon type="bell" />
         </a-col>
        <a-col :span="1" class="headerIcon">
        <a-icon type="user" />
         </a-col>
          <a-col :span="3">
        <a-button type="primary">投稿</a-button>
         </a-col>
      </a-row>
      </header>
    <nuxt/>
     <footer class="footer"><span style="margin-right: 20px;">Powered by</span> <img src="~assets/img/titleLogo.png" style="width: 122px;height: 23px;"></footer>
  </div>
</template>
<script>

export default {
  components: {
  }
}
</script>

<style>
.headerIcon{
  margin-right: 10px;
}
.footer{
   font-weight: Regular;
  font-size: 12px;
  font-family: "Noto Sans CJK JP";
  color:#777777
}
.container
{
  margin: 0;
  width: 100%;
  padding: 100px 0;
  text-align: center;
}
.header {
  background: white;
  padding: 0 20%;
  height: 64px;
  line-height: 64px;
  font-weight: bold;
  font-size: 20px;
  font-family: "Noto Sans CJK JP";
  color: #333333;
}
.button, .button:visited
{
  display: inline-block;
  color: black;
  letter-spacing: 1px;
  background-color: #fff;
  border: 2px solid #000;
  text-decoration: none;
  text-transform: uppercase;
  padding: 15px 45px;
}

.button:hover, .button:focus
{
  color: #fff;
  background-color: #000;
}

.title
{
  color: #000;
  font-weight: 300;
  font-size: 2.5em;
  margin: 0;
}

</style>
