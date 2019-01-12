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
        <p>表单上传文件（设置form的action，method，enctype属性）</p>
        <p>优点: 所有浏览器均支持该方式</p>
        <p>缺点：上传后，页面会发生跳转，如果需求本身在上传文件成功或者失败后会跳转到新页面可以采用</p>
        <form action="/api/saveFileByForm" method="post" enctype="multipart/form-data">
          <input type="file" name="saveFileByForm" multiple="multiple" />
          <button type="submit">上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，iframe捕获返回提示（设置form的action，method，enctype属性）</p>
        <p>优点: 所有浏览器均支持该方式</p>
        <p>缺点：需要后台配合，返回js脚本，通过iframe去处理结果，调用第三方接口可能没法配合</p>
        <form action="/api/saveFileByFormNoFresh" method="post" enctype="multipart/form-data" target="stop_route">
          <input type="file" id='saveFileByFormNoFresh' name="saveFileByForm" multiple="multiple" />
          <button type="submit">上传</button>
        </form>
        <iframe id="stop_route" name="stop_route" style="display:none;"></iframe>
      </li>
      <li>
        <p>表单上传文件，jqueryForm(ajax+iframe)</p>
        <p>优点: 所有浏览器均支持该方式</p>
        <p>问题点：jqueryForm采用了blob技术结合form的降级方案，后台返回值需要做兼容处理，IE9下返回的json信息会变成json文件下载</p>
        <form id='ajaxForm'>
          <input type="file" id='saveFileByAjaxForm' name="saveFileByForm" multiple="multiple" />
          <button type="button" @click='submitAjaxFormHandle'>上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，axios，FormData</p>
        <p>优点: 有elementui封装好的插件可用</p>
        <p>问题点：不支持FormData对象的浏览器不可用</p>
        <form name='formAxios'>
          <input type="file" id='saveFileByAxios' name="saveFileByForm" multiple="multiple" />
          <button type="button" @click='submitAxiosHandle'>上传</button>
        </form>
      </li>
      <li>
        <p>表单上传文件，axios，FileReader(arrayBuffer)</p>
        <p>优点: 不清楚</p>
        <p>问题点：文件大小在MB级别就有点卡顿了</p>
        <form name='formAxios'>
          <input type="file" id='saveFileByFileReader' name="saveFileByForm" />
          <button type="button" @click='submitFileReaderHandle'>上传</button>
        </form>
      </li>
    </ul>
    <br />
    <ul>
      <li>
        <p>a标签配合download属性（这里我是将文件写在程序内部的静态资源目录里了，当然实际项目中不能这么做）</p>
        <p>下载自己上传的图片或者txt文档（IE下能直接打开的文件都会预览，chrome和firefox能下载）</p>
        <p>下载自己上传的office文件、pdf等（都能下载）</p>
        <p v-for='(f, i) in fileList' :key='i'><a :href="'../static/files/' + f" download>{{f}}</a> </p>
      </li>
      <li>
        <p>a标签配合download属性，下载第三方的图片, 表现形式各异</p>
        <a href="https://common.cnblogs.com/images/wechat.png" download>第三方图片</a>
      </li>
      <li>
        <p>form表单实现下载（表现稳定）</p>
        <form v-for='(f, i) in fileList' :key='i' action="../api/getFile" method="GET">
          <input type="text" name='FileName' style="border:none;" readonly :value='f' />
          <button type="submit">下载</button>
        </form>
      </li>
      <li>
        <p>从接口读取流，下载, 不支持IE9及以下</p>
        <p v-for='(f, i) in fileList' :key='i'>{{f}}<button type="button" @click="downfile(f)">下载</button></p>
      </li>
      <li>
        <p>利用iframe下载（表现稳定）</p>
        <iframe name="downloadIframe" style="display:none;"></iframe>
        <p v-for='(f, i) in fileList' :key='i'>{{f}}<button type="button" @click="downfileByIframe(f)">下载</button></p>
      </li>
    </ul>
  </header>
  <script src="/static/js/polyfill.js"></script>
  <script src="/static/js/reload.min.js"></script>
  <script src="/static/js/axios.min.js"></script>
  <script src="/static/js/vue.min.js"></script>
  <script src="/static/js/element.js"></script>
  <script src="/static/js/jquery2.10.js"></script>
  <script src="/static/js/jqueryForm.js"></script>
  <script>
    var showMessageFunc = null;
    function clearValueById(id) {
      var file = document.getElementById(id)
      file.outerHTML = file.outerHTML;
    }
    function showRes (msg, type) {
      clearValueById('saveFileByFormNoFresh');
      showMessageFunc(msg, type);
    }
    var app = new Vue({
      el: '#app',
      data: {
        fileList: []
      },
      created: function () {
        showMessageFunc = this.showMessage
      },
      mounted: function() {
        this.getFilesList()
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
            success: function(res) {
              if (res.Code === 200) {
                self.showMessage(res.Msg, 'success')
                self.getFilesList()
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
                self.showMessage(response.data.Msg, 'success')
                self.getFilesList()
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
            if (filesList.length) {
              for (var i = 0; i < filesList.length; i++) {
                this.FileUpload(filesList[i]);
              }
            } else {
              this.showMessage('GetFile failed', 'error')
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
                  self.showMessage(res.Msg, 'success')
                  clearValueById('saveFileByFileReader')
                  self.getFilesList()
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
        },
        getFilesList: function(){
          var self = this
          axios.get('/api/getFilesList?t='+(new Date()).getTime())
          .then(function (response) {
            if (response.data.Code === 200) {
              self.fileList = response.data.Data.FilesList
            }else{
              self.showMessage('更新文件列表失败', 'error')
            }
          })
          .catch(function (error) {
            self.showMessage('网络或者服务异常', 'error')
          });
        },
        downfile: function(FileName) {
          var self = this
          axios.get('/api/getFile?FileName=' + encodeURI(FileName),  {responseType: 'blob'})
          .then(function (response) {
            try {
              var blob = response.data
              if (window.navigator.msSaveOrOpenBlob) {
                navigator.msSaveBlob(blob, FileName);
              } else {
                blob.type = "application/octet-stream";
                var url = URL.createObjectURL(blob);
                var a = document.createElement('a');
                a.href = url;
                a.download = FileName;
                var evt = document.createEvent("MouseEvents");  
                evt.initEvent("click",true,true);  
                a.dispatchEvent(evt);
                window.URL.revokeObjectURL(url);
              }
            }catch(err){
              self.showMessage('不支持', 'error')
            }
          })
          .catch(function (error) {
            self.showMessage('网络或者服务异常', 'error')
          });
        },
        downfileByIframe: function(FileName) {
          var elemIF = window.frames['downloadIframe'].frameElement;
          elemIF.src = '/api/getFile?FileName=' + encodeURI(FileName);
        }
      }
    })
  </script>
</body>
</html>
