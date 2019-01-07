<html>

	<head>
		<meta charset="utf-8">
		<title>王氏</title>
		<!-- 引入样式 -->
		<!-- <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css"> -->
		<!-- 引入组件库 -->
		<!-- <script src="https://unpkg.com/element-ui/lib/index.js"></script> -->
		<style type="text/css">
			body,html{
				margin: 0;
				padding: 0;
				border: 0;
				background: #f7f7f7;
				font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;

			}
			ul,li{
				border: 0;margin: 0;padding: 0;
				list-style: none;
			}

			
			input[type="text"] {
			    -webkit-appearance: none;
			    background-color: #fff;
			    background-image: none;
			    border-radius: 4px;
			    border: 1px solid #dcdfe6;
			    box-sizing: border-box;
			    color: #606266;
			    display: inline-block;
			    font-size: inherit;
			    height: 40px;
			    line-height: 40px;
			    outline: none;
			    padding: 0 15px;
			    transition: border-color .2s cubic-bezier(.645,.045,.355,1);
			    /*width: 100%;*/
			}
			.family{
				display: flex;
				padding: 0 10%;
				flex-direction: column;
			}
			.node-item{
				padding: 10px;
				background: #409EFF;
				margin: 10px 0;
				cursor: pointer;
				color: #fff;
			}
			.opr{
				position:fixed;
				width: 400px;
				height: 300px;
				top: 50%;
				left: 50%;
				margin-left: -250px;
				margin-top: -150px;
				border-radius: 10px;
				padding: 20px;
				z-index: 10;
				background: #fff;
				display: flex;
			    flex-direction: column;
			    align-items: center;
			}
			.opr span{
				display: block;
				margin: 10px 0;
			}
			.cover{
				position: absolute;
				top: 0;
				left: 0;
				background: #000;
				opacity: .5;
				width: 100vw;
				height: 100vh; 
			}
			span.add-ok{
				/*text-decoration: none;*/
			    margin: 10px 0;
			    
			    border-color: #409eff;
			    display: inline-block;
			    line-height: 1;
			    white-space: nowrap;
			    cursor: pointer;
			    background: #fff;
			    border: 1px solid #dcdfe6;
			    border-color: #dcdfe6;
			    color: #606266;
			    -webkit-appearance: none;
			    text-align: center;
			    box-sizing: border-box;
			    outline: none;
			    margin: 0;
			    transition: .1s;
			    font-weight: 500;
			    -moz-user-select: none;
			    -webkit-user-select: none;
			    -ms-user-select: none;
			    padding: 12px 20px;
			    font-size: 14px;
			    border-radius: 4px;
			    background: #409eff;
			    color: #fff;
			}

		</style>
	</head>
	<body>
		
		<div class="family">
			<div class="menu"></div>
			<div>
				<span><a href="javascript:void(0);">添加</a></span>
			</div>
			<div class="list">
				<ul>
					<li class="node-item" data-id="123" >
						<span>王某</span>
						<span></span>
						<span></span>
					</li>
					<li class="node-item" data-id="123">
						<span>王某</span>
						<span></span>
						<span></span>
					</li>
				</ul>
			</div>
		</div>

		<div class="opr">
			<span>
				<label>姓名：</label><input type="text" name="">
				
			</span>
			<span>
				<label>性别：</label><input type="text" name="">
				
			</span>
			<span>
				<label>生日：</label><input type="text" name="">
				
			</span>
			<span>
				<label>教育：</label><input type="text" name="">
			</span>
			<span>
				<span class="add-ok">提交</span> 
			</span>
		</div>
		<div class="cover">
		</div>

	</body>
</html>