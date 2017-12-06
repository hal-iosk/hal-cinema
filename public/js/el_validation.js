/**
 *   el_validation.js
 *
 *   Copyright (c) 2013 ellsif.com
 *
 *   This software is released under the MIT License.
 *   http://opensource.org/licenses/mit-license.php
 *
 *   フォームの入力チェック用のバリデーションクラスです。
 *   使い方は下記を参照ください。
 *   http://ellsif.com/blog/?p=1
 */

if (typeof $.el_valid === typeof undefined) {
	$.el_valid = {};
	$.el_valid.setting = {};

	$.el_valid.MESSAGES = {
			'required' : "必須項目です。",
			'min_length' : "%s 文字以上で入力してください。",
			'max_length' : "%s 文字以内で入力してください。",
			'exact_length' : "%s 文字で入力してください。",
			'alpha' : "半角アルファベット以外は入力できません。",
			'alpha_numeric' : "半角英数以外は入力できません。",
			'alpha_dash' : "半角英数字、アンダースコア(_)、ハイフン(-)以外は入力できません。",
			'numeric' : "数字以外は入力できません。",
			'hankaku': '半角文字以外は入力できません。',
			'integer' : "整数以外は入力できません。",
			'natural' : "正の整数以外は入力できません。",
			'less_than' : "%s より小さい数しか入力できません。",
			'greater_than' : "%s より大きな数しか入力できません。",
			'less_eq' : "%s 以下の数しか入力できません。",
			'greater_eq' : "%s 以上の数しか入力できません。",
			'url' : "正しいURLを入力してください。",
			'email' : "正しいメールアドレスを入力してください。",
			'tel' : "正しい電話番号を入力してください。",
			'date' : '正しい日付を入力してください。',
			'time' : '正しい時間を入力してください。',
            'regex_match' : "正しい形式ではありません。",
            'zenkaku' : "全角で入力してください。"
	};

	// バリデーションの初期化
	$.el_valid.init = function(arg) {
		$.el_valid.setting = arg;

		// イベントに対する入力チェックを追加
		for ( var key in arg) {
			var e = 'blur keyup';
			if (arg[key].on) {
				e = arg[key].on;
			}
			$('#' + key).on(e, function(event) {
				if (event.keyCode != 9 && event.keyCode != 244) {	// Tabキー、半角／全角キーは無視
					$.el_valid.valid($(this).attr('id'));
				}
			});
		}
	}

	// 全チェック
	$.el_valid.all = function() {
		var result = true;
		for ( var key in $.el_valid.setting) {
			var res = $.el_valid.valid(key);
			if (!res) {
				result = false;
			}
		}
		return result;
	}

	// バリデーション処理
	$.el_valid.valid = function(name) {
		var val = $('#' + name).val();
		var setting = $.el_valid.setting[name];

		for (var i = 0; i < setting.rule.length; i++) {
			var rule = setting.rule[i];
			var arg = null;

			// []で分割
			var match = rule.match(/(.*?)\[(.*)\]/);
			if (match) {
				if (match.length > 2) {
					rule = match[1];
					arg = match[2];
				}
			}

			if (typeof $.el_valid[rule] === 'function') {
				if (!$.el_valid[rule](val, arg)) {
					var msg = $.el_valid.MESSAGES[rule];

					if (arg) {
						msg = msg.replace(/%s/, arg);
					}

					$.el_valid.err(name, msg, setting);
					return false;
				}
			}
		}

		// エラーなし
		$.el_valid.clear_err(name, setting);
		return true;
	};

	// エラー処理
	$.el_valid.err = function(id, msg, setting) {

		// メッセージの設定
		if (setting.err) {
			$('#'+setting.err).show();
			$('#'+setting.err).html(msg);
		} else if ($('#e_'+id)[0]) {
			$('#e_'+id).show();
			$('#e_'+id).html(msg);
		}

		// スタイルの設定
		if (setting.err_class) {
			$('#' + id).addClass(setting.err_class);
		}
	};

	// エラー全クリア処理（フォームの値もクリアする）
	$.el_valid.clear_all = function() {
		for ( var key in $.el_valid.setting) {
			$.el_valid.clear_err(key, $.el_valid.setting[key]);
			$('#'+key).val('');
		}
	}

	// エラークリア処理
	$.el_valid.clear_err = function(id, setting) {

		// メッセージの設定
		if (setting.err) {
			$('#'+setting.err).show();
			$('#'+setting.err).html('');
		} else if ($('#e_'+id)[0]) {
			$('#e_'+id).show();
			$('#e_'+id).html('');
		}

		// スタイルのクリア
		if (setting.err_class) {
			$('#' + id).removeClass(setting.err_class);
		}
	};

	/*********************************************
	 *  各種バリデーション処理
	 *********************************************/

	// 必須
	$.el_valid.required = function(val) {
		return (val);
	};

	// 最小長
	$.el_valid.min_length = function(val, min) {
		return (val.length == 0 || val.length >= min);
	};

	// 最大長
	$.el_valid.max_length = function(val, max) {
		return (val.length == 0 || val.length <= max);
	};

	// 固定長
	$.el_valid.exact_length = function(val, len) {
		return (val.length == 0 || val.length == len);
	};

	// 半角英字
	$.el_valid.alpha = function(val) {
		return $.el_valid.regex_match(val, "^[a-zA-Z]+$");
	};

	// 半角英数
	$.el_valid.alpha_numeric = function(val) {
		return $.el_valid.regex_match(val, "^[a-zA-Z0-9]+$");
	};

	// 半角英数、アンダースコア、ハイフン
	$.el_valid.alpha_dash = function(val) {
		return $.el_valid.regex_match(val, "^[a-zA-Z0-9_-]+$");
	};

	// 数字
	$.el_valid.numeric = function(val) {
		return $.el_valid.regex_match(val, "^[0-9]+$");
	};

	// 半角
	$.el_valid.hankaku = function(val) {
		if (val) {
			for (var i = 0; i < val.length; i++) {
				var c = val.charCodeAt(i);
				if ( c > 0x80 ) {
					return false;
				}
			}
		}
		return true;
	};

    $.el_valid.zenkaku = function(val) {
		if (val) {
			for (var i = 0; i < val.length; i++) {
				var c = val.charCodeAt(i);
				if ( c > 0x80 ) {
					return true;
				}
			}
		}
		return false;
	};
	// 整数
	$.el_valid.integer = function(val) {
		return val == '0' || $.el_valid.regex_match(val, "^[1-9][0-9]*$") || $.el_valid.regex_match(val, "^-[1-9][0-9]*$");
	};

	// 正の整数
	$.el_valid.natural = function(val) {
		return $.el_valid.regex_match(val, "^[1-9][0-9]*$");
	};

	// より小さい
	$.el_valid.less_than = function(val, arg) {
		if (val) {
			var num = Number(val);
			var comp_num = Number(arg);

			if (num == NaN) {
				return false;
			} else if (comp_num != NaN) {
				return num < comp_num;
			}
		}
		return true;
	};

	// より大きい
	$.el_valid.greater_than = function(val, arg) {
		if (val) {
			var num = Number(val);
			var comp_num = Number(arg);

			if (num == NaN) {
				return false;
			} else if (comp_num != NaN) {
				return num > comp_num;
			}
		}
		return true;
	};

	// 以下
	$.el_valid.less_eq = function(val, arg) {
		if (val) {
			var num = Number(val);
			var comp_num = Number(arg);

			if (num == NaN) {
				return false;
			} else if (comp_num != NaN) {
				return num <= comp_num;
			}
		}
		return true;
	};

	// 以上
	$.el_valid.greater_eq = function(val, arg) {
		if (val) {
			var num = Number(val);
			var comp_num = Number(arg);

			if (num == NaN) {
				return false;
			} else if (comp_num != NaN) {
				return num >= comp_num;
			}
		}
		return true;
	};

	// URL
	$.el_valid.url = function(val) {
		return $.el_valid.regex_match(val, "^(http|https):\/\/.+$");
	}

	// メールアドレス
	$.el_valid.email = function(val) {
		if (val) {
			if (val.match(/^([a-z0-9\+_\-]+)(\.[a-z0-9\+_\-]+)*@([a-z0-9\-]+\.)+[a-z]{2,6}$/i)) {
				return true;
			}
			return false;
		}
		return true;
	};

	// 電話番号
	$.el_valid.tel = function(val) {
		if (val) {
			// ハイフン無し
			if (val.match(/^0\d{9,10}$/)) {
				return true;
			} else {
				var tel = val.split("-");
				if (tel.length != 3) {
					return false;
				}
				if (tel[0].match(/^0\d{1,3}$/) &&
					tel[1].match(/^\d{1,4}$/) &&
					tel[2].match(/^\d{4}$/) &&
					val.replace(/\-/g, '').match(/^0\d{9,10}$/)) {
					return true;
				}
			}
			return false;
		}
		return true;
	};

	// 日付（yyyy/mm/dd）
	$.el_valid.date = function(val) {
		if (val) {
			var tmp = val.split('/');

			if (tmp.length != 3) {
				return false;
			}

			var y = parseInt(tmp[0], 10);
			var m = parseInt(tmp[1], 10);
			var d = parseInt(tmp[2], 10);

			var di = new Date(y, m - 1, d);
			if (di.getFullYear() == y && di.getMonth() == m - 1
					&& di.getDate() == d) {
				return true;
			}
			return false;
		}
		return true;
	};

	// 時間（hh:mm）
	$.el_valid.time = function(val) {
		if (val) {
			var tmp = val.split(':');

			if (tmp.length < 2 || tmp.length > 3) {
				return false;
			}

			var h = parseInt(tmp[0], 10);
			var m = parseInt(tmp[1], 10);
			var s = 0;

			if (tmp.length == 3) {
				s = parseInt(tmp[2], 10);
			}

			if (h >= 0 && h < 24 && m >= 0 && m < 60 && s >= 0 && s < 60) {
				return true;
			}
			return false;
		}
		return true;
	};

	// 正規表現
	$.el_valid.regex_match = function(val, arg) {
		if (val) {
			var reg = new RegExp(arg);
			if (!val.match(reg)) {
				return false;
			}
		}
		return true;
	};
}
