<html>
<head>
    <meta charset="UTF-8"/>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, minimum-scale=0.5, maximum-scale=2.0, user-scalable=yes" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge"/>
    <title>家谱登录</title>
    <link rel="shortcut icon" href="/static/img/home.png" type="image/x-icon"/>

    <link rel="stylesheet" type="text/css" href="/static/css/zzsc.css">
    <link rel="stylesheet" type="text/css" href="/static/css/txt.wav.css">


    <style>
        body{
            margin: 0;padding: 0;border: 0;
            background: #;
            height: 100vh;
        }
        .flex-center {
            height: 100%;
            display: flex;
            align-items: center;
            /*justify-content: center;*/
        }
        span{
            display: block;
        }
        input{
            height: 40px;
            margin: 5px 0;
            width: 90%;
            outline: none;
            font-size: 16px;
        }
        .ok_button {

            -webkit-appearance: none;

            border: 0;
            /*padding: 10px 20px;*/
            background: #1e88e5;
            height: 40px;
            /*line-height: 40px;*/
            color: #fff;
            cursor: pointer;
            border-radius: 3px;
        }
        .slogan{
            text-align: center;
        }
        form{
            text-align: center;
        }
        input[type="text"],input[type="password"]  {
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
                /*line-height: 44px;*/
                outline: none;
                padding: 0 15px;
                transition: border-color .2s cubic-bezier(.645,.045,.355,1);
                /*width: 60%;*/
        }
        .RNNXgb {
                background: #fff;
                display: flex;
                /*border-radius: 2px;*/
                border: none;
                box-shadow: 0 2px 2px 0 rgba(0,0,0,0.16), 0 0 0 1px rgba(0,0,0,0.08);
                z-index: 3;
                height: 44px;
                margin: 5px 0;
                /*width: 484px;*/
            }
        .forhead{
            margin: 50px 10%;
            height: 500px;
            display: -webkit-box;
            /*width: 40%;*/
        }
        .forhead div{
            writing-mode: tb-rl;
            color: #a2a2a2;
            font-size: 16px;
            font-weight: 200;
            margin: 0 3px;
            letter-spacing: 1px;
            text-align: justify;
            line-height: 30px;
        }

        @media screen and (max-width: 768px) {
            .forhead{
                display: none;
            }
        }

    </style>
</head>
<body>
    <div class="flex-center">
        <div class="forhead zzsc">
            <div class="txtwav slow">尝观天地之间 物各有主 始水莫不有其源 木莫不有本 昭昭然矣 况人为万物之灵 可不知其所由始乎 予生不肖 赖祖宗之荫蔽 蒙父师之教育 年未及冠 叨入黉宫 因念祖宗之功德 难报子孙之追荐 宜诚无奈世远年湮 未有家谱 遂不觉於我心有戚戚焉 幸有父老传言 若合一口 今日来自山西洪洞县大槐树迁移於此 离沙河城十五里西南尚贤村安居 又有轴可考证 
               
                上书 
            高祖王管垢 次书 二祖王德 又书 三祖王玉 又书 四祖王朝生四子 长曰聪 二曰昇 三曰学 四曰举 又书 五祖王聪生四子 长曰亭章 二曰亭道 三曰月 四曰本奇 
            

                                      由是观之 或者因此有老四门 小四门之说焉 分门别宗 亭章爷为小四门头门 自子孙元曾 昆云仍来之派行 及予身 乃亭章爷之耳孙也 迄今有一户有一户之轴 一门有一门之轴 参观互考 源流分明 不得不誌之以传 勿替云 然莫为之前 虽美弗彰 莫为之后 虽盛弗传 若夫子子孙孙继继承承 犹有待于后人之缵绪也夫 
                           时 
                                                                            大清道光五年 岁次乙酉 邑庠生王廷辅 题于忠如堂
            </div>
        </div>
        
        <form action="/member/login" method="post">
            <!-- <h2 class="slogan">应用管理</h2> -->
            <input type="text" name="username" class="RNNXgb" placeholder="用户名"/>
            <input type="password" name="password" class="RNNXgb" placeholder="密码"/>
            <input type="submit" class="ok_button" value="登录"/>
        </form>
    </div>
    <script src="/static/js/txt.wav.min.js"></script>
    <!-- <script src="/gg_bd_ad_720x90.js" type="text/javascript"></script> -->
    <!-- <script src="/follow.js" type="text/javascript"></script> -->
</body>