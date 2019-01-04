<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
  <header id='app'>
    <ul>
      <li>
        <p>表单上传文件</p>
        <form action="/api/saveFileByForm" name='form1' method="post" enctype="multipart/form-data" onsubmit="form1.submit();return false;">
          <input type="file" name="saveFileByForm" />
          <button type="submit">上传</button>
        </form>
      </li>
    </ul>
  </header>
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/axios.min.js"></script>
  <script src="/static/js/vue.min.js"></script>
  <script>
    new Vue({
      el: '#app',
      data: {
        a: 1
      },
      created: function () {
        // `this` 指向 vm 实例
        console.log('a is: ' + this.a)
      }
    })
  </script>
</body>
</html>
