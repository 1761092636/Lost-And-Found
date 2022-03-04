/*
 * SYSBOX弹框盒子
 * 799129700@qq.com SYSHUXL
 * Reserved head available commercial use
 * Universal background system interface framework
 */
var SYSBOX = {};
function BOX$(i) {
	return document.getElementById(i)
} //简化document.getElementById方法

SYSBOX.box = function() {
	var p, m, b, ic, iu, iw, ih, ia, ititle,iLock, f = 0;
	return {
		show: function(c, u, w, h, a, title,Lock, t) {
			if(!f) {
				p = document.createElement('div');
				p.id = 'sysbox';
				m = document.createElement('div');
				m.id = 'sysmask';
				b = document.createElement('div');
				b.id = 'syscontent';
				document.body.appendChild(m);
				document.body.appendChild(p);
				dt = document.createElement('div');
				dt.id = 'systitle';
				tm = document.createElement('span');
				tm.className = 'sys-title-name';
				tb = document.createElement('span');
				tb.className = 'sys-oper-box';
				close = document.createElement('a');
				close.className = 'sys-close-oper sys-oper'; //
				magnify = document.createElement('a');
				magnify.className = 'sys-magnify-oper sys-oper';
				shrink = document.createElement('a');
				shrink.className = 'sys-shrink-oper sys-oper';
				p.appendChild(dt).className = 'sys-title';
				dt.appendChild(tm).innerText = title;
				dt.appendChild(tb);
				tb.appendChild(magnify).dataset.name = 'magnify';
				tb.appendChild(shrink).dataset.name = 'shrink';
				tb.appendChild(close).dataset.name = 'close';
				tb.appendChild(shrink).href = 'javascript:SYSBOX.box.size("sysbox",' + w + ',' + h + ',"magnify")';
				tb.appendChild(magnify).href = 'javascript:SYSBOX.box.size("sysbox","","","shrink")';
				tb.appendChild(close).href = 'javascript:SYSBOX.box.hide()';
				p.appendChild(b).className = "sys-content";
				Lock ? m.onclick = SYSBOX.box.hide:'';
				title==""? dt.remove():'';
				dt = dt || p;
				dt.onmousedown = SYSBOX.box.dollybox;
				window.onresize = SYSBOX.box.resize;
				f = 1
			}
			
			var mobile_flag = this.isMobile();
			mobile_flag ? p.className = " mobileStyle" : '';
			if(!a && !u) {
				p.style.width = w ? w + 'px' : 'auto';
				p.style.height = h ? h + 'px' : 'auto';
				p.style.backgroundImage = 'none';
				b.innerHTML = c;
			} else {
				b.style.display = 'none';
				p.style.width = p.style.height = '100px'
			}
			this.mask();
			ic = c;
			iu = u;
			iw = w;
			ih = h;
			ia = a;
			iLock=Lock;
			ititle = title;
			this.alpha(m, 1, 80, 3);
			if(t) {
				setTimeout(function() {
					SYSBOX.box.hide()
				}, 1000 * t)
				magnify.remove();
				p.style.cssText = "margin-left:" + (-p.offsetWidth / 2) + "px" + ";" + "margin-top:" + (-p.offsetHeight / 2) + "px";
			};
		},
		fill: function(c, u, w, h, a, title,Lock) {
			if(u) {
				p.style.backgroundColor = '#F2F2F2';
				var x = window.XMLHttpRequest ? new XMLHttpRequest() : new ActiveXObject('Microsoft.XMLHTTP');
				x.onreadystatechange = function() {
					if(x.readyState == 4 && x.status == 200) {
						SYSBOX.box.psh(x.responseText, w, h, a, title.Lock)
					}
				};
				x.open('GET', c, 1);
				x.send(null)
			} else {
				this.psh(c, w, h, a, title)
			}
		},
		psh: function(c, w, h, a, title,Lock) {
			if(a) {
				if(!w || !h) {
					var x = p.style.width,
						y = p.style.height;
					b.innerHTML = c;
					p.style.width = w ? w + 'px' : '';
					p.style.height = h ? h + 'px' : '';
					b.style.display = '';
					w = parseInt(b.offsetWidth);
					h = parseInt(b.offsetHeight);
					b.style.display = 'none';
					p.style.width = x;
					p.style.height = y;
				} else {
					b.innerHTML = c
				}
				this.size(p, w, h)
			} else {
				p.style.backgroundImage = 'none'
			}
		},
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
		hide: function() {
			SYSBOX.box.alpha(p, -1, 0, 3);
		},
		resize: function() {
			SYSBOX.box.pos();
			SYSBOX.box.mask()
		},
		mask: function() {
			m.style.height = SYSBOX.page.total(1) + 'px';
			m.style.width = '';
			m.style.width = SYSBOX.page.total(0) + 'px';

		},
		pos: function() {
			var t = (SYSBOX.page.height() / 2) - (p.offsetHeight / 2);
			t = t < 10 ? 10 : t;
			p.style.top = (t + SYSBOX.page.top()) + 'px';
			p.style.left = (SYSBOX.page.width() / 2) - (p.offsetWidth / 2) + 'px';

		},
		alpha: function(e, d, a) {
			clearInterval(e.ai);
			if(d == 1) {
				e.style.opacity = 0;
				e.style.filter = 'alpha(opacity=0)';
				e.style.display = 'block';
				this.pos()
			}
			e.ai = setInterval(function() {
				SYSBOX.box.ta(e, a, d)
			}, 20)
		},
		ta: function(e, a, d) {
			var o = Math.round(e.style.opacity * 100);
			if(o == a) {
				clearInterval(e.ai);
				if(d == -1) {
					e.style.display = 'none';
					e == p ? SYSBOX.box.alpha(m, -1, 0, 2) : b.innerHTML = p.style.backgroundImage = ''
				} else {
					e == m ? this.alpha(p, 1, 100) : SYSBOX.box.fill(ic, iu, iw, ih, ia, ititle,iLock)
				}
			} else {
				var n = Math.ceil((o + ((a - o) * .5)));
				n = n == 1 ? 0 : n;
				e.style.opacity = n / 100;
				e.style.filter = 'alpha(opacity=' + n + ')'
			}
		},
		size: function(e, w, h, s) {
			e = typeof e == 'object' ? e : BOX$(e);
			clearInterval(e.si);
			var ow = e.offsetWidth,
				oh = e.offsetHeight,
				wo = ow - parseInt(e.style.width),
				ho = oh - parseInt(e.style.height);
			var wd = ow - wo > w ? 0 : 1,
				hd = (oh - ho > h) ? 0 : 1;
			SYSBOX.box.ts(e, w, wo, wd, h, ho, hd);
			if(s == "shrink") {
				magnify.style.display = "none";
				shrink.style.display = "block"

			} else if(s == "magnify") {
				shrink.style.display = "none";
				magnify.style.display = "block";
			}
		},
		ts: function(e, w, wo, wd, h, ho, hd) {
			var ow = e.offsetWidth - wo,
				oh = e.offsetHeight - ho;
			if(w != "" && h != "") {
				if(ow == w && oh == h) {
					clearInterval(e.si);
					p.style.backgroundImage = 'none';
					b.style.display = 'block'
				} else {
					var mobile_flag = this.isMobile();
					if(!mobile_flag) {
						if(ow != w) {
							e.style.width = wd ? Math.ceil(w) + 'px' : Math.floor(w) + 'px'
						}
						if(oh != h) {
							e.style.height = hd ? Math.ceil(h) + 'px' : Math.floor(h) + 'px'
						}
					} else {
						if(ow != w) {
							e.style.width = wd ? Math.ceil('90') + '%' : Math.floor('90') + '%'
						}
						if(oh != h) {
							e.style.height = hd ? Math.ceil(h) + 'px' : Math.floor(h) + 'px'
						}
					}
					var ch = e.style.height = hd ? Math.ceil(h) : Math.floor(h);
					var hh = BOX$('systitle').offsetHeight;
					b.style.cssText = "height:" + (ch - hh) + "px;position: relative;overflow: auto;";
					this.pos();
				}
			} else {
				if(ow != w) {
					var n = e.style.width = 100 + '%'
				}
				if(oh != h) {
					var n = e.style.height = 100 + '%'
				}
				p.style.top = (SYSBOX.page.height() / 2) - (p.offsetHeight / 2) + 'px';
				p.style.left = (SYSBOX.page.width() / 2) - (p.offsetWidth / 2) + 'px';
				b.style.cssText = "height:100%;position: relative;overflow: auto;";
				return false
			}

		},
		dollybox: function(event) {
			var disX = dixY = 0;
			this.style.cursor = "move";
			var event = event || window.event;
			disX = event.clientX - p.offsetLeft;
			disY = event.clientY - p.offsetTop;
			document.onmousemove = function(event) {
				var event = event || window.event;
				var iL = event.clientX - disX;
				var iT = event.clientY - disY;
				var maxL = document.documentElement.clientWidth - p.offsetWidth;
				var maxT = document.documentElement.clientHeight - p.offsetHeight;
				iL <= 0 && (iL = 0);
				iT <= 0 && (iT = 0);
				iL >= maxL && (iL = maxL);
				iT >= maxT && (iT = maxT);
				p.style.left = iL + "px";
				p.style.top = iT + "px";
				return false
			};
			document.onmouseup = function() {
				document.onmousemove = null;
				document.onmouseup = null;
				this.releaseCapture && this.releaseCapture()
			};
			this.setCapture && this.setCapture();
			return false
		},
	}
}();

SYSBOX.page = function() {
	return {
		top: function() {
			return document.documentElement.scrollTop || document.body.scrollTop
		},
		width: function() {
			return self.innerWidth || document.documentElement.clientWidth || document.body.clientWidth
		},
		height: function() {
			return self.innerHeight || document.documentElement.clientHeight || document.body.clientHeight
		},
		total: function(d) {
			var b = document.body,
				e = document.documentElement;
			return d ? Math.max(Math.max(b.scrollHeight, e.scrollHeight), Math.max(b.clientHeight, e.clientHeight)) :
				Math.max(Math.max(b.scrollWidth, e.scrollWidth), Math.max(b.clientWidth, e.clientWidth))
		}
	}
}();