var QueryString = {  
  parse: function(text, sep, eq, isDecode) {
    // デフォルト値を設定
    text = text || location.search.substr(1); // 省略時は URL パラメータ
    sep = sep || '&'; // 省略時は &
    eq = eq || '=';   // 省略時は =
    // デコードするかどうか
    var decode = (isDecode) ? decodeURIComponent : function(a) { return a; };
    // sep(&) で split した配列を reduce で回した結果を返す
    return text.split(sep).reduce(function(obj, v) {
      // eq(=) で split して一番目を key 二番目を value として obj に代入
      var pair = v.split(eq);
      obj[pair[0]] = decode(pair[1]);
      return obj;
    }, {});
  },
  stringify: function(value, sep, eq, isEncode) {
    // デフォルト値を設定
    sep = sep || '&'; // 省略時は &
    eq = eq || '=';   // 省略時は =
    // エンコードするかどうか
    var encode = (isEncode) ? encodeURIComponent : function(a) { return a; };
    // Object.keys で key 配列を取得し, それを map で回した結果を sep(&) で join して返す
    return Object.keys(value).map(function(key) {
      // key, value を eq(=) で連結した文字列を返す
      return key + eq + encode(value[key]);
    }).join(sep);
  },
};
