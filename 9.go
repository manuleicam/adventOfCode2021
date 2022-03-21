package main

import (
	"fmt"
	"sort"
	"strconv"
)

type deep struct {
	deepLenght     int
	somethingLower bool
	checkBasin     bool
}

func main() {

	lenInput := len(Input_list)

	heatMap := make([][]deep, lenInput)

	totLowPoints := 0

	for i := 0; i < lenInput; i++ {

		heatMap[i] = stringToArray(Input_list[i])

	}

	for y, row := range heatMap {

		for x, _ := range row {
			//fmt.Println(x, heatMap[y][x])
			//if !heatMap[y][x].somethingLower { // if checkLow is true means that the number is either bigger than the one behind or above
			totLowPoints += checkLowPoint(&heatMap, x, y)
			//}
			//fmt.Println(heatMap[y][x])
			//fmt.Println(totLowPoints)

		}
	}

	fmt.Println(totLowPoints)

	basinSize := 0
	basinTots := make([]int, 0)
	for y, row := range heatMap {

		for x, hole := range row {

			if !hole.somethingLower {
				basinSize++
				getBasinSize(&heatMap, x, y, &basinSize)
				//fmt.Println(basinSize)
			}

			if basinSize != 0 {
				basinTots = append(basinTots, basinSize)
			}
			basinSize = 0

		}
	}

	fmt.Println(getThebiggest(basinTots, 3))
	res := 1
	for _, val := range getThebiggest(basinTots, 3) {
		res *= val
	}

	fmt.Println(res)

}

func getThebiggest(basinTots []int, numberOfReturns int) []int {

	sort.Ints(basinTots)

	basinTots = basinTots[len(basinTots)-(numberOfReturns) : len(basinTots)]

	return basinTots

}

func checkIfSmallerAround(heatMap *[][]deep, x int, y int) int {
	return 0
}

func catchOutBoundExp() {
	if r := recover(); r != nil {
		//return 0
	}
}

func getBasinSize(heatMap *[][]deep, x int, y int, basinSize *int) {

	checkSmaller(heatMap, x+1, y, (*heatMap)[y][x].deepLenght, basinSize)

	checkSmaller(heatMap, x-1, y, (*heatMap)[y][x].deepLenght, basinSize)

	checkSmaller(heatMap, x, y+1, (*heatMap)[y][x].deepLenght, basinSize)

	checkSmaller(heatMap, x, y-1, (*heatMap)[y][x].deepLenght, basinSize)

}

func checkSmaller(heatMap *[][]deep, x int, y int, firstPost int, basinSize *int) {

	defer catchOutBoundExp()

	//fmt.Println((*heatMap)[y][x].deepLenght > firstPost)

	if (*heatMap)[y][x].deepLenght > firstPost && (*heatMap)[y][x].deepLenght < 9 && !(*heatMap)[y][x].checkBasin {
		*basinSize++
		(*heatMap)[y][x].checkBasin = true
		getBasinSize(heatMap, x, y, basinSize)
	}

	//return 0

}

func checkLowPoint(heatMap *[][]deep, xCoord int, yCoord int) int {

	lowPointVal := 0

	// last line
	if yCoord == len((*heatMap))-1 {
		lowPointVal += checkLastLine(heatMap, xCoord, yCoord) + 1
	} else if xCoord == len((*heatMap)[yCoord])-1 {
		lowPointVal += checkLastCol(heatMap, xCoord, yCoord) + 1
	} else {
		lowPointVal += checkOtherRegs(heatMap, xCoord, yCoord) + 1
	}

	return lowPointVal

}

func checkOtherRegs(heatMap *[][]deep, xCoord int, yCoord int) int {

	smaller := true

	if (*heatMap)[yCoord][xCoord].deepLenght < (*heatMap)[yCoord][xCoord+1].deepLenght {

		(*heatMap)[yCoord][xCoord+1].somethingLower = true

	} else {
		smaller = false
	}

	if (*heatMap)[yCoord][xCoord].deepLenght < (*heatMap)[yCoord+1][xCoord].deepLenght {
		(*heatMap)[yCoord+1][xCoord].somethingLower = true
	} else {
		smaller = false
	}

	if smaller && !(*heatMap)[yCoord][xCoord].somethingLower {
		return (*heatMap)[yCoord][xCoord].deepLenght
	}

	(*heatMap)[yCoord][xCoord].somethingLower = true
	return -1

}

func checkLastCol(heatMap *[][]deep, xCoord int, yCoord int) int {

	if yCoord == len((*heatMap))-1 {
		return -1
	}

	if (*heatMap)[yCoord][xCoord].deepLenght < (*heatMap)[yCoord+1][xCoord].deepLenght {
		(*heatMap)[yCoord+1][xCoord].somethingLower = true
		if !(*heatMap)[yCoord][xCoord].somethingLower {
			return (*heatMap)[yCoord][xCoord].deepLenght
		}

	}

	(*heatMap)[yCoord][xCoord].somethingLower = true
	return -1

}

func checkLastLine(heatMap *[][]deep, xCoord int, yCoord int) int {

	if xCoord == len((*heatMap)[yCoord])-1 {
		if !(*heatMap)[yCoord][xCoord].somethingLower {
			return (*heatMap)[yCoord][xCoord].deepLenght
		} else {
			return -1
		}
	}

	if (*heatMap)[yCoord][xCoord].deepLenght < (*heatMap)[yCoord][xCoord+1].deepLenght {
		(*heatMap)[yCoord][xCoord+1].somethingLower = true
		if !(*heatMap)[yCoord][xCoord].somethingLower {
			return (*heatMap)[yCoord][xCoord].deepLenght
		}
	}

	(*heatMap)[yCoord][xCoord].somethingLower = true
	return -1

}

func stringToArray(numbers string) []deep {

	intArr := make([]deep, 0)

	for _, char := range numbers {

		deepHole := deep{0, false, false}

		number, _ := strconv.Atoi(string(char))
		deepHole.deepLenght = number

		intArr = append(intArr, deepHole)

	}

	return intArr

}

var Input_list_teste = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678"}

var Input_list = []string{
	"8654434789432446987654321056789235678953245798764212456789656568977654232457898754567898765431335999",
	"8743125678931234599987643238993123567894345989542101234678943459865432101345789643479999878540129877",
	"9651056799320125678998764369895644678965956976543233458799202568986574519499996532445989998761298956",
	"6543234789632347899219865456789795789879899898754755567893212456998986798987897621234678979872987645",
	"7887645696543456799423989978998989893989767789865676679995343689879497986596789432349899865989876434",
	"8998967789999577897999799999457678902497645699876799798789459798767329876455678943456798764699987569",
	"9999878899878988956986678953234567893989534578997998987689567976543213984323589894967899643209898978",
	"0987999999767899549875439870146789999875323466899886596578979987954201987312398789898978965799769989",
	"2996567899656795423964321981269893987983212345689765434499898799865332975201297598789467979987654292",
	"9875456998945789019875992984378942496543201456789897521345679689965449964319976465694356797998893101",
	"9984335987939892125999879876459921998654312345899998787657789578986598765698764344589578956789989392",
	"8893219876898993234569765987867899899765543459999659999789993459997679877999985233467999345999978989",
	"7789109765456789349996543398978976799876654678998544569896432367989789998999876101256789134579768678",
	"6689997653337868998889421239989235789988796899987632355999553489879993459789987212345693013498657569",
	"4577898942123456987679310298790124678999987999874821254598978598767892345678984345678932125987545456",
	"3456999866345568976589429397651245789547898998763210123987997679998921556899976456799894296986434347",
	"1467898765468689876459598999542356898936799899854321999576789889349743677999987887898795989976321234",
	"2348999976567899765398967788943467897425656789866459878455678993219654569887899998954689879865430125",
	"3567899987898998543257956697894879986314346995987999867234599995398765698776598769323569768976521289",
	"4678997898989987664134943476789989875401235794398987956145698789999876987632459954201678956965434378",
	"9789976999979876543245792445678998994322346789129986541016799697899988998543569877352789249896776456",
	"9999865698765988664356890237889767895443456791099998432123978546789899899987679865467892198789887768",
	"9898954349994398789867891346893456986775767892989876549939765434679765678998989976878999299679998979",
	"8787894210989239898978932457922345697896889999978987897899964323579654567979493197989998989597899995",
	"8676795349879199976989654568901234798998992698967998986799899104568963459854321019999876679476889754",
	"6555689598968989765698765689429349899019893497658999875689778915689654598765432156898995432345678943",
	"9434599697657978974569897897698998942198789985432987664573567923798975999896543247997986541236789432",
	"7323458986543767893789998999987987893989698764321096543212458954567899899987755459986598762445699999",
	"5434567894322459954699769998976446999876569875463497854101348965678989789999876598765439643596789888",
	"9567898943210347895987654987654235798654429876775698543212457899789879578954997899874329764589994667",
	"8979969768731236789999432196542125699543212997987899658323468949899765466893298999994210975678923556",
	"7993459879544345699987691029931013987654343569998999987569678957998754345789129598765499796789012345",
	"6989567988655468979999989298754323498785654678999998998678989767897643236789012499876989698995433456",
	"4568978999766579757989978989895434679896765989789997899799599898998894356892134989987976549996764567",
	"3657899879878989645679765678976656989998876797569876789895431999979965467894349879799865439987765698",
	"2345789764989596436798654578989867893459987895498765656989549987659879878965698765569876328999896789",
	"1456799753295432224689768999993978921237998902359854245678998765432989999979999843459876437943987890",
	"2345678932129421015699879789932989890126899213479863136899439987554599887897897654699976546892198931",
	"3456989321098939123689989645791296789245698924589854256789921398668999765256798865789987987893029642",
	"4569895443987698934579896535899345678959997945699767767896790989779998754345789996896598998932198763",
	"6998789569986467896989765423678976799898786899999878978945679879889999865456999989965439239993499884",
	"9895679698765356998992989513467899899768645997898989989896798765996889978597898677989910198989989995",
	"8794598976573245789431999764578964999853129876797699995789899854324578987679923456898799987875679976",
	"7683987996432127897699898975789653298654012965679539876896998943212459899789412346789678976664778988",
	"7572176789321046799988767896898762198763129754568920987987987892101398799897323467896589765453467999",
	"5421065679433125679876459987899873987654298643457891299899876999212987697976534568997678954342349654",
	"6572123999745234567989578998979989899765459432346999976789765878929876585989745699898989765101578943",
	"7683939789657656778998689879567898769876798621245798745899874367949965464599656989789996543212567892",
	"8899895689969867899498799865449987954987987540134689534699543248998754223498769875678997655423456789",
	"9989764599899879994319899979323496543498997621234578945988992124679897101239898764989998787935678995",
	"9876543487789998789423989998934987662349898742345679899876789013498765312457999433897899899987889453",
	"8985432276568987678935678987895999871245698755457789767945678924599896543468998421766789994398994342",
	"7654310123457976567896789896799876997659989866669997643134689536789987654679877610545798689459653210",
	"8766521245667895476789898785989875349798879878778999721025678949892398785989765431236789569599764332",
	"9987632345679954345699987684678983235987954989989899832127889898921569896798999932347992398989875964",
	"9898745456798765456789976543796899345976543293496798753456799787890199997896778893458954987878989875",
	"8789876767899876567998765432345778957987652101245987654567987656789989998985867789569869876567899986",
	"6577997978945989678979876321234567969999763412356799867689976347999978999874345678978989795456989987",
	"5456798999533499989569998974346679878969874329867899998799865457899766798765456789989998674345678998",
	"4399899764312579392398989765458989989349876498978999019892976568987654599876587891299987553239789989",
	"5987999765201679209987679976569990191245997987899998934943987689998765789987679943479999432138999879",
	"9896789954312469398976567899678999910123999876799987996799998799879886899998789656567998521016799968",
	"8765679895423578987865455678989998899239897545689976789987969899865998998999899767678987672125678956",
	"7654457789534679876644344579998987678999765434578965678986856999954529987896964979989898743234599345",
	"9643235678945798765432123678997864564789894323589654567975445698543219986965432193199799954355891256",
	"8964015989976899976321012389986773243698998645789765688983234987654397654899542095987659865476789345",
	"7895123799988919875432124567895432101566789787899878899954195999978598943678953989999549878989999956",
	"6899434568999901986553234789976544512345899899974989967899989899988679862567899878999532999893139899",
	"5678965689323892987794545898989796923556999998765697656898576789999897531459988867987321246793098788",
	"4599987899904679598865666987899989894689989999896789245987455567899986432398876756896410235932199697",
	"3989998978899889379979897896789978789799879899987892168986343456789497643987754345894321376893988575",
	"2467899565678996467998959945898769699899765789899943239874212398894398959876543256795434487975976454",
	"1248932354578999589876543234987854569998654564767896549954101989943239767988654467976578598999994323",
	"2867890123499998995998655129876543998999972123556897698763222478932129899998769679987889789678986414",
	"3988943234689987894698776334976642866789893012346798999986434567893434998999878798998996993569864202",
	"4799654356799976993999897957985421345693799224897929899996545678997549997898989897799875401698765313",
	"5679765667998765789898999898995430235954568935689909698998786799987698986567899966698764213899897455",
	"9889876898987764679767898789876554397895679876789898567899997978998997754378987654559879984569998767",
	"6993987939876543198657999678989765498989789987898787456789998966789876543245699863245998976898769878",
	"5432398923994321098767987567999896789777995498999656345678969345999865420135679943123987897899954989",
	"4321299012986532129878986459899987892656894319689843210199653236791986434678989991019876899939869899",
	"6430989194597643534989598349789998921245689323598764322388964378892397956789999989198765998921989789",
	"6549879989798757657896459998689999210123459934569765434567965567954498867899989678987654357899997678",
	"7659766467989898798965349887567894339735567895678976545678996678966789988999976589996542123987754567",
	"8798654245678939989876298765456995687545689996789987898789789799987896599998665457789753239876543468",
	"9899650126789549879989349656245789898689797987992198999895678999998987434987574336678968959971012389",
	"9998743245679698767898956932127899959789986498993999998954789998989876323998432124569999898765433499",
	"8999854345791987656767899873246789349899987329989899876543599896579874209876321018678989789976785567",
	"7998765456890989433456799965368990234999995439878788985432456789459964313995432167899768679987896679",
	"6899979877899874322369899876878921345678987598764567896431235679298765939989943245987654568998999793",
	"5798989988989965101998932987899432456899799987653468984320356789109899898767894556895723567899599892",
	"4376799999569754319887891298976543597897698765432157976547589893212999765458789698954312568923478901",
	"3284567893498765498776989999597654789989539976754245697969678954923998774345679899985623459944569892",
	"2123458912349879989655678895498799999878920989985358789899899769899876543234678999876535567895798789",
	"3434567893499998976544686789579988998767899899876469996799909898798765442123789879996546778989998678",
	"4645678999978977965432345678989867899843498789987896565989324987679876321012899868997657899578965459",
	"6657899498769356996544456789898755689765987649798965434878939878568985432124789656569767933467954334",
	"8767934349653249889767567898759434567986796536679976326567998765457976543234578942349898912356892123",
	"9899321296542135679889798999843215789987987424567893212456987654345987654575678931235969901867963344",
	"9965410987321024589999899998765676891299876535678975324667897643245698867686789210123456892378954467"}
