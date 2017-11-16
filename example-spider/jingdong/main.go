package main

import (
	"YiSpider/spider"
	"YiSpider/spider/model"
	spider2 "YiSpider/spider/spider"
	"fmt"
	"strings"
)

func main() {
	goodsType := `电子书刊|电子书|网络原创|数字杂志|多媒体图书|音像|音乐|影视|教育音像|英文原版|少儿|商务投资|英语学习与考试|文学|传记|励志|文艺|小说|文学|青春文学|传记|艺术|少儿|少儿|0-2岁|3-6岁|7-10岁|11-14岁|人文社科|历史|哲学|国学|政治/军事|法律|人文社科|心理学|文化|社会科学|经管励志|经济|金融与投资|管理|励志与成功|生活|生活|健身与保健|家庭与育儿|旅游|烹饪美食|科技|工业技术|科普读物|建筑|医学|科学与自然|计算机与互联网|电子通信|教育|中小学教辅|教育与考试|外语学习|大中专教材|字典词典|港台图书|艺术/设计/收藏|经济管理|文化/学术|少儿|其他|工具书|杂志/期刊|套装书|手机通讯|手机|对讲机|运营商|合约机|选号中心|装宽带|办套餐|手机配件|移动电源|电池/移动电源|蓝牙耳机|充电器/数据线|苹果周边|手机耳机|手机贴膜|手机存储卡|充电器|数据线|手机保护套|车载配件|iPhone 配件|手机电池|创意配件|便携/无线音响|手机饰品|拍照配件|手机支架|大 家 电|平板电视|空调|冰箱|洗衣机|家庭影院|DVD/电视盒子|迷你音响|冷柜/冰吧|家电配件|功放|回音壁/Soundbar|Hi-Fi专区|电视盒子|酒柜|厨卫大电|燃气灶|油烟机|热水器|消毒柜|洗碗机|厨房小电|料理机|榨汁机|电饭煲|电压力锅|豆浆机|咖啡机|微波炉|电烤箱|电磁炉|面包机|煮蛋器|酸奶机|电炖锅|电水壶/热水瓶|电饼铛|多用途锅|电烧烤炉|果蔬解毒机|其它厨房电器|养生壶/煎药壶|电热饭盒|生活电器|取暖电器|净化器|加湿器|扫地机器人|吸尘器|挂烫机/熨斗|插座|电话机|清洁机|除湿机|干衣机|收录/音机|电风扇|冷风扇|其它生活电器|生活电器配件|净水器|饮水机|个护健康|剃须刀|剃/脱毛器|口腔护理|电吹风|美容器|理发器|卷/直发器|按摩椅|按摩器|足浴盆|血压计|电子秤/厨房秤|血糖仪|体温计|其它健康电器|计步器/脂肪检测仪|五金家装|电动工具|手动工具|仪器仪表|浴霸/排气扇|灯具|LED灯|洁身器|水槽|龙头|淋浴花洒|厨卫五金|家具五金|门铃|电气开关|插座|电工电料|监控安防|电线/线缆|摄影摄像|数码相机|单电/微单相机|单反相机|摄像机|拍立得|运动相机|镜头|户外器材|影棚器材|冲印服务|数码相框|数码配件|存储卡|读卡器|滤镜|闪光灯/手柄|相机包|三脚架/云台|相机清洁/贴膜|机身附件|镜头附件|电池/充电器|移动电源|数码支架|智能设备|智能手环|智能手表|智能眼镜|运动跟踪器|健康监测|智能配饰|智能家居|体感车|其他配件|智能机器人|无人机|影音娱乐|MP3/MP4|智能设备|耳机/耳麦|便携/无线音箱|音箱/音响|高清播放器|收音机|MP3/MP4配件|麦克风|专业音频|苹果配件|电子教育|学生平板|点读机/笔|早教益智|录音笔|电纸书|电子词典|复读机|虚拟商品|延保服务|杀毒软件|积分商品|家纺|桌布/罩件|地毯地垫|沙发垫套/椅垫|床品套件|被子|枕芯|床单被罩|毯子|床垫/床褥|蚊帐|抱枕靠垫|毛巾浴巾|电热毯|窗帘/窗纱|布艺软饰|凉席|灯具|台灯|节能灯|装饰灯|落地灯|应急灯/手电|LED灯|吸顶灯|五金电器|筒灯射灯|吊灯|氛围照明|生活日用|保暖防护|收纳用品|雨伞雨具|浴室用品|缝纫/针织用品|洗晒/熨烫|净化除味|家装软饰|相框/照片墙|装饰字画|节庆饰品|手工/十字绣|装饰摆件|帘艺隔断|墙贴/装饰贴|钟饰|花瓶花艺|香薰蜡烛|创意家居|宠物生活|宠物主粮|宠物零食|医疗保健|家居日用|宠物玩具|出行装备|洗护美容|电脑整机|笔记本|超极本|游戏本|平板电脑|平板电脑配件|台式机|服务器/工作站|笔记本配件|一体机|电脑配件|CPU|主板|显卡|硬盘|SSD固态硬盘|内存|机箱|电源|显示器|刻录机/光驱|散热器|声卡/扩展卡|装机配件|组装电脑|外设产品|移动硬盘|U盘|鼠标|键盘|鼠标垫|摄像头|手写板|硬盘盒|插座|线缆|UPS电源|电脑工具|游戏设备|电玩|电脑清洁|网络仪表仪器|游戏设备|游戏机|游戏耳机|手柄/方向盘|游戏软件|游戏周边|网络产品|路由器|网卡|交换机|网络存储|4G/3G上网|网络盒子|网络配件|办公设备|投影机|投影配件|多功能一体机|打印机|传真设备|验钞/点钞机|扫描设备|复合机|碎纸机|考勤机|收款/POS机|会议音频视频|保险柜|装订/封装机|安防监控|办公家具|白板|文具/耗材|硒鼓/墨粉|墨盒|色带|纸类|办公文具|学生文具|财会用品|文件管理|本册/便签|计算器|笔类|画具画材|刻录碟片/附件|服务产品|上门安装|延保服务|维修保养|电脑软件|京东服务|烹饪锅具|炒锅|煎锅|压力锅|蒸锅|汤锅|奶锅|锅具套装|煲类|水壶|火锅|刀剪菜板|菜刀|剪刀|刀具套装|砧板|瓜果刀/刨|多功能刀|厨房配件|保鲜盒|烘焙/烧烤|饭盒/提锅|储物/置物架|厨房DIY/小工具|水具酒具|塑料杯|运动水壶|玻璃杯|陶瓷/马克杯|保温杯|保温壶|酒杯/酒具|杯具套装|餐具|餐具套装|碗/碟/盘|筷勺/刀叉|一次性用品|果盘/果篮|酒店用品|自助餐炉|酒店餐具|酒店水具|茶具/咖啡具|整套茶具|茶杯|茶壶|茶盘茶托|茶叶罐|茶具配件|茶宠摆件|咖啡具|其他|清洁用品|纸品湿巾|衣物清洁|清洁工具|驱虫用品|家庭清洁|皮具护理|一次性用品|面部护肤|洁面|乳液面霜|面膜|剃须|套装|精华|眼霜|卸妆|防晒|防晒隔离|T区护理|眼部护理|精华露|爽肤水|身体护理|沐浴|润肤|颈部|手足|纤体塑形|美胸|套装|精油|洗发护发|染发/造型|香薰精油|磨砂/浴盐|手工/香皂|洗发|护发|染发|磨砂膏|香皂|口腔护理|牙膏/牙粉|牙刷/牙线|漱口水|套装|女性护理|卫生巾|卫生护垫|私密护理|脱毛膏|其他|洗发护发|洗发|护发|染发|造型|假发|套装|美发工具|脸部护理|香水彩妆|香水|底妆|腮红|眼影|唇部|美甲|眼线|美妆工具|套装|防晒隔离|卸妆|眉笔|睫毛膏|女装|T恤|衬衫|针织衫|雪纺衫|卫衣|马甲|连衣裙|半身裙|牛仔裤|休闲裤|打底裤|正装裤|小西装|短外套|风衣|毛呢大衣|真皮皮衣|棉服|羽绒服|大码女装|中老年女装|婚纱|打底衫|旗袍/唐装|加绒裤|吊带/背心|羊绒衫|短裤|皮草|礼服|仿皮皮衣|羊毛衫|设计师/潮牌|男装|衬衫|T恤|POLO衫|针织衫|羊绒衫|卫衣|马甲/背心|夹克|风衣|毛呢大衣|仿皮皮衣|西服|棉服|羽绒服|牛仔裤|休闲裤|西裤|西服套装|大码男装|中老年男装|唐装/中山装|工装|真皮皮衣|加绒裤|卫裤/运动裤|短裤|设计师/潮牌|羊毛衫|内衣|文胸|女式内裤|男式内裤|睡衣/家居服|塑身美体|泳衣|吊带/背心|抹胸|连裤袜/丝袜|美腿袜|商务男袜|保暖内衣|情侣睡衣|文胸套装|少女文胸|休闲棉袜 |大码内衣|内衣配件|打底裤袜|打底衫|秋衣秋裤|情趣内衣|洗衣服务|服装洗护|服饰配件|太阳镜|光学镜架/镜片|围巾/手套/帽子套装|袖扣|棒球帽|毛线帽|遮阳帽|老花镜|装饰眼镜|防辐射眼镜|游泳镜|女士丝巾/围巾/披肩|男士丝巾/围巾|鸭舌帽|贝雷帽|礼帽|真皮手套|毛线手套|防晒手套|男士腰带/礼盒|女士腰带/礼盒|钥匙扣|遮阳伞/雨伞|口罩|耳罩/耳包|假领|毛线/布面料|领带/领结/领带夹|钟表|男表|瑞表|女表|国表|日韩表|欧美表|德表|儿童手表|智能手表|闹钟|座钟挂钟|钟表配件|流行男鞋|商务休闲鞋|正装鞋|休闲鞋|凉鞋/沙滩鞋|男靴|功能鞋|拖鞋/人字拖|雨鞋/雨靴|传统布鞋|鞋配件|帆布鞋|增高鞋|工装鞋|定制鞋|时尚女鞋|高跟鞋|单鞋|休闲鞋|凉鞋|女靴|雪地靴|拖鞋/人字拖|踝靴|筒靴|帆布鞋|雨鞋/雨靴|妈妈鞋|鞋配件|特色鞋|鱼嘴鞋|布鞋/绣花鞋|马丁靴|坡跟鞋|松糕鞋|内增高|防水台|奶粉|婴幼奶粉|孕妈奶粉|营养辅食|益生菌/初乳|米粉/菜粉|果泥/果汁|DHA|宝宝零食|钙铁锌/维生素|清火/开胃|面条/粥|尿裤湿巾|婴儿尿裤|拉拉裤|婴儿湿巾|成人尿裤|喂养用品|奶瓶奶嘴|吸奶器|暖奶消毒|儿童餐具|水壶/水杯|牙胶安抚|围兜/防溅衣|辅食料理机|食物存储|洗护用品|宝宝护肤|洗发沐浴|奶瓶清洗|驱蚊防晒|理发器|洗澡用具|婴儿口腔清洁|洗衣液/皂|日常护理|座便器|童车童床|婴儿推车|餐椅摇椅|婴儿床|学步车|三轮车|自行车|电动车|扭扭车|滑板车|婴儿床垫|寝居服饰|婴儿外出服|婴儿内衣|婴儿礼盒|婴儿鞋帽袜|安全防护|家居床品|睡袋/抱被|爬行垫|妈妈专区|妈咪包/背婴带|产后塑身|文胸/内裤|防辐射服|孕妈装|孕期营养|孕妇护肤|待产护理|月子装|防溢乳垫|童装童鞋|套装|上衣|裤子|裙子|内衣/家居服|羽绒服/棉服|亲子装|儿童配饰|礼服/演出服|运动鞋|皮鞋/帆布鞋|靴子|凉鞋|功能鞋|户外/运动服|安全座椅|提篮式|安全座椅|增高垫|潮流女包|钱包|手拿包|单肩包|双肩包|手提包|斜挎包|钥匙包|卡包/零钱包|精品男包|男士钱包|男士手包|卡包名片夹|商务公文包|双肩包|单肩/斜挎包|钥匙包|功能箱包|电脑包|拉杆箱|旅行包|旅行配件|休闲运动包|拉杆包|登山包|妈咪包|书包|相机包|腰包/胸包|礼品|火机烟具|礼品文具|军刀军具|收藏品|工艺礼品|创意礼品|礼盒礼券|鲜花绿植|婚庆节庆|京东卡|美妆礼品|礼品定制|京东福卡|古董文玩|奢侈品|箱包|钱包|服饰|腰带|太阳镜/眼镜框|配件|鞋靴|饰品|名品腕表|高档化妆品|婚庆|婚嫁首饰|婚纱摄影|婚纱礼服|婚庆服务|婚庆礼品/用品|婚宴|进口食品|饼干蛋糕|糖果/巧克力|休闲零食|冲调饮品|粮油调味|牛奶|地方特产|其他特产|新疆|北京|山西|内蒙古|福建|湖南|四川|云南|东北|休闲食品|休闲零食|坚果炒货|肉干肉脯|蜜饯果干|糖果/巧克力|饼干蛋糕|无糖食品|粮油调味|米面杂粮|食用油|调味品|南北干货|方便食品|有机食品|饮料冲调|饮用水|饮料|牛奶乳品|咖啡/奶茶|冲饮谷物|蜂蜜/柚子茶|成人奶粉|食品礼券|月饼|大闸蟹|粽子|卡券|茗茶|铁观音|普洱|龙井|绿茶|红茶|乌龙茶|花草茶|花果茶|养生茶|黑茶|白茶|其它茶|时尚饰品|项链|手链/脚链|戒指|耳饰|毛衣链|发饰/发卡|胸针|饰品配件|婚庆饰品|黄金|黄金吊坠|黄金项链|黄金转运珠|黄金手镯/手链/脚链|黄金耳饰|黄金戒指|K金饰品|K金吊坠|K金项链|K金手镯/手链/脚链|K金戒指|K金耳饰|金银投资|投资金|投资银|投资收藏|银饰|银吊坠/项链|银手镯/手链/脚链|银戒指|银耳饰|足银手镯|宝宝银饰|钻石|裸钻|钻戒|钻石项链/吊坠|钻石耳饰|钻石手镯/手链|翡翠玉石|项链/吊坠|手镯/手串|戒指|耳饰|挂件/摆件/把件|玉石孤品|水晶玛瑙|项链/吊坠|耳饰|手镯/手链/脚链|戒指|头饰/胸针|摆件/挂件|彩宝|琥珀/蜜蜡|碧玺|红宝石/蓝宝石|坦桑石|珊瑚|祖母绿|葡萄石|其他天然宝石|项链/吊坠|耳饰|手镯/手链|戒指|铂金|铂金项链/吊坠|铂金手镯/手链/脚链|铂金戒指|铂金耳饰|木手串/把件|小叶紫檀|黄花梨|沉香木|金丝楠|菩提|其他|橄榄核/核桃|檀香|珍珠|珍珠项链|珍珠吊坠|珍珠耳饰|珍珠手链|珍珠戒指|珍珠胸针|维修保养|机油|正时皮带|添加剂|汽车喇叭|防冻液|汽车玻璃|滤清器|火花塞|减震器|柴机油/辅助油|雨刷|车灯|后视镜|轮胎|轮毂|刹车片/盘|维修配件|蓄电池|底盘装甲/护板|贴膜|汽修工具|改装配件|车载电器|导航仪|安全预警仪|行车记录仪|倒车雷达|蓝牙设备|车载影音|净化器|电源|智能驾驶|车载电台|车载电器配件|吸尘器|智能车机|冰箱|汽车音响|车载生活电器|美容清洗|车蜡|补漆笔|玻璃水|清洁剂|洗车工具|镀晶镀膜|打蜡机|洗车配件|洗车机|洗车水枪|毛巾掸子|汽车装饰|脚垫|座垫|座套|后备箱垫|头枕腰靠|方向盘套|香水|空气净化|挂件摆件|功能小件|车身装饰件|车衣|安全自驾|安全座椅|胎压监测|防盗设备|应急救援|保温箱|地锁|摩托车|充气泵|储物箱|自驾野营|摩托车装备|汽车服务|清洗美容|功能升级|保养维修|油卡充值|车险|加油卡|ETC|驾驶培训|赛事改装|赛事服装|赛事用品|制动系统|悬挂系统|进气系统|排气系统|电子管理|车身强化|赛事座椅|运动鞋包|跑步鞋|休闲鞋|篮球鞋|板鞋|帆布鞋|足球鞋|乒羽网鞋|专项运动鞋|训练鞋|拖鞋|运动包|运动服饰|羽绒服|棉服|运动裤|夹克/风衣|卫衣/套头衫|T恤|套装|乒羽网服|健身服|运动背心|毛衫/线衫|运动配饰|骑行运动|折叠车|山地车/公路车|电动车|其他整车|骑行服|骑行装备|平衡车|垂钓用品|鱼竿鱼线|浮漂鱼饵|钓鱼桌椅|钓鱼配件|钓箱鱼包|其它|游泳用品|泳镜|泳帽|游泳包防水包|女士泳衣|男士泳衣|比基尼|其它|户外鞋服|冲锋衣裤|速干衣裤|滑雪服|羽绒服/棉服|休闲衣裤|抓绒衣裤|软壳衣裤|T恤|户外风衣|功能内衣|军迷服饰|登山鞋|雪地靴|徒步鞋|越野跑鞋|休闲鞋|工装鞋|溯溪鞋|沙滩/凉拖|户外袜|户外装备|帐篷/垫子|睡袋/吊床|登山攀岩|户外配饰|背包|户外照明|户外仪表|户外工具|望远镜|旅游用品|便携桌椅床|野餐烧烤|军迷用品|救援装备|滑雪装备|极限户外|冲浪潜水|健身训练|综合训练器|其他大型器械|哑铃|仰卧板/收腹机|其他中小型器材|瑜伽舞蹈|甩脂机|踏步机|武术搏击|健身车/动感单车|跑步机|运动护具|体育用品|羽毛球|乒乓球|篮球|足球|网球|排球|高尔夫|台球|棋牌麻将|轮滑滑板|其他|适用年龄|0-6个月|6-12个月|1-3岁|3-6岁|6-14岁|14岁以上|遥控/电动|遥控车|遥控飞机|遥控船|机器人|轨道/助力|毛绒布艺|毛绒/布艺|靠垫/抱枕|娃娃玩具|芭比娃娃|卡通娃娃|智能娃娃|模型玩具|仿真模型|拼插模型|收藏爱好|健身玩具|炫舞毯|爬行垫/毯|户外玩具|戏水玩具|动漫玩具|电影周边|卡通周边|网游周边|益智玩具|摇铃/床铃|健身架|早教启智|拖拉玩具|积木拼插|积木|拼图|磁力棒|立体拼插|DIY玩具|手工彩泥|绘画工具|情景玩具|创意减压|减压玩具|创意玩具|乐器|钢琴|电子琴/电钢琴|吉他/尤克里里|打击乐器|西洋管弦|民族管弦乐器|乐器配件|电脑音乐|工艺礼品乐器|口琴/口风琴/竖笛|手风琴||机票|国内机票|酒店|国内酒店|酒店团购|旅行|度假|景点|租车|火车票|旅游团购|充值|手机充值|游戏|游戏点卡|QQ充值|票务|电影票|演唱会|话剧歌剧|音乐会|体育赛事|舞蹈芭蕾|戏曲综艺|产地直供|水果|苹果|橙子|奇异果/猕猴桃|车厘子/樱桃|芒果|蓝莓|火龙果|葡萄/提子|柚子|香蕉|牛油果|梨|菠萝/凤梨|桔/橘|柠檬|草莓|桃/李/杏|更多水果|水果礼盒/券|猪牛羊肉|牛肉|羊肉|猪肉|内脏类|海鲜水产|鱼类|虾类|蟹类|贝类|海参|海产干货|其他水产|海产礼盒|禽肉蛋品|鸡肉|鸭肉|蛋类|其他禽类|冷冻食品|水饺/馄饨|汤圆/元宵|面点|火锅丸串|速冻半成品|奶酪黄油|熟食腊味|熟食|腊肠/腊肉|火腿|糕点|礼品卡券|饮品甜品|冷藏果蔬汁|冰激凌|其他`
	//goodsType := `猫`
	goodsTypes := strings.Split(goodsType, "|")
	reqs := []*model.Request{}

	for _, ty := range goodsTypes {
		reqs = append(reqs, &model.Request{
			Method:      "get",
			Url:         fmt.Sprintf("https://search.jd.com/Search?keyword=%s&enc=utf-8&page={1-5,1}", ty),
			ProcessName: "jingdong-list",
		})
	}

	task := &model.Task{
		Id:      "jingdong",
		Name:    "jingdong",
		Request: reqs,
		Process: []model.Process{
			{
				Name: "jingdong-list",
				Type: "template",
				TemplateRule: model.TemplateRule{
					Rule: map[string]string{
						"node":        "array|.gl-item",
						"img":         "attr.src|.err-product",
						"price":       "text|.p-price strong i",
						"goods_name":  "text|.p-name a em",
						"desc":        "text|.p-name a i",
						"comment_num": "text|.p-commit strong a",
						"shop_addr":   "attr.href|.curr-shop",
						"shop_name":   "attr.title|.curr-shop",
						"goods_id":    "attr.data-sku|.J_focus",
					},
				},
				AddQueue: []*model.Request{
					{
						Method:      "get",
						Url:         "https://sclub.jd.com/comment/productPageComments.action?productId={$goods_id}&score=0&sortType=5&page=0&pageSize=10",
						ProcessName: "jingdong-comment-first",
					},
				},
			},
			{
				//
				Name: "jingdong-comment-first",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						"max_page": "maxPage",
						"id":       "productCommentSummary.productId",
					},
				},
				AddQueue: []*model.Request{
					{
						Method:      "get",
						Url:         "https://sclub.jd.com/comment/productPageComments.action?productId={$id}&score=0&sortType=5&page={0-$max_page,1}&pageSize=10",
						ProcessName: "jingdong-comments",
					},
				},
			},
			{
				Name: "jingdong-comments",
				Type: "json",
				JsonRule: model.JsonRule{
					Rule: map[string]string{
						"node":            "array|comments",
						"comment_id":      "id",
						"content":         "content",
						"create_time":     "creationTime",
						"image_count":     "imageCount",
						"isMobile":        "isMobile",
						"productColor":    "productColor",
						"productSize":     "productSize",
						"productId":       "referenceId",
						"score":           "score",
						"replyCount":      "replyCount",
						"usefulVoteCount": "usefulVoteCount",
						"userClient":      "userClient",
						"userClientShow":  "userClientShow",
						"userLevelId":     "userLevelId",
						"userLevelName":   "userLevelName",
						"userProvince":    "userProvince",
						"nickname":        "nickname",
					},
				},
			},
			{
				Name: "jingdong-type",
				Type: "template",
				TemplateRule: model.TemplateRule{
					Rule: map[string]string{
						"type": "texts|.items a",
					},
				},
			},
		},

		Pipline: "file",
	}

	app := spider.New()
	app.AddSpider(spider2.InitWithTask(task))
	app.Run()
}
