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
        <form action="/api/saveFileByForm" name='form1' method="post" enctype="multipart/form-data" target="stop_route">
          <input type="file" name="saveFileByForm" />
          <button type="submit">上传</button>
        </form>
        <iframe style="display:none;" id="stop_route" name="stop_route"></iframe>
      </li>
    </ul>
  </header>
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/axios.min.js"></script>
  <script src="/static/js/vue.min.js"></script>
  <script>
    function showRes (msg) {
      alert(msg)
    }
  </script>
  <script>
    new Vue({
      el: '#app',
      data: {
      },
      created: function () {
      },
      methods: {
      }
    })
  </script>
</body>
</html>
