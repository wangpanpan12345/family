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
			.family{
				display: flex;
				padding: 0 10%;
				flex-direction: column;
				/*justify-content: center;*/
				/*align-items: center;*/
				height: 100vh;
			}
			input[type="text"] {
			    -webkit-appearance: none;
			    background-color: #fff;
			    background-image: none;
			    border-radius: 2px;
			    /*border: 1px solid #dcdfe6;*/
			    box-sizing: border-box;
			    color: #606266;
			    display: inline-block;
			    font-size: inherit;
			    /*height: 40px;*/
			    line-height: 44px;
			    outline: none;
			    padding: 0 15px;
			    transition: border-color .2s cubic-bezier(.645,.045,.355,1);
			    width: 60%;
			}
			::-webkit-input-placeholder{
		        color: #d2d2d2;
		        font-weight: 300;
			}
			::-moz-placeholder{  
		        color: #d2d2d2; 
		        font-weight: 300;       
			}
			:-ms-input-placeholder{  
		        color: #d2d2d2; 
		        font-weight: 300;       
			}
			.search {
				width: 100%;
				margin: 20px 0;
				margin-top: 50px;
			}
			.search-input{
				display: block;
				width: 100%;
				text-align: center;
			}
			.RNNXgb {
			    background: #fff;
			    display: flex;
			    /*border-radius: 2px;*/
			    border: none;
			    box-shadow: 0 2px 2px 0 rgba(0,0,0,0.16), 0 0 0 1px rgba(0,0,0,0.08);
			    z-index: 3;
			    height: 44px;
			    margin: 0 auto;
			    /*width: 484px;*/
			}
			.search-list{
				display: flex;
				/*align-items: center;*/
				justify-content: center;
				margin:4px 0;
				padding:0 20%;
			}
			.el-autocomplete-suggestion{
				background: #fff;
				width: 100%

			}
			.el-scrollbar{

			}
			.el-autocomplete-suggestion li {
			    padding: 0 20px;
			    margin: 0;
			    line-height: 34px;
			    cursor: pointer;
			    color: #606266;
			    font-size: 14px;
			    list-style: none;
			    white-space: nowrap;
			    overflow: hidden;
			    text-overflow: ellipsis;
			    display: flex;
			    align-items: center;
			}

			.avatar {
				width: 30px;
				height: 30px;
				border-radius: 50%;
			}

			.el-autocomplete-suggestion li span{
				margin: 5px 3px;
			}
			.panel{
				display: flex;
				flex-direction: row;

			}
			.member-info{
				width: 40%;
				height: auto;
				background: #ffff;
				border-radius: 5px;
				padding: 20px;
			}
			.member-info span{
				display: block;
			}
						
		</style>
	</head>
	<body>

		<div class="family">
			<div class="search">
				<span class="search-input">
					<input type="text" class="RNNXgb"  placeholder="请输入姓名" name="">
				</span>
				<span class="search-list">
					<div class="el-autocomplete-suggestion">
						<ul class="el-scrollbar">
							<!-- <li class="">
								<span><img class="avatar" src="/static/img/test.png" /></span>
								<span class="addr">长宁区新渔路144号</span>
							</li> -->
							<!-- <li class="">
								<span><img class="avatar" src="/static/img/test.png" /></span>
								<span class="addr">上海市长宁区淞虹路661号</span>
							</li> -->
						</ul>
					</div>
				</span>
			</div>

			<div class="panel">
				<div class="member-info">
					<span class="" id="name">姓名:</span>
					<span class="" id="sex">性别:</span>
					<!-- <span class="" id="bf">辈分:23世</span> -->
					<span class="" id="info">简介:</span>
				</div>
				<div class="member-compute-tobe">
					<span></span>
					<span></span>
				</div>
				
			</div>

		</div>

		<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
		<script type="text/javascript">
			
			function _debounce(func, wait, immediate) {
                var timeout, args, context, timestamp, result;

                var later = function () {
                    // 据上一次触发时间间隔
                    var last = Date.now() - timestamp;

                    // 上次被包装函数被调用时间间隔last小于设定时间间隔wait
                    if (last < wait && last > 0) {
                        timeout = setTimeout(later, wait - last);
                    } else {
                        timeout = null;
                        // 如果设定为immediate===true，因为开始边界已经调用过了此处无需调用
                        if (!immediate) {
                            result = func.apply(context, args);
                            if (!timeout) context = args = null;
                        }
                    }
                };

                return function () {
                    context = this;
                    args = arguments;
                    timestamp = Date.now();
                    var callNow = immediate && !timeout;
                    // 如果延时不存在，重新设定延时
                    if (!timeout) timeout = setTimeout(later, wait);
                    if (callNow) {
                        result = func.apply(context, args);
                        context = args = null;
                    }

                    return result;
                };
            };
      

      		var set_interval = _debounce(Search, 1000);
	            $('.RNNXgb').on('input', function () {
	                
	                var name = $(this).val();
	               	if (name ==""){
	               		return false;
	               	}
	                set_interval("/member/q?name="+name, $(this));

            });	

	        $(".search-list").on("click",".member-item",function(){
	        	$('.RNNXgb').val($(this).children(".info").text());
	        	$(".el-scrollbar").children().remove();
	        });
			function Search(url){
	        	 $.ajax({
	                method: "Get",
	                url: url,
	                // contentType: "application/json;charset=utf-8",
	                // dataType: "json",
	                // data: JSON.stringify(param),
	                success: function (data) {
	                    if (data.code == "0") {
	                       	console.log(data)
	                       	var html = ""
	                       	
	                       	$.each(data.data,function(n,item){
	                       		html = html + `<li class="member-item" data-id= `+item.id+`>
									<span><img class="avatar" src="/static/img/test.png" /></span>
									<span class="info">`+item.name+`</span>
								</li>`
	                       	})
	                       	
	                       	$(".el-scrollbar").html(html)
	                      	
	                    } else {
	                        console.log(data);
	                    }
	                },
	                error: function (data) {
	                    alert("服务器出错！请联系管理员！");
	                }
	            });
	        }
		</script>
		
	</body>
</html>