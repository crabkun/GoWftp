{{define "admin_user"}}
<script src="/static/js/admin_user.js"></script>
<div class="panel panel-default">
  <div class="panel-heading">用户管理</div>
  <table class="table">
    <thead>
    <tr>
      <th>ID</th>
      <th>用户名</th>
      <th>状态</th>
      <th>已用空间</th>
      <th>权限</th>
      <th colspan="2">操作</th>
    </tr>
  </thead>
  <tbody>
    {{.UserList}}
  </tbody>
  </table>
</div>
<div class="modal fade" id="delModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title" id="myModalLabel">警告</h4>
            </div>
            <div class="modal-body delmodal-body">确定要删除此用户？</div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-danger" onclick="deleteuser()">确认</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>
<div class="modal fade" id="editModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
    <div class="modal-dialog">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
                <h4 class="modal-title" id="editModalLabel">修改</h4>
            </div>
            <div class="modal-body editmodal-body">
              
            <form class="form-horizontal" id="editform" style="margin-left:20px;" action="/admin/user/edit"  method="post">
                <fieldset>
                <div class="control-group">
                    <label class="control-label">角色</label>
                    <div class="controls">
                <!-- Multiple Radios -->
                <label class="radio">
                    <input type="radio" value="0" name="usertype" checked="checked">
                    <input type="hidden" value="" name="id" id="sel">
                    未激活
                </label>
                <label class="radio">
                    <input type="radio" value="1" name="usertype">
                    普通用户
                </label>
                <label class="radio">
                    <input type="radio" value="2" name="usertype">
                    管理员
                </label>
            </div>

                    </div>
                 <br/>       
                <div class="control-group">

                    <!-- Text input-->
                    <label class="control-label" for="input01">用户目录</label>
                    <div class="controls">
                        <input type="text" placeholder="" class="input-xlarge" name="path">
                        <p class="help-block">（请输入文件夹名字而不是整个路径）</p>
                    </div>
                    </div>

                <div class="control-group">

                    <!-- Text input-->
                    <label class="control-label" for="input01">密码</label>
                    <div class="controls">
                        <input type="text" placeholder="" class="input-xlarge" name="password">
                        <p class="help-block">留空为不修改</p>
                    </div>
                </div>
                <div class="control-group">
                        <label class="control-label">权限</label>
                        <div class="controls">
                    <!-- Multiple Checkboxes -->
                    <label class="checkbox">
                        <input type="checkbox" name="read" value="1">
                        读
                    </label>
                    <label class="checkbox">
                        <input type="checkbox" name="write" value="1">
                        写
                    </label>
                    <label class="checkbox">
                        <input type="checkbox" name="delete" value="1">
                        删
                    </label>
                    </div>
                </div>
                </fieldset>
            </form>


            
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
                <button type="button" class="btn btn-danger" onclick="edituser()">确认</button>
            </div>
        </div><!-- /.modal-content -->
    </div><!-- /.modal -->
</div>
{{end}}