<!DOCTYPE html>

<html>
<head>
  <title>文件上传与下载总结</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link rel="stylesheet" type="text/css" href="../static/css/element.css" />
</head>

<body>
  <el-button :plain="true" @click="openCenter"></el-button>
  <header id='app'>
    <ul>
      <li>
        <p>表单上传文件</p>
        <form action="/api/saveFileByForm" method="post" enctype="multipart/form-data">
          <input type="file" name="saveFileByForm" multiple="multiple" />
          <button type="submit">上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，iframe提示(模拟类似ajax等异步上传，IE9及以下)</p>
        <form action="/api/saveFileByFormNoFresh" method="post" enctype="multipart/form-data" target="stop_route">
          <input type="file" id='saveFileByFormNoFresh' name="saveFileByForm" multiple="multiple" />
          <button type="submit">上传</button>
        </form>
        <iframe id="stop_route" name="stop_route" style="display:none;"></iframe>
      </li>
      <li>
        <p>表单上传文件，jqueryForm(ajax+iframe)</p>
        <form id='ajaxForm' method="post" enctype="multipart/form-data">
          <input type="file" id='saveFileByAjaxForm' name="saveFileByForm" multiple="multiple" />
          <button type="button" @click='submitAjaxFormHandle'>上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，axios，FormData(elementUI的上传也是基于这个对象完成的，不支持IE9)</p>
        <form name='formAxios'>
          <input type="file" id='saveFileByAxios' name="saveFileByForm" multiple="multiple" />
          <button type="button" @click='submitAxiosHandle'>上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，axios，FileReader(arrayBuffer, 文件MB级别就有点卡顿了)</p>
        <form name='formAxios'>
          <input type="file" id='saveFileByFileReader' name="saveFileByForm" multiple="multiple" />
          <button type="button" @click='submitFileReaderHandle'>上传</button>
        </form>
      </li>
    </ul>
  </header>
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/axios.min.js"></script>
  <script src="/static/js/vue.min.js"></script>
  <script src="/static/js/element.js"></script>
  <script src="/static/js/jquery2.10.js"></script>
  <script src="/static/js/jqueryForm.js"></script>
  <script>
    var showMessageFunc = null;
    function showRes (msg, type) {
      clearValueById('saveFileByFormNoFresh');
      showMessageFunc(msg, type);
    }
    function clearValueById(id) {
      document.getElementById(id).value = ''
    }
    var app = new Vue({
      el: '#app',
      data: {
      },
      created: function () {
        showMessageFunc = this.showMessage
      },
      methods: {
        showMessage: function(msg, type) {
          this.$message({
            message: msg,
            center: true,
            type: type
          })
        },
        submitAjaxFormHandle: function () {
          var self = this
          var option = {
            url: '/api/saveFileByAjaxForm',
            type: 'POST',
            success: function(data) {
              var res
              if (data && data.match(/\{[^\}]+\}/)[0]) {
                res = JSON.parse(data.match(/\{[^\}]+\}/)[0])
              }
              if (res.Code === 200) {
                self.showMessage('上传成功', 'success')
              } else {
                self.showMessage(res.Msg, 'error')
              }
              $("#ajaxForm").resetForm()
            },
            error: function(data) {
              self.showMessage('网络或者服务异常', 'error')
            }
          };
          $("#ajaxForm").ajaxSubmit(option);
          return false;
        },
        submitAxiosHandle: function  () {
          var self = this;
          try {
            var formData = new FormData();
            var filseList = document.getElementById('saveFileByAxios').files;
            for (var i = 0; i < filseList.length; i++) {
              formData.append("saveFileByForm", filseList[i]);
            }
            axios.post('/api/saveFileByAxios', formData, {
              headers: {'content-type': 'multipart/form-data'}
            })
            .then(function (response) {
              if (response.data.Code === 200) {
                self.showMessage('上传成功', 'success')
              } else {
                self.showMessage(response.data.Msg, 'error')
              }
              clearValueById('saveFileByAxios')
            })
            .catch(function (error) {
              self.showMessage('网络或者服务异常', 'error')
            });
          }catch(err){
            this.showMessage('不支持DataForm对象', 'error')
            throw new Error('不支持DataForm对象')
          }
        },
        submitFileReaderHandle: function() {
          try{
            var filesList = document.getElementById("saveFileByFileReader").files;
          
            for (var i = 0; i < filesList.length; i++) {
              this.FileUpload(filesList[i]);
            }
          }
          catch(err){
            this.showMessage('IE9及以下不支持', 'error')
          }
        },
        FileUpload: function(file) {
          var reader = new FileReader();  
          var xhr = new XMLHttpRequest();
          var self = this;
          xhr.onreadystatechange=function(){
            if(xhr.readyState==4){
              if(xhr.status==200){
                var res = JSON.parse(xhr.response)
                if (res.Code === 200) {
                  self.showMessage('上传成功', 'success')
                } else {
                  self.showMessage(res.Msg, 'error')
                }
              }
            }
          }
          xhr.onerror=function(e){
            console.log("error!");
          }
          xhr.open("POST", "/api/saveFileByFileReader");
          xhr.setRequestHeader('content-type', 'application/json; charset=UTF-8');
          reader.onload = function(evt) {
            if (evt.target.result.length > Number.MAX_SAFE_INTEGER) {
              self.showMessage('文件太大', 'error')
              return false;
            }
            xhr.send(JSON.stringify({Name: file.name, Data: Array.prototype.slice.call(new Uint8Array(evt.target.result))}));
          };
          reader.readAsArrayBuffer(file);
        }
      }
    })
  </script>
</body>
</html>
