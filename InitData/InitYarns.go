package InitData

import (
	"github.com/mustafasaahin/Draper/config"
	"github.com/mustafasaahin/Draper/models"
)

func InitYarns() {
	var yarns models.Yarns
	config.DB.
		Where("id<>0").
		Delete(&models.Yarns{})
	yarns = models.Yarns{
		Code:      "001",
		Name:      "13/1 KAPLANLAR 30 WOOL 70 PES",
		UnitPrice: 7.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "002",
		Name:      "70dn lycra(75f36+40f1)  115f37   COMBAT 18.10.2021",
		UnitPrice: 4.29,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "003",
		Name:      "150 Denye Koyu Tintin",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "004",
		Name:      "150 F 72 Tekstüre  Melanj     18.10.2021",
		UnitPrice: 3.33,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "005",
		Name:      "450 F 144 Tekstüre Melanj     18.10.2021",
		UnitPrice: 2.76,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "006",
		Name:      "28/1boyalıK.ELYAF flam + 40/1 renkli likra",
		UnitPrice: 6.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "007",
		Name:      "150 F 48T Teksture Img - SPR       18.10.2021",
		UnitPrice: 2.44,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "008",
		Name:      "300 F 96 S Texture Img Siyah   18.10.2021",
		UnitPrice: 2.58,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "009",
		Name:      "300 F 120 Teks Melanj       18.10.2021",
		UnitPrice: 2.78,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "010",
		Name:      "270 Denye Yün Efekli (SMT)",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "011",
		Name:      "28/1 renkli +40/1 polyvison likra",
		UnitPrice: 5.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "012",
		Name:      "14/1 VİSKON FLAM",
		UnitPrice: 4.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "013",
		Name:      "100 F 36 Tekstüre-İMG Ekru",
		UnitPrice: 2.24,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "014",
		Name:      "150 F 72X2  (  150/144 )M.Streç Teks.kat 18.10.21",
		UnitPrice: 2.87,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "015",
		Name:      "70  dn lycra siyah (50f36+20f1)",
		UnitPrice: 5.19,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "016",
		Name:      "300 F 96T Tekstüre-İMG SPR       18.10.2021",
		UnitPrice: 2.24,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "017",
		Name:      "300 Denye Koyu Tintin",
		UnitPrice: 3.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "018",
		Name:      "300 Denye Düz Katyonik",
		UnitPrice: 3.45,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "019",
		Name:      "120/500 T Floş",
		UnitPrice: 13.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "020",
		Name:      "150f/48s dn siyah tekstüze 18.10.2021",
		UnitPrice: 2.42,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "021",
		Name:      "150 F 48 S Texture IMG Siyah puntalı18.10.2021",
		UnitPrice: 2.61,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "022",
		Name:      "150 F 36 - 48 x 2 teks-İmg   (300/ 96punta )18.10",
		UnitPrice: 2.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "023",
		Name:      "#40/2  POLY VİSKON SİYAH",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "024",
		Name:      "150 F 144 Teks Soft İmg             18.10.2021",
		UnitPrice: 2.39,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "025",
		Name:      "150 F 167 X 84 CE Melanj F-Set   özel üretim",
		UnitPrice: 6.7,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "026",
		Name:      "150 F 167 x 48 R Siyah Melanj     özel üretim",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "027",
		Name:      "150 F288 x 2 Mikro Flaman Tek.  (300/576)18.10",
		UnitPrice: 2.39,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "028",
		Name:      "Efek İpligi",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "029",
		Name:      "26/1 KETEN BOYALI AÇIK RENK BTMZ",
		UnitPrice: 17,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "030",
		Name:      "50/2 FLAM CD.İTAL",
		UnitPrice: 7,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "031",
		Name:      "150 DNY POLY.KATYONİK",
		UnitPrice: 3.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "032",
		Name:      "100 F 36 TEKTÜRE TURLU 500 tmz",
		UnitPrice: 3.4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "033",
		Name:      "140/96 YÜN EFEKLİ",
		UnitPrice: 5.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "034",
		Name:      "135f132 DNY HAM TURLU EKRU",
		UnitPrice: 3.74,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "035",
		Name:      "75F36 S.P.FDY 350TM/2           Harput",
		UnitPrice: 3.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "036",
		Name:      "20/1 VİSKON",
		UnitPrice: 5.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "037",
		Name:      "100 F 48T TEKS BÜK SP İMG 550TMZ",
		UnitPrice: 5.44,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "038",
		Name:      "28/1 K.ELYAF  BOYALI FLAM+28/1 PV SİYAH",
		UnitPrice: 6.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "039",
		Name:      "100 F 48 T TEKS BÜKÜM 350 TM",
		UnitPrice: 4.64,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "040",
		Name:      "20/1 MOMOFLAMAN SÜPER PARLAK",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "041",
		Name:      "150F96 EKRU KORTEKS",
		UnitPrice: 2.19,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "042",
		Name:      "12/1 VİSKON",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "043",
		Name:      "150 FDY S.PARLAK RENKLİ",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "044",
		Name:      "100 F 36 S.P FDY.350 TM/Z HAM TURLU",
		UnitPrice: 4.64,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "045",
		Name:      "150 F 48 S.P.FDY.250TM/Z",
		UnitPrice: 2.87,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "046",
		Name:      "150 F 48 SİYAH TURLU",
		UnitPrice: 3.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "047",
		Name:      "450/156 CD+PES IMG A.MEL",
		UnitPrice: 4.95,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "048",
		Name:      "334/96 H.TEKS S.P DUB.HVY",
		UnitPrice: 6.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "049",
		Name:      "20/1 PAMUK POLYESTER KARDE ( İTAL)  GLOBTEKS",
		UnitPrice: 2.85,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "050",
		Name:      "300 CD+PES",
		UnitPrice: 3.83,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "051",
		Name:      "150 F48 T FDY.MAVİ 001  S.P",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "052",
		Name:      "150 F 48 T FDY SARI 009  S.P",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "053",
		Name:      "28/4            19202 nin ipliği   KAVAK TEKSTİL",
		UnitPrice: 4.7,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "054",
		Name:      "1000 Denye tekstüre pes",
		UnitPrice: 2.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "055",
		Name:      "300/576 TEKSTÜRE MİCRO",
		UnitPrice: 2.24,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "056",
		Name:      "20/1 TENSEL FLAM",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "057",
		Name:      "30/1 PAMUK OPENENT",
		UnitPrice: 2.15,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "058",
		Name:      "50/2 KATYONİK PES İTHAL GAYRET MODİANO EKRU",
		UnitPrice: 7,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "059",
		Name:      "300 F 72 TEKSTÜRE İMG",
		UnitPrice: 1.94,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "060",
		Name:      "#40/2 EKRU P/V",
		UnitPrice: 5.4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code: "061",
		Name: "30/1 PAMUK DE TRİKO",
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "062",
		Name:      "300 FDY KAHVE TURLU",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "063",
		Name:      "450 FDY SİYAH TURLU",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "064",
		Name:      "300 F 96T FDY SİYAH S.PARLAK",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "065",
		Name:      "8/1 PAMUK OPENED",
		UnitPrice: 3.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "066",
		Name:      "150 F 48 TEKSTÜRE",
		UnitPrice: 1.86,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "067",
		Name:      "170 DENYE WOOL EFECT (SMT)",
		UnitPrice: 4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "068",
		Name:      "150 F 84 MELANJ TURLU",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "069",
		Name:      "150 F 84 PARLAK MELANJ TURLU",
		UnitPrice: 6.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "070",
		Name:      "200 İMG BEYAZ TURLU",
		UnitPrice: 3.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "071",
		Name:      "20/1  P.V.SİYAH",
		UnitPrice: 4.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "072",
		Name:      "30/1  P.V. SİYAH",
		UnitPrice: 4.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "073",
		Name:      "50/2 P.V. SİYAH",
		UnitPrice: 6.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "074",
		Name:      "200 F 72 Tekstüre-İMG EKRU",
		UnitPrice: 2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "075",
		Name:      "30/1 P.V. EKRU",
		UnitPrice: 3.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "076",
		Name:      "300/75 LİNEN",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "077",
		Name:      "90 DENYE 350 TURLU SİYAH",
		UnitPrice: 5.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "078",
		Name:      "100 F 36 S.P FDY TURLU SİYAH",
		UnitPrice: 5.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "079",
		Name:      "100 F 48 S.P FDY TURLU BEYAZ",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "080",
		Name:      "20/1 P.V EKRU",
		UnitPrice: 3.45,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "081",
		Name:      "7/2 STR.MULİNE LACİ-A.GRİ (BERNİ)",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "082",
		Name:      "20/2 P.V",
		UnitPrice: 5.01,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "083",
		Name:      "30/1 P.V CD",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "084",
		Name:      "20/2 P.V.LYCRA",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "085",
		Name:      "150/48 PES LYCRA",
		UnitPrice: 4.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "086",
		Name:      "100 F 36 S  TEKSTÜRE SİYAH-İMG",
		UnitPrice: 3.4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "087",
		Name:      "75/36 POLYESTER İMG KORTEKS",
		UnitPrice: 2.34,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "088",
		Name:      "167/150/84 İMG MEL S.P",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "089",
		Name:      "10/1 BUKLET GÜRAY FANTAZİ",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "090",
		Name:      "200 F 96 T TEKS.ASG.S.PARLAK",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "091",
		Name:      "*28/2 POLY VİSKON RENKLİ",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "092",
		Name:      "40/2 P.V SİYAH LYCRA",
		UnitPrice: 6.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "093",
		Name:      "150 denye boyatma batmaz gri",
		UnitPrice: 4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "094",
		Name:      "50/1 P.V SİYAH",
		UnitPrice: 5.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "095",
		Name:      "28/1 P.V TEK KAT SİYAH",
		UnitPrice: 3.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "096",
		Name:      "10/1 BERNİ STRAYGAN YÜNLÜ",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "097",
		Name:      "6/1  OPENEND ÇÖZGÜLÜK SİYAH",
		UnitPrice: 2.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "098",
		Name:      "70 F 37 S E.PUNTALI SİYAH KAP",
		UnitPrice: 6.77,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "099",
		Name:      "*28/2  POLY.VİSKON SİYAH",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "100",
		Name:      "*28/2 POLY.VİSKON EKRU",
		UnitPrice: 4.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "101",
		Name:      "40/1 P.V. EKRU",
		UnitPrice: 4.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "102",
		Name:      "28/1 P.V. EKRU FLAMLI",
		UnitPrice: 4.45,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "103",
		Name:      "50/2 EKRU Z BÜKÜM",
		UnitPrice: 8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "104",
		Name:      "50/2 Z BÜKÜM LYCRA",
		UnitPrice: 8.4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "105",
		Name:      "150/96 FDY BÜKÜM YM 550 TM",
		UnitPrice: 4.17,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "106",
		Name:      "68 F 24 FDY BÜKÜM YM 800 TM",
		UnitPrice: 4.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "107",
		Name:      "80/20 W00L KAŞMİR EKRU 13/1 KAPLANLAR",
		UnitPrice: 10.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "108",
		Name:      "26/1 KASARLI KETEN",
		UnitPrice: 16,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "109",
		Name:      "150 -(159/48/70)LİKRA SİYAH KAPLAMALI",
		UnitPrice: 4.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "110",
		Name:      "150-(159/144/70)LİKRA EKRU",
		UnitPrice: 3.24,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "111",
		Name:      "28/2 RENKLİ+FLAM MULİNE DİCLE 17033",
		UnitPrice: 4.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "112",
		Name:      "#40/2 POLY VİSKON RENKLİ",
		UnitPrice: 6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "113",
		Name:      "28/1 KESİK ELYAF DİCLE KAR TEKS",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "114",
		Name:      "28/1 EKRU VİSKON+28/1 VİSKON FLAM (19325 ) KVK",
		UnitPrice: 4.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "115",
		Name:      "26/1 POLYVİSKON TEK KAT",
		UnitPrice: 4.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "116",
		Name:      "32/1 KATYONİK VİSKON TEK KAT LYCRA",
		UnitPrice: 5.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "117",
		Name:      "30/1 KARDE ŞANTUK BEREN 46 ERDEM TEKSTİL",
		UnitPrice: 5.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "118",
		Name:      "28/2 MULİNE POLY VİSKON",
		UnitPrice: 5.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "119",
		Name:      "28/2 EKRU FLAM DİCLE",
		UnitPrice: 4.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "120",
		Name:      "*28/2 PES-VİS. LYCRA SİYAH",
		UnitPrice: 5.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "121",
		Name:      "28/1 PES VİS TEKKAT RENKLİ",
		UnitPrice: 4.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "122",
		Name:      "100f36  renkli  (batmaz boyatma)",
		UnitPrice: 3.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "123",
		Name:      "100f36 puntalı korteks",
		UnitPrice: 2.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "124",
		Name:      "190DN 211 f 97R YM EKRU LYCRA",
		UnitPrice: 3.09,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "125",
		Name:      "20/1 PAMUK",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "126",
		Name:      "20/1 KARDE PAMUK LİKRA",
		UnitPrice: 3.85,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "127",
		Name:      "20/1 PENYE PAMUK LİKRA",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "128",
		Name:      "16/1 O.E. PAMUK",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "129",
		Name:      "300 DENYE HAKİ BATMAZ BOYANAN",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "130",
		Name:      "100f36+40f1 polyester likralı KORTEKS",
		UnitPrice: 2.55,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "131",
		Name:      "34/2 28/1FLAMK.E.40/1 P.V.RENKLİ MUL.",
		UnitPrice: 4.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "132",
		Name:      "28/1 PV+28/1 KESİK E.FLAM    (28/2 FLAM)",
		UnitPrice: 4.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "133",
		Name:      "#40/2 LYCR RENKLİ POLYVİSK",
		UnitPrice: 7,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "134",
		Name:      "16/1..80/20 wool polyamit kaplanlar",
		UnitPrice: 18.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "135",
		Name:      "50/2 RENKLİ POLYVİSKON",
		UnitPrice: 7.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "136",
		Name:      "28/2 SİYAH+AZTEK+LYCRA (MULİNE)",
		UnitPrice: 5.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "137",
		Name:      "#40/2 LYCRA SİYAH ANT",
		UnitPrice: 6.35,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "138",
		Name:      "28/2 POLYVİSK+AZTEK MULİNE",
		UnitPrice: 5.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "139",
		Name:      "40/1 RENKLİ+28/1 FLAM BÜKÜMLÜ",
		UnitPrice: 5.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "140",
		Name:      "*28/2 POLYVİSKON LYCRA RENKLİ",
		UnitPrice: 6.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "141",
		Name:      "ATLANTİS SİYAH -MULİNE",
		UnitPrice: 5.35,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "142",
		Name:      "ATLANTİS RENKLİ MULİNE",
		UnitPrice: 5.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "143",
		Name:      "28/2 (S65+SİYAH)   MULİNE LYCRA",
		UnitPrice: 5.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{

		Code:      "144",
		Name:      "10/1 BUKLET KAVAK",
		UnitPrice: 6.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "145",
		Name:      "28/4 POLY VİSKON MULİNE(KAVAK)",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "146",
		Name:      "ALTINORDU SİYAH 10/1",
		UnitPrice: 4.85,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "147",
		Name:      "ALTINORDU RENKLİ 10/1",
		UnitPrice: 4.85,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "148",
		Name:      "150/dn polyester likra renkli",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "149",
		Name:      "50f72+20f1 PES .LİKRA KORTEKS",
		UnitPrice: 5.39,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "150",
		Name:      "#40/2 LYCRA  PV EKRU",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "151",
		Name:      "190f49 LYCRA SİYAH (150/48+40f1)",
		UnitPrice: 3.34,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "152",
		Name:      "40/2 KESİK ELYAF LYCRA",
		UnitPrice: 5.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "153",
		Name:      "40/1 ekru poly visk+28/1 kesik e. Flam",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "154",
		Name:      "28/2 İTHAL BÜKÜMLÜ (SAGNAM)",
		UnitPrice: 4.32,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "155",
		Name:      "28/4 2ekru+1siyah+5545 renk",
		UnitPrice: 4.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "156",
		Name:      "*28/2 LİKRA EKRU P.V.",
		UnitPrice: 5.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "157",
		Name:      "7/1 PAMUK FLAM GÜRAY",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "158",
		Name:      "28/2 ATLANTİS LİKRA SİYAH",
		UnitPrice: 5.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "159",
		Name:      "28/2 ATLANTİS LİKRA RENKLİ",
		UnitPrice: 5.95,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "160",
		Name:      "28/2 MOZART LİKRA SİYAH",
		UnitPrice: 5.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "161",
		Name:      "28/2 MOZART LİKRA RENKLİ",
		UnitPrice: 5.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "162",
		Name:      "5 NM MAKARNA (SİYAH-GRİ) KAVAK",
		UnitPrice: 5.15,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "163",
		Name:      "28/4 DÜZ MULİNE RENKLİ+EKRU",
		UnitPrice: 6.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "164",
		Name:      "14/1 KAPLANLAR(40 WOOL 60 POLYAMİD",
		UnitPrice: 8.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "165",
		Name:      "40/1 6103+28/1 İNT 1015 MÜLİNE LİKRA",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "166",
		Name:      "40/1 6103 +28/1 İNT 1015 MULİNE PV",
		UnitPrice: 6.1,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "167",
		Name:      "28/2 BOYALI PV.19292 İÇİN KİREMİT",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "168",
		Name:      "75/72 fdy yarımat Z 800 tur (Memory SİYAH çözgü",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "169",
		Name:      "28/1 RENKLİ+28/1 KESİK ELYAF FLAM(KAVAK) 18105",
		UnitPrice: 4.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "170",
		Name:      "28/1 Kırçıllı(KONFÜÇYÜS)+28/1 EKRU-DİCLE-- 19323",
		UnitPrice: 4.55,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "171",
		Name:      "28/1 KIRÇILLI(KONFÜÇYÜS+28/1 RENKLİ P.V.",
		UnitPrice: 5.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "172",
		Name:      "68/24 S 600 tur Memory montluk Atkılığı BEYAZ",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "173",
		Name:      "20/1 POLYPAMUK 50/50 KARDE",
		UnitPrice: 3.55,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "174",
		Name:      "30/1 PAMUKPOLYESTER 50/50 KARDE",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "175",
		Name:      "40/1 POLYCOTON 50/50 PENYE",
		UnitPrice: 3.9,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "176",
		Name:      "45/1 POLYCOTON PENYE 50/50",
		UnitPrice: 4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "177",
		Name:      "305/120 YÜN EFECT POLYESTER-İSOTEX",
		UnitPrice: 3.4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "178",
		Name:      "28/1 PV RENKLİ+26/1 KETEN MULİNE DİCLE(19332)",
		UnitPrice: 12,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "179",
		Name:      "28/1 RENKLİ+28/1 KESİK ELYAF DÜZ BÜKÜM MULİNE -KVK",
		UnitPrice: 4.65,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "180",
		Name:      "50/1 PENYE COMPACT ALMER KAYSERİ",
		UnitPrice: 5.85,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "181",
		Name:      "140/148+40 BEYAZ NAYLON LİKRA SINGLE COVER OMAFİL",
		UnitPrice: 9.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "182",
		Name:      "50/1 AVUSTURYA HARMANI PENYE COMPACT",
		UnitPrice: 6.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "183",
		Name:      "70/1-72 FLAM+ELESTAN(NAYLON 6) OMAFİL",
		UnitPrice: 12,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "184",
		Name:      "70/1 24 FLAM+70/1 68 NYL 6+70 DENYE",
		UnitPrice: 10,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "185",
		Name:      "75f 36 +40 elestan (115d 127f 37 r ekru tex.COMBAT",
		UnitPrice: 4.14,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "186",
		Name:      "150 D 167F 48 R YARIMAT TEXTÜRE SİYAH",
		UnitPrice: 2.46,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "187",
		Name:      "100 dn likra (140f37) 100f 36+40f1  KORTEKS",
		UnitPrice: 2.47,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "188",
		Name:      "20/1 PENYE PAMUK- ERDEM TEKSTİL USA 100%",
		UnitPrice: 4.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "189",
		Name:      "17 Nm Fantezi pes KAVAK tekstil (20129)",
		UnitPrice: 5.3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "190",
		Name:      "3OO dn LİKRA BEYAZ (300f72+70f1) korteks",
		UnitPrice: 1.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "191",
		Name:      "300 dn LİKRA   SİYAH (300f96+70f1) korteks",
		UnitPrice: 2.2,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "192",
		Name:      "28/2 RENKLİ MULİNE -KAVAK  (20138 )",
		UnitPrice: 5.05,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "193",
		Name:      "28/1 6 İĞNE LASE VEYSEL ÇALDIRAN",
		UnitPrice: 4.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "194",
		Name:      "3 NUMARA BUKLET SİYAH",
		UnitPrice: 5.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "195",
		Name:      "10/2 SHTRAYGARN DURSUN BÜKÜM 2.75 TL",
		UnitPrice: 4.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "196",
		Name:      "300 likra ekru korteks",
		UnitPrice: 2.54,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "197",
		Name:      "150/96 img puntalı POLYTEKS",
		UnitPrice: 2.19,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "198",
		Name:      "16/1  50/50 PAMUK POLYESTER",
		UnitPrice: 2.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "199",
		Name:      "16/1 65/35 PAMUK POLYESTER",
		UnitPrice: 3,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "200",
		Name:      "16/1  35/65 PAMUK POLYESTER",
		UnitPrice: 2.75,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "201",
		Name:      "28/1 EKRU POLYVİSKON",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "202",
		Name:      "28/1 EKRU POLYVİSKON TEKKAT",
		UnitPrice: 3.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "203",
		Name:      "50/2 EKRU DİCLE NORMAL",
		UnitPrice: 5.8,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "204",
		Name:      "28/1 SİYAH P.V+ 75 DNY BEYAZ(MULİNE )",
		UnitPrice: 4.95,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "205",
		Name:      "40/2 SİYAH +S 848 RENKLİ",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "206",
		Name:      "10/1 PAMUK AMERİKAN",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "207",
		Name:      "36/1 COMPACT PAMUK",
		UnitPrice: 4.6,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "208",
		Name:      "150/288+40 LİKRA  SINGLE COVER USTA GİPE",
		UnitPrice: 5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "209",
		Name:      "40/1 PV S59 +40/1 S65 DÜZ MULİNE  ( 21235 )",
		UnitPrice: 5.95,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "210",
		Name:      "40/1 MULİNE LİKRA S59+S65          (21235 )",
		UnitPrice: 6.55,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "211",
		Name:      "40/1 POLYVİSKON TEKKAT",
		UnitPrice: 5.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "212",
		Name:      "40/1 POLYVİS+40/1 POLVİS MULİNE +LİKRA RENKLİ",
		UnitPrice: 7.25,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "213",
		Name:      "40/1 POLYVİSKON+40/1 POLYVİSKON RENKLİ DÜZ BÜKÜM",
		UnitPrice: 6.5,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "214",
		Name:      "300 DENYE YÜN EFECT LİLİ   BARIŞ FİLAMENT",
		UnitPrice: 4,
	}
	config.DB.Create(&yarns)

	yarns = models.Yarns{
		Code:      "215",
		Name:      "150 DN RENKLİ POLYESTER LİSA İPLİK",
		UnitPrice: 3.5,
	}
	config.DB.Create(&yarns)

}
