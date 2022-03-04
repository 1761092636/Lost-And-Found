$(document).ready(function() {
	//时间转换,template插件扩展功能
	template.defaults.imports.timestamp = function(date, format) {
		if(typeof date === "string") {
			var mts = date.match(/(\/Date(\d+)\/)/);
			if(mts && mts.length >= 3) {
				date = parseInt(mts[2]);
			}
		}
		date = new Date(parseInt(date * 1000));
		if(!date || date.toUTCString() == "Invalid Date") {
			return "";
		}
		var map = {
			"M": date.getMonth() + 1, //月份
			"d": date.getDate(), //日
			"h": date.getHours(), //小时
			"m": date.getMinutes(), //分
			"s": date.getSeconds(), //秒
			"q": Math.floor((date.getMonth() + 3) / 3), //季度
			"S": date.getMilliseconds() //毫秒
		};
		format = format.replace(/([yMdhmsqS])+/g, function(all, t) {
			var v = map[t];
			if(v !== undefined) {
				if(all.length > 1) {
					v = '0' + v;
					v = v.substr(v.length - 2);
				}
				return v;
			} else if(t === 'y') {
				return(date.getFullYear() + '').substr(4 - all.length);
			}
			return all;
		});
		return format;
	};
	//分页插件定义构造函数
	function Paging(el, options) {
		this.el = el;
		this.options = {
			pageNo: options.initPageNo || 1, // 初始页码
			totalPages: options.totalPages || 1, //总页数
			totalCount: options.totalCount || '', // 条目总数
			slideSpeed: options.slideSpeed || 0, // 缓动速度
			jump: options.jump || false, // 支持跳转
			callback: options.callback || function() {} // 回调函数
		};
		this.init();
	}
	// 给实例对象添加公共属性和方法
	Paging.prototype = {
		constructor: Paging,
		init: function() {
			this.createDom();
			this.bindEvents();
		},
		createDom: function() {
			var that = this,
				ulDom = '',
				jumpDom = '',
				content = '',
				liWidth = 50, // li的宽度
				totalPages = that.options.totalPages, // 总页数
				wrapLength = 0;
			totalPages > 5 ? wrapLength = 5 * liWidth : wrapLength = totalPages * liWidth;
			for(var i = 1; i <= that.options.totalPages; i++) {
				i != 1 ? ulDom += '<li>' + i + '</li>' : ulDom += '<li class="sel-page">' + i + '</li>';
			}
			that.options.jump ? jumpDom = '<input type="text" placeholder="1" class="jump-text" id="jumpText"><button type="button" class="jump-button" id="jumpBtn">跳转</button>' : jumpDom = '';
			content = '<button type="button" id="firstPage" class="turnPage first-page">首页</button>' +
				'<button class="turnPage" id="prePage">上一页</button>' +
				'<div class="pageWrap" style="width:' + wrapLength + 'px">' +
				'<ul id="pageSelect" style="transition:all ' + that.options.slideSpeed + 'ms">' +
				ulDom +
				'</ul></div>' +
				'<button class="turnPage" id="nextPage">下一页</button>' +
				'<button type="button" id="lastPage" class="last-page">尾页</button>' +
				jumpDom +
				'<p class="total-pages">共&nbsp;' +
				that.options.totalPages +
				'&nbsp;页</p>' +
				'<p class="total-count">' +
				that.options.totalCount +
				'</p>';
			that.el.html(content);
		},
		bindEvents: function() {
			var that = this,
				pageSelect = $('#pageSelect'), // ul
				lis = pageSelect.children(), // li的集合
				liWidth = lis[0].offsetWidth, // li的宽度
				totalPages = that.options.totalPages, // 总页数
				pageIndex = that.options.pageNo, // 当前选择的页码
				distance = 0, // ul移动距离
				prePage = $('#prePage'),
				nextPage = $('#nextPage'),
				firstPage = $('#firstPage'),
				lastPage = $('#lastPage'),
				jumpBtn = $('#jumpBtn'),
				jumpText = $('#jumpText');

			prePage.on('click', function() {
				pageIndex--;
				if(pageIndex < 1) pageIndex = 1;
				handles(pageIndex);
			})

			nextPage.on('click', function() {
				pageIndex++;
				if(pageIndex > lis.length) pageIndex = lis.length;
				handles(pageIndex);
			})

			firstPage.on('click', function() {
				pageIndex = 1;
				handles(pageIndex);
			})

			lastPage.on('click', function() {
				pageIndex = totalPages;
				handles(pageIndex);
			})

			jumpBtn.on('click', function() {
				var jumpNum = parseInt(jumpText.val().replace(/\D/g, ''));
				if(jumpNum && jumpNum >= 1 && jumpNum <= totalPages) {
					pageIndex = jumpNum;
					handles(pageIndex);
					jumpText.val(jumpNum);
				}
			})

			lis.on('click', function() {
				pageIndex = $(this).index() + 1;
				handles(pageIndex);
			})

			function handles(pageIndex) {
				lis.removeClass('sel-page').eq(pageIndex - 1).addClass('sel-page');
				if(totalPages <= 5) {
					that.options.callback(pageIndex);
					return false;
				}
				if(pageIndex >= 3 && pageIndex <= totalPages - 2) distance = (pageIndex - 3) * liWidth;
				if(pageIndex == 2 || pageIndex == 1) distance = 0;
				if(pageIndex > totalPages - 2) distance = (totalPages - 5) * liWidth;
				pageSelect.css('transform', 'translateX(' + (-distance) + 'px)');
				pageIndex == 1 ? firstPage.attr('disabled', true) : firstPage.attr('disabled', false);
				pageIndex == 1 ? prePage.attr('disabled', true) : prePage.attr('disabled', false);
				pageIndex == totalPages ? lastPage.attr('disabled', true) : lastPage.attr('disabled', false);
				pageIndex == totalPages ? nextPage.attr('disabled', true) : nextPage.attr('disabled', false);
				that.options.callback(pageIndex);
			}

			handles(that.options.pageNo); // 初始化页码位置
		}
	}
	$.fn.paging = function(options) {
		return new Paging($(this), options);
	}

});
/*************图片上传*************/
function imgUpload(api, type) {
	if(type == undefined) {
		type = 'sy_thumb';
	}
	//点击
	$('.addImg').click(function() {
		$(this).parent().find('.upload_input').click();
	})
	//删除
	$('.delete').click(function() {
		$(this).parent().find('input').val('');
		$(this).parent().find('img.preview').attr("src", "");
		//IE9以下
		$(this).parent().find('img.preview').css("filter", "");
		$(this).hide();
		$(this).parent().find('.addImg').show();
	})
	//选择图片
	$('.upload_input').change(function() {
		//预览
		var pic = $(this).parent().find(".preview");
		//添加按钮
		var addImg = $(this).parent().find(".addImg");
		//删除按钮
		var deleteImg = $(this).parent().find(".delete");

		var picContainer = $(this).parent().find('.preBlock');

		var ext = this.value.substring(this.value.lastIndexOf(".") + 1).toLowerCase();

		// gif在IE浏览器暂时无法显示
		if(ext != 'png' && ext != 'jpg' && ext != 'jpeg') {
			if(ext != '') {
				layer.msg("图片的格式必须为png或者jpg或者jpeg格式！");
			}
			return;
		}
		html5Reader(this, pic, addImg, deleteImg, picContainer, type);
	});
	//H5渲染
	function html5Reader(file, pic, addImg, deleteImg, picContainer, type) {
		var files = file.files[0];
		var reader = new FileReader();
		reader.onload = function(e) {
			picContainer.css({
				"z-index": 1
			});
			pic.attr("src", this.result);
			//图片读取完成的回调函数（必须加上否则数据读入不完整导致出错！）
			pic[0].onload = function() {
				var fd = new FormData();
				fd.append('dosubmit', '1');
				fd.append('userid', '1');
				fd.append('thumb_width', '0');
				fd.append('thumb_height', '0');
				fd.append('watermark_enable', '1');
				fd.append('Filedata', files);
				$.ajax({
					url: api + type, //api接口地址
					type: "POST",
					processData: false,
					contentType: false,
					data: fd,
					success: function(ret) {
						var path = JSON.parse(ret);
						path = path.split(',');
						$(file).parent().find("#img").val(path[1]);
					}
				});
			}
		};
		//将文件已Data URL的形式读入页面
		reader.readAsDataURL(files);
		addImg.hide();
		deleteImg.show();
	}
}

function del_space(elem) {

	var elem_child = elem.childNodes; //得到参数元素的所有子元素

	for(var i = 0; i < elem_child.length; i++) { //遍历子元素
		if(elem_child.nodeName == "#text" && !/\S/.test(elem_child.nodeValue)) {
			elem.removeChild(elem_child)
		}
	}

}
/****************************************设置文本框字数*****************************************************/
function Wordcount(name, number) {
	this.Wcount = this.$(name);
	this.Length(arguments);
}
// 简写document.getElementById方法. 
Wordcount.prototype.$ = function(element) {
	return document.getElementById(element);
}
Wordcount.prototype.Length = function(number) {
	var h = document.createElement("span");
	this.Wcount.parentNode.appendChild(h).className ="word_count";
	this.Wcount.onkeyup = this.GetFunction(this, "keyup", number[1]);
	count(number[1], this.Wcount);
}
function count(number, name) {
	var w = name.parentNode.querySelector(".word_count");
	w.innerHTML = "剩余字数 :<em class='colorred number'>" + number + "</em>字符";
	var d = name.parentNode.querySelector(".number");
	if(name.value.length > number) {
		var b = document.body.querySelector(".box_Bullet");
		if(!b) {
			var box = document.createElement("div");
			document.body.appendChild(box).className = "box_Bullet";
			document.body.appendChild(box).innerHTML = "您输入的字数超过限制!";
			box.style.cssText="margin-left:"+(-box.offsetWidth/2)+"px"+";"+"margin-top:"+(-box.offsetHeight/2)+"px";
			var int = setInterval(function() {clock()}, 1000);
			var num =2;
			function clock() {
				num > 0 ? num-- : clearInterval(int);
				if(num == 0) {
					clearInterval(int);
					document.body.removeChild(box);
					return
				}
			}
		}
		// 超过限制的字数了就将 文本框中的内容按规定的字数 截取
		name.value = name.value.substring(0, number);
		d.innerHTML = 0;
		return false;
	} else {
		var curr = number - name.value.length; //减去 当前输入的	
		d.innerHTML = curr.toString();
		return true;
	}

}
//onkeyup方法
Wordcount.prototype.keyup = function(col) {
	count(col, this.Wcount);
};
Wordcount.prototype.GetFunction = function(variable, method, number) {
	return function() {
		variable[method](number);
	}
}
