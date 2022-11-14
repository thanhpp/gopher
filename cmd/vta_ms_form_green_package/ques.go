package main

type Ans struct {
	Id  int
	Key string
}

type QuesGroup struct {
	Questions []string
	Answers   []Ans
}

var (
	qChoice = []QuesGroup{q1, q2, q3}
	q1      = QuesGroup{
		Questions: []string{
			"r04f8a63450614a439783eee1d64fcd8e",
			"r1acd5486eaba4640812695314c8ddb50",
			"r282e11efca4746699b77b2a97877337e",
			"r4caefb532498457eae9bdc147533935b",
			"r54dbbad30f684d0fbae9c944b07701c5",
			"r9056b52a98d64550b4efca6ef6a850e0",
			"r9ddef0a0dbaf40cfba7fb639f820d377",
			"rc20fc8e50fd04464af9367e88df79e85",
			"rc2efd2b723f543cdb80e604ec5caa597",
			"rece97f3a671447f4b795293b533e76b4",
		},
		Answers: []Ans{
			{1, "1af6d609-de0e-4a44-8760-bc6b87f48a10"},
			{2, "1d70d21c-934e-47b5-9d52-e593942de6e3"},
			{3, "86897558-7fb7-40c9-9add-107b230b6473"},
			{4, "3b52502e-1070-4412-af70-744f544cc8f5"},
			{5, "7f68b283-0cd7-49f3-a827-fb4b982f23c2"},
			{6, "e51337b2-28d9-4e39-bfcc-eff66163b9e7"},
			{7, "f3925523-16a6-4697-b707-7c40df45a670"},
		},
	}

	q2 = QuesGroup{
		Questions: []string{
			"rde79c7395b4e46eb84605b658ac698f6",
			"r0dc878569da646d781fe630bf8071f85",
			"rd352b222501f4c0391ddf89a630e05eb",
			"ra98e685c3af44904b5389c64a16e4913",
			"r968bebf2a04247a6bfa910a16d9a4aa8",
			"re1718eb4321d4e76baecd4e894a21e22",
			"rf5a143b617fa4b2cb4fe222b92a3cd3d",
			"r20db2bf8cc6a4027bdda06a667bfddb7",
		},
		Answers: []Ans{
			{1, "a0f3757a-2390-43a4-8411-0e0086188714"},
			{2, "871bb7d8-b45c-48e2-a1a6-fa984f347a06"},
			{3, "f3323ccd-3911-40f8-a1e7-5cfd5ce1fca2"},
			{4, "c87a56c7-f200-4331-8d97-40818eb60747"},
			{5, "441201e0-0427-4e54-bb59-5c7676256132"},
			{6, "2009cff6-e7c3-4cae-9569-caf57d9417d8"},
			{7, "66670ac7-c3e6-45c9-8ae2-45a1a8dc727a"},
		},
	}

	q3 = QuesGroup{
		Questions: []string{
			"r6090b752758a44f08ea8edc5cd83089a",
			"r89ca726b97504cd79434726227b11c01",
			"r979bd382cb9e43a18ee27852b1a3d6db",
			"ra883a11e27224ecaa8d6ec4981aaf533",
			"rbbe81ab5579f43e499d31f5ccad24b39",
			"rd096d60cddad4b4da5fed238d689769c",
			"rebfc2d78bf40457386fa5e7d0778565d",
		},
		Answers: []Ans{
			{1, "8c7e2bce-da6b-42d9-bd8a-618780af863d"},
			{2, "14a42472-bfe5-4309-bbd2-5553de86d868"},
			{3, "522bdb6d-b154-477a-9626-9be7aa401436"},
			{4, "4e91e9a7-bf0a-4626-8c96-0ac2d7fe14ba"},
			{5, "c61c0522-3444-4acd-ab69-d6378906ad23"},
			{6, "14949038-49d5-4d4d-8b9e-40ceecedefe5"},
			{7, "fa47408f-f7ed-45f3-8963-7f2efc1b559a"},
		},
	}
)

type QuesMul struct {
	Id  string
	Ans []string
}

var (
	quesMul = QuesMul{
		Id:  "raaccbea3b3e449b3bc42cdd356c1f98d",
		Ans: []string{"Tự nhiên", "Hóa học"},
	}
)

type Ques struct {
	Id  string
	Ans []string
}

var (
	quesTxt = []Ques{
		{
			Id:  "rbf9a69dacb244700a9de97a7e4912735",
			Ans: mailList,
		},
		{
			Id:  "r10492d18f1574977a9eb1d6af40625d0",
			Ans: []string{"Có"},
		},
		{
			Id:  "r34373a1a64084a1db540399d473d9a17",
			Ans: []string{"Đã dùng"},
		},
		{
			Id:  "rc807d0a7d8fe4296a530c6dd1a17815f",
			Ans: []string{"Có, mình rất muốn thử", "Có, mình rất muốn thử", "Có, mình rất muốn thử", "Có, mình rất muốn thử", "Không, mình không muốn thử"},
		},
		{
			Id:  "rd6a0c0b218b143e392397870ff2ee411",
			Ans: []string{"cocoon", "Cocoon", "Frudia", "cỏ mềm", "Cỏ mềm", "decumar"},
		},
		{
			Id:  "rd5416c02dd4c48cd8fdda73e7d8f7f22",
			Ans: []string{"Nữ", "Nữ", "Nữ", "Nữ", "Nữ", "Nam"},
		},
		{
			Id:  "rc13060d56e494305955a823cdfbf2287",
			Ans: []string{"Sinh viên Đại học", "Sinh viên Đại học", "Sinh viên Đại học", "Tốt nghiệp Đại học", "Đã đi làm"},
		},
		{
			Id:  "rbdc2c061351045f9927bf93916fa53d6",
			Ans: []string{"Độc thân", "Độc thân", "Đang trong một mối quan hệ"},
		},
		{
			Id: "rffd89e998912452bbeab03083a82475f",
			Ans: []string{
				"18-20 tuổi",
				"18-20 tuổi",
				"18-20 tuổi",
				"18-20 tuổi",
				"21-25 tuổi",
				"21-25 tuổi",
				"21-25 tuổi",
				"25-30 tuổi",
			},
		},
		{
			Id: "r44179e41db714d99b2107d9abb687e57",
			Ans: []string{
				"Dưới 3 triệu VNĐ",
				"Dưới 3 triệu VNĐ",
				"Dưới 3 triệu VNĐ",
				"Dưới 3 triệu VNĐ",
				"Từ 3 đến 5 triệu VNĐ",
				"Từ 3 đến 5 triệu VNĐ",
				"Trên 5 triệu VNĐ",
				"Trên 5 triệu VNĐ",
			},
		},
	}
)

var (
	mailList = []string{
		"buikhoaha@yahoo.com.vn",
		"nvhai.dhxd@gmail.com",
		"kientrucac37@gmail.com",
		"thanhhavff@yahoo.com",
		"thuphuong1898@gmail.com",
		"sonhv@vingroup.net",
		"ngotrunghai2009@gmail.com",
		"giangvh.hk@gmail.com",
		"dt.tung@vietinbank.vn",
		"phuongntm@oceanbank.vn",
		"nmphuong75@yahoo.com",
		"thuhoaipham102@yahoo.com",
		"ngothuphuong@vietinbank.vn",
		"trangthu.lee@gmail.com",
		"thinhrudn@yahoo.com",
		"baolinh.preciousbell@yahoo.com",
		"thtqcuong@yahoo.com.vn",
		"thanhlamtran2004@yahoo.com.vn",
		"viet.hoang.0302@gmail.com",
		"luucongthanh@gmail.com",
		"thangeodesy@yahoo.com",
		"huongntt1@vincomjsc.com",
		"thailm@bidv.com.vn",
		"huyenbtk40@yahoo.com.vn",
		"hoanggialocxd@gmail.com",
		"oanhtran@hsbc.com.vn",
		"ng_mthag@yahoo.com",
		"gianght@vietinbank.vn",
		"lanvan78@yahoo.com",
		"binh.nt@vietinbank.vn",
		"nguyenvietcuong@vietinbank.vn",
		"huongltb@vietinbank.vn",
		"lthoa@vietinbank.vn",
		"maiphuong.dang@yahoo.com",
		"lamdh@vietinbank.vn",
		"fullmoon159@yahoo.com.vn",
		"anhdv@vietinbank.vn",
		"thoadt@vietinbank.vn",
		"dailc@vietinbank.vn",
		"haingotkiev@gmail.com",
		"ngocnb@pvfc.com.vn",
		"ngoclong1612@yahoo.com",
		"giang.kien@kiwifood.net",
		"smakiazji@gmail.com",
		"ichhien@gmail.com",
		"tranghahn2000@yahoo.com",
		"ntduong2006@yahoo.ca",
		"hoantt.bic@bidv.com.vn",
		"huongnguyen_quoc@vietfracht.com.vn.com.vn",
		"haln@bidv.com.vn",
		"hangfashion_vn@yahoo.com.vn",
		"bm_phuong@yahoo.com",
		"thuynt6@bidv.com.vn",
		"honganhcn@gmail.com",
		"htmdch21@yahoo.com",
		"nguyenthanhson093@gmail.com",
		"toantn@bidv.com.vn",
		"hanhna@bidv.com.vn",
		"duyanhhs@gmail.com",
		"tue.nga.chi@gmail.com",
		"thanhvx@bidv.com.vn",
		"quocthang48@gmail.com",
		"luuvan_2000@yahoo.com",
		"dungeniags@yahoo.com.vn",
		"giangpham482@yahoo.com",
		"linhnguyenthuy1410@yahoo.com",
		"cuongle@tdtco.com.vn",
		"minhnh@hsbc.com.vn",
		"kiss_minh_honey@yahoo.com",
		"lechiphuc@gmail.com",
		"cat706194@yahoo.com",
		"hien.nguyen389@gmail.com",
		"duonghuyentrang1808@gmail.com",
		"thongndm@epu.edu.vn",
		"yen7475@yahoo.com",
		"trucquynh@phuongdonginox.vn",
		"tuyetnganht@yahoo.com",
		"construc.pdc@gmail.com",
		"tran_vankhang@yahoo.com",
		"thaianhvu@gmail.com",
		"duong.thuyphuong@gmail.com",
		"longnh@cisco.com",
		"namnotl@yahoo.com",
		"hoangson76@gmail.com",
		"tranbaongoc2000@yahoo.com",
		"catvina.cuong@gmail.com",
		"thanhquyet@yahoo.com",
		"hoanghl@163.com",
		"phandunghbb@gmail.com",
		"hailinhngoc@yahoo.com",
		"llinhbau0207@yahoo.com",
		"nhuuchinh@yahoo.com",
		"honda_gialoc@yahoo.com.vn",
		"vuthuthuy63@gmail.com",
		"phuonglien1955@yahoo.com.vn",
		"daotghien@yahoo.com",
		"dungnnpl@yahoo.com.vn",
		"tp_quynhnhu@yahoo.com",
		"ngocanhn72@yahoo.com.vn",
		"anhmai_hntv@yahoo.com",
		"nguyenthikynam@yahoo.com",
		"hongphong28@gmail.com",
		"phamvantuan792000@yahoo.com",
		"vibv@foopro.com.vn",
		"lannt@marina-logistics.com",
		"cogaikinhbac01@yahoo.com",
		"duyquang@fpt.vn",
		"giangdieplinh@gmail.com",
		"thangle.mb@yahoo.com",
		"oho_aitai@yahoo.com.vn",
		"sandythutrang@gmail.com",
		"daphacobr@hn.vnn.vn",
		"nguyen.bhue@gmail.com",
		"lanphuongvtl@yahoo.com",
		"lethu4549@gmail.com",
		"tuanla6@bidv.com.vn",
		"tienphong122@gmail.com",
		"phamvanninh75@vnn.vn",
		"quocnb@evnfc.com",
		"maianh889@yahoo.com",
		"anhptd@vinnataba.com.vn",
		"sunlogo@hn.vnn.vn",
		"tonycaozz6688@yahoo.com",
		"lalinh@cmc.com.vn",
		"le.nguyen@vtc.vn",
		"l.h.anh@fpt.vn",
		"hiepcv.mobilems@gmail.com",
		"longhd_ice@mail.hut.edu.vn",
		"tieuchien@yahoo.com",
		"phuong@126.com",
		"kotobukibg@yahoo.com",
		"vmduc1979@yahoo.com",
		"levietnga@gmail.com",
		"huongb@bidv.com.vn",
		"ngahieu0703@yahoo.com",
		"xuanbt@bidv.com.vn",
		"lhthuy@gmail.com",
		"nhungnp@bidv.com.vn",
		"levudiemhang@gmail.com",
		"vinhung@gmail.com",
		"lamnp@bidv.com.vn",
		"thuhientn301@yahoo.com",
		"huhoang@deloitle.com",
		"datnh@bidv.com.vn",
		"ltdungvnpar@yahoo.com",
		"ngatt@bidv.com.vn",
		"letra75@yahoo.com",
		"dinhthuylinh8999@gmail.com",
		"kimhuong11091964@yahoo.com",
		"trungtse@gmail.com",
		"haichautran@yahoo.com",
		"caominhchau@yahoo.com",
		"iceicebaby2510@gmail.com",
		"thuhanh@vietinbank.vn",
		"dunglan63@gmail.com",
		"longzin3110@gmail.com",
		"oriole0207@yahoo.com",
		"canhdn.ho@techcombank.com.vn",
		"canhdn@yahoo.com",
		"quoccuong@sovicoholdings.com",
		"duyn80@gmail.com",
		"tranlethanh2009@yahoo.com",
		"truonggiangoto@yahoo.com",
		"chu.kim.chi@gmail.com",
		"haminhhiep@gmail.com",
		"phuongkhoashop_hn@yahoo.com",
		"vunguyenhonganh@yahoo.com",
		"vanxuanvn@vnn.vn",
		"nthien99@yahoo.com",
		"youthjsc@yahoo.com",
		"mviethoa2002@yahoo.com",
		"thanhtl15@fpt.vn",
		"hongle1952@mail.ru",
		"daothuyhoa63@yahoo.com",
		"lexuantra@gmail.com",
		"hanoi1501@yahoo.com",
		"tranngoc61@gmail.com",
		"thanggd386@gmail.com",
		"chuonglv@gmail.com",
		"quyln@bidv.com.vn",
		"hiepdt@bidv.com.vn",
		"menam0903@yahoo.com.vn",
		"toanxq@bidv.com.vn",
		"inmascoknight@yahoo.com",
		"thunga002000@yahoo.com",
		"cuonghai_bn@yahoo.com",
		"cuong.bv@bidv.com.vn",
		"quynhsky@gmail.com",
		"dinhbichthuy68@yahoo.com",
		"truyenthonghtv@gmail.com",
		"duongvt@bidv.com.vn",
		"binhdt@bidv.com.vn",
		"khanhpn@bidv.com.vn",
		"huylq@bsc.com.vn",
		"thanhhuong20383@gmail.com",
		"duongvesdi74@yahoo.com",
		"hai_ttcp75@yahoo.com.vn",
		"dungthao21@yahoo.com",
		"baohero1009@yahoo.com",
		"kmt0085@yahoo.com",
		"thuhong.nguyentatx@gmail.com",
		"dungcv81@gmail.com",
		"pttunganh@gmail.com",
		"huett@pvep.com.vn",
		"yendong73@yahoo.com.vn",
		"hoadp_sveta@yahoo.com",
		"hoaibeo@gmail.com",
		"ngmnga@yahoo.com",
		"quynhdn@bidv.com.vn",
		"quochoang35@yahoo.com",
		"hoangxuan_vnx@yahoo.com",
		"phanhuongduong@gmail.com",
		"nxtien08@gmail.com",
		"hungld@moj.gov.vn",
		"ncaoluan@yahoo.com",
		"dungdx@gmail.com",
		"tinhlb@gmail.com",
		"caupt@vincomjsc.com",
		"navakoi@yahoo.com",
		"tuanle-civil@yahoo.com",
		"quoclamcb@gmail.com",
		"nthuhuong@vietinbank.vn",
		"huy_krt@yahoo.com",
		"bs.m.nguyen@yahoo.com",
		"khanhhoa14b@vnn.vn",
		"vuyen1282@yahoo.com",
		"anhnkd@gmail.com",
		"ncphuonghanoizoo@yahoo.com",
		"dr.lequangvinh@yahoo.com",
		"minh@hamalinclothing.com",
		"giangnx@vinatranshp.com.vn",
		"phgiangnguyen@yahoo.com.br",
		"nguyethdm@yahoo.com",
		"qlinh87@yahoo.com",
		"nkimphi@yahoo.com",
		"tranthutrang_1910@yahoo.com",
		"toccungngo39@yahoo.com",
		"havt@act.com.vn",
		"hanoisamac@yahoo.com",
		"hoangvanbach@yahoo.com",
		"huongdang1974@yahoo.com",
		"sunt.bic@bidv.com.vn",
		"liberty_tm@vnn.vn",
		"nghiatt.bic@bidv.com.vn",
		"oanhvtk@bidv.com.vn",
		"vuhoang01@gmail.com",
		"doanha7608@yahoo.com",
		"httbinh@bidv.com.vn",
		"linhpt.bic@bidv.com.vn",
		"dungtien167@gmail.com",
		"ntson65@yahoo.com",
		"truongnguyen0809@gmail.com",
		"thuyttt.bic@bidv.com.vn",
		"anh_lemai@yahoo.com",
		"mta.vnn@gmail.com",
		"khanh.db@seabank.com.vn",
		"ctnga58@yahoo.com.vn",
		"giangthaibang72@yahoo.com.vn",
		"trungnghia_pharma@yahoo.com.vn",
		"icvietnam@yahoo.com",
		"vanhdoan@mitalabvn.com",
		"quydxp23@gmail.com",
		"nguyenanhpl@yahoo.com",
		"sn.nc@seabank.com.vn",
		"landao117@yahoo.com",
		"thanhhq_daisy@yahoo.com",
		"lethihang_20@yahoo.com",
		"benie.auto@gmail.com",
		"tuyen.tc@thuanphatgroup.com.vn",
		"huyenbecky@gmail.com",
		"fgchanoi@gmail.com",
		"thao.nt@vietinbank.vn",
		"tuyenbtk@yahoo.com",
		"quynhnga36@yahoo.com",
		"longtrong85@yahoo.com.vn",
		"kthoang88@yahoo.com.vn",
		"thanhmaiph@fpt.vn",
		"maitungson1311@yahoo.com",
		"anhtrankim1x@gmail.com",
		"hoan_giang@163.com",
		"dung@idp.vn",
		"kimdungmedi@gmail.com",
		"phantheduy@baoviet.com.vn",
		"trinhla23@gmail.com",
		"tuanhm@viettel.com.vn",
		"hmtuanhut@yahoo.com",
		"president@vkttc.gov.vn",
		"inectethn@yahoo.com",
		"hanoidaotao@vnn.vn",
		"tranthuybk60@yahoo.com",
		"thanhhai842000@gmail.com",
		"phuong_stylish_9885@yahoo.com",
		"thienkh@hvnh.edu.vn",
		"anly@inteco.com.vn",
		"thanhdongytq@gmail.com",
		"thumiuberlin@yahoo.com",
		"thanhnn@moit.gov.vn",
		"nghiemquochungbm@gmail.com",
		"vuviquoc@yahoo.com",
		"lunnova@yahoo.com",
		"phuongmmr@gmail.com",
		"huongtramtct@yahoo.com",
		"thaoanhhd_2002@yahoo.com",
		"luongnguyetminh83@gmail.com",
		"ptphuong69@yahoo.com.vn",
		"baotrung86@gmail.com",
		"tuananhmit78@yahoo.com",
		"nguyenhongvan@ducphatsteel.com.vn",
		"tham_1967@gmail.com.vn",
		"cc123ngoc@yahoo.com.vn",
		"nguyenmanhtuan@coninco.com.vn",
		"quocnb@evnfc.vn",
		"danghaitdc46@yahoo.com",
		"dungcandelta@yahoo.com.vn",
		"tuandeltahn@yahoo.com",
		"hangbibi81@yahoo.com",
		"phamvantuan_delta@yahoo.com.vn",
		"chigaivn03@yahoo.com",
		"bathang1988@gmail.com",
		"anphuong_hl@yahoo.com",
		"adtltd@yahoo.com",
		"mai_khac_nhu@yahoo.com",
		"tranghop1421@yahoo.com",
		"hiepnt74@yahoo.com",
		"trieuminhngoc108@yahoo.com",
		"chungphong_vn@yahoo.com.vn",
		"mailinhgiang@yahoo.com.vn",
		"thieunv@trustbank.com.vn",
		"DTTVinh67@yahoo.com",
		"thaoptb@vincomjsc.com",
		"quachhoangngan@yahoo.com",
		"vu.thu_ha@yahoo.com",
		"pth_hdu@yahoo.com",
		"anhnhg@bidv.com.vn",
		"nguyetnhi999@yahoo.de",
		"tanviettien@gmail.com",
		"dongmn.bic@bidv.com.vn",
		"hailp@bidv.com.vn",
		"mauyen1@yahoo.com",
		"nguyetttt@bidv.com.vn",
		"trannguyet80@yahoo.com",
		"maihuongnt85@gmail.com",
		"hapt2@bidv.com.vn",
		"lemontree16102001@yahoo.com",
		"thuyptt5@bidv.com.vn",
		"anhuv@bidv.com.vn",
		"hanhbtt@bidv.com.vn",
		"oanhnt@bidv.com.vn",
		"hongngabidv88@yahoo.com.vn",
		"hvq80@yahoo.com",
		"minhnguyetdvt@gmail.com",
		"datanh_a9@yahoo.com.vn",
		"chungthuy1806@yahoo.com",
		"truongquochung73@gmail.com",
		"kts-tangx82x@yahoo.com.vn",
		"lanlp@vietinbank.vn",
		"levq@vietinbank.vn",
		"tohaininh@vietinbank.cvn",
		"manhtd@vietinbank.vn",
		"ngocvtm@vietinbank.vn",
		"dienlien@yahoo.com",
		"trinhthanhnhan@vietinbank.vn",
		"haolv@vietinbank.vn",
		"lethidung@vietinbank.vn",
		"thuha2268@gmail.com",
		"vuthuylinh74@hn.vnn.vn",
		"khanhpq@vietinbank.vn",
		"lanltm@vietinbank.vn",
		"bangnh@vietinbank.vn",
		"tuanbn@vietinbank.vn",
		"quynhdtt@vietinbank.vn",
		"tuanhna@vietinbank.vn",
		"hapth@vietinbank.vn",
		"tuyetdung@techcombank.com.vn",
		"anhmn.hcm@techcombank.com.vn",
		"tranglh.qtrr@techcombank.com.vn",
		"thanghq@techcombank.com.vn",
		"yennt.ho@techcombank.com.vn",
		"yenlth@techcombank.com.vn",
		"phuongpt.khl@techcombank.com.vn",
		"anhnv.ho@techcombank.com.vn",
		"thanhvan.nguyen@yahoo.com",
		"ttrungsan1980@yahoo.com",
		"hiendv@techcombank.com.vn",
		"bangbt@techcombank.com.vn",
		"lananh007@yahoo.com",
		"hungescape@yahoo.com",
		"the_guardian8888@yahoo.com",
		"thongbh@gmail.com",
		"chibuikim@yahoo.com",
		"dangducchinh@veam-motor.com",
		"hung.escape@yahoo.com",
		"huunam@hnb.vn",
		"daongocthu2003@gmail.com",
		"cuong6c@gmail.com.vn",
		"dtd060@gmail.com",
		"ngocanh2510@gmail.com",
		"kien_vncc@yahoo.com",
		"vuducthaiani@gmail.com",
		"white_house_68@yahoo.com.vn",
		"anhdlm@vincomjsc.com",
		"ngocdtb@canifa.com",
		"nguyenthanhtruc@ajc.com.vn",
		"hangmtt@systech.com.vn",
		"quangvankts@gmail.com",
		"tiendanb@yahoo.com",
		"vansbv73@yahoo.com",
		"thaolp@techcombank.com.vn",
		"nguyetlm@dap.vn",
		"tuttm75@yahoo.com",
		"phamhai.khcn@gmail.com",
		"hangxinheco@yahoo.com",
		"huongcts@gmail.com",
		"vincovn@yahoo.com",
		"dang.toannh@gmail.com",
		"lienntp@fpt.com.vn",
		"duclm2010@gmail.com",
		"luongnab@yahoo.com.vn",
		"nga.vnexpress@gmail.com",
		"phungvinhbtg@gmail.com",
		"hung.td@vietinbank.vn",
		"tranngocdiep0211@gmail.com",
		"binh2201@yahoo.com",
		"vixuco@fpt.com.vn",
		"tathihien@hn.vnn.vn",
		"dunghpcd@gmail.com",
		"hoa17med@gmail.com",
		"daodung_vncb@yahoo.com",
		"quoctien80@gmail.com",
		"hiennam68@yahoo.com",
		"thovnu@gmail.com",
		"aciapacificocean@gmail.com",
		"huonghtt.hqv@mbbank.com.vn",
		"leminhthu_1908@yahoo.com",
		"duongngoclan.duong@gmail.com",
		"nguyetdao01@gmail.com",
		"hoangxuanthuy1975@yahoo.comm.vn",
		"phuongthaoplastic@gmail.com",
		"bien_sang_17@yahoo.com",
		"dangluonganh@gmail.com",
		"nuidt@agribank.com.vn",
		"btkhue1974@gmail.com",
		"loanajc2009@gmail.com",
		"trangnguyen.3188@yahoo.com",
		"vuong_ngoc85@yahoo.com",
		"ngoc@francopacific.com",
		"forever7783@yahoo.com",
		"kimlan0263@yahoo.com",
		"hienquang@fpt.vn",
		"huongnguyen@vietinbank.vn",
	}
)