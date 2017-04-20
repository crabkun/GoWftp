var nowsel=-1
function deleteuserconfirm(a)
{
$('.delmodal-body').html("确定要删除此用户？")
userid=$(a).parent().siblings()[0].innerHTML
username=$(a).parent().siblings()[1].innerHTML
size=$(a).parent().siblings()[3].innerHTML
nowsel=userid
$('.delmodal-body').append("<br/>用户名："+username+"<br/>已用空间："+size)
$.ajax( {    
    url:'/admin/user/checkpath',   
    data:{id:userid},    
    type:'post',    
    cache:false,      
    success:function(result) {    
      $('.delmodal-body').append("<br/><br/>"+result)
     },    
     error : function() {
        alert("服务器连接失败，无法列出此用户的目录关联信息．")  
     }    
    });  
}
function deleteuser()
{
$.post("/admin/user/delete",{id:nowsel},function(result){
    alert(result)
    $("#delModal").modal("hide")
    location.reload()
  });
}
function edituserconfirm(a)
{
  userid=$(a).parent().siblings()[0].innerHTML
  username=$(a).parent().siblings()[1].innerHTML
  size=$(a).parent().siblings()[3].innerHTML
  $("#sel").val(userid)
   $('#editModalLabel').html("编辑"+username+"(id:"+userid+")")
   $.ajax( {    
    url:'/admin/user/getinfo',   
    data:{id:userid},    
    type:'post',    
    cache:false, 
    async : false,     
    success:function(result) {    
      var tmp=result.split("|")
      switch(tmp[0]) {
        case "0":$("input[name='usertype']").eq(0).click();break;
        case "1":$("input[name='usertype']").eq(1).click();break;
        case "2":$("input[name='usertype']").eq(2).click();break;
      }
      if(tmp[2]=="1")
        $("input[name='read']").attr("checked","checked");
      else
        $("input[name='read']").removeAttr("checked");
      if(tmp[3]=="1")
        $("input[name='write']").attr("checked","checked");
      else
        $("input[name='write']").removeAttr("checked");
      if(tmp[4]=="1")
        $("input[name='delete']").attr("checked","checked");
      else
        $("input[name='delete']").removeAttr("checked");
      $("input[name='path']").val(tmp[1])
     },    
     error : function() {
        alert("服务器连接失败，无法列出此用户的目录关联信息．")  
        location.reload()
     }    
    });  
}
function edituser()
{
$('#editform').submit()
}