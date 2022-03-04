/*
 * SYSUI框架1.2测试
 * 799129700@qq.com SYSHUXL
 * Reserved head available commercial use
 * Universal background system interface framework
 */
(function($) {
	"use strict";
	$.fn.SYSFrame = function(options) {
		var frame = this;
		var defaults = {
			hrederheight: 50, //顶部高度
			footerheight: 40, //底部高度
			sortmodeWidth: 60, //该参数为竖排菜单栏设置图标菜单时的宽度值
			menuheight: 80, //菜单栏高度（该方法只使用菜单栏显示在顶部横排时有效）
			menuWidth: 200, //当差菜单栏显示在左侧或右侧竖排显示时该方法有效
			cookieDate: 7, //设置皮肤显示保存时间（天单位）
			minStatue:'', //默认显示或隐藏菜单栏  (true/false)
			data: null, //数据
			rightclick: false, //是否启用右键功能
			titlename: '菜单栏目', //设置菜单栏名称
			header: '.header-section', //框架顶部
			footer: '.footer-section', //框架底部
			content: '.content-section', //框架内容区
			message: '#iframe_box', //设置窗体显示位置
			menuModule: "#bk-con-menu", //菜单模块
			menu: '.menu-section', //菜单栏
			templateStyle:'#Template_style',//设置模版风格显示
			menuopt: '', //栏目显示选择模式
			closebtn: '.close_btn', //点击隐藏
			showbtn: '.show_btn', //点击显示	
			slide: '#breadcrumb',
			Bombboxinfo: '.Bombbox-info', //消息框
			boxname: '', //设置你需要弹框模块的名称
			Promp: '.PrompClose ', //提示盒子
			homepage: "home.html", // 默认显示页
			prompt_btn: '.dropdown',
			iframe: '.show_iframe',
			fullscreen: '#fullscreen',
			tabs: '#tab-navigation',
			iconfont: 'iconfont', //字体库显示图标名称(自定义字体库名称)
			iconSkin: 'icon-duigou', //皮肤选中状态图标
			iconsort: '&#xe618', //图文展示栏目
			iconsortg: '&#xe658', //图标展示栏目
			iconExpandicon: '&#xe653', //全屏图标
			iconShrink: '&#xe654', //全屏返回图标
			iconDelete: '&#xe70a', //删除图标
			iconHome: '&#xe626', //默认首页显示图标
			icontoplarrow:'&#xe6d0',//
			icontoprarrow:'&#xe6c9',//
			mouIconOpen: '&#xe615', // 菜单 下箭头
			mouIconClose: '&#xe630', // 菜单 上箭头
			scrollbar: function(e) {}, //设置滚动模式
			expand: function() {} //扩展
		};
		var settings = $.extend({}, defaults, options); //将一个空对象做为第一个参数
		return this.each(function() {
			var thisBox = $(this), //当前对象
				headerform = thisBox.find(settings.header),
				footerform = thisBox.find(settings.footer),
				mmodule = thisBox.find(settings.menuModule),
				contentform = thisBox.find(settings.content),
				tabsform = thisBox.find(settings.tabs),
				iframeform = thisBox.find(settings.iframe),
				Bombboxform = thisBox.find(settings.Bombboxinfo),
				Prompname = thisBox.find(settings.Promp),
				Bombbox = thisBox.find(settings.boxname),
				messageform = thisBox.find(settings.message),
				closemenu = thisBox.find(settings.closebtn),
				showmenu = thisBox.find(settings.showbtn),
				menuform = thisBox.find(settings.menu),
				temsye=thisBox.find(settings.templateStyle);
//*****************************************************************************事件处理
			headerform.css({
				height: (settings.hrederheight),
				position: "relative",
				lineHeight: (settings.hrederheight) + 'px'
			});
			settings.expand(thisBox, settings);
			$.learuntab = {
				//全屏模式
				requestFullScreen: function() {
					var de = document.documentElement;
					if(de.requestFullscreen) {
						de.requestFullscreen();
					} else if(de.mozRequestFullScreen) {
						de.mozRequestFullScreen();
					} else if(de.webkitRequestFullScreen) {
						de.webkitRequestFullScreen();
					}
				},
				exitFullscreen: function() {
					var de = document;
					if(de.exitFullscreen) {
						de.exitFullscreen();
					} else if(de.mozCancelFullScreen) {
						de.mozCancelFullScreen();
					} else if(de.webkitCancelFullScreen) {
						de.webkitCancelFullScreen();
					}
				},
				//判断是手机还是pc
				isMobile: function(mobile_flag) {
				    var userAgentInfo = navigator.userAgent;
				    var mobileAgents = ["Android", "iPhone", "SymbianOS", "Windows Phone", "iPad", "iPod"];
				    var mobile_flag = false;
				    //根据userAgent判断是否是手机
				    for(var v = 0; v < mobileAgents.length; v++) {
					    if(userAgentInfo.indexOf(mobileAgents[v]) > 0) {
						mobile_flag = true;
						break;
					    }
				    }
					var screen_width = window.screen.width;
					var screen_height = window.screen.height;
					//根据屏幕分辨率判断是否是手机
					if(screen_width < 500 && screen_height < 800) {
						mobile_flag = true;
					}
					return mobile_flag;
				},
			    //判断移动端还是pc端的模块显示处理
				mobileflag:function(framename) {
				    var mobile_flag = $.learuntab.isMobile(mobile_flag);
					if(mobile_flag) {
					    /* 设计图文档宽度 */
					    var docWidth = 750;
					    var doc = window.document,
					        docEl = doc.documentElement,
					        resizeEvt = 'orientationchange' in window ? 'orientationchange' : 'resize';
					    var recalc = (function refreshRem () {
					        var clientWidth = docEl.getBoundingClientRect().width;
					        /* 8.55：小于320px不再缩小，11.2：大于420px不再放大 */
					        docEl.style.fontSize = Math.max(Math.min(20 * (clientWidth / docWidth), 11.2), 8.55) * 5 + 'px';
					
					        return refreshRem;
					    })();
					    /* 添加倍屏标识，安卓倍屏为1 */
					    docEl.setAttribute('data-dpr', window.navigator.appVersion.match(/iphone/gi) ? window.devicePixelRatio : 1);
					
					    if (/iP(hone|od|ad)/.test(window.navigator.userAgent)) {
					        /* 添加IOS标识 */
					        doc.documentElement.classList.add('ios');
					        /* IOS8以上给html添加hairline样式，以
					         * 便特殊处理 */
					        if (parseInt(window.navigator.appVersion.match(/OS (\d+)_(\d+)_?(\d+)?/)[1], 10) >= 8)
					            doc.documentElement.classList.add('hairline');
					    }
					    if (!doc.addEventListener) return;
					    window.addEventListener(resizeEvt, recalc, false);
					    doc.addEventListener('DOMContentLoaded', recalc, false);
					    /******************************************************************************/
						footerform.css({
						    height: (settings.footerheight),
					    	bottom: 0,
						    width: "100%"
					    }).removeClass("footer-section");
						contentform.css({height: $(window).height() - ((settings.hrederheight) * 2 + (settings.footerheight))});
						closemenu.hide();
						$(".timehtml").remove();
						messageform.animate({left: 0,width: $(window).width()}, 300);
						showmenu.show().css("display", "block");
						if(framename == "left") {
							menuform.css({
							height: $(window).height() - $(".title-menu").height(),
							width:settings.menuWidth,
							left:""
						});
						}if(framename == "top"){
							showmenu.hide().css("display", "none");
							closemenu.show().css("display", "block");
							menuform.css({height: (settings.menuheight)/2,width:""});
							$.learuntab.tabNavallwidth(framename);
							$.learuntab.Menumove();	
						}
					} else {
						if(framename == "left"){
							footerform.css({
							top: (settings.hrederheight)
						}).addClass("footer-top");
						contentform.css({
							top: (settings.footerheight)
						});
						}if(framename == "top"){
							$("#tab-column").css({
							width: $(window).width() - $("#operation").width() - settings.menuWidth - 1,
							height: (settings.footerheight) - 1,
							marginTop: settings.hheight,
							position: "relative"
						});	
						}
						$(".timehtml").css({
							width: settings.menuWidth,
							float: "left",
							height: (settings.footerheight) + 'px',
							"text-align": "center"
						});						
						
					}
				},
				skinchange: function(skinName) {
					$("#skin").attr("href", "skin/" + skinName + "/skin.css");
					$.cookie("MyCssSkin", skinName, {
						path: '/',
						expires: settings.cookieDate
					});
				},
				fstructure: function(structureName) {	
				   temsye.attr("data-name",structureName);
					$.cookie("My", structureName, {
						path: '/',
						expires:settings.cookieDate
					});	
				},
				//设置框架显示模式
				framemode: function(framename) {
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					footerform.find('.timehtml').remove();
					$(settings.menuModule).find(".title-menu").remove();
					if(framename == "left") {
						var iconsort = settings.iconsort;
						var timehtml = '<div class="timehtml"></div>';
						var titlename = '<div class="title-menu"><span class="' + settings.iconfont + ' icon-menu" id="menuclick" ischek="false">' + iconsort + '</span>' + settings.titlename + '</div>';
						footerform.prepend(timehtml);
						$(settings.menuModule).prepend(titlename);				
						$(settings.menuModule).removeClass("topModule").addClass("leftModule").css({
							width: settings.menuWidth,
							height: ($(window).height() - ((settings.hrederheight) + (settings.footerheight))),
							float: "left",
						});
						messageform.removeClass("topMessage").addClass("leftMessage").css({
							float: "left",
							left: settings.menuWidth + "px",
							position: "absolute",
							width: $(window).width() - settings.menuWidth,
							height: $(window).height() - ((settings.hrederheight) + (settings.footerheight)),
							"z-index": 1
						});	
						contentform.css({height: $(window).height() - ((settings.hrederheight) + (settings.footerheight)),top:"",position: "relative"});
						$(settings.menuModule).find(".radius").remove();
							var menuclick=$("#menuclick");
						    menuclick.on('click',function(){
							    var iconsort = settings.iconsort,
							        iconsortg=settings.iconsortg;
			           	        var ischek = $(this).attr("ischek");
								if(ischek == 'false') {
									$(settings.slide).addClass("extendstyle");
									$(this).attr("ischek", 'true').addClass('stretch').html(iconsortg);
									$(settings.menuModule).animate({width: settings.sortmodeWidth}, 200);
									menuform.animate({width: settings.sortmodeWidth}, 200).find(".menu-text").hide().parent().parent().find(".submenu").addClass("elasticStyle");
									$(".elasticStyle").css({left: settings.sortmodeWidth});
									menuform.find(".arrow").css("right", "0")
									$(".title-menu").contents().filter(function() {return this.nodeType === 3}).remove();
									messageform.animate({width: $(window).width() - settings.sortmodeWidth, left: settings.sortmodeWidth}, 200);
								} else {
									$(settings.slide).removeClass("extendstyle");
									$(this).attr("ischek", 'false').removeClass('stretch').html(iconsort);
									$(".title-menu").prepend(settings.titlename);
									$(settings.menuModule).animate({width: settings.menuWidth}, 200);
									$(".elasticStyle").css({left: 0});
									menuform.animate({width: settings.menuWidth}, 200).find(".menu-text").show().parent().parent().find(".submenu").removeClass("elasticStyle");
									messageform.animate({width: $(window).width()-settings.menuWidth, left: settings.menuWidth}, 200);
									menuform.find(".arrow").css("right", "");
								}
							    menuform.find('.submenu .menu-text').show();
					        });
					    settings.scrollbar($(settings.slide));	
						menuform.css({
						height: $(window).height() - (settings.hrederheight + settings.footerheight) - $(".title-menu").height(),
						width: settings.menuWidth,
						left:"" 
						}).attr("id", "menuList");	;
					};
					if(framename == "top") {
						$(settings.menuModule).removeClass("leftModule").addClass("topModule").css({ width:"",height:settings.menuheight, float: ""});
						messageform.removeClass("leftMessage").addClass("topMessage").css({position: "relative",float: "",left:"",width:""});
						$(".title-menu").remove();
						footerform.find(".timehtml").remove();
						contentform.css({height: $(window).height() - ((settings.hrederheight) + (settings.footerheight)),top:""});
						footerform.css({top:""}).removeClass("footer-top");
						menuform.css({height: (settings.menuheight),width:""});
						$(settings.menuModule).append('<a id="js-tabNav-next" class="radius btn-default right_roll" href="javascript:void(0);"><i class="'+settings.iconfont+'">'+ settings.icontoprarrow+'</i></a>');
						$(settings.menuModule).prepend('<a id="js-tabNav-prev" class="radius btn-default left_roll" href="javascript:void(0);"><i class="'+settings.iconfont +'">'+ settings.icontoplarrow+'</i></a>');	
						$.learuntab.Menumove();	
						$.learuntab.tabNavallwidth(framename);
					};
					//判断是否是电脑还是移动手机
						if(mobile_flag) {
							thisBox.addClass("mobileStyle");
							//移动端顶部高度变化
							headerform.css({
								height: (settings.hrederheight) * 2
							});
							if(framename == "left") {
								messageform.css({
									height: $(window).height() - ((settings.hrederheight) * 2 + (settings.footerheight)),
									left: 0,
									width: $(window).width(),
									"z-index": 11
								});
								$(settings.menuModule).css({
									position: "absolute",
									width: 0,
									"z-index": 111,
									height: $(window).height(),
									top:-(settings.hrederheight) * 2
								});
								$(settings.menuModule).find('.title-menu').hide();

							} else {
								menuform.css({height: (settings.menuheight) / 2}).addClass("wap-menulist");
								messageform.css({
									height: ($(window).height() - ((settings.hrederheight) * 2 + (settings.footerheight) + (settings.menuheight) / 2))
								});
								$(settings.menuModule).css({
									height:(settings.menuheight) / 2
									
								});
							}
							$(".operate-header , .bk-herder-logo").addClass("wap-header");
							$(".wap-header").css({
								height: (settings.hrederheight),
								lineHeight: (settings.hrederheight) + 'px',
								position: "relative"
							});
							$(settings.fullscreen).css("display", "none");
							Bombboxform.addClass("wap-Bombbox");

						} else {
							$(".operate-header , .bk-herder-logo").removeClass("wap-header").addClass("pc-header");
							$(settings.fullscreen).css("display", "block");
							Bombboxform.removeClass("wap-Bombbox");
					}
						 $.learuntab.pattern(framename);
						 $.learuntab.mobileflag(framename);
				},
				tabNavallwidth:function(framename){
					if(framename == "top") {
						var topWindow = $(window.parent.document);
						var taballwidth = 0,
							$tabNav = topWindow.find(menuform),
							$tabNavWp = topWindow.find("#bk-con-menu"),
							$tabNavitem = topWindow.find(menuform).find(" > li.column-name"),
							$tabNavmore = $(".Hui-tabNav-more");
						if(!$tabNav[0]) {
							 return
						}
						$tabNavitem.each(function(index, element) {
							taballwidth += Number(parseFloat($(this).width()));
						});
						$tabNav.width(taballwidth);
						var w = $tabNavWp.width();
						if(taballwidth > w) {
							$tabNavmore.show();
						} else {
							$tabNavmore.hide();
								$tabNav.css({
								left: 0
							})
						}
				    }	
				},
				//菜单栏溢出左右移动显示
				Menumove:function(framename){
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					var topWindow = $(window.parent.document);
					var num = 0,
					tabNavnext=$('#js-tabNav-next'),
					tabNavprev=$('#js-tabNav-prev');
					var breadcrumb = $(settings.slide).width();
					var menuwidth = menuform.width();
					var tabNavitem = topWindow.find(menuform).find(" > li.column-name");
					if(breadcrumb <= menuwidth) {
						tabNavnext.on('click',function() {
							if(tabNavitem.length > num + 1) {
								num === menuform.find('li.column-name').length - 1 ? num = menuform.find('li.column-name').length - 1 : num++;
								toNavPos();
							} else {
								return false;
							}
						});
					} else {
						return false;
					}
					tabNavprev.on('click',function() {
						num === 0 ? num = 0 : num--;
						toNavPos();
					});
					function toNavPos() {
						var wm=menuform.find('li.column-name').width();
						var w="";
						mobile_flag? w=wm:w=wm;
						if(menuwidth <= (num * w)) {
							return false;
						} else {
							menuform.stop().animate({'left': -num * w}, 100);
						}
					}
				},
				//设置框架样式
				pattern: function(framename) {
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					var menuheight = 0;
					if(framename == "top") {
						menuheight = settings.menuheight;
						$("#tab-column").css({width: $(window).width() - $("#operation").width()});
						footerform.css({
						height: (settings.footerheight),
						bottom: 0,
						width: "100%"
					});
					contentform.css({
						width: $(window).width(),
						height: ($(window).height() - ((settings.hrederheight) + (settings.footerheight))),
						marginTop: settings.hheight,
						position: "relative"
					});
					}
					if(framename == "left" ) {
						
						$("#tab-column").css({
							width: $(window).width() - $("#operation").width() - settings.menuWidth - 1
						});
						Bombboxform.css({
							height: $(window).height() - (settings.hrederheight)
						});
						settings.scrollbar(Bombboxform);
						if($(settings.menuModule).width() == 0) {
							messageform.css({
								width: $(window).width()
							});
						}
						if($(settings.menuModule).width() == settings.sortmodeWidth) {
							messageform.css({
								width: $(window).width() - settings.sortmodeWidth
							})
						} else {
							messageform.css({
								width: $(window).width() - settings.menuWidth
							})
						}
						
					} else {
						messageform.css({
							width: $(window).width(),
							height: ($(window).height() - ((settings.hrederheight) + (settings.footerheight) + menuheight)),
							marginTop: settings.hheight
						});
						Bombboxform.css({
							height: $(window).height() - (settings.hrederheight) - (settings.footerheight)
						});
						settings.scrollbar(Bombboxform);
					}
					if(mobile_flag) {
						$("#tab-column").css({
							width: $(window).width() - $("#operation").width(),
							height: settings.footerheight - 1,
							marginTop: settings.hheight,
							position: "relative"
						});
					}
					
					
				},
				//设置菜单点击
				closemenus: function() {
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					var framename=$(temsye).attr("data-name");
					$(this).hide();
					$('#bk-con-menu .radius').hide();
					if( framename == "left") {
						$(settings.menuModule).stop(true, true).delay(0).animate({
							width: 0,
							position: "relative"
						}, 300);
						messageform.animate({
							left: 0,
							width: $(window).width()
						}, 300);
						$(settings.menuModule).find('.title-menu').hide();
						if(mobile_flag) {
							$('#colselayer').remove();
						}
					} else {
						menuform.stop(true, true).delay(0).animate({
							height: 0,
							position: "relative"
						}, 300).hide(300);
						if(mobile_flag) {
							messageform.css({
								height: ($(window).height() - ((settings.hrederheight) * 2 + (settings.footerheight)))
							});
						} else {
							messageform.css({
								height: ($(window).height() - ((settings.hrederheight) + (settings.footerheight)))
							});
						$(settings.menuModule).css({height:0});
						}
					}
					showmenu.show().css("display", "block");
				},
				//显示对象点击事件
				showmenus: function() {
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					var framename=$(temsye).attr("data-name");
					$(this).hide();
					closemenu.show();
					if( framename == "left") {
						$(settings.menuModule).stop(true, true).delay(0).animate({
							width: settings.menuWidth,
							position: "relative"
						}, 200);
						mobile_flag ? frame.append("<div class='baglayer' id='colselayer'></div>") : '';
						var iconsort = settings.iconsort;
						$(".title-menu").contents().filter(function() {
							return this.nodeType === 3
						}).remove();
						$(".title-menu").prepend(settings.titlename);
						$(settings.slide).removeClass("extendstyle");
						$('#menuclick').attr("ischek", 'false').removeClass('stretch').html(iconsort);
						$(".elasticStyle").css({left: 0});
						menuform.find('.menu-text').show();
						menuform.find(".submenu").removeClass("elasticStyle");
						menuform.find(".arrow").css("right", "");
						menuform.animate({width: settings.menuWidth,}, 200);
						$(settings.menuModule).find('.title-menu').show();
						if(mobile_flag){
							messageform.animate({width: $(window).width()}, 200);
						}else{
							messageform.animate({width: $(window).width() - settings.menuWidth,left: settings.menuWidth}, 200);
						}
					} else {
						if(mobile_flag){
							messageform.css({
								height: ($(window).height() - ((settings.hrederheight) * 2 + (settings.footerheight) + (settings.menuheight) / 2))
							});
							menuform.css({height: (settings.menuheight) / 2}, 300).show(300);
						} else {
							messageform.css({
								height: ($(window).height() - ((settings.hrederheight) + (settings.footerheight) + (settings.menuheight)))
							});
							menuform.css({
								height: (settings.menuheight)
							}, 300).show(300);
							$(settings.menuModule).css({height:settings.menuheight});
						}
					}
					$('#bk-con-menu .radius').show();
				},
				sessionStorage: function() {
					var menu = JSON.parse(window.sessionStorage.getItem("menu"));
					for(var i = 0; i < menu.length; i++) {
						if($("#tab-column li.active span").attr('data-title') == menu[i].title) {
							menu.splice(0, menu.length, menu[i]);
							window.sessionStorage.setItem("menu", JSON.stringify(menu));
						}
					}
				},
				refreshTab: function() {
					var currentId = tabsform.find('.active span').attr('data-href');
					var target = $('.simei_iframe[data-href="' + currentId + '"]');
					var url = target.attr('src');
					target.attr('src', url).load(function() {});
				},
				tabCloseAll: function(e) {
					var num = 0;
					if($(this).hasClass("disabled")) {
						event.preventDefault();
					} else {
						var $tabCloseAll = $("#tab-column .tab-navigation>li");
						if($tabCloseAll.not(".name").length) {
							$($tabCloseAll).not(".home").remove(); //关闭除首页以外的全部栏目			
							$("#tab-column").find('.tab-navigation>li').addClass("active");
							messageform.find(settings.iframe).not(".index_iframe").remove(); //关闭除首页以外的全部栏目
							$(".index_iframe").css("display", "block");
							window.sessionStorage.removeItem("menu");
							menu = [];
							var sessionArr = [];
							num == 0 ? num = 0 : num--;
						}
						$tabCloseAll.length > 1?$(this).addClass("disabled"):$(this).removeClass("disabled");
					}
				},
				tabCloseOther: function() {
					var num = 0;
					$("#tab-column").find('.tab-navigation>li').not(".active,.home").remove(); //关闭打开栏目以外的栏目（除首页）	  
					messageform.find(settings.iframe).not(".selected,.index_iframe").remove();
					messageform.find(settings.iframe).not(".index_iframe").css("display", "block");
					$.learuntab.sessionStorage();
					var sessionArr = [];
					num === 0 ? num == 0 : num--;	
				},
				tabCloseCurrent: function() {
					var num = 0;
					var quantity = $("#tab-column").find(' .tab-navigation li');
					for(var i = 0; i < quantity.length; i++) {
						if(i!=0){
						    if(quantity.eq(i).hasClass("active")){
							    quantity[i].remove();
							    $(settings.iframe).eq(i).remove();
							    Storage(i);
						    }
						}
					}
					tabsform.find("li:last").addClass("active");  
					messageform.find(settings.iframe+":last").css("display", "block");
					function Storage(i) {
						var id = i - 1;
						var menu = JSON.parse(window.sessionStorage.getItem("menu"));
						for(var s = 0; s < menu.length; s++) {
							if($(id).find('span').attr('data-title') == menu[s].titleName) {
								menu.splice(id, 1);
								window.sessionStorage.setItem("menu", JSON.stringify(menu));
							}
						}
					}
					var sessionArr = [];
					num === 0 ? num == 0 : num--;	
				},
				getTabList: function() {
					var tablist = [];
					frame.find('.tab-navigation  li span').each(function() {
						tablist.push({
							id: $(this).attr('data-href'),
							title: $(this).attr('data-title')
						});
					});
					return tablist;
				},
				onTabList: function() {
					var tablist = $.learuntab.getTabList();
					var html = [];
					$.each(tablist, function(key, val) {
						if(key === 0) {
							html.push('<li class="toggle-tab" data-href="' + val.id + '"><span>' + val.title + '</span></li>');
						} else {
							html.push('<li class="toggle-tab" data-href="' + val.id + '"><span>' + val.title + '</span><em class="' + settings.iconfont + '">' + settings.iconDelete + '</em></li>');
						}
					});
					if(tablist.length >= 4) {
						settings.scrollbar(frame.find(".tab-list").html(html.join('')).css({
							height: 300
						}));
					} else {
						frame.find(".tab-list").html(html.join(''));
					}
					$.learuntab.onTabListToggle();
					$.learuntab.tabDeleteColumn();
				},
				onTabListToggle: function() {
					$('.toggle-tab').on('click', function() {
						var name = $("#tab-navigation li");
						var bStopIndex = $(this).index();
						name.removeClass("active").eq(bStopIndex).addClass("active");
						messageform.find(settings.iframe).hide().removeClass("selected").eq(bStopIndex).show().addClass("selected");
					});
				},
				tabDeleteColumn: function() {
					$('.toggle-tab em').on('click', function() {
						var num = 0;
						var name = $("#tab-navigation li");
						var aCloseIndex = $(this).parents("li").index();
						$(this).parent().remove();
						name.removeClass("active").eq(aCloseIndex).remove().addClass("active");
						messageform.find(settings.iframe).eq(aCloseIndex).remove();
						var id = aCloseIndex - 1;
						messageform.find(settings.iframe).eq(id).css("display", "block");
						var tabli = name.removeClass("active").eq(aCloseIndex);
						var menu = JSON.parse(window.sessionStorage.getItem("menu"));
						for(var s = 0; s < menu.length; s++) {
							if(tabli.find('span').attr('data-title') == menu[s].title) {
								menu.splice(id, 1);
								window.sessionStorage.setItem("menu", JSON.stringify(menu));
							}
						}
						num === 0 ? num == 0 : num--;
						aCloseIndex > 1 ? $(".tabCloseAll").removeClass("disabled") : $(".tabCloseAll").addClass("disabled");
						
					});
				},
				prompt: function() {
					var box = $(this).parents(".Promp_plate");
					box.stop(true, true).delay(0).remove(); //提示框
				},
				clickBombbox: function(e) {
					var framename=temsye.attr("data-name");
					var target = $(e.target);
					var p = target.attr('data-position'); //方向
					var t = target.attr('data-type'); //箭头
					var name = $('.clickBombbox');
					var x = target.attr('data-id'); //
					var $btn = $(".memu_title");
					if(target.is('.clickBombbox') || target.parent().is('.clickBombbox')) {
						if(target.hasClass("selected")) {
							target.next('.Bombbox').css({
								"right": -(settings.prompt_width),
								"display": "none",
							}).animate({
								"top": settings.hrederheight
							}, 500);
							target.removeClass("selected");
							target.parent().removeClass("selected");
							(t === 'arrow') ? target.find('i').attr('class', settings.iconfont + ' arrow').html(settings.mouIconClose): target.find('i');
						} else {
							target.next('.Bombbox').css({
								display: "block",
							}).animate({
								"top": settings.hrederheight
							}, 500).parent().siblings().children('.Bombbox').hide();
							target.addClass("selected").parent().siblings().children('a').removeClass('selected').find('i').attr('class', settings.iconfont).html(settings.mouIconClose);
							(t === 'arrow') ? target.find('i').attr('class', settings.iconfont + ' arrow').html(settings.mouIconOpen): target.find('i');
							$btn.eq(x).addClass('selected').siblings().removeClass('selected');
							$('.mode').attr('data-id', x);
							$('.tab-box > div ').eq(x).show().siblings().hide();
							var index = $btn.index(target);
							if(index >= 0) {
								$btn.eq(index).addClass("selected");
								$('.tab-box > div ').eq(index).show().siblings().hide();
							}
							x?$(".amount").remove():'';
							
						}
					} else if(!target.is('.clickBombbox') || !target.parent().is('.clickBombbox')) {
						$('.Bombbox').animate({}, 500).css({
							"display": "",
							"top": ""
						});
						name.find('i.arrow').attr('class', settings.iconfont).html(settings.mouIconClose);
						name.removeClass("selected");
						name.find('a').removeClass("selected");
					}
					if(target.is(settings.prompt_btn)) {
						if($(settings.prompt_btn).hasClass("promptclose")) {
							$('.dropdown-menu').css({
								"right": -(settings.prompt_width),
								"display": "none",
								top: $('.dropdown-menu').outerHeight(),
							}).animate({
								"right": -($('.dropdown-menu').outerWidth()),}, 500);
							$(settings.prompt_btn).attr('class', 'dropdown ' + settings.iconfont).html(settings.mouIconClose).removeClass("promptclose");
						} else {
							if(footerform.hasClass('footer-top')) {
								$('.dropdown-menu').css({
									display: "block",
									top: (settings.footerheight),
									margin: 0,
									borderBottom: "1px solid #cccccc"
								}).animate({"right": 0,}, 500);

							} else {
								$('.dropdown-menu').css({
									display: "block",
									top: -($('.dropdown-menu').outerHeight() + 2),
								}).animate({"right": 0,}, 500);

							}
							$(settings.prompt_btn).attr('class', 'dropdown ' + settings.iconfont).html(settings.mouIconOpen).addClass("promptclose");
						}
					}
					if(!target.is(settings.prompt_btn)) {
						$('.dropdown-menu').animate({
							"right": -($('.dropdown-menu').outerWidth()),
						}, 500).css({"display": "none"});
						$(settings.prompt_btn).attr('class', 'dropdown  ' + settings.iconfont).html(settings.mouIconClose).removeClass("promptclose");
					}
					if(target.is('.baglayer')) {
						showmenu.show();
						target.remove();
						$('#bk-con-menu .radius').hide();
						if(framename== "left") {
							$(settings.menuModule).stop(true, true).delay(0).animate({
								width: 0,
								position: "relative"}, 300);
							messageform.animate({
								left: 0,
								width: $(window).width()}, 300);
							$(settings.menuModule).find('.title-menu').hide();
						}
					}
				},
				//设置菜单栏显示还是隐藏
				minStatues:function(framename){
					if(settings.minStatue) {
						$.learuntab.framemode(framename);
						$(settings.menuModule).css({width: 0});
						showmenu.show().css("display", "block");
						messageform.animate({left: 0,width: $(window).width()}, 300);
						$(settings.menuModule).find('.title-menu').hide();
					}
					else{
						$.learuntab.framemode(framename);	
					}
				},
				init: function() {
					$.learuntab.bindEvents();
					var mobile_flag = $.learuntab.isMobile(mobile_flag);
					var li = $('.skin-section').find('li a');
					$('.tabCloseAll').on('click', $.learuntab.tabCloseAll); //除首页以外全部关闭
					$('.tabCloseOther').on('click', $.learuntab.tabCloseOther);
					$('.tabCloseCurrent').on('click', $.learuntab.tabCloseCurrent);
					$('.tabReload').on('click', $.learuntab.refreshTab);
					frame.on('click', $.learuntab.clickBombbox); //点击弹出层
					Prompname.on('click', $.learuntab.prompt); //提示框关闭
                    closemenu.on('click', $.learuntab.closemenus); //隐藏对象点击事件
					showmenu.on('click', $.learuntab.showmenus); //显示对象点击事件	
					$(settings.prompt_btn).on('click', $.learuntab.onTabList);
					$(settings.fullscreen).on('click', function() {
						if(!$(this).attr('data-fullscreen')) {
							$(this).attr('data-fullscreen', 'true').find("i").attr('class', '' + settings.iconfont).html(settings.iconShrink);
							$.learuntab.requestFullScreen();
						} else {
							$(this).removeAttr('data-fullscreen').find("i").attr('class', '' + settings.iconfont).html(settings.iconExpandicon);
							$.learuntab.exitFullscreen();
						}
					});
					//换肤确认
					li.on('click', function() {
						$(this).addClass("selected " + settings.iconfont + " " + settings.iconSkin + "").parent().siblings().children('a').removeClass("selected " + settings.iconfont + " " + settings.iconSkin + "");
						var v = $(this).attr("data-val");
						$.learuntab.skinchange(v);
					});
					var cookie_skin = $.cookie("MyCssSkin");
					if(cookie_skin) {
						var li = $('.skin-section').find('li');
						$.learuntab.skinchange(cookie_skin);
						for(var i = 0; i < li.length; i++) {
							var v = li.eq(i).find('a').attr("data-val");
							if(v == cookie_skin) {
								li.eq(i).find('a').addClass("selected " + settings.iconfont + " " + settings.iconSkin + "").parent().siblings().children('a').removeClass("selected " + settings.iconfont + " " + settings.iconSkin + "");
							}
						}
					};
				},
				bindEvents:function(){
					$('.templatename').on('click', function(e){
						var structureName = $(this).attr("data-style");
						$(this).addClass("selectmobe").siblings('a').removeClass("selectmobe");
						$.learuntab.fstructure(structureName);
						$.learuntab.framemode(structureName);
					}); 
					//设置框架结构模式
					var cookie_sframe = $.cookie("My");
					if(cookie_sframe){
						var name = $('.templatelist').find('a');
						for(var i=0;i<name.length;i++){
						var styles=name.eq(i).attr("data-style");
						styles==cookie_sframe ? name.eq(i).addClass("selectmobe"):name.eq(i).removeClass("selectmobe");
						}
						$.learuntab.fstructure(cookie_sframe);	
						$.learuntab.tabNavallwidth(cookie_sframe);
					    $.learuntab.pattern(cookie_sframe);
					    $.learuntab.minStatues(cookie_sframe);
					    $(window).resize(function(){					
						    $.learuntab.framemode(cookie_sframe);
					    });
					}else{
						var framename=temsye.attr("data-name");
						$.learuntab.tabNavallwidth(framename);
					    $.learuntab.pattern(framename);
					    $.learuntab.fstructure(framename);
						$.learuntab.framemode(framename);
						$.learuntab.minStatues(framename);
						$(window).resize(function(){
						    $.learuntab.framemode(framename);
						});
					}
				}				
			};		
			//右键功能模块		
			if(settings.rightclick === true) {
				$(document).on('contextmenu', thisBox, function(e) {
					e.preventDefault();
					e.stopPropagation();
				}); //禁用右键功能
			} else if(settings.rightclick === false) {}
	       //鼠标移动显示按钮操作
			$(settings.menuModule).hover(function() {
				var menuwidth = $(settings.slide).width();
				if(menuwidth >= settings.menuWidth) {
					$(this).addClass("display_btn");
					closemenu.css("display", "block");
				} else {
					$(this).removeClass("display_btn");
					closemenu.css("display", "none");
				}
			}, function() {
				$(this).removeClass("display_btn");
				closemenu.css("display", "none");
			});
			//设置时间显示
			function currentTime() {
				var weekday = new Array(7);
				weekday[0] = "星期一";
				weekday[1] = "星期二";
				weekday[2] = "星期三";
				weekday[3] = "星期四";
				weekday[4] = "星期五";
				weekday[5] = "星期六";
				weekday[6] = "星期日";
				var d = new Date(),
				    h=d.getHours(),
				    m=d.getMinutes(),
				    s=d.getSeconds(),
				    mot=d.getMonth() + 1,
				    day=d.getDate(),
				    hours="",minutes="",seconds="", month="", daydata="";
				h>=10? hours=h : hours='0'+h;
				m>=10? minutes=m :minutes='0'+m;
				s>=10? seconds=s:seconds='0'+s;
				mot>=10? month=mot:month='0'+mot;
				day>=10? daydata=day:daydata='0'+day;
				var str = '';
				    str += '<div class="time_display clearfix">';
					str += '<span class="h com-time ">' + hours+ '</span>' + '<b>：</b>';
					str += '<span class="m com-time ">' + minutes + '</span>' + '<b>：</b>';
					str += '<span class="s com-time ">' + seconds + '</span></div>';
					str += '<div class="date_display clearfix"><span class="years">' + d.getFullYear() + '年' + '</span>';
					str += '<span class="mot">' + month + '月' + '</span>';
					str += '<span class="day">' + daydata + '</span>';
					str += '<span>' + weekday[d.getDay() - 1] + '</span></div>';
					return str;
			}
			setInterval(function() {$(".timehtml").html(currentTime());}, 1000);
			//获取栏目设置
			var template = '<div class="tab-list-container"><div class="title-name">选项卡栏目</div><ul class="tab-list"></ul></div>';
			$('.inte-operation').find('.tab-menu-list').append(template);
			//初始化事件
			var fileData = (settings.data); //获取数据
			var JSONTreeArr = proJSON(fileData, 0);
			menuform.html = menuHtml(JSONTreeArr);
			//生成树形JSON数据
			function proJSON(oldArr, pid) {
				var newArr = [];
				if(oldArr === null) {
					return false;
				} else {
					oldArr.map(function(item) {
						if(item.pid === pid) {
							var obj = {
								id: item.id,
								name: item.name,
								icon: item.icon,
								url: item.url
							};
							var child = proJSON(oldArr, item.id);
							if(child.length > 0) {
								obj.child = child;
							}
							newArr.push(obj);
						}
					});
				}
				return newArr;
			}
			var munlist = menuform.html;
			var template = munlist;

			function menuHtml(data) {
				var ulHtml = "";
				if(data === false) {
					return false;
				} else {
					data.map(function(item) {

						var lihtml = "<li class='column-name' ischek='false'>";
						var a = '';
						item.url == '' ? a = 0 : a = item.url.charCodeAt(2);
						if(item.child && item.child.length > 0) {
							lihtml += "<a href='javascript:void(0)'  class='menuUl_title' data-uid='" + a + "' name='" + item.url + "' title='" + item.name + "'><i ischek='false' class='menuicon " + item.icon + "'></i>" + "<span data-id='" + item.id + "'class='menu-text'>" + item.name + "</span><i ischek='false' class='arrow " + settings.iconfont + "'>" + settings.mouIconOpen + "</i></a>";
							var _ul = menuHtml(item.child);
							lihtml += "<ul class='submenu clearfix'>" + _ul + "</ul></li>";
						} else {
							lihtml += "<a href='javascript:void(0)' ischek='false' name='" + item.url + "' data-uid='" + a + "'  title='" + item.name + "' class='name menuUl_title' data-id='" + item.id + "'><i ischek='false' class='menuicon " + item.icon + "'></i>" + "<span data-id='" + item.id + "' class='menu-text'>" + item.name + "</span></a>";
						}
						ulHtml += lihtml;
					});
				}
				ulHtml += "";
				return ulHtml;
			}

			function menusess(name) {
				var menu = JSON.parse(window.sessionStorage.getItem("menu"));
				if(menu != null) {
					for(var i = 0; i < menu.length; i++) {
						if($(name).attr('title') == menu[i].title) {
							menu[i].nate == 2 ? menu[i].nate = 2 : menu[i].nate = 1;
						} else {
							menu[i].nate = 0;
						}
						window.sessionStorage.setItem("menu", JSON.stringify(menu));
					}
				}
				var uid = $(name).attr('data-id');
				iframeinfo(name, uid); //显示设置内容iframe
			}
			//设置菜单栏事件
			function bindEvent() {
				var column = frame.find(menuform).append(template);
				if(settings.menuopt === 'click') {
					frame.find(".menulist  li").on("click", function() {
						$(this).find('a').on("click", function() {
							menusess(this);
						});
						Emode(this); //菜单栏样式
						return false;
					});
				}
				if(settings.menuopt === 'hover') {
					frame.find(".menulist  li").hover(function() {
						$(this).find('a').on("click", function() {
							menusess(this);
						});
						Emode(this);
						return false;
					});
				}
				//当默认菜单栏打开时执行该事件
				frame.find(".menulist  li a").on("click", function() {
					menusess(this);
					Emode(this);
				});
				return column;
			}
			bindEvent();

			function Emode(name) {
				var labeul = $(name).parent("ul");
				var parents = $(name).parent();
				var ischek = $(name).attr("ischek");
				var menuUl = $(name).find(".submenu");
				var munlist = labeul.children("li:visible");
				if(ischek === 'false') {
					munlist.each(function(i) {
						var x = $(name).index();
						if(i == x) {
							munlist.eq(i).addClass("open").slideDown(300).attr("ischek", 'true');
							munlist.eq(i).find(".arrow").html('').html(settings.mouIconClose);
							munlist.eq(i).children(".submenu").slideDown(300);
							munlist.eq(i).find("i").attr("ischek", 'true');
						} else {
							munlist.eq(i).removeClass("open").attr("ischek", 'false');
							munlist.eq(i).children(".submenu").slideUp(300);
							munlist.eq(i).find(".arrow").html('').html(settings.mouIconOpen);
							munlist.eq(i).find("i").attr("ischek", 'false');
						}
					});
				} else if(ischek === 'true') {
					var x = $(name).index();
					var urlname = $(name).eq(x).find('a').attr('name');
					if(urlname == null || urlname == "") {
						$(name).removeClass("open");
						$(name).find(".arrow").html('').html(settings.mouIconOpen);
						menuUl.slideUp(300);
						$(name).attr("ischek", 'false');
						$(name).find('i').attr("ischek", 'false');
					} else {
						return false
					}
				}
			}
			//设置默认显示页
			var slidelist = $(settings.slide).find("li");
			slidelist.each(function(i) {
				var home = $(this).eq(i).find('a').attr('name');
				if(settings.homepage == home) {
					var htmls = $(this).eq(i).html();
					iframeinfo(htmls, 2);
				}
			});

			function iframeinfo(name, uid) {
				if($(name).attr('name')) {
					var bStop = false;
					var bStopIndex = 0;
					var _href = $(name).attr('name');
					var _titleName = $(name).text();
					var _pageid = $(name).attr('data-id');
					var topWindow = $(window.parent.document);
					var show_navLi = topWindow.find("#tab-navigation li");
					show_navLi.each(function() {
						if($(this).find('span').attr("data-id") === uid) {
							bStop = true;
							bStopIndex = show_navLi.index($(this));
							return false;
						}
					});
					if(!bStop) {
						show_navLi.removeClass("active");
						if(uid == 2) {
							creatIframe(_href, _titleName, uid, _pageid);
						} else {
							creatIframe(_href, _titleName, bStopIndex, _pageid);
						}
						min_titleList();
					} else {
						show_navLi.removeClass("active").eq(bStopIndex).addClass("active");
						var iframe_box = topWindow.find(messageform);
						iframe_box.find(settings.iframe).hide().eq(bStopIndex).show().find("iframe").attr({
							"src": _href,
							"data-href": _href
						});
						if(uid == 2) {
							sessionname(_href, _titleName, uid, _pageid);
						}
					}
				}
			}

			function min_titleList() {
				var topWindow = $(window.parent.document);
				var show_nav = topWindow.find(tabsform, "#dropdown_menu");
				var aLi = show_nav.find("li");
				var iframe_box = topWindow.find(messageform);
			}
			$(settings.tabs).find('li').css({
				height: settings.footerheight,
				lineHeight: settings.footerheight + "px"
			});

			function creatIframe(href, titleName, id, pageid) {
				var topWindow = $(window.parent.document);
				var show_nav = topWindow.find(settings.tabs);
				var id_name = $(settings.iframe).each(function(i) {
					$(this).attr('id', "Sort_link_" + i);
					if(i == 0) {
						$(this).addClass("index_iframe");
					}
				});
				var url = href;
				var iframe_box = topWindow.find(settings.message);
				show_nav.find('li').removeClass("active");
				if(id == 2) {
					show_nav.append('<li class="home active"><span class="' + settings.iconfont + '" data-id="' + pageid + '" data-title="' + titleName + '" data-href="' + url + '">' + settings.iconHome + '</span></li>');
				} else {
					show_nav.append('<li class="name active"><span data-id="' + pageid + '" data-title="' + titleName + '" data-href="' + url + '">' + titleName + '</span><em class="close_icon"></em></li>');
				}
				show_nav.find('li').css({
					height: settings.footerheight,
					lineHeight: settings.footerheight + "px"
				});
				var iframeBox = iframe_box.find('.show_iframe:visible');
				iframeBox.hide().removeClass("selected");
				iframe_box.append('<div class="show_iframe selected" id="' + id_name + '" data-id="' + pageid + '" date-name="' + titleName + '" data-href="' + url + '"><div class="loading">' + loadinghtml + '</div><iframe class="simei_iframe" frameborder="0" src="' + href + '" data-href="' + url + '"></iframe></div>');
				var showBox = iframe_box.find('.show_iframe:visible');
				var loadinghtml = showBox.find('.loading').html(settings.loading);
				showBox.find('iframe').attr("src", href).load(function() {
					showBox.find('.loading').hide();
				})
				if(id != 2) {
					sessionname(href, titleName, id, pageid);
				} else {
					iframe_box.addClass('index_iframe');
				}
				var ftabNavitem = show_nav.find(">li");
				ftabNavitem.length > 1 ? $(".tabCloseAll").removeClass("disabled") : $(".tabCloseAll").addClass("disabled");
			}
			//记录保存
			var sessionArr = [];

			function sessionname(href, titleName, id, pageid) {
				var obj = {
					title: titleName,
					url: href,
					nate: id,
					id: pageid
				};
				sessionArr.push(obj);
				var jsonStr = JSON.stringify(sessionArr);
				window.sessionStorage.setItem("menu", jsonStr);
			}
			var num = 0;
			$(document).on("click", "#tab-navigation li", function() {
				var bStopIndex = $(this).index();
				if(bStopIndex == 0) {
					var menu = JSON.parse(window.sessionStorage.getItem("menu"));
					var titleName = $(this).find('span').attr("data-title");
					var href = $(this).find('span').attr("data-href");
					var pageid = $(this).find('span').attr("data-id");
					var id = 2;
					sessionname(href, titleName, id, pageid);
					for(var i = 0; i < menu.length; i++) {
						var name = menu[i].title;
						if(name == titleName) {
							menu.splice(i + 1, 1);
							window.sessionStorage.setItem("menu", JSON.stringify(menu));
						}
					}
				}
				$("#tab-navigation li").removeClass("active").eq(bStopIndex).addClass("active");
				messageform.find(settings.iframe).hide().removeClass("selected").eq(bStopIndex).show().addClass("selected");
				var menu = JSON.parse(window.sessionStorage.getItem("menu"));
				var munlist = $("#menuList > .column-name").next();
				if(menu != null) {
					for(var i = 0; i < menu.length; i++) {
						if($(this).find('span').attr('data-title') == menu[i].title) {
							for(var t = 0; t < munlist.length; t++) {
								var ischek = munlist.eq(t).attr("ischek");
								var submenu = munlist.eq(t).find(".submenu li");
								for(var e = 0; e < submenu.length; e++) {
									var grade = submenu.eq(e).find('a').attr('title');
									if(grade == menu[i].title) {
										submenu.eq(e).addClass("open").slideDown(300).attr("ischek", 'true');
										submenu.eq(e).find('.submenu').slideDown(300);
										submenu.eq(e).find("i").attr("ischek", 'true');
										munlist.eq(t).find('.submenu').slideDown(300);
										munlist.eq(t).addClass("open").slideDown(300).attr("ischek", 'true');
										munlist.eq(t).find("i").attr("ischek", 'true');
										munlist.eq(t).find('.arrow').html(settings.mouIconClose);
										munlist.eq(t).siblings().removeClass("open").find('.submenu').slideUp(300).attr("ischek", 'false');
									} else {
										submenu.eq(e).removeClass("open").attr("ischek", 'false');
										submenu.eq(e).find('.submenu').slideUp(300);
										submenu.eq(e).find("i").attr("ischek", 'false');
										munlist.eq(t).find("i").attr("ischek", 'false');
										munlist.eq(t).find('.arrow').html(settings.mouIconOpen);
									}
								}
							}
							menu[i].nate == 2 ? menu[i].nate = 2 : menu[i].nate = 1;
						} else {
							menu[i].nate = 0;
						}
						window.sessionStorage.setItem("menu", JSON.stringify(menu));
					}
				}
			});
			$(document).on("click", "#tab-navigation li em.close_icon", function(e) {
				var aCloseIndex = $(this).parents("li").index();
				$(this).parent().remove();
				messageform.find(settings.iframe).eq(aCloseIndex).remove();
				num === 0 ? num == 0 : num--;
				var menu = JSON.parse(window.sessionStorage.getItem("menu"));
				var munlist = $("#menuList").find("li");
				if(menu != null) {
					for(var i = 0; i < menu.length; i++) {
						if($(this).parents("li").find('span').attr('data-title') == menu[i].title) {
							munlist.each(function(e) {
								var title = munlist.find('a').eq(e).attr('title');
								if(title == menu[i].title) {
									munlist.eq(e).removeClass("open").attr("ischek", 'false');
									munlist.eq(e).find("i").attr("ischek", 'false');
								}
							});
							menu.splice(i, 1);
							if(menu != 0) {
								menu.length == 1 ? menu[0].nate = 1 : menu[i - 1].nate = 1;
							}
						}
						window.sessionStorage.setItem("menu", JSON.stringify(menu));
					}
					sessionArr = [];
				}
				aCloseIndex > 1 ? $(".tabCloseAll").removeClass("disabled") : $(".tabCloseAll").addClass("disabled");
			});
			$(document).on("dblclick", "#tab-navigation li", function(e) {
				var menu = JSON.parse(window.sessionStorage.getItem("menu"));
				var aCloseIndex = $(this).index();
				if(aCloseIndex > 0) {
					$(this).remove();
					$(tabsform).find(settings.iframe).eq(aCloseIndex).remove();
					num === 0 ? num == 0 : num--;
					$("#tab-navigation li").removeClass("active").eq(aCloseIndex - 1).addClass("active");
					messageform.find(settings.iframe).hide().eq(aCloseIndex - 1).show();
					for(var i = 0; i < menu.length; i++) {
						if($(this).find('span').attr('data-title') == menu[i].title) {
							menu.splice(i, 1);
							window.sessionStorage.setItem("menu", JSON.stringify(menu));
						}
					}
					sessionArr = [];
				} else {
					return false;
				}
				aCloseIndex > 1 ? $(".tabCloseAll").removeClass("disabled") : $(".tabCloseAll").addClass("disabled");
			});
			//刷新后还原打开的窗口
			if(window.sessionStorage.getItem("menu") != null) {
				var menu = JSON.parse(window.sessionStorage.getItem("menu"));
				for(var i = 0; i < menu.length; i++) {
					var href = menu[i].url;
					var titleName = menu[i].title;
					var id = menu[i].nate;
					var pageid = menu[i].id;
					id != 2 ? creatIframe(href, titleName, id, pageid) : '';

				}
				var munlist = $("#menuList").find("li");
				$.each(menu, function(i) {
					var id = menu[i].nate;
					var name = menu[i].title;
					if(id == 1) {
						$("#tab-navigation li.name").removeClass("active").eq(i).addClass("active");
						messageform.find(settings.iframe).not('.index_iframe').hide().removeClass("selected").eq(i).show().addClass("selected");
						munlist.each(function(e) {
							var title = munlist.find('a').eq(e).attr('title')
							if(title == name) {
								munlist.eq(e).addClass("open").slideDown(300).attr("ischek", 'true');
								munlist.eq(e).find(".submenu").slideDown(300);
								munlist.eq(e).find("i").attr("ischek", 'true');
								munlist.eq(e).parent().parent().addClass("open").slideDown(300).attr("ischek", 'true');
								munlist.eq(e).parent().parent().find("i").attr("ischek", 'true');
								munlist.eq(e).parent().slideDown(300);
								munlist.eq(e).parent().parent().find('.arrow').html(settings.mouIconClose);
							}
						})
					}
					if(id == 2) {
						$("#tab-navigation li").removeClass("active").eq(0).addClass("active");
						messageform.find(settings.iframe).hide().removeClass("selected").eq(0).show().addClass("selected");
						munlist.eq(0).addClass("open").slideDown(300).attr("ischek", 'true');
					}
				});
			}			
			//底部栏目溢出滚动
			var footermenu = $("#tab-column").width();
			var fmenuwidth = tabsform.width();
			var ftopWindow = $(window.parent.document);
			var ftabNavitem = ftopWindow.find("#tab-navigation > li");
			if(footermenu <= fmenuwidth) {
				tabsform.addClass("mobilemun");
				settings.scrollbar(tabsform);
			}
			//初始化加载方法
			$(function() {
				$.learuntab.init();
			});
		});
	};
})(jQuery);