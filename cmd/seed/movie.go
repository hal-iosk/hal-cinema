package seed

import (
	"time"

	"github.com/hal-iosk/hal-cinema/model"
	"github.com/hal-iosk/hal-cinema/service"
)

func CreateMovie() {
	service.Movie.Create(model.Movie{
		MovieName: "君の名は。",
		Details:   "秒速5センチメートル』（07年）、『言の葉の庭』（13年）など意欲的な作品を数多く作り出してきた気鋭のアニメーション映画監督・新海誠。その新海誠監督の待望の新作となる『君の名は。』は、夢の中で“入れ替わる”少年と少女の恋と奇跡の物語。世界の違う二人の隔たりと繋がりから生まれる「距離」のドラマを圧倒的な映像美とスケールで描き出す。誰もが経験したことのない、アニメーションの新領域。新たな“不朽の名作”が誕生する！",
		StartDate: time.Date(2017, 1, 18, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2017, 6, 28, 0, 0, 0, 0, loc),
		ImagePath: "https://rr.img.naver.jp/mig?src=http%3A%2F%2Fwww.kiminona.com%2Fimages%2Fcommon%2Fog_image.jpg&twidth=1000&theight=0&qlt=80&res_format=jpg&op=r",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "美女と野獣",
		Details:   "ディズニーが製作した大ヒットアニメ『美女と野獣』を実写化した、ファンタジーロマンス。「ハリー・ポッター」シリーズで人気のエマ・ワトソンを主演に、ユアン・マクレガー、ルーク・エヴァンス、ケヴィン・クラインほか、豪華実力派キャストが集結。巨匠アラン・メンケンによる音楽に彩られ、壮大で、華麗な世界が甦る。観る者の期待と想像を遥かに超え、世界中の人々を魅了した、永遠に語り継がれる真実の愛の物語。",
		StartDate: time.Date(2017, 4, 9, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2017, 9, 20, 0, 0, 0, 0, loc),
		ImagePath: "http://www.disney.co.jp/content/dam/disney/images/studio/beautyandbeast/ogp/ogp_bab_01.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "ワイルド・スピード ICE BREAK",
		Details:   "世界的なヒットを記録したカーアクション『ワイルド・スピード』シリーズの第8弾。ヴィン・ディーゼルふんする主人公ドミニクの裏切りによって、強固な絆で結ばれていたファミリーが崩壊の危機にひんするさまを描く。ドウェイン・ジョンソン、ミシェル・ロドリゲスといった続投組のほか、オスカー女優のシャーリーズ・セロンとヘレン・ミレン、クリント・イーストウッドの息子スコット・イーストウッドら豪華キャストが新たに参戦。意表を突く波乱の展開に加え、巨大潜水艦まで登場する氷上カーチェイスにも注目。「ハリー・ポッター」シリーズで人気のエマ・ワトソンを主演に、ユアン・マクレガー、ルーク・エヴァンス、ケヴィン・クラインほか、豪華実力派キャストが集結。巨匠アラン・メンケンによる音楽に彩られ、壮大で、華麗な世界が甦る。観る者の期待と想像を遥かに超え、世界中の人々を魅了した、永遠に語り継がれる真実の愛の物語。",
		StartDate: time.Date(2017, 12, 6, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 3, 2, 0, 0, 0, 0, loc),
		ImagePath: "https://m.media-amazon.com/images/S/aplus-media/sota/e3f181e3-33ff-4f6f-98b8-d52ee6d3a038._SR970,300_.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "海賊とよばれた男",
		Details:   "第10回本屋大賞を受賞した百田尚樹のベストセラー小説を、『永遠の0』の監督＆主演コンビ、山崎貴と岡田准一のタッグで実写映画化。明治から昭和にかけて数々の困難を乗り越え石油事業に尽力した男の生きざまを、戦後の復興、そして世界の市場を牛耳る石油会社との闘いを軸に描く。日本人の誇りを胸に、周囲の仲間との絆を重んじた主人公・国岡鐡造の青年期から老年期までを、主演の岡田が一人でこなす。共演は吉岡秀隆、鈴木亮平、綾瀬はるか、堤真一ら豪華俳優陣がそろう。",
		StartDate: time.Date(2017, 11, 25, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 2, 25, 0, 0, 0, 0, loc),
		ImagePath: "http://www.akiradrive.com/wp-content/uploads/2016/12/kaizoku.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "ズートピア",
		Details:   "あらゆる動物が住む高度な文明社会を舞台にした、ディズニーによるアニメーション。大きさの違いや、肉食・草食にかかわらず、動物たちが共に暮らすズートピアで、ウサギの新米警官とキツネの詐欺師が隠された衝撃的な事件に迫る。製作総指揮をジョン・ラセターが務め、監督を『塔の上のラプンツェル』などのバイロン・ハワードと『シュガー・ラッシュ』などのリッチ・ムーアが共同で担当。製作陣がイマジネーションと新たな解釈で誕生させたという、動物が生活する世界のビジュアル。",
		StartDate: time.Date(2018, 2, 2, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 6, 14, 0, 0, 0, 0, loc),
		ImagePath: "http://www.disney.co.jp/content/dam/disney/images/studio/zootopia/ogp/ogp.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "セーラー服と機関銃 -卒業-",
		Details:   "赤川次郎の小説「セーラー服と機関銃・その後－卒業－」を実写化した異色の青春ムービー。ヤクザの組長をしていた過去を持つ女子高生が、友人たちを巻き込む詐欺の黒幕と対峙（たいじ）する姿を追い掛ける。メガホンを取るのは、『夫婦フーフー日記』などの前田弘二。Rev.from DVLの橋本環奈が主演を務め、共演に『鈴木先生』シリーズなどの長谷川博己のほか、安藤政信、伊武雅刀、武田鉄矢といった実力派が名を連ねている。叫びながら機関銃を撃ちまくるなど、これまでのかれんなイメージを打ち破る橋本の力演に注目。",
		StartDate: time.Date(2018, 3, 5, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 7, 2, 0, 0, 0, 0, loc),
		ImagePath: "https://magazine.maho.jp/pc/img/movie/kikanjuu/visual.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "貞子vs伽椰子",
		Details:   "世界でも評価の高いJホラーの2大キャラクター、『リング』シリーズの貞子と『呪怨』シリーズの伽椰子の対決を描くホラー。貞子と伽椰子、さらには『呪怨』シリーズの俊雄も絡み、恐怖のキャラクターたちによる衝突を活写する。主演は、『東京PRウーマン』などの山本美月。監督には『戦慄怪奇ファイル』シリーズや『ボクソール★ライドショー～恐怖の廃校脱出！～』などの白石晃士。Jホラー界を代表するキャラクターたちの対決の行く末に期待。",
		StartDate: time.Date(2018, 1, 10, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 5, 26, 0, 0, 0, 0, loc),
		ImagePath: "http://imgc.nxtv.jp/img/info/tit/00027/SID0027450.png",
		WatchTime: 90,
	})

	service.Movie.Create(model.Movie{
		MovieName: "HiGH＆LOW THE MOVIE",
		Details:   "ならず者たちが激しい抗争を繰り広げる“SWORD地区”を舞台に、彼らの友情や絆を描いたハードな青春アクション。勢力争いを展開する5つのグループと、かつての支配者による血で血を洗う戦いが始まる。AKIRAやTAKAHIROらEXILE TRIBEのメンバーに加え、窪田正孝、小泉今日子、BIGBANGのV.Iなどバラエティに富んだ共演陣にも注目。",
		StartDate: time.Date(2018, 2, 25, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 6, 19, 0, 0, 0, 0, loc),
		ImagePath: "http://dews365.com/wp-content/uploads/2016/06/maxresdefault.jpg",
		WatchTime: 120,
	})

	service.Movie.Create(model.Movie{
		MovieName: "メッセージ",
		Details:   "SFファンから絶大な支持を受けるテッド・チャンの短編小説を映画化し、第89回アカデミー賞で8部門にノミネートされ、音響編集賞に輝いたSFドラマ。突然、地球に襲来した異星人との交流を通して言語学者が娘の喪失から立ち直っていく姿が描かれる。主人公の言語学者をアカデミー賞では常連の演技派エイミー・アダムスが演じる。",
		StartDate: time.Date(2016, 12, 25, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2017, 4, 6, 0, 0, 0, 0, loc),
		ImagePath: "https://i.ytimg.com/vi/1GMGMzHRE4Q/maxresdefault.jpg",
		WatchTime: 60,
	})

	service.Movie.Create(model.Movie{
		MovieName: "映画 ビリギャル",
		Details:   "学年ビリのギャルがたった1年で偏差値を40も上げて慶應大学に現役合格した実話をつづり、ベストセラーとなった書籍を『ストロボ・エッジ』の有村架純主演で映画化した青春ストーリー。素行不良で何度も停学を経験したヒロインが塾の講師と運命的な出会いを果たし、仲間に支えられながら、慶應大学合格という大きな目標に挑む姿が描かれる。",
		StartDate: time.Date(2017, 10, 14, 0, 0, 0, 0, loc),
		EndDate:   time.Date(2018, 1, 23, 0, 0, 0, 0, loc),
		ImagePath: "https://i0.wp.com/xn--u9j1gsa8mmgt69o30ec97apm6bpb9b.com/wp-content/uploads/2015/11/3a4caea05cf70495644dac951b10ba50.jpg?resize=640%2C507",
		WatchTime: 90,
	})

}
