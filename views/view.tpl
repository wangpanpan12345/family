<html>

	<head>
		<meta charset="utf-8">
		<link rel="icon" href="/static/img/home.png" type="image/x-icon" />
		<link rel="shortcut icon" href="/static/img/home.png" type="image/x-icon"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=0.5, maximum-scale=2.0, user-scalable=yes" />

		<title>尚贤王氏族谱</title>
		<!-- 引入样式 -->
		<!-- <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css"> -->
		<!-- 引入组件库 -->
		<!-- <script src="https://unpkg.com/element-ui/lib/index.js"></script> -->
		<script src="https://cdn.bootcss.com/echarts/4.2.0-rc.2/echarts.min.js"></script>
		<style type="text/css">
			body,html{
				margin: 0;
				padding: 0;
				border: 0;
				background: #fff;
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
				/*margin: 20px 0;*/
				margin-top: 50px;
			}
			.search-input{
				position: relative;
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
				position: relative;
				display: flex;
				/*align-items: center;*/
				justify-content: center;
				margin:1px 0;
				padding:0 20%;
			}
			.el-autocomplete-suggestion{
				position: absolute;
				background: #fff;
				/*width: 100%*/
			    width: 60%;
			    z-index: 20;

			}
			.el-scrollbar{

				border: 1px solid #d7d7d7;

			}
			.el-autocomplete-suggestion li {
			    padding: 0 20px;
			    /*margin: 0 20px;*/
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

			.el-autocomplete-suggestion li:hover{
				background: #f7f7f7;
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
				/*display: flex;*/
				/*flex-direction: row;*/

			}
			.member-info{
				/*width: 40%;*/
				/*height: auto;*/
				/*height: 1000px;*/
				/*background: #ffff;*/
				border-radius: 5px;
				padding: 20px;
				margin:0 0%;
			}
			.member-info span{
				display: block;
			}
			.compute-button{
				position: absolute;
			    right: 20%;
			    top: 50%;
			    margin-top: -7px;
			}
			.compute{
				padding: 5px 20px;
			    background: #409eff;
			    color: #fff;
			    /*border-radius: 5px;*/
			    text-decoration: none;
			    height: 20px;
			    display: inline-block;
			    line-height: 20px;
			    display: inline-block;
			    /*line-height: 1;*/
			    white-space: nowrap;
			    cursor: pointer;
			    -webkit-appearance: none;
			    text-align: center;
			    outline: none;
			    transition: .1s;
			    font-weight: 500;
			    -moz-user-select: none;
			    -webkit-user-select: none;
			    -ms-user-select: none;
			    font-size: 14px;
			    /*border-radius: 4px;*/
			}
			.cation{
				width: 60%;
			    display: -webkit-inline-box;
			    text-align: left;
			    opacity: .7;
			    color: #a2a2a2;
			    font-weight: 300;
			    font-size: 12px;
			}

			@media screen and (max-width: 768px) {
			    .family{
			    	padding: 0;
			    }
			    .member-info{
			    	padding: 0;
			    	margin: 0 3%;
			    }
			    input[type="text"]{
			    	width: 90%;
			    }
			    .compute-button{
			    	right: 5%;
			    }
			    .el-autocomplete-suggestion{
			    	width: 90%;
			    }
			}
									
		</style>
	</head>
	<body>

		<div class="family">
			<div class="search">

				<span class="search-input">
					<span class="cation">注：数据更新至2014年</span>
					<input type="text" class="RNNXgb" id="c3" placeholder="请输入姓名" name="">
					<input type="text" class="RNNXgb" id="c2"  placeholder="请输入姓名" name="">
					<div class="compute-button"><a href="javascript:void(0);" class="compute">计算关系</a></div>
				</span>

				<span class="search-list">
					<div class="el-autocomplete-suggestion">
						<ul class="el-scrollbar">
							
						</ul>
					</div>
				</span>
			</div>

			<div class="panel">
				<div class="member-info" id="main">
				
				</div>
			
			</div>
			

		</div>



		<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
		<script type="text/javascript">

			var myChart = echarts.init(document.getElementById('main'));

			function toHTML(param){
				return "<span style='font-size:20px;'>"+param+"</span>";
			}

			function getData(url,name1,name2){


				myChart.showLoading();

				myChart.clear();

				$.get(url, function (data) {

					var diff = data.diff

					var ht = diff*(150-4*diff)
					// console.log(height);
					// $(".member-info").css({"height":height});

				console.log(data)
				data = data.data
			    myChart.hideLoading();
			    myChart.resize(opts = {
			    	height : ht
			    });
			    myChart.setOption(option = {
			        tooltip: {
			            trigger: 'item',
			            triggerOn: 'mousemove',
			            formatter: function (params,ticket,callback) {
				            // console.log(params)
				            var birth = params.data.birth
				            var info = params.data.info 
				            if (birth == ""){
				            	birth = "暂无"
				            }
				            if( info == ""){
				            	info = "暂无"
				            }
				            var res = '<div style="width:100%;text-align:center;"><img style="width:30px;height:30px;border-radius:50%;" src="http://imgsrc.baidu.com/imgad/pic/item/cc11728b4710b91250f4e5dbc9fdfc03924522d9.jpg"/></div><br/>出生日期 : '+birth+'<br/>' + '简介 : '+ info;
 				            setTimeout(function (){
				                // 仅为了模拟异步回调
				                callback(ticket, res);
				            }, 500)
				            return 'loading';
				        }
			        },
			        series:[
			            {
			                type: 'tree',

			                data: [data],

			                left: '2%',
			                right: '2%',
			                top: '8%',
			                bottom: '20%',



			                // symbol: 'circle',

			                symbol:"path://M512 538.1c130.9 0 237-106.1 237-237s-106.1-237-237-237-237 106.1-237 237 106.1 237 237 237z m0 110.6c-218.2 0-395.1 69.7-395.1 155.6S293.8 960 512 960s395.1-69.7 395.1-155.6S730.2 648.7 512 648.7z",
			               
			                symbolSize : 15,

			                borderColor : 'rgba(64,158,255,0.8)',

			                orient: 'vertical',

			                expandAndCollapse: true,
			                initialTreeDepth : -1,
			                

			                itemStyle : {

				                color: {
								    type: 'linear',
								    x: 0,
								    y: 0,
								    x2: 0,
								    y2: 1,
								    colorStops: [{
								        offset: 0, color: 'blue' // 0% 处的颜色
								    }, {
								        offset: 1, color: 'blue' // 100% 处的颜色
								    }],
								    globalCoord: false // 缺省为 false
								}


			                },


			                
			                label: {


			                    normal: {
			                    	formatter: function (value) {
			                			// console.log(value,name1);
			                			var name = value.data.name;
			                			arr=name.split("|");
			                			// console.log(arr[0]);
			                			var cen = "第"+arr[1]+"世"
			                			var sex = value.data.sex
			                			var sex_str = sex == 1?"(男)":"(女)"
			                			if (arr[1]==1){
			                				cen = "高祖"
			                			}
			                			if (arr[1]==2){
			                				cen = "二祖"
			                			}
			                			if (arr[1]==3){
			                				cen = "三祖"
			                			}
			                			if (arr[0] == name1){			            		
											return '{nameh|' + arr[0]+sex_str + '}\n{centryh|' + cen + '}';
			                			}else{
			                				return '{name|' + arr[0]+sex_str + '}\n{centry|' + cen + '}';
			                			}
			                			// return value.data.name;
					                	
					            	},
					            	rich: {
						            	name: {
						            		fontSize: 14,
						                    lineHeight: 20,
						                    align: 'center',
						                   
						                },
						                nameh: {
						                	fontSize: 14,
						                    lineHeight: 20,
						                    align: 'center',
						                    color: 'rgba(103,194,58,0.8)'
						                },
						                centry: {
						                    // height: 20,
						                    fontSize: 12,
						                    align: 'center'
						                },
						                centryh: {
						                	fontSize: 12,
						                	align: 'center',
						                	color: 'rgba(103,194,58,0.8)'
						                }
						            },
			                    	fontWeight : 500,
			                        position: 'top',
			                        rotate: 0,
			                        verticalAlign: 'top',
			                        align: 'right',
			                        fontSize: 16,
			                        padding : [ 0,  // 上
											    15, // 右
											    0,  // 下
											    10, // 左
											   ],
			                        // color: 'rgba(64,158,255,0.8)'
			                        color: 'rgba(0,0,0,0.8)'


			                    },
			                   
			                },

			                leaves: {
			                    label: {
			                        normal: {
			                            position: 'top',
			                            rotate: 0,
			                            verticalAlign: 'top',
			                            align: 'center',		                       
			                            padding :20
			                        }
			                    }
			                },

			                animationDurationUpdate: 750
			            }
			        ]
			    });
			});
			}
			

			
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
	        	// $(this).parent(".el-scrollbar").parent("")
	        	var iid = $(this).attr("iid");
	        	// alert(iid);
	        	$("#"+iid).val($(this).children(".info").text());
	        	// $('.RNNXgb').val($(this).children(".info").text());
	        	var id = $(this).attr("data-id");
	        	var bf = $(this).attr("bf");
	        	$("#"+iid).attr("data-id",id);
	        	$("#"+iid).attr("bf",bf);
	        	var name = $(this).children(".info").text();

	        	getData('/member/node/ref?id='+id,name,"");
	        	$(".el-scrollbar").children().remove();
	        });

	        $(".compute").click(function(){

	        	var id_ = $("#c2").attr("data-id");
	        	// console.log(id_)
	        	var id2 = $("#c3").attr("data-id");
	        	
        		if (typeof(id_) =="undefined" || typeof(id2)=="undefined"){
        			alert("两个姓名都要输入");
               		return false;
               	}
               	url = '/member/node/comp?id1='+id_+'&id2='+id2;
              
               	getData(url);
	            // Compute("/node/comp?id1="+id1+"&id2="+id2, $(this));


	        });

	        function Compute(url,dom){
	        	 $.ajax({
	                method: "Get",
	                url: url,
	                // contentType: "application/json;charset=utf-8",
	                // dataType: "json",
	                // data: JSON.stringify(param),
	                success: function (data) {
	                    if (data.code == "0") {
	                       	
	                      	
	                    } else {
	                        console.log(data);
	                    }
	                },
	                error: function (data) {
	                    alert("服务器出错！请联系管理员！");
	                }
	            });
	        }

			function Search(url,dom){
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
	                       	
	                       	var id = dom.attr("id")
	                       	$.each(data.data,function(n,item){
	                       		html = html + `<li class="member-item" data-id= `+item.id+` iid = `+id+` bf= `+item.bf+` >
									<span><img class="avatar" src="/static/img/test.png" /></span>
									<span class="info">`+item.name+`</span>
								</li>`
	                       	})
	                       	dom.parent(".search-input").siblings(".search-list").find(".el-scrollbar").html(html)
	                      	
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