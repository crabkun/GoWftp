function save() {
    var o=new Array();
    o[0]={Id:0,Name:"canreg",Value:"false"}
    if($("input[name=reg]:checked").val()=="1")
        o[0].Value="true";
    $.ajax( {
        url:'/admin/sys/set',
        type:'post',
        cache:false,
        data:{d:JSON.stringify(o)},
        success:function(result) {
           window.location.reload();
        },
        error : function() {
            alert("服务器连接失败．")
        }
    });
}
$(function(){
    $.ajax( {
        url:'/admin/sys/get',
        type:'get',
        cache:false,
        dataType:'json',
        success:function(result) {
            if(result[0].Value=="true")
                $('#reg_on').click()
            else
                $('#reg_off').click()
        },
        error : function() {
            alert("服务器连接失败．")
        }
    });
    $('#save').click(function(){
        save();
    });
})