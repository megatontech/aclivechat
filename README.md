# aclivechat
用于避难所广播的直播互动栏

## 感谢：
* 前端来自： https://github.com/xfgryujk/blivechat
* 后端弹幕获取： https://github.com/orzogc/acfundanmu
* 原始项目：https://github.com/ShigemoriHakura/aclivechat
## 使用
* 运行aclivechat.exe
* 浏览器打开：[http://localhost:12451](http://localhost:12451)


### 从源代码编译
1. 编译前端（需要安装Node.js和npm【或者用cnpm更快】）：
   ```sh
   cd frontend
   npm i 【cnpm i】
   npm run build
   ```
   
2. 编译后端（需要安装go）
   ```sh
   go build
   ```
   
3. 正确放置文件
   ```sh
   前端 /dist
   后端 /
   ```

4. 浏览器打开[http://localhost:12451](http://localhost:12451)
