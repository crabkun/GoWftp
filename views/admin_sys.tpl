{{define "admin_sys"}}
<script src="/static/js/admin_sys.js"></script>
<div class="panel panel-default">
  <div class="panel-heading">系统设置</div>
  <table class="table">
    <thead>
    <tr>
      <th>配置项</th>
      <th>操作</th>
    </tr>
    
  </thead>
  <tbody>
  <tr>
      <td>允许注册</td>
      <td><input type="radio" id="reg_on" name="reg" value="1">开　　　　<input type="radio" id="reg_off" name="reg" value="0">关</td>
    </tr>

  </tbody>
  </table>
  <button type="button" class="btn btn-default" id="save">保存</button>
</div>
{{end}}