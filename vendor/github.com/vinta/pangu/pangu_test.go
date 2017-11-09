package pangu_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/vinta/pangu"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type PanguTestSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPanguTestSuite(t *testing.T) {
	suite.Run(t, new(PanguTestSuite))
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func md5Of(filename string) string {
	var result []byte

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	checkError(err)

	checksum := hex.EncodeToString(hash.Sum(result))

	return checksum
}

func (suite *PanguTestSuite) TestSpacingText() {
	suite.Equal(pangu.SpacingText(`新八的構造成分有95%是眼鏡、3%是水、2%是垃圾`), `新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`)
	suite.Equal(pangu.SpacingText(`所以,請問Jackey的鼻子有幾個?3.14個!`), `所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`)
	suite.Equal(pangu.SpacingText(`JUST WE就是JUST WE，既不偉大也不卑微！`), `JUST WE 就是 JUST WE，既不偉大也不卑微！`)
	suite.Equal(pangu.SpacingText(`搭載MP3播放器，連續播放時數最長達到124小時的超強利刃……菊一文字RX-7!`), `搭載 MP3 播放器，連續播放時數最長達到 124 小時的超強利刃…… 菊一文字 RX-7!`)
	suite.Equal(pangu.SpacingText(`V`), `V`)
}

func (suite *PanguTestSuite) TestLatin1Supplement() {
	suite.Equal(pangu.SpacingText(`中文Ø漢字`), `中文 Ø 漢字`)
	suite.Equal(pangu.SpacingText(`中文 Ø 漢字`), `中文 Ø 漢字`)
}

func (suite *PanguTestSuite) TestGeneralPunctuation() {
	suite.Equal(pangu.SpacingText(`中文•漢字`), `中文 • 漢字`)
	suite.Equal(pangu.SpacingText(`中文 • 漢字`), `中文 • 漢字`)
}

func (suite *PanguTestSuite) TestNumberForms() {
	suite.Equal(pangu.SpacingText(`中文Ⅶ漢字`), `中文 Ⅶ 漢字`)
	suite.Equal(pangu.SpacingText(`中文 Ⅶ 漢字`), `中文 Ⅶ 漢字`)
}

func (suite *PanguTestSuite) TestCJKRadicalsSupplement() {
	suite.Equal(pangu.SpacingText(`abc⻤123`), `abc ⻤ 123`)
	suite.Equal(pangu.SpacingText(`abc ⻤ 123`), `abc ⻤ 123`)
}

func (suite *PanguTestSuite) TestKangxiRadicals() {
	suite.Equal(pangu.SpacingText(`abc⾗123`), `abc ⾗ 123`)
	suite.Equal(pangu.SpacingText(`abc ⾗ 123`), `abc ⾗ 123`)
}

func (suite *PanguTestSuite) TestHiragana() {
	suite.Equal(pangu.SpacingText(`abcあ123`), `abc あ 123`)
	suite.Equal(pangu.SpacingText(`abc あ 123`), `abc あ 123`)
}

func (suite *PanguTestSuite) TestKatakana() {
	suite.Equal(pangu.SpacingText(`abcア123`), `abc ア 123`)
	suite.Equal(pangu.SpacingText(`abc ア 123`), `abc ア 123`)
}

func (suite *PanguTestSuite) TestBopomofo() {
	suite.Equal(pangu.SpacingText(`abcㄅ123`), `abc ㄅ 123`)
	suite.Equal(pangu.SpacingText(`abc ㄅ 123`), `abc ㄅ 123`)
}

func (suite *PanguTestSuite) TestEnclosedCJKLettersAndMonths() {
	suite.Equal(pangu.SpacingText(`abc㈱123`), `abc ㈱ 123`)
	suite.Equal(pangu.SpacingText(`abc ㈱ 123`), `abc ㈱ 123`)
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographsExtensionA() {
	suite.Equal(pangu.SpacingText(`abc㐂123`), `abc 㐂 123`)
	suite.Equal(pangu.SpacingText(`abc 㐂 123`), `abc 㐂 123`)
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographs() {
	suite.Equal(pangu.SpacingText(`abc丁123`), `abc 丁 123`)
	suite.Equal(pangu.SpacingText(`abc 丁 123`), `abc 丁 123`)
}

func (suite *PanguTestSuite) TestCJKCompatibilityIdeographs() {
	suite.Equal(pangu.SpacingText(`abc車123`), `abc 車 123`)
	suite.Equal(pangu.SpacingText(`abc 車 123`), `abc 車 123`)
}

func (suite *PanguTestSuite) TestTilde() {
	suite.Equal(pangu.SpacingText(`前面~後面`), `前面~ 後面`)
	suite.Equal(pangu.SpacingText(`前面 ~ 後面`), `前面 ~ 後面`)
	suite.Equal(pangu.SpacingText(`前面~ 後面`), `前面~ 後面`)
}

func (suite *PanguTestSuite) TestBackQuote() {
	suite.Equal("前面 ` 後面", pangu.SpacingText("前面`後面"))
	suite.Equal("前面 ` 後面", pangu.SpacingText("前面 ` 後面"))
	suite.Equal("前面 ` 後面", pangu.SpacingText("前面` 後面"))
}

func (suite *PanguTestSuite) TestExclamationMark() {
	suite.Equal(pangu.SpacingText(`前面!後面`), `前面! 後面`)
	suite.Equal(pangu.SpacingText(`前面 ! 後面`), `前面 ! 後面`)
	suite.Equal(pangu.SpacingText(`前面! 後面`), `前面! 後面`)
}

func (suite *PanguTestSuite) TestAt() {
	// https://twitter.com/vinta
	suite.Equal(pangu.SpacingText(`前面@vinta後面`), `前面 @vinta 後面`)
	suite.Equal(pangu.SpacingText(`前面 @vinta 後面`), `前面 @vinta 後面`)

	// http://weibo.com/vintalines
	suite.Equal(pangu.SpacingText(`前面@陳上進 後面`), `前面 @陳上進 後面`)
	suite.Equal(pangu.SpacingText(`前面 @陳上進 後面`), `前面 @陳上進 後面`)
	suite.Equal(pangu.SpacingText(`前面 @陳上進tail`), `前面 @陳上進 tail`)

	// TODO
	// suite.Equal(pangu.SpacingText(`陳上進@地球`), `陳上進@地球`)
}

func (suite *PanguTestSuite) TestHash() {
	suite.Equal(pangu.SpacingText(`前面#H2G2後面`), `前面 #H2G2 後面`)
	suite.Equal(pangu.SpacingText(`前面#銀河便車指南 後面`), `前面 #銀河便車指南 後面`)
	suite.Equal(pangu.SpacingText(`前面#銀河便車指南tail`), `前面 #銀河便車指南 tail`)
	suite.Equal(pangu.SpacingText(`前面#銀河公車指南 #銀河拖吊車指南 後面`), `前面 #銀河公車指南 #銀河拖吊車指南 後面`)

	suite.Equal(pangu.SpacingText(`前面#H2G2#後面`), `前面 #H2G2# 後面`)
	suite.Equal(pangu.SpacingText(`前面#銀河閃電霹靂車指南#後面`), `前面 #銀河閃電霹靂車指南# 後面`)
}

func (suite *PanguTestSuite) TestDollar() {
	suite.Equal(pangu.SpacingText(`前面$後面`), `前面 $ 後面`)
	suite.Equal(pangu.SpacingText(`前面 $ 後面`), `前面 $ 後面`)

	suite.Equal(pangu.SpacingText(`前面$100後面`), `前面 $100 後面`)

	// TODO
	// suite.Equal(pangu.SpacingText(`前面$一百塊 後面`), `前面 $一百塊 後面`)
}

func (suite *PanguTestSuite) TestPercent() {
	suite.Equal(pangu.SpacingText(`前面%後面`), `前面 % 後面`)
	suite.Equal(pangu.SpacingText(`前面 % 後面`), `前面 % 後面`)

	suite.Equal(pangu.SpacingText(`前面100%後面`), `前面 100% 後面`)
}

func (suite *PanguTestSuite) TestCarat() {
	suite.Equal(pangu.SpacingText(`前面^後面`), `前面 ^ 後面`)
	suite.Equal(pangu.SpacingText(`前面 ^ 後面`), `前面 ^ 後面`)
}

func (suite *PanguTestSuite) TestAmpersand() {
	suite.Equal(pangu.SpacingText(`前面&後面`), `前面 & 後面`)
	suite.Equal(pangu.SpacingText(`前面 & 後面`), `前面 & 後面`)

	suite.Equal(pangu.SpacingText(`Vinta&Mollie`), `Vinta&Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta&陳上進`), `Vinta & 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進&Vinta`), `陳上進 & Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A&B的結果`), `得到一個 A&B 的結果`)
}

func (suite *PanguTestSuite) TestAsterisk() {
	suite.Equal(pangu.SpacingText(`前面*後面`), `前面 * 後面`)
	suite.Equal(pangu.SpacingText(`前面 * 後面`), `前面 * 後面`)

	suite.Equal(pangu.SpacingText(`Vinta*Mollie`), `Vinta*Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta*陳上進`), `Vinta * 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進*Vinta`), `陳上進 * Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A*B的結果`), `得到一個 A*B 的結果`)
}

func (suite *PanguTestSuite) TestParenthesis() {
	// suite.Equal(pangu.SpacingText(`前面(後面`), `前面 ( 後面`)
	// suite.Equal(pangu.SpacingText(`前面 ( 後面`), `前面 ( 後面`)

	// suite.Equal(pangu.SpacingText(`前面)後面`), `前面 ) 後面`)
	// suite.Equal(pangu.SpacingText(`前面 ) 後面`), `前面 ) 後面`)

	suite.Equal(pangu.SpacingText(`前面(中文123漢字)後面`), `前面 (中文 123 漢字) 後面`)
	suite.Equal(pangu.SpacingText(`前面(中文123)後面`), `前面 (中文 123) 後面`)
	suite.Equal(pangu.SpacingText(`前面(123漢字)後面`), `前面 (123 漢字) 後面`)
	suite.Equal(pangu.SpacingText(`前面(中文123漢字) tail`), `前面 (中文 123 漢字) tail`)
	suite.Equal(pangu.SpacingText(`head (中文123漢字)後面`), `head (中文 123 漢字) 後面`)
	suite.Equal(pangu.SpacingText(`head (中文123漢字) tail`), `head (中文 123 漢字) tail`)
}

func (suite *PanguTestSuite) TestMinus() {
	suite.Equal(pangu.SpacingText(`前面-後面`), `前面 - 後面`)
	suite.Equal(pangu.SpacingText(`前面 - 後面`), `前面 - 後面`)

	suite.Equal(pangu.SpacingText(`Vinta-Mollie`), `Vinta-Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta-陳上進`), `Vinta - 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進-Vinta`), `陳上進 - Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A-B的結果`), `得到一個 A-B 的結果`)
}

func (suite *PanguTestSuite) TestUnderscore() {
	suite.Equal(pangu.SpacingText(`前面_後面`), `前面_後面`)
	suite.Equal(pangu.SpacingText(`前面 _ 後面`), `前面 _ 後面`)
}

func (suite *PanguTestSuite) TestPlus() {
	suite.Equal(pangu.SpacingText(`前面+後面`), `前面 + 後面`)
	suite.Equal(pangu.SpacingText(`前面 + 後面`), `前面 + 後面`)

	suite.Equal(pangu.SpacingText(`Vinta+Mollie`), `Vinta+Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta+陳上進`), `Vinta + 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進+Vinta`), `陳上進 + Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A+B的結果`), `得到一個 A+B 的結果`)

	suite.Equal(pangu.SpacingText(`得到一個C++的結果`), `得到一個 C++ 的結果`)

	// TODO
	// suite.Equal(pangu.SpacingText(`得到一個A+的結果`), `得到一個 A+ 的結果`)
}

func (suite *PanguTestSuite) TestEqual() {
	suite.Equal(pangu.SpacingText(`前面=後面`), `前面 = 後面`)
	suite.Equal(pangu.SpacingText(`前面 = 後面`), `前面 = 後面`)

	suite.Equal(pangu.SpacingText(`Vinta=Mollie`), `Vinta=Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta=陳上進`), `Vinta = 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進=Vinta`), `陳上進 = Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A=B的結果`), `得到一個 A=B 的結果`)
}

func (suite *PanguTestSuite) TestBrace() {
	// suite.Equal(pangu.SpacingText(`前面{後面`), `前面 { 後面`)
	// suite.Equal(pangu.SpacingText(`前面 { 後面`), `前面 { 後面`)

	// suite.Equal(pangu.SpacingText(`前面}後面`), `前面 } 後面`)
	// suite.Equal(pangu.SpacingText(`前面 } 後面`), `前面 } 後面`)

	suite.Equal(pangu.SpacingText(`前面{中文123漢字}後面`), `前面 {中文 123 漢字} 後面`)
	suite.Equal(pangu.SpacingText(`前面{中文123}後面`), `前面 {中文 123} 後面`)
	suite.Equal(pangu.SpacingText(`前面{123漢字}後面`), `前面 {123 漢字} 後面`)
	suite.Equal(pangu.SpacingText(`前面{中文123漢字} tail`), `前面 {中文 123 漢字} tail`)
	suite.Equal(pangu.SpacingText(`head {中文123漢字}後面`), `head {中文 123 漢字} 後面`)
	suite.Equal(pangu.SpacingText(`head {中文123漢字} tail`), `head {中文 123 漢字} tail`)
}

func (suite *PanguTestSuite) TestBracket() {
	// suite.Equal(pangu.SpacingText(`前面[後面`), `前面 [ 後面`)
	// suite.Equal(pangu.SpacingText(`前面 [ 後面`), `前面 [ 後面`)

	// suite.Equal(pangu.SpacingText(`前面]後面`), `前面 ] 後面`)
	// suite.Equal(pangu.SpacingText(`前面 ] 後面`), `前面 ] 後面`)

	suite.Equal(pangu.SpacingText(`前面[中文123漢字]後面`), `前面 [中文 123 漢字] 後面`)
	suite.Equal(pangu.SpacingText(`前面[中文123]後面`), `前面 [中文 123] 後面`)
	suite.Equal(pangu.SpacingText(`前面[123漢字]後面`), `前面 [123 漢字] 後面`)
	suite.Equal(pangu.SpacingText(`前面[中文123漢字] tail`), `前面 [中文 123 漢字] tail`)
	suite.Equal(pangu.SpacingText(`head [中文123漢字]後面`), `head [中文 123 漢字] 後面`)
	suite.Equal(pangu.SpacingText(`head [中文123漢字] tail`), `head [中文 123 漢字] tail`)
}

func (suite *PanguTestSuite) TestPipe() {
	suite.Equal(pangu.SpacingText(`前面|後面`), `前面 | 後面`)
	suite.Equal(pangu.SpacingText(`前面 | 後面`), `前面 | 後面`)

	suite.Equal(pangu.SpacingText(`Vinta|Mollie`), `Vinta|Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta|陳上進`), `Vinta | 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進|Vinta`), `陳上進 | Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A|B的結果`), `得到一個 A|B 的結果`)
}

func (suite *PanguTestSuite) TestBackslash() {
	suite.Equal(pangu.SpacingText(`前面\後面`), `前面 \ 後面`)
}

func (suite *PanguTestSuite) TestColon() {
	suite.Equal(pangu.SpacingText(`前面:後面`), `前面: 後面`)
	suite.Equal(pangu.SpacingText(`前面 : 後面`), `前面 : 後面`)
	suite.Equal(pangu.SpacingText(`前面: 後面`), `前面: 後面`)
}

func (suite *PanguTestSuite) TestSemicolon() {
	suite.Equal(pangu.SpacingText(`前面;後面`), `前面; 後面`)
	suite.Equal(pangu.SpacingText(`前面 ; 後面`), `前面 ; 後面`)
	suite.Equal(pangu.SpacingText(`前面; 後面`), `前面; 後面`)
}

func (suite *PanguTestSuite) TestQuote() {
	// suite.Equal(pangu.SpacingText(`前面"後面`), `前面 " 後面`)
	// suite.Equal(pangu.SpacingText(`前面""後面`), `前面 "" 後面`)
	// suite.Equal(pangu.SpacingText(`前面" "後面`), `前面 " " 後面`)

	suite.Equal(pangu.SpacingText(`前面"中文123漢字"後面`), `前面 "中文 123 漢字" 後面`)
	suite.Equal(pangu.SpacingText(`前面"中文123"後面`), `前面 "中文 123" 後面`)
	suite.Equal(pangu.SpacingText(`前面"123漢字"後面`), `前面 "123 漢字" 後面`)
	suite.Equal(pangu.SpacingText(`前面"中文123漢字" tail`), `前面 "中文 123 漢字" tail`)
	suite.Equal(pangu.SpacingText(`head "中文123漢字"後面`), `head "中文 123 漢字" 後面`)
	suite.Equal(pangu.SpacingText(`head "中文123漢字" tail`), `head "中文 123 漢字" tail`)

	// \u201c and \u201d
	suite.Equal(pangu.SpacingText(`前面“中文123漢字”後面`), `前面 “中文 123 漢字” 後面`)
}

func (suite *PanguTestSuite) TestSingleQuote() {
	// suite.Equal(pangu.SpacingText(`前面'後面`), `前面 ' 後面`)
	// suite.Equal(pangu.SpacingText(`前面''後面`), `前面 '' 後面`)
	// suite.Equal(pangu.SpacingText(`前面' '後面`), `前面 ' ' 後面`)

	suite.Equal(pangu.SpacingText(`前面'中文123漢字'後面`), `前面 '中文 123 漢字' 後面`)
	suite.Equal(pangu.SpacingText(`前面'中文123'後面`), `前面 '中文 123' 後面`)
	suite.Equal(pangu.SpacingText(`前面'123漢字'後面`), `前面 '123 漢字' 後面`)
	suite.Equal(pangu.SpacingText(`前面'中文123漢字' tail`), `前面 '中文 123 漢字' tail`)
	suite.Equal(pangu.SpacingText(`head '中文123漢字'後面`), `head '中文 123 漢字' 後面`)
	suite.Equal(pangu.SpacingText(`head '中文123漢字' tail`), `head '中文 123 漢字' tail`)

	suite.Equal(pangu.SpacingText(`陳上進 likes 林依諾's status.`), `陳上進 likes 林依諾's status.`)
}

func (suite *PanguTestSuite) TestLessThan() {
	suite.Equal(pangu.SpacingText(`前面<後面`), `前面 < 後面`)
	suite.Equal(pangu.SpacingText(`前面 < 後面`), `前面 < 後面`)

	suite.Equal(pangu.SpacingText(`Vinta<Mollie`), `Vinta<Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta<陳上進`), `Vinta < 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進<Vinta`), `陳上進 < Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A<B的結果`), `得到一個 A<B 的結果`)

	suite.Equal(pangu.SpacingText(`前面<中文123漢字>後面`), `前面 <中文 123 漢字> 後面`)
	suite.Equal(pangu.SpacingText(`前面<中文123>後面`), `前面 <中文 123> 後面`)
	suite.Equal(pangu.SpacingText(`前面<123漢字>後面`), `前面 <123 漢字> 後面`)
	suite.Equal(pangu.SpacingText(`前面<中文123漢字> tail`), `前面 <中文 123 漢字> tail`)
	suite.Equal(pangu.SpacingText(`head <中文123漢字>後面`), `head <中文 123 漢字> 後面`)
	suite.Equal(pangu.SpacingText(`head <中文123漢字> tail`), `head <中文 123 漢字> tail`)
}

func (suite *PanguTestSuite) TestComma() {
	suite.Equal(pangu.SpacingText(`前面,後面`), `前面, 後面`)
	suite.Equal(pangu.SpacingText(`前面 , 後面`), `前面 , 後面`)
	suite.Equal(pangu.SpacingText(`前面, 後面`), `前面, 後面`)
}

func (suite *PanguTestSuite) TestGreaterThan() {
	suite.Equal(pangu.SpacingText(`前面>後面`), `前面 > 後面`)
	suite.Equal(pangu.SpacingText(`前面 > 後面`), `前面 > 後面`)

	suite.Equal(pangu.SpacingText(`Vinta>Mollie`), `Vinta>Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta>陳上進`), `Vinta > 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進>Vinta`), `陳上進 > Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A>B的結果`), `得到一個 A>B 的結果`)
}

func (suite *PanguTestSuite) TestPeriod() {
	suite.Equal(pangu.SpacingText(`前面.後面`), `前面. 後面`)
	suite.Equal(pangu.SpacingText(`前面 . 後面`), `前面 . 後面`)
	suite.Equal(pangu.SpacingText(`前面. 後面`), `前面. 後面`)

	// … is \u2026
	suite.Equal(pangu.SpacingText(`前面…後面`), `前面… 後面`)
	suite.Equal(pangu.SpacingText(`前面……後面`), `前面…… 後面`)
}

func (suite *PanguTestSuite) TestQuestionMark() {
	suite.Equal(pangu.SpacingText(`前面?後面`), `前面? 後面`)
	suite.Equal(pangu.SpacingText(`前面 ? 後面`), `前面 ? 後面`)
	suite.Equal(pangu.SpacingText(`前面? 後面`), `前面? 後面`)
}

func (suite *PanguTestSuite) TestSlash() {
	suite.Equal(pangu.SpacingText(`前面/後面`), `前面 / 後面`)
	suite.Equal(pangu.SpacingText(`前面 / 後面`), `前面 / 後面`)

	suite.Equal(pangu.SpacingText(`Vinta/Mollie`), `Vinta/Mollie`)
	suite.Equal(pangu.SpacingText(`Vinta/陳上進`), `Vinta / 陳上進`)
	suite.Equal(pangu.SpacingText(`陳上進/Vinta`), `陳上進 / Vinta`)

	suite.Equal(pangu.SpacingText(`得到一個A/B的結果`), `得到一個 A/B 的結果`)

	// TODO
	// suite.Equal(pangu.SpacingText(`陳上進/Vinta/Mollie`), `陳上進 / Vinta / Mollie`)
}

func (suite *PanguTestSuite) TestSpacingFile() {
	input := "_fixtures/test_file.txt"
	output := "_fixtures/test_file.pangu.txt"

	fw, err := os.Create(output)
	checkError(err)
	defer fw.Close()

	err = pangu.SpacingFile(input, fw)
	suite.Nil(err)
	suite.Equal(md5Of(output), md5Of("_fixtures/test_file.expected.txt"))
}

func (suite *PanguTestSuite) TestSpacingFileNoNewlineAtEOF() {
	input := "_fixtures/test_file_no_eof_newline.txt"
	output := "_fixtures/test_file_no_eof_newline.pangu.txt"

	fw, err := os.Create(output)
	checkError(err)
	defer fw.Close()

	err = pangu.SpacingFile(input, fw)
	suite.Nil(err)
	suite.Equal(md5Of(output), md5Of("_fixtures/test_file_no_eof_newline.expected.txt"))
}

func (suite *PanguTestSuite) TestSpacingFileNoSuchFile() {
	input := "_fixtures/none.exist"

	err := pangu.SpacingFile(input, ioutil.Discard)
	suite.EqualError(err, "open _fixtures/none.exist: no such file or directory")
}
