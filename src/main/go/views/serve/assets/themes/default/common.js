//退出登录
function logOut() {
	$.ajax({
		type : "POST",
		data : "",
		url : "logOut.html",
		success : function(data) {
			window.location = "login.html";
		}
	});
}

// base64加密开始
var keyStr = "ABCDEFGHIJKLMNOP" + "QRSTUVWXYZabcdef" + "ghijklmnopqrstuv"
		+ "wxyz0123456789+/" + "=";
function encode64(input) {
	var output = "";
	var chr1, chr2, chr3 = "";
	var enc1, enc2, enc3, enc4 = "";
	var i = 0;
	do {
		chr1 = input.charCodeAt(i++);
		chr2 = input.charCodeAt(i++);
		chr3 = input.charCodeAt(i++);
		enc1 = chr1 >> 2;
		enc2 = ((chr1 & 3) << 4) | (chr2 >> 4);
		enc3 = ((chr2 & 15) << 2) | (chr3 >> 6);
		enc4 = chr3 & 63;
		if (isNaN(chr2)) {
			enc3 = enc4 = 64;
		} else if (isNaN(chr3)) {
			enc4 = 64;
		}
		output = output + keyStr.charAt(enc1) + keyStr.charAt(enc2)
				+ keyStr.charAt(enc3) + keyStr.charAt(enc4);
		chr1 = chr2 = chr3 = "";
		enc1 = enc2 = enc3 = enc4 = "";
	} while (i < input.length);
	return output;
}

/**
 * 时间转换工具
 */
var Util = function() {
	var FtSeats = function(str) {
		/**
		 * 数字补位
		 * **/
		if (str < 10) {
			str = "0" + String(str);
		}
		return str;
	};

	return {
		DateToUnix : function(year, month, day, hour, minute, second) {
			/**
			 * 日期 转换为 Unix时间戳
			 * @param <int> year    年
			 * @param <int> month   月
			 * @param <int> day     日
			 * @param <int> hour    时
			 * @param <int> minute  分
			 * @param <int> second  秒
			 * @return <int>        unix时间戳(秒)
			 */
			if (arguments.length > 4) {
				var oDate = new Date(Date.UTC(parseInt(year), parseInt(month),
						parseInt(day), parseInt(hour), parseInt(minute),
						parseInt(second)));
			} else {
				var oDate = new Date(year);
			}

			return (Math.floor(oDate.getTime() / 1000));
		},
		UnixToDate : function(unixTime, isFull, timeZone) {
			/**
			 * 时间戳转换日期
			 * @param <int> unixTime    待时间戳(秒)
			 * @param <bool> isFull    返回完整时间(Y-m-d 或者 Y-m-d H:i:s)
			 * @param <int>  timeZone   时区
			 */
			if (typeof (timeZone) == 'number') {
				unixTime = parseInt(unixTime) + parseInt(timeZone) * 60 * 60;
			}

			if (isNaN(unixTime)) {
				return unixTime;
			} else if (unixTime.length < 10) {
				return "-";
			} else if (unixTime.length == 10) {
				unixTime += "000";
			}
			var time = new Date(Math.floor(unixTime / 1000) * 1000);
			var ymdhis = "";
			ymdhis += time.getFullYear() + "/";
			ymdhis += FtSeats(time.getMonth() + 1) + "/";
			ymdhis += FtSeats(time.getDate());
			if (isFull === true) {
				ymdhis += " " + FtSeats(time.getHours()) + ":";
				ymdhis += FtSeats(time.getMinutes()) + ":";
				ymdhis += FtSeats(time.getSeconds());
			}
			return ymdhis;
		},
		DatetoDate : function(date, isDull, timeZone) {
			var unix = this.DateToUnix(date).toString();
			return this.UnixToDate(unix, isDull, timeZone);
		}

	}
}();

/**
 * 身份证校验
 * @param card
 * @returns true or false
 */
function isIdCard(obj) {
	var card = obj.value;
	// 身份证号码为15位或者18位，15位时全为数字，18位前17位为数字，最后一位是校验位，可能为数字或字符X  
	var reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
	if(card.trim() == ''){
		return false;
	}
	if(reg.test(card)){
		return true;
	}else{
		alert('请输入18位合法身份证号码！！！');
		//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
		return false;
	}
}
/**
 * 身份证校验
 * @param card
 * @returns true or false
 */
function isIdCardByValue(value) {
	// 身份证号码为15位或者18位，15位时全为数字，18位前17位为数字，最后一位是校验位，可能为数字或字符X  
	var reg = /(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/;
	if(value.trim() == ''){
		return false;
	}
	if(reg.test(value)){
		return true;
	}else{
		return false;
	}
}

/**
 *检查输入手机号码是否正确
 *@returns true or false
 */
function checkMobile(obj) {
	var mobile = obj.value;
	if(mobile.trim() == ''){
		return false;
	}
//	var regu = /^((13[0-9])|(147)|(15[0-9])|(18[0-9]{1}))\\d{8}$/;
	var regu = /^(1+\d{10})$/; 
	var re = new RegExp(regu);
	if (re.test(mobile)) {
		return true;
	} else {
		alert('手机号码格式不正确！！');
		//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
		return false;
	}
}
/**
 * 检查输入手机号码是否正确
 * @returns true or false
 */
function checkMobileByValue(value) {
	var mobile = value;
	if (mobile.trim() == '') {
		return true;
	}
//	var regu =  /^((13[0-9])|(147)|(15[0-9])|(18[0-9]{1}))\\d{8}$/;
	var regu =  /^(1+\d{10})$/; 
	var re = new RegExp(regu);
	if (re.test(mobile)) {
		return true;
	} else {
		return false;
	}
}

/**
 * 检查输入的电话号码格式是否正确
 * 
 * @param strPhone
 * @returns {Boolean} 如果通过验证返回true,否则返回false
 */
function checkPhone(obj) {
	var strPhone = obj.value;
	if(strPhone.trim() == ''){
		return false;
	}
	var phoneRegWithArea = /^[0][1-9]{2,3}-[0-9]{5,10}$/;
	var phoneRegNoArea = /^[1-9]{1}[0-9]{5,8}$/;
	//var prompt = "您输入的电话号码不正确!"
	if (strPhone.length > 9) {
		if (phoneRegWithArea.test(strPhone)) {
			return true;
		} else {
			alert('电话号码不正确');
			//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
		}
	} else {
		if (phoneRegNoArea.test(strPhone)) {
			return true;
		} else {
			alert('电话号码不正确');
			//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
		}

	}
}

/**
 * 银行卡号校验
 * @param bankCard
 * @returns true or false
 */
function isBankCard(obj) {
	var card = obj.value;
	var reg = /^(\d{16}|\d{19})$/;
	if(card.trim() == ''){
		return false;
	}
	if(reg.test(card)){
		return true;
	}else{
		alert('请输入合法银行卡号！！！');
		//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
	}
}

/**
 * 检查输入的字符是否具有特殊字符
 * @param str
 * @returns {Boolean} true表示包含特殊字符
 */
function checkName(obj) {
	var str = obj.value;
	var items = new Array("~", "`", "!","！", "@", "#", "$", "%", "^", "&", "*",
			"{", "}", "[", "]", "(", ")");
	items.push(":", ";", "'", "|", "\\", "<", ">", "?", "/", "<<", ">>", "||",
			"//","、","‘","’","“","”","\"","·");
	items.push("admin", "administrators", "administrator", "管理员", "系统管理员");
	items.push("select", "delete", "update", "insert", "create", "drop",
			"alter", "trancate");
	str = str.toLowerCase();
	for (var i = 0; i < items.length; i++) {
		if (str.indexOf(items[i]) >= 0) {
			alert('请输入合法的姓名！！');
			//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
			return true;
		}
	}
	return false;
}
function checkNameByValue(str) {
	var items = new Array("~", "`", "!", "！","@", "#", "$", "%", "^", "&", "*",
			"{", "}", "[", "]", "(", ")");
	items.push(":", ";", "'", "|", "\\", "<", ">", "?", "/", "<<", ">>", "||",
			"//","、","‘","’","“","”","\"","·");
	items.push("admin", "administrators", "administrator", "管理员", "系统管理员");
	items.push("select", "delete", "update", "insert", "create", "drop",
			"alter", "trancate");
	str = str.toLowerCase();
	if(str==""){
		return true;
	}
	for (var i = 0; i < items.length; i++) {
		if (str.indexOf(items[i]) >= 0) {
			//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
			return false;
		}
	}
	return true;
}

/**
 * 是否数字
 * @param s
 * @returns {Boolean}
 */
function isNumber(obj) {
	var num = obj.value;
	if(num.trim() == ''){
		return false;
	}
	var re = new RegExp("^[0-9]*$");  
	if (re.test(num)) {
		return true;
	} else {
		alert('请输入纯数字');
		//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
	}
}
/**
 * 是否数字
 * @param s
 * @returns {Boolean}
 */
function isNumberByValue(value) {
	var num = value;
	if(num.trim() == ''){
		return true;
	}
	var re = new RegExp("^[0-9]*$");  
	if (re.test(num)) {
		return true;
	} else {
		return false;
	}
}
/**
 * 检查输入字符串是否符合金额格式
 * @param s
 * @returns {Boolean}
 */
function isMoney(obj) {
	var money = obj.value;
	if(money.trim() == ''){
		return false;
	}
	var isNum=/^(([1-9][0-9]*)|(([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2})))$/;
	if (isNum.test(money)) {
		return true;
	} else {
		alert('请输入正确金额');
		//window.setTimeout( function(){ document.getElementById(obj.id).focus(); }, 0);
	}
}
/**
 * 检查输入字符串是否符合金额格式
 * @param s
 * @returns {Boolean}
 */
function isMoneyByValue(value) {
	var money = value;
	if(money.trim() == ''){
		return true;
	}
	var isNum=/^(([1-9][0-9]*)|(([0]\.\d{1,2}|[1-9][0-9]*\.\d{1,2})))$/;
	if (isNum.test(money)) {
		return true;
	} else {
		return false;
	}
	}
