<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <title>登陆</title>
        <link rel="stylesheet" href="/static/bootstrap/css/bootstrap.min.css">
		<link rel="stylesheet" href="/static/css/form-elements.css">
        <link rel="stylesheet" href="/static/css/style.css">
		<script src="/static/js/jquery-1.11.1.min.js"></script>
        <script src="/static/bootstrap/js/bootstrap.min.js"></script>
        <script src="/static/js/jquery.backstretch.min.js"></script>
        <script src="/static/js/scripts.js"></script>
    </head>

    <body>

        <!-- Top content -->
        <div class="top-content">
        	
            <div class="inner-bg">
                <div class="container">
                    <div class="row">
                        <div class="col-sm-8 col-sm-offset-2 text">
                            <h1>{{title}}</h1>
                        </div>
                    </div>
                    <div class="row">
                        <div class="col-sm-6 col-sm-offset-3 form-box">
                        	<div class="form-top">
                        		<div class="form-top-left">
                                    <img src="static/img/{{.icon}}.png">
                        		</div>
                        		<div class="form-top-right">
                                    <h3>{{.title}}</h3>
                            		<p>没有帐号？<a href='/register' style="color:white;text-decoration:underline">点我注册</a></p>
                        		</div>
                            </div>
                            <div class="form-bottom">
			                    <form role="form" action="login" method="post" class="login-form">
			                    	<div class="form-group">
			                    		<label class="sr-only" for="form-username">用户名</label>
			                        	<input type="text" name="username" placeholder="用户名..." class="form-username form-control" id="form-username">
			                        </div>
			                        <div class="form-group">
			                        	<label class="sr-only" for="form-password">密码</label>
			                        	<input type="password" name="password" placeholder="密码..." class="form-password form-control" id="form-password">
			                        </div>
			                        <button type="submit" class="btn">登陆</button>
			                    </form>
		                    </div>
                        </div>
                    </div>
					<div class="row">
					<p style="margin-top:30px;">version：{{ver}}</p>
                    <p>Powered by GO and Beego © 2017 螃蟹</p>
					</div>
                </div>
            </div>
            
        </div>
        <!--[if lt IE 10]>
            <script src="static/js/placeholder.js"></script>
        <![endif]-->

    </body>

</html>